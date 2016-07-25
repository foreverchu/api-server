package partySrv

import (
	"errors"
	"strings"

	"models"

	"github.com/astaxie/beego/orm"

	"github.com/chinarun/utils"
	"strconv"
)

const (
	PARTY_TYPE_OWN    = 1
	PARTY_TYPE_SIGN   = 2
	PARTY_TYPE_FOLLOW = 3
)

const (
	DEF_VALID_STATE_WAIT    = 0
	DEF_VALID_STATE_PASS    = 1
	DEF_VALID_STATE_NO_PASS = 2
	DEF_VALID_STATE_ALL     = 3
)

const (
	PARTY_PHOTO_TYPE_PARTY       = 1
	PARTY_PHOTO_TYPE_GAME        = 2
	PARTY_PHOTO_TYPE_PARTY_ROUTE = 3
)

// Usage:
// plq := partySrv.NewPartyListQuery().PageNo(1).PageSize(10).City("上海")
// order := make(map[string]bool{
//		"reg_time":true,
//		"close_time":false,
//	})
//
// plq.Order(order)
// err := plq.Do()
// if err != nil {
//		return plq.ErrorMsg()
//	}
// result := plq.Result()
//
type PartyListQuery struct {
	//是否查询过了
	did_query bool

	result    []*models.PartyAndDetail
	total     int64
	page_no   int
	page_size int
	err       error
	msg       string

	// sql涉及到的table
	queryTables string
	// 统计个数sql
	queryCount string
	// sql的where
	queryWhere string
	// sql的order
	queryOrder string
	// 参数
	params []interface{}
}

func NewPartyListQuery() *PartyListQuery {
	plq := PartyListQuery{}
	plq.Init()

	return &plq
}

func (plq *PartyListQuery) Init() *PartyListQuery {
	plq.queryTables = "select p.*, d.*, p2.url as photo_url FROM " + models.DB_TABLE_PARTY + " as p " +
		" inner join " + models.DB_TABLE_PARTY_DETAIL + " as d on p.id=d.party_id " +
		" left join " + models.DB_TABLE_PHOTO + " as p2 on p.id=p2.rel_id and p2.type = " + strconv.Itoa(PARTY_PHOTO_TYPE_PARTY) + " "

	plq.queryCount = "select p.id FROM " + models.DB_TABLE_PARTY + " as p "

	plq.queryWhere = " where 1=1 "

	return plq
}

func (plq *PartyListQuery) ValidState(n int) *PartyListQuery {
	if n < DEF_VALID_STATE_ALL {
		plq.queryWhere += " and p.valid_state = ? "
		plq.params = append(plq.params, n)
	}

	return plq
}

func (plq *PartyListQuery) PageNo(n int) *PartyListQuery {
	plq.page_no = n
	return plq
}

func (plq *PartyListQuery) PageSize(n int) *PartyListQuery {
	plq.page_size = n
	return plq
}

func (plq *PartyListQuery) Country(country string) *PartyListQuery {
	if country != "" {
		plq.queryWhere += " and p.country = ? "
		plq.params = append(plq.params, country)
	}

	return plq
}

func (plq *PartyListQuery) Province(province string) *PartyListQuery {
	if province != "" {
		plq.queryWhere += " and p.province = ? "
		plq.params = append(plq.params, province)
	}

	return plq
}

func (plq *PartyListQuery) City(city string) *PartyListQuery {
	if city != "" {
		plq.queryWhere += " and p.city = ? "
		plq.params = append(plq.params, city)
	}

	return plq
}

func (plq *PartyListQuery) Month(month int) *PartyListQuery {
	if month != 0 {
		plq.queryWhere += " and month(p.start_time) = ? "
		plq.params = append(plq.params, month)
	}

	return plq
}

func (plq *PartyListQuery) Tags(tags string) *PartyListQuery {
	if tags == "" || tags == "," {
		return plq
	}

	joinSql := " inner join " + models.DB_TALBE_PARTY_TAG_MAP + " as ptm on ptm.party_id=p.id " +
		" inner join " + models.DB_TALBE_TAG + " on tag.id=ptm.tag_id "

	plq.queryTables += joinSql
	plq.queryCount += joinSql

	tagArr := strings.Split(tags, ",")
	tagLen := len(tagArr)
	if tagLen > 0 {
		plq.queryWhere += " and tag.name in ( " + strings.TrimRight(strings.Repeat("?,", tagLen), ",") + " ) "

		for _, tag_name := range tagArr {
			plq.params = append(plq.params, tag_name)
		}
	}

	return plq
}

func (plq *PartyListQuery) IncludeClose(flag bool) *PartyListQuery {
	if !flag {
		plq.queryWhere += " and p.close_time is NULL "
	}

	return plq
}

func (plq *PartyListQuery) SetUserPartyType(userId string, partyType int) *PartyListQuery {
	switch partyType {
	// user_id创建的赛事
	case PARTY_TYPE_OWN:
		plq.joinUserOwn(userId)

	// user_id报名参加的赛事
	case PARTY_TYPE_SIGN:
		plq.joinUserSign(userId)

	// user_id关注的赛事
	case PARTY_TYPE_FOLLOW:
		plq.joinUserFollow(userId)
	}

	return plq
}

// user_id创建的赛事
func (plq *PartyListQuery) joinUserOwn(userId string) {
	plq.queryWhere += " and p.user_id = ? "
	plq.params = append(plq.params, userId)
}

// user_id报名参加的赛事
func (plq *PartyListQuery) joinUserSign(userId string) {

	joinSql := " inner join " + models.DB_TABLE_GAME + " as g on p.id=g.party_id " +
		" inner join `" + models.DB_TABLE_ORDER + "` as o on o.game_id=g.id and o.user_id=" + userId

	plq.queryTables += joinSql
	plq.queryCount += joinSql

}

// user_id关注的赛事
func (plq *PartyListQuery) joinUserFollow(userId string) {
	joinSql := " inner join " + models.DB_TABLE_USER_PARTY + " as u on u.party_id=p.id "

	plq.queryTables += joinSql
	plq.queryCount += joinSql

	plq.queryWhere += " and p.user_id = ? "
	plq.params = append(plq.params, userId)
}

func (plq *PartyListQuery) OrderBy(column map[string]bool) *PartyListQuery {
	orderLen := len(column)
	if orderLen < 1 {
		return plq
	}

	plq.queryOrder = " order by "

	sepStr := ""

	mOrder := map[bool]string{
		true:  " asc ",
		false: " desc ",
	}

	for orderName, orderType := range column {
		plq.queryOrder += sepStr + orderName + mOrder[orderType]

		if sepStr == "" {
			sepStr = " , "
		}
	}

	return plq
}

func (plq *PartyListQuery) Do() error {
	utils.Logger.Debug("partySrv.PartyListQuery.Do: start: sql - plq = %v", plq)

	utils.Logger.Debug("SQL----->%s", plq.queryTables+plq.queryWhere+" group by p.id "+plq.queryOrder+" LIMIT ? OFFSET ?", plq.params,
		plq.page_size, plq.page_no*plq.page_size)
	num, err := models.Orm.Raw(plq.queryTables+plq.queryWhere+" group by p.id "+plq.queryOrder+" LIMIT ? OFFSET ?", plq.params,
		plq.page_size, plq.page_no*plq.page_size).QueryRows(&plq.result)

	utils.Logger.Debug("partySrv.PartyListQuery.Do: num = %v, err = %v", num, err)

	if err == orm.ErrNoRows {
		plq.err = err
	}

	plq.count()

	plq.did_query = true

	return plq.err
}

func (plq *PartyListQuery) ErrorMsg() string {
	if plq.err != nil {
		return plq.err.Error()
	}
	return ""
}

func (plq *PartyListQuery) Result() []*models.PartyAndDetail {
	if !plq.did_query {
		plq.err = errors.New("partySrv.PartyListQuery.Result : call Do() before Result()")
	}
	return plq.result
}

func (plq *PartyListQuery) Total() int64 {
	return plq.total
}

func (plq *PartyListQuery) Reset() *PartyListQuery {
	plq = nil
	return NewPartyListQuery()
}

// 赛事列表个数统计
func (plq *PartyListQuery) count() {
	var maps []orm.Params
	num, err := models.Orm.Raw(plq.queryCount+plq.queryWhere, plq.params).Values(&maps)
	utils.Logger.Debug("partySrv.PartyListQuery.count: num = %v, err = %v", num, err)

	plq.total = num
}

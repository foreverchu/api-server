package models

import "github.com/astaxie/beego/orm"

const (
	DB_REG_META_TABLE_NAME       = "reg_meta"
	DB_PARTY_REG_META_TABLE_NAME = "party_reg_meta"
	DB_REG_META_DATA_TABLE_NAME  = "reg_meta_data"
)

//'数据类型，0：单行正文，1：多行正文，2：整数，3：浮点数，4：金额，5：日期，6：时间，7：日期时间， 8：复选框， 9：经纬度坐标, 10: 路径(是个经纬度坐标数组), 11: 图片（图片url), 12: 多个图片(图片（图片urlURL数组)',

const (
	REG_META_TYPE_SINGLELINE_TEXT = 0
	REG_META_TYPE_MULTILINE_TEXT  = 1
	REG_META_TYPE_INT             = 2
	REG_META_TYPE_FLOAT           = 3
	REG_META_TYPE_U_INT           = 4 //无符号整型
	REG_META_TYPE_U_FLOAT         = 5 //无符号浮点型
	REG_META_TYPE_CURRENCY        = 6
	REG_META_TYPE_DATE            = 7
	REG_META_TYPE_TIME            = 8
	REG_META_TYPE_DATETIME        = 9
	REG_META_TYPE_OPTION          = 10 //复选框
	REG_META_TYPE_RADIO           = 11 //单选框
	REG_META_TYPE_LAT_LONG        = 12 //经纬度
	REG_META_TYPE_PATH            = 13 //路径，经纬度数组
	REG_META_TYPE_IMAGE           = 14 //图像url
	REG_META_TYPE_IMAGES          = 15 //图像url数组
	REG_META_TYPE_RICH_TEXT       = 16 //富文本
	REG_META_TYPE_HTML            = 17
	REG_META_TYPE_MARKDOWN        = 18
)

type RegMeta struct {
	Id uint32

	Name   string `orm:"size(64)"`
	Type   uint8
	ExData string `orm:"type(text)"`
}

type PartyRegMeta struct {
	Id uint32

	Party  *Party   `orm:"rel(fk);rel(one)"`
	MetaId *RegMeta `orm:"rel(fk);rel(one)"`
	flags  uint32
}

type RegMetaData struct {
	Id uint32

	Party *Party        `orm:"rel(fk);rel(one)"`
	Reg   *Registration `orm:"rel(fk);rel(one);column(reg_id)"`
	Meta  *RegMeta      `orm:"rel(fk);rel(one)"`
	Value string        `orm:"type(text)"`
}

func GetRegMetaById(id uint32) (*RegMeta, error) {
	reg_meta := RegMeta{
		Id: id,
	}
	err := Orm.Read(&reg_meta)
	if err == orm.ErrNoRows {
		return nil, err
	}
	return &reg_meta, nil
}

func GetAllRegMetas() ([]RegMeta, error) {
	var metas []RegMeta

	queryselector := Orm.QueryTable(DB_REG_META_TABLE_NAME)
	_, err := queryselector.All(&metas)
	if err == orm.ErrNoRows {
		return metas, nil
	} else if err != nil {
		return metas, err
	}

	return metas, nil
}

func GetAllMetasOfParty(party_id string) ([]PartyRegMeta, error) {
	var party_metas []PartyRegMeta

	queryselector := Orm.QueryTable(DB_PARTY_REG_META_TABLE_NAME)
	_, err := queryselector.Filter(DB_GAME_FN_PARTY_ID, party_id).All(&party_metas)

	if err == orm.ErrNoRows {
		return party_metas, nil
	} else if err != nil {
		return party_metas, err
	}

	return party_metas, nil
}

func GetAllMetasDataOfPartyReg(party_id string, reg_id uint32) ([]RegMetaData, error) {
	var reg_meta_datas []RegMetaData

	queryselector := Orm.QueryTable(DB_REG_META_DATA_TABLE_NAME)
	_, err := queryselector.Filter(DB_GAME_FN_PARTY_ID, party_id).Filter(DB_REG_META_DATA_REG_ID, reg_id).All(&reg_meta_datas)

	if err != nil {
		return reg_meta_datas, err
	}

	return reg_meta_datas, nil
}

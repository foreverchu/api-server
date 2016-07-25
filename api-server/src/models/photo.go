package models

const (
	DB_PHOTO_TYPE   = "type"
	DB_PHOTO_REL_ID = "rel_id"
)

// Photo type
const (
	PHOTO_TYPE_PARTY       = 1 // 赛事图片
	PHOTO_TYPE_GAME        = 2 // 比赛图片
	PHOTO_TYPE_PARTY_ROUTE = 3 // 赛事路线图片
)

type Photo struct {
	Id    uint32
	RelId uint32
	Url   string `orm:"size(128)"`
	Type  uint8
}

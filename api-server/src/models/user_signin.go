package models

import "time"

type UserSignIn struct {
	Id        uint
	UserId    uint32
	Source    uint8     //登入来源, 比如web, android, ios, h5 etc.
	CreatedAt time.Time `orm:"type(datetime)"`
	UpdatedAt time.Time `orm:"type(datetime)"`
}

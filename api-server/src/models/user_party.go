package models

import (
	"errors"
	"time"
)

type UserParty struct {
	Id        uint32
	UserId    uint32
	PartyId   uint32
	CreatedAt time.Time `orm:"type(datetime)"`
}

var (
	ErrFollow   = errors.New("models.UserParty.Follow : db insert failed")
	ErrUnfollow = errors.New("models.UserParty.Unfollow : db delete failed")
)

func (up *UserParty) Follow() error {
	if _, err := Orm.Insert(up); err != nil {
		return ErrFollow
	}
	return nil
}

func (up *UserParty) Unfollow() error {
	if _, err := Orm.Delete(up); err != nil {
		return ErrUnfollow
	}
	return nil
}

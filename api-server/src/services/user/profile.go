package userSrv

import (
	"errors"
	"time"

	"models"
)

var ErrProfile = errors.New("用户档案没有建立")

func (u *User) setProfile(params map[string]interface{}) {
	u.profile.UpdatedAt = time.Now().Local()
	for k, v := range params {
		switch k {
		case "address":
			u.profile.Address = v.(string)
		case "gender":
			u.profile.Gender = v.(uint8)
		case "constellation":
			u.profile.Constellation = v.(uint8)
		case "profession":
			u.profile.Profession = v.(string)
		case "about":
			u.profile.About = v.(string)
		default:
			continue
		}
	}
}
func (u *User) NewProfile(params map[string]interface{}) {
	u.profile = new(models.Profile)
	u.profile.UserId = u.user.Id
	u.profile.CreatedAt = time.Now().Local()
	for k, v := range params {
		switch k {
		case "address":
			u.profile.Address = v.(string)
		case "gender":
			u.profile.Gender = v.(uint8)
		case "constellation":
			u.profile.Constellation = v.(uint8)
		case "profession":
			u.profile.Profession = v.(string)
		case "about":
			u.profile.About = v.(string)
		default:
			continue
		}
	}
}

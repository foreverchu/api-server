package userSrv

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"

	"models"
)

var (
	ErrUserInfo  = errors.New("无法获取用户资料")
	ErrNotSignIn = errors.New("您还未登录")
)

type User struct {
	userId  string
	user    *models.User
	profile *models.Profile
}

func NewUser(userId string) (u *User, err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("userSrv.User.NewUser : %s", err.Error())
		} else {
			utils.Logger.Debug("userSrv.User.NewUser : success uid= %s", u.userId)
		}
	}()
	//id, _ := strconv.Atoi(u.userId)

	u = &User{
		userId:  userId,
		user:    new(models.User),
		profile: new(models.Profile),
	}

	userQueryCond := map[string]interface{}{
		"id": u.userId,
	}
	if err := u.user.FindBy(userQueryCond); err == models.ErrUserNotFound {
		return nil, ErrUserInfo
	}

	profileQueryCond := map[string]interface{}{
		"user_id": u.userId,
	}
	if err = u.profile.FindBy(profileQueryCond); err != nil && err != models.ErrUserProfileNotFound {
		return nil, ErrUserInfo
	}

	return u, nil
}

func (u *User) setUser(params map[string]interface{}) {
	for k, v := range params {
		switch k {
		case "name":
			u.user.Name = v.(string)
		case "avatar":
			u.user.Avatar = v.(string)
		default:
			continue
		}
	}

}

func (u *User) Update(params map[string]interface{}) error {
	Orm := orm.NewOrm()
	err := Orm.Begin()
	if err != nil {
		return err
	}
	count, err := Orm.QueryTable("profile").Filter("user_id", u.user.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		u.NewProfile(params)
		_, err = Orm.Insert(u.profile)
		if err != nil {
			return err
		}
	} else {
		u.setProfile(params)
		_, err = Orm.Update(u.profile)
		if err != nil {
			return err
		}
	}
	u.setUser(params)

	_, err = Orm.Update(u.user)
	if err != nil {
		err = Orm.Rollback()
		if err != nil {
			return err
		} else {
			return errors.New("RollBack")
		}
	}
	return Orm.Commit()
}

func (u *User) User() *models.User {
	return u.user
}

func (u *User) Profile() *models.Profile {
	return u.profile
}

func FindNameById(id uint32) (name string) {
	u, err := NewUser(strconv.FormatUint(uint64(id), 32))
	user := u.User()
	if err != nil {
		return ""
	}
	return user.Name
}

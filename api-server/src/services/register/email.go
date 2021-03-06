package registerSrv

import (
	"models"

	"github.com/astaxie/beego/validation"
)

type Email struct {
	email string
	valid *validation.Validation
}

func NewEmail(email string) *Email {
	return &Email{
		email: email,
		valid: &validation.Validation{},
	}
}

func (e *Email) EmailAddress() string {
	return e.email
}

func (r *Email) IsExists() bool {
	return models.IsValueExists(models.DB_TABLE_USER, models.DB_USER_EMAIL, r.email)
}

func (e *Email) IsValid() bool {
	if v := e.valid.Email(e.email, "email"); v.Ok {
		return true
	}
	return false
}

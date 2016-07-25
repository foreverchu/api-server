package registerSrv

import (
	"errors"
	"fmt"
	"math"
	"models"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/chinarun/utils"
)

var ErrSMSDurationTooShort = errors.New("请求短信验证码过于频繁")
var ErrSendSMS = errors.New("发送短信验证码失败")
var ErrSMSCodeNotExists = errors.New("无效的短信验证码")
var ErrSMSCodeUsed = errors.New("短信验证码被使用过")

const (
	SMS_REGISTER_TEMPLATE_ID = 313
	SMS_REQ_DURATION         = 1 * time.Minute //验证码请求间隔
	DEFAULT_SMS_CODE_LENGTH  = 6
)

type SMSer interface {
	Send(templateId int, phone string, vars map[string]string) error
}

type Phone struct {
	phone   string
	valid   *validation.Validation
	code    string
	smscode *models.Smscode
}

func NewPhone(phone string) *Phone {
	return &Phone{
		phone:   phone,
		smscode: &models.Smscode{},
		valid:   &validation.Validation{},
	}
}

func (p *Phone) PhoneNumber() string {
	return p.phone
}

func (p *Phone) IsValid() bool {
	if v := p.valid.Mobile(p.phone, "phone"); v.Ok {
		return true
	}
	return false
}

func (p *Phone) IsExists() bool {
	return models.IsValueExists(models.DB_TABLE_USER, models.DB_USER_PHONE, p.phone)
}

func (p *Phone) ValidSMSCode(code string) error {
	cond := map[string]interface{}{
		models.DB_SMSCODE_PHONE: p.phone,
		models.DB_SMSCODE_CODE:  code,
	}
	if err := p.smscode.FindBy(cond); err == models.ErrSMSCodeNotFound {
		return ErrSMSCodeNotExists
	}

	if p.smscode.IsUsed() {
		return ErrSMSCodeUsed
	}
	return nil
}

func genCode(length int) string {
	n := math.Pow(10, float64(length))
	return fmt.Sprintf("%d", utils.Rander.Intn(int(n)))
}

func (p *Phone) GenCode(length int) string {
	p.code = genCode(length)
	return p.code
}

// 防止单一手机号码无限次数被请求
func (p *Phone) IsRequestedCode() bool {
	p.smscode.Phone = p.phone
	return p.smscode.IsRequestedCode()
}

func (p *Phone) IsOkToRqeustCode() bool {
	if !p.IsRequestedCode() {
		return true
	}
	if time.Since(p.smscode.CreatedAt) > SMS_REQ_DURATION {
		return true
	}
	return false
}

// Send 发送短信验证码
func (p *Phone) SendSMS(sms SMSer) (err error) {
	var actualErr error
	defer func() {
		if err != nil {
			utils.Logger.Error("registerSrv.Phone.Send : %s", actualErr.Error())
		} else {
			utils.Logger.Debug("registerSrv.Phone.Send : %v", p)
		}
	}()

	if !p.IsValid() {
		err = ErrInvalidPhone
		actualErr = err
		return
	}

	if p.IsExists() {
		err = ErrPhoneExists
		actualErr = err
		return
	}

	if !p.IsOkToRqeustCode() {
		err = ErrSMSDurationTooShort
		actualErr = err
		return
	}

	code := p.GenCode(DEFAULT_SMS_CODE_LENGTH)
	data := map[string]string{
		"%Code%": code,
	}
	err = sms.Send(SMS_REGISTER_TEMPLATE_ID, p.phone, data)
	if err != nil {
		actualErr = err
		err = ErrSMSCode
		return
	}

	err = p.record()
	if err != nil {
		actualErr = err
		return ErrSMSCode
	}

	return nil
}

func (p *Phone) record() error {
	p.smscode.Code = p.code
	p.smscode.CreatedAt = time.Now()
	p.smscode.Insert()
	return nil
}

func (p *Phone) UseCode() error {
	p.smscode.UsedAt = time.Now()
	return p.smscode.Update()
}

package noticeSrv

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/chinarun/utils"
)

var ErrSmsTemplateArguments = errors.New("发送短信参数错误")

type Sms map[string]string

func NewSms() Sms {
	return Sms{
		"smsUser": smsUser,
	}
}

func (s Sms) setVars(vars string) Sms {
	s["vars"] = vars
	return s
}

func (s Sms) setPhone(phone string) Sms {
	s["phone"] = phone
	return s
}

func (s Sms) setTemplateId(templateId int) Sms {
	s["templateId"] = strconv.Itoa(templateId)
	return s
}

func (s Sms) generateSign() Sms {
	keys := make([]string, 0, len(s))
	for key := range s {
		if key == "signature" {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var urlQueryParams string
	urlQueryParams += smsKey

	for _, key := range keys {
		value := s[key]
		if value == "" {
			continue
		}
		urlQueryParams += fmt.Sprintf("&%s=%s", key, value)
	}
	urlQueryParams += "&" + smsKey

	signature := utils.GetMd5(urlQueryParams)

	s["signature"] = signature //"f627f253577e777d448001895ea62bcf" //signature
	return s
}

func (s Sms) encode() string {
	if s == nil {
		return ""
	}
	var urlQueryParams bytes.Buffer
	keys := make([]string, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := s[key]
		prefix := url.QueryEscape(key) + "="
		if urlQueryParams.Len() > 0 {
			urlQueryParams.WriteByte('&')
		}
		urlQueryParams.WriteString(prefix)
		urlQueryParams.WriteString(url.QueryEscape(value))
	}
	return urlQueryParams.String()
}

/*


Exa:
     xx := make(map[string]string)
     xx["%user%"] = "xuhang"
     xx[`%party%`] = "北马"
     xx[`%score%`] = "100"
     err = noticeSrv.NewSms().SendSms(315, "15250416847", xx)
     if err != nil {
         fmt.Printf("err----->%v\n", err)
     }
*/
func (s Sms) Send(templateId int, phone string, vars map[string]string) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("noticeSrv.Sms.Send : error %v", err)
		} else {
			utils.Logger.Debug("noticeSrv.Sms.Send : send to phone: %s Success.", phone)
		}
	}()

	utils.Logger.Debug("noticeSrv.Sms.Send : vars : %v", vars)

	s = s.setPhone(phone).setTemplateId(templateId)
	if vars != nil {
		varsBytes, err := json.Marshal(vars)
		if err != nil {
			return ErrSmsTemplateArguments
		}
		s = s.setVars(string(varsBytes))
	}
	req := s.generateSign().encode()
	utils.Logger.Debug("noticeSrv.Sms.Send : req: %s", req)

	postBody := strings.NewReader(req)

	utils.Logger.Debug("noticeSrv.Sms.Send : postBody: %s", postBody)
	utils.Logger.Debug("noticeSrv.Sms.Send : url: %s", sendSmsUrl)

	js, err := httpSend(sendSmsUrl, postBody)
	if err != nil {
		return
	}
	message, err := js.Get("message").String()
	if err != nil {
		return
	}
	if message != "请求成功" {
		return fmt.Errorf("短信发送失败")
	}

	return
}

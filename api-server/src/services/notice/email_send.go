package noticeSrv

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/chinarun/utils"
)

type Email map[string]string

func NewEmail() Email {
	return Email{
		"api_user": apiUser,
		"api_key":  apiKey,
		"from":     from,
		"fromname": from,
	}
}

//html可以为邮件内容，或者邮件模板
func (e Email) setHtml(html string) Email {
	e["html"] = html
	return e
}

func (e Email) setSubject(subject string) Email {
	e["subject"] = subject
	return e
}

func (e Email) setTo(to string) Email {
	e["to"] = to
	if strings.Contains(to, ";") {
		e["x_smtpapi"] = fmt.Sprintf(`{"to": %s}`, to)
	}
	return e
}

func (e Email) setTemplate(template_name string) Email {
	e["template_invoke_name"] = template_name
	return e
}

//设置替换变量，包括收件人地址。substitution_vars示例：{"to": ["ben@ifaxin.com", "joe@ifaxin.com"],"sub":{"%name%": ["Ben", "Joe"],"%money%":[288, 497]}}。此时不需要设置to
func (e Email) setSubstitutionVars(substitution_vars string) Email {
	e["substitution_vars"] = substitution_vars
	return e
}

//当use_maillist 为true时，需要设置to，不能使用变量替换
//当use_maillist 为false时，需要设置substitution_vars
//默认为false
func (e Email) setUseMaillist(use_maillist string) Email {
	e["use_maillist"] = use_maillist
	return e
}

func (e Email) encode() string {
	if e == nil {
		return ""
	}
	var urlQueryParams bytes.Buffer
	keys := make([]string, 0, len(e))
	for key := range e {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if urlQueryParams.Len() > 0 {
			urlQueryParams.WriteByte('&')
		}
		value := e[key]
		prefix := url.QueryEscape(key) + "="
		urlQueryParams.WriteString(prefix)
		urlQueryParams.WriteString(url.QueryEscape(value))
	}
	return urlQueryParams.String()
}

func (e Email) Send(html, subject, to string) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("noticeSrv.Email.Send : error: %v", err)
		} else {
			utils.Logger.Debug("noticeSrv.Email.Send : SendEmail success")
		}
	}()

	req := e.setHtml(html).setSubject(subject).setTo(to).encode()

	postBody := bytes.NewBufferString(req)

	js, err := httpSend(sendEmailUrl, postBody)
	if err != nil {
		return
	}
	message, err := js.Get("message").String()
	if err != nil {
		return
	}
	if message != "success" {
		return fmt.Errorf("邮件发送失败")
	}
	return
}

type EmailData struct {
	To   string
	Data map[string]interface{}
}

//设置替换变量，包括收件人地址。substitution_vars示例：{"to": ["ben@ifaxin.com", "joe@ifaxin.com"],"sub":{"%name%": ["Ben", "Joe"],"%money%":[288, 497]}}。此时不需要设置to
func formatEmailData(datas []*EmailData) (string, error) {
	tos := []string{}
	subs := map[string][]interface{}{}
	for _, v := range datas {
		tos = append(tos, v.To)
		for key, data := range v.Data {
			subs[key] = append(subs[key], data)
		}
	}
	combine := map[string]interface{}{
		"to":  tos,
		"sub": subs,
	}
	combineBytes, err := json.Marshal(combine)
	if err != nil {
		return "", err
	}
	return string(combineBytes), nil
}

func (e Email) SendByTemplate(template_name string, vars []*EmailData) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("SendEmailByTemplate error: %v", err)
		} else {
			utils.Logger.Debug("SendEmailByTemplate  Success.")
		}
	}()

	if len(vars) < 1 {
		err = errors.New("No Vars")
		return
	}
	if template_name == "" {
		err = errors.New("No TamplateName")
		return
	}
	formattedVars, err := formatEmailData(vars)
	if err != nil {
		return
	}

	utils.Logger.Debug("noticeSrv.Email.SendByTemplate : formatEmailData : %s", formattedVars)

	req := e.setTemplate(template_name).setSubstitutionVars(formattedVars).encode()
	utils.Logger.Debug("noticeSrv.Email.SendByTemplate : req : %s", req)

	postBody := strings.NewReader(req)
	js, err := httpSend(sendEmailByTempUrl, postBody)
	if err != nil {
		return
	}
	message, err := js.Get("message").String()
	if err != nil {
		return
	}

	if message != "success" {
		return fmt.Errorf("邮件发送失败")
	}
	return

}

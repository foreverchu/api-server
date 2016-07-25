package noticeSrv

import "github.com/chinarun/utils"

var (
	apiUser            string
	apiKey             string
	from               string
	sendEmailUrl       string
	sendEmailByTempUrl string

	sendSmsUrl    string
	smsUser       string
	smsKey        string
	smsExpireTime int

	Key string
)

func NoticeParamsInit() (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("noticeSrv.NoticeParamsInit : error : %s", err.Error())
		}
	}()

	Key, err = utils.Cfg.GetString("signature", "key")
	if err != nil {
		return err
	}
	//email params
	apiUser, err = utils.Cfg.GetString("notice", "api_user")
	if err != nil {
		return err
	}
	apiKey, err = utils.Cfg.GetString("notice", "api_key")
	if err != nil {
		return err
	}
	from, err = utils.Cfg.GetString("notice", "from")
	if err != nil {
		return err
	}
	sendEmailUrl, err = utils.Cfg.GetString("notice", "email_url")
	if err != nil {
		return err
	}

	sendEmailByTempUrl, err = utils.Cfg.GetString("notice", "email_tmp_url")
	if err != nil {
		return err
	}

	//sms params
	sendSmsUrl, err = utils.Cfg.GetString("notice", "sms_url")
	if err != nil {
		return err
	}

	smsUser, err = utils.Cfg.GetString("notice", "smsUser")
	if err != nil {
		return err
	}

	smsKey, err = utils.Cfg.GetString("notice", "smsKey")
	if err != nil {
		return err
	}
	return err
}

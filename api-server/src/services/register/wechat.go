package registerSrv

import (
	"encoding/json"
	"errors"
	"fmt"
	"models"
	"net/url"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/chinarun/utils"
)

var ErrGetCode = errors.New("获取code出错")
var ErrGetAccessToken = errors.New("获取access_token出错")
var ErrRefreshToken = errors.New("获取access_token出错")
var ErrUserinfo = errors.New("获取用户信息出错")
var ErrThridPartyInfo = errors.New("无法保存第三方用户信息")

const (
	ThirdyPartyRegisterTypeWechat = 1 //wechat
)

var (
	appid         string
	secret        string
	redirectUrl   string
	wechatAPIHost string
)

type WechatAPI struct {
	host       string
	code       string
	appid      string
	secret     string
	state      string
	token      *WechatToken
	wechatUser *WechatUser
	reg        *models.ThirdPartyRegister
	user       *models.User
}

type WechatToken struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
	Openid       string        `json:"openid"`
	Scope        string        `json:"scope"`
	Unionid      string        `json:"unionid"`
}

type WechatUser struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        uint     `json:"sex"`
	Country    string   `json:"country"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Headimgurl string   `json:"headimgurl"` //下载到本地
	Privilege  []string `json:"privilege"`
}

// how to remove the config dependence
func initConfig() {
	appid, _ = utils.Cfg.GetString("wechat", "appid")
	secret, _ = utils.Cfg.GetString("wechat", "secret")
	baseUrl, _ := utils.Cfg.GetString("beego", "base_url")
	redirectUrl = url.QueryEscape(baseUrl + "/wechat")
	wechatAPIHost = "https://api.weixin.qq.com"
}

func NewWechatAPI() *WechatAPI {
	initConfig()
	return &WechatAPI{
		host:   wechatAPIHost,
		appid:  appid,
		secret: secret,
		reg:    &models.ThirdPartyRegister{},
		user:   &models.User{},
		state:  generateRandomString(16),
	}
}

func (r *WechatAPI) SetCode(code string) {
	r.code = code
}

func (r *WechatAPI) GetCode() error {
	url := fmt.Sprintf("https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=%s", r.appid, redirectUrl, r.state)
	_, err := httplib.Get(url).Bytes()
	if err != nil {
		return ErrGetCode
	}
	return nil
}

func (r *WechatAPI) Signin() error {
	if err := r.GetAccessToken(); err != nil {
		return err
	}
	if err := r.GetWechatUserInfo(); err != nil {
		return err
	}
	if err := r.Update(); err != nil {
		return err
	}
	return nil
}

func (r *WechatAPI) UserInfo() (*models.User, error) {
	utils.Logger.Debug("%s", utils.Sdump(r))
	if r.user == nil {
		return nil, ErrUserinfo
	}
	return r.user, nil
}

func (r *WechatAPI) GetAccessToken() error {
	url := fmt.Sprintf("%s/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", r.host, r.appid, r.secret, r.code)
	utils.Logger.Debug("%s", url)
	resp, err := httplib.Get(url).Bytes()
	utils.Logger.Debug("%s", resp)
	if err != nil {
		utils.Logger.Debug("%s", err.Error())
		return ErrGetAccessToken
	}

	//marshal response
	tokenInfo := &WechatToken{}
	err = json.Unmarshal(resp, tokenInfo)
	if err != nil {
		return err
	}
	utils.Logger.Debug("%v", tokenInfo)
	r.token = tokenInfo
	return nil
}

// 在用户获取用户信息成功后, 马上RefreshToken, access_token不变, 但expiresIn变长, 为30天
func (r *WechatAPI) RefreshToken() error {
	url := fmt.Sprintf("%s/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", r.host, r.appid, r.token.RefreshToken)
	_, err := httplib.Get(url).Bytes()
	if err != nil {
		return ErrRefreshToken
	}
	return nil
}

func (r *WechatAPI) ValidAccessToken() error {
	return nil
}

func (r *WechatAPI) GetWechatUserInfo() error {
	url := fmt.Sprintf("%s/sns/userinfo?access_token=%s&openid=%s", r.host, r.token.AccessToken, r.token.Openid)
	resp, err := httplib.Get(url).Bytes()
	utils.Logger.Debug("%s", resp)
	if err != nil {
		utils.Logger.Debug("%s", err.Error())
		return ErrUserinfo
	}

	//marshal response
	userInfo := &WechatUser{}
	err = json.Unmarshal(resp, userInfo)
	if err != nil {
		return err
	}
	r.wechatUser = userInfo
	return nil
}

func (r *WechatAPI) Update() error {
	if !r.IsRegistered() {
		if err := r.CreateUser(); err != nil {
			return err
		}
		utils.Logger.Debug("current user :%v", r.user)
		if err := r.createThirdPartyRegister(); err != nil {
			return err
		}
		utils.Logger.Debug("current third party info :%v", r.reg)
	} else {
		if err := r.updateThridPartyRegister(); err != nil {
			return err
		}
	}
	return nil
}

func (r *WechatAPI) IsRegistered() bool {
	cond := map[string]interface{}{
		models.DB_THRID_PARTY_REGISTER_FROM_ID: r.token.Openid,
	}
	if err := r.reg.FindBy(cond); err == models.ErrUserNotFound {
		utils.Logger.Debug("is register :  %s", err.Error())
		return false
	}
	utils.Logger.Debug("current third party info :%v", r.reg)
	r.user.Id = r.reg.UserId

	return true
}

func (r *WechatAPI) setThirdPartyRegister() {
	r.reg.UserId = r.user.Id
	r.reg.Type = ThirdyPartyRegisterTypeWechat
	r.reg.FromId = r.token.Openid
	r.reg.AccessToken = r.token.AccessToken
	r.reg.RefreshToken = r.token.RefreshToken
	r.reg.ExpiredAt = time.Now().Add(r.token.ExpiresIn * time.Second)
}

func (r *WechatAPI) createThirdPartyRegister() error {
	r.setThirdPartyRegister()
	r.reg.CreatedAt = time.Now()
	_, err := models.Orm.Insert(r.reg)
	if err != nil {
		return ErrThridPartyInfo
	}
	return nil
}

func (r *WechatAPI) updateThridPartyRegister() error {
	r.setThirdPartyRegister()
	r.reg.UpdatedAt = time.Now()
	_, err := models.Orm.Update(r.reg)
	if err != nil {
		return ErrThridPartyInfo
	}
	return nil
}

func (r *WechatAPI) CreateUser() error {
	r.user.Name = r.wechatUser.Nickname
	r.user.Avatar = r.wechatUser.Headimgurl
	r.user.ComeFrom = ComeFromWechatInt

	id, err := models.Orm.Insert(r.user)
	if err != nil {
		return ErrorCreateUserFailed
	}
	r.user.Id = uint32(id)
	return nil
}

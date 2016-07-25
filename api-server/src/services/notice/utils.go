package noticeSrv

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/chinarun/utils"
)

func httpSend(url string, body io.Reader) (js *simplejson.Json, err error) {
	responseHandler, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		utils.Logger.Error("httpSend error : %s", err.Error())
		return nil, err

	}
	defer responseHandler.Body.Close()
	bodyByte, err := ioutil.ReadAll(responseHandler.Body)
	if err != nil {
		return nil, err

	}
	js, err = simplejson.NewJson(bodyByte)
	return js, err

}

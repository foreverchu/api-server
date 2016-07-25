package qrcodeSrv

import (
	"errors"

	"github.com/chinarun/utils"
	qrcode "github.com/skip2/go-qrcode"
)

var ErrContent = errors.New("不正确的内容")

type Qrcode struct {
	content string
	size    int
}

func NewQrcode(content string, size int) *Qrcode {
	return &Qrcode{
		content: content,
		size:    size,
	}
}

func (qr *Qrcode) Generate() (png []byte, err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("qrcodeSrv.Qrcode.Generate : %s", err.Error())
		} else {
			utils.Logger.Debug("qrcodeSrv.Qrcode.Generate : %v", qr)
		}
	}()

	if len(qr.content) <= 0 {
		err = ErrContent
		return
	}
	png, err = qrcode.Encode(qr.content, qrcode.Medium, qr.size)
	return
}

package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

// 从指定位置截取string，当start ＋ length 超出末尾后，返回长度根据实际长度返回，如Substr（“chinarun”， 7， 3），返回为n
// start为负值时反向（从最末向前）截取.比如Substr（“chinarun”， -2， 3），返回为nar
func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)

	if start >= rl || length < 1 {
		return ""
	}

	end := 0

	if start < 0 {
		end = rl + start
		if end < 0 {
			return ""
		}

		start = end - length
		if start < 0 {
			start = 0
		}
	} else {
		end = start + length
		if end > rl {
			end = rl
		}
	}

	return string(rs[start:end])
}

func _sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func _md5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对string进行md5 sha1 编码
func GetMd5Sha1(s string) string {
	var md5_sha1 string
	md5_sha1 = _sha1(_md5(s))
	return md5_sha1

}

func GetMd5(s string) string {

	t := md5.New()
	io.WriteString(t, s)
	return fmt.Sprintf("%x", t.Sum(nil))
}

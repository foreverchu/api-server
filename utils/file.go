package utils

import (
	"fmt"
	"time"
    "unicode"
    "strings"
)

func GetRandomName() string {
	t := time.Now()
	filename := fmt.Sprintf("%04d%02d%02d%02d%02d%02d%03d%06d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), int(t.Nanosecond()/1000000), Rander.Intn(1000000))
	return filename
}

// 驼峰命名转成下划线或短横等命名
// src - 要转换的驼峰字符串
// sep - 分隔符，传空为下划线，可指定其他字符
// 返回转换后的字符串
func SepSplitName(src, sep string) (dst string) {
    if sep == "" {
        sep = "_"
    }

    for _, s := range src {
        if unicode.IsUpper(s) {
            dst += sep + strings.ToLower(string(s))
        } else {
            dst += strings.ToLower(string(s))
        }
    }

    if string(dst[0]) == sep {
        return dst[1:]
    }

    return
}

// 下划线或短横等命名转成驼峰命名
// src - 要转换的下划线字符串
// sep - 分隔符，传空为下划线，可指定其他字符
// 返回转换后的驼峰字符串，首字母大写
func CamelName(src, sep string) (dst string) {
    if sep == "" {
        sep = "_"
    }

    strs := strings.Split(src, sep)
    for _, s := range strs {
        dst += strings.ToUpper(string(s[0])) + SubStr(s, 1, len(s) - 1)
    }

    return
}

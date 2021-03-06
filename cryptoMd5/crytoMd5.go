package cryptomd5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Auto(pw string,st string) bool {
	h := md5.New()
	io.WriteString(h, pw)
	//指定1个 salt： salt1 = @#$%
	alt1 := "@#$%"
	io.WriteString(h, alt1)
	pwdmd5 := fmt.Sprintf("%x", h.Sum(nil))
	//fmt.Println(pwdmd5)
	if pwdmd5 == st {
		return true
	} else {
		return false
	}
}

func WriteAuto2(pw string) (p string) {
	h := md5.New()
	io.WriteString(h, pw)
	//指定1个 salt： salt1 = @#$%
	alt1 := "@#$%"
	io.WriteString(h, alt1)
	p = fmt.Sprintf("%x", h.Sum(nil))
	return
}

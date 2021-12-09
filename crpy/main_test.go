package main

import (
	"testing"
)

func TestCheckPasswd(t *testing.T) {
	b := checkPasswd("xfsweb")
	if b {
		t.Log("密码正确")
	} else {
		t.Log("密码错误")
	}

}

package cryptomd5

import "testing"

func TestAuto(t *testing.T) {
	pwd := "1234567"
	ok := Auto(pwd)
	if ok {
		t.Log("测试成功")
	} else {
		t.Fatal("测试失败")
	}
}

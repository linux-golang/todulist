package cryptomd5

import (
	"fmt"
	"testing"
)

// func TestAuto(t *testing.T) {
// 	pwd := "1234567"
// 	ok := Auto(pwd)
// 	if ok {
// 		t.Log("测试成功")
// 	} else {
// 		t.Fatal("测试失败")
// 	}
// }

func TestWriteAuto2(t *testing.T) {
	pwd := "1234567"
	p := WriteAuto2(pwd)
	fmt.Println(p)

}

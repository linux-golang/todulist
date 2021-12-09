package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

var (
	pwd string = "NQsaGX$$f1c6c104cc2eff9b244bff33ea4bb216"
)

func GetSalt() string {
	// 定义一个集合
	salt := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNMVBN"
	//使用随机的方式取六个集合的字母
	rand.Seed(time.Now().UnixNano())
	var salt1 []byte
	//rand.Intn()
	for i := 0; i < 6; i++ {
		slice := []byte(salt)[rand.Intn(len([]byte(salt)))]
		salt1 = append(salt1, slice)
	}

	return string(salt1)

}

// 生成md5 盐+传入的参数

func createMd5(salt, passwd string) string {
	//salt_passwd := fmt.Sprintln(salt + passwd)
	h := md5.New()
	io.WriteString(h, salt)
	io.WriteString(h, passwd)
	return fmt.Sprintf("%s$$%x", salt, h.Sum(nil))
}

// checkPasswd

func checkPasswd(passwd string) bool {
	i := strings.Index(pwd, "$")
	slat := []byte(pwd)[:i]
	md5pass := createMd5(string(slat), passwd)
	is := strings.Index(md5pass, "$")
	if is == -1 {
		fmt.Println("未找到")
		return false
	}

	//slice := []byte(md5pass)[i+2:]
	if md5pass != pwd {
		return false
	}
	return true
}

func main() {

	// 1. 创建参数
	crypt := flag.String("crypt", "", "encrypt string")
	help := flag.Bool("h", false, "help")
	passwd := flag.String("passwd", "", "proving passwd")

	flag.Usage = func() {
		fmt.Println("passwd encrypt tools")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help {
		flag.Usage()
	}

	if *passwd != "" {
		fmt.Println(createMd5(GetSalt(), *passwd))
	} else if *crypt != "" {
		b := checkPasswd(*crypt)
		if b {
			fmt.Println("密码正确")
		} else {
			fmt.Println("密码错误")
		}
	}

}

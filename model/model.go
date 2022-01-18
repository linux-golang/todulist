package model

import (
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"strconv"
	cryptomd5 "todoList/cryptoMd5"
	"todoList/fileModel"
	"todoList/vars"
)


// 自动获取任务名ID
func GetId() func() int {

	// 第一步判断文件中的行哈，如果有则返回
	id := 0
	tasks, b := fileModel.ReadFileTodolist()
	fmt.Println(tasks)
	if  (len(tasks) != 0 ) && (tasks != nil) {
		if b {
			id1, err := strconv.Atoi(tasks[len(tasks)-1].Id)
			if err != nil {
				id = 1
			} else {
				id = id1
			}

		}
	}


	return func() int {
		id++
		return id
	}
}


// 读取密码

func readpwd() string {
	b := make([]byte, 1024)
	f, err := os.OpenFile(vars.Pwdfile, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)

	}

	defer f.Close()
	n, _ := f.Read(b)
	str := b[:n]

	return string(str)

}




// 验证三次密码

func CheckPasswd() {

	for conut := 1; conut <= 3; conut++ {
		fmt.Println("请输入您的密码:")
		st, _ := gopass.GetPasswdMasked()
		ok := cryptomd5.Auto(string(st), readpwd())
		if ok {
			vars.Sku = true
			break
		} else if conut == 3 {
			fmt.Println("输入密码3次全部失败退出")
			os.Exit(1)
		}
	}

}



package fileModel

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"io"
	"os"
	"strings"
	cryptomd5 "todoList/cryptoMd5"
	"todoList/vars"
)
//readtodolist

func ReadFileTodolist() (tasks []*vars.Task, b bool) {
	// 判断文件是否存在
	ok := Exists(vars.Todolistfile)
	if  ok {
		// 读取文件
		tasks,err := Readfile(vars.Todolistfile)
		if err !=nil {
			return nil ,false
		}

		return tasks,true

	}
	return tasks,true

}

// 判断文件或者目录是否存在
func Exists(path string) bool {
	_,err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return true
	}
	return  true
}

// 读取tolist文件
func Readfile(path string) (tasks []*vars.Task,err error) {
	file, err := os.OpenFile(vars.Todolistfile, os.O_RDONLY, 0777)
	// 关闭文件
	defer file.Close()
	if err != nil {
		return nil,err
	}
	// 创建bufio
	read := bufio.NewScanner(file)
	for read.Scan(){
		stringline := read.Text()

		stringslice:= strings.Split(stringline,",")
		if len(stringslice) != 7 {
			return nil ,err
		}


		v1 := &vars.Task{
			Id:          stringslice[0],
			Name:        stringslice[1],
			Detailed:    stringslice[2],
			Create_time: stringslice[3],
			Finish_time: stringslice[4],
			Status:      stringslice[5],
			Founder:     stringslice[6],
		}

		// append 增加切片
		tasks = append(tasks,v1)
	}
	return tasks,err
}



// 初始化密码
func Initpasswd(f *os.File) error {
	fmt.Println("请输入初始化密码")
	st, _ := gopass.GetPasswdMasked()
	p := cryptomd5.WriteAuto2(string(st))
	_, err := f.Write([]byte(p))
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

// 检查密码文件是否存在
func ChechFileIsExit(path string) (error) {

	ok := Exists(path)

	if !ok {
		file,err := os.Create(path)
		defer func() {
			file.Close()
		}()
		if err !=nil {
			return err
		}

		err = Initpasswd(file)
		if err != nil {
			return err
		}
	}

	// 如果文件存在则读取文件中的加密后的数据
	file,err := os.Open(path)

	if err != nil {
		return  err
	}
	ctxpwd := make([]byte,1024)

	for {
		_,err := file.Read(ctxpwd)

		if err !=nil {
			if err == io.EOF {
				break
			}
			return  err
		}
	}
	//  判断读取的内容是否有字节
	if len(ctxpwd) < 10 {
		return  fmt.Errorf("密码文件少于10字节请删除:%s 在重启程序",path)
	}

	return nil
}



func Writetodolis(todo []*vars.Task) bool {

	f, err := os.OpenFile(vars.Todolistfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()

	buff := bufio.NewWriter(f)

	for _, v := range todo {
		s := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\r\n", v.Id, v.Name, v.Detailed, v.Create_time, v.Finish_time, v.Status, v.Founder)
		buff.Write([]byte(s))
	}

	buff.Flush()

	return true
}

package fileModel

import (
	"bufio"
	"encoding/json"
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
		tasks,err := Readfile()
		//fmt.Println("----",len(tasks))
		//fmt.Println(tasks)
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

// 读取tolist文件 // 使用json.NewDecoder 方式
func Readfile_1() ([]*vars.Task,error) {
	//file, err := os.OpenFile(vars.Todolistfile, os.O_RDONLY, 0777)
	file,err := os.Open(vars.Todolistfile)

	// 关闭文件
	defer file.Close()
	if err != nil {
		return nil,err
	}
	vars.Tasks = make([]*vars.Task,0)
	//var tt *vars.Task
	scn := bufio.NewScanner(file)
	for scn.Scan() {
		var tt *vars.Task
		d := json.NewDecoder(strings.NewReader(scn.Text()))
		d.Decode(&tt)
		vars.Tasks = append(vars.Tasks,tt)
	}

	//vars.Tasks = append(vars.Tasks,tt...)
	return vars.Tasks,err
}


// 使用json Umarshal
func Readfile() ([]*vars.Task,error) {
	//file, err := os.OpenFile(vars.Todolistfile, os.O_RDONLY, 0777)
	file,err := os.Open(vars.Todolistfile)

	// 关闭文件
	defer file.Close()
	if err != nil {
		return nil,err
	}
	vars.Tasks = make([]*vars.Task,0)
	//var tt *vars.Task
	scn := bufio.NewScanner(file)
	for scn.Scan() {
		var tt *vars.Task
		 json.Unmarshal(scn.Bytes(),&tt)
		vars.Tasks = append(vars.Tasks,tt)
	}

	//vars.Tasks = append(vars.Tasks,tt...)
	return vars.Tasks,err
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


// 使用json.Newnecoder 方式
func Writetodolis(todo *vars.Task) bool {


	f, err := os.OpenFile(vars.Todolistfile, os.O_CREATE|os.O_RDWR|os.O_APPEND , 0777)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()


	//json.NewEncoder(f).Encode(vars.Tasks)
	json.NewEncoder(f).Encode(todo)

	return true
}



// 使用json.Newnecoder 方式
func ModifyWritTodoList(task []*vars.Task) {
	f, err := os.OpenFile(vars.Todolistfile, os.O_CREATE|os.O_RDWR|os.O_TRUNC , 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	//b := make([]byte,1024)
	//buf := bufio.NewWriter(f)
	defer f.Close()
		for _,v := range task {
			json.NewEncoder(f).Encode(v)

		}
}


// 使用json.Marshal 方式


func WritMarshal(todo *vars.Task) bool {
	f, err := os.OpenFile(vars.Todolistfile, os.O_CREATE|os.O_RDWR|os.O_APPEND , 0777)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()

	b,err := json.Marshal(todo)
	if err != nil {
		return false
	}
	buf := bufio.NewWriter(f)
	buf.Write(b)
	buf.Flush()

	return true
}
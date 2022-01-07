package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	cryptomd5 "todoList/cryptoMd5"
	"todoList/fromt"
	"todoList/vars"

	"github.com/howeyc/gopass"
)

// 创建task 函数
var (
	task = make([]vars.Task, 0)

	sku bool

	pathstr string = "C:/todolist2/"
	pwdfile string = pathstr + "/pwd.pwd"
)

const (
	statusRun    = "run"
	statusDelete = "delete"
	statusNone   = "none"
	statusFinish = "finish"
)

func GetId() func() int {
	id := 0
	defer func() {
		err := recover()
		if err != nil {
			//fmt.Println(err)
		}

	}()
	va, b := readtodolist()

	if b {
		id1, err := strconv.Atoi(va[len(va)-1].Id)
		if err != nil {
			id = 1
		} else {
			id = id1
		}

	}

	return func() int {
		id++
		return id

	}
}

// 将任务列表写入文件

func writetodolis(todo []vars.Task) bool {

	f, err := os.OpenFile(fmt.Sprint(pathstr+"/todo.txt"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
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

//(vtask []*vars.Task)
func readtodolist() (vtask []vars.Task, err bool) {

	f, b := os.OpenFile(fmt.Sprint(pathstr+"/todo.txt"), os.O_RDONLY, 0777)
	if b != nil {
		return nil, false
	}
	buf := bufio.NewScanner(f)
	ss1 := make([][]byte, 0)
	ss := make([]byte, 10)
	for buf.Scan() {
		//ss = append(ss, buf.Bytes())
		ss = buf.Bytes()
		ss1 = append(ss1, ss)

	}
	//fmt.Println(ss1)
	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}

	var v1 vars.Task
	// s1 := make([]byte, 2)
	for _, v := range ss1 {
		str := strings.Split(string(v), ",")
		v1 = vars.Task{
			Id:          str[0],
			Name:        str[1],
			Detailed:    str[2],
			Create_time: str[3],
			Finish_time: str[4],
			Status:      str[5],
			Founder:     str[6],
		}
		vtask = append(vtask, v1)
	}
	return vtask, true

}

func splic(s string) *vars.Task {
	a := strings.Split(s, ",")
	return &vars.Task{
		Id:          a[0],
		Name:        a[1],
		Detailed:    a[2],
		Create_time: a[3],
		Finish_time: a[4],
		Status:      a[5],
		Founder:     a[6],
	}

}

func createTask(idInt func() int, name, detailed, finish_time, founder, create_time string) vars.Task {
	idi := idInt()
	ids := strconv.Itoa(idi)
	// return map[string]string{
	// 	"id":          ids,
	// 	"name":        name,
	// 	"detailed":    detailed,
	// 	"create_time": create_time,
	// 	"finish_time": finish_time,
	// 	"status":      statusNone,
	// 	"founder":     founder,
	// }

	return vars.Task{
		Id:          ids,
		Name:        name,
		Detailed:    detailed,
		Create_time: create_time,
		Finish_time: finish_time,
		Status:      statusNone,
		Founder:     founder,
	}
}

func Addtask(idInt func() int, name, detailed, finish_time, founder, time string) []vars.Task {
	add := createTask(idInt, name, detailed, finish_time, founder, time)
	task = append(task, add)

	return task
}

// 验证任务是否存在
func checkTask(id string) (idn int, err error) {
	idn = -1
	//fmt.Println("====", task)
	for i, v := range task {
		//c, _ := v["id"]
		c := v.Id
		//fmt.Println("-----", v[id])
		if c == id {
			idn = i
			//fmt.Println("找到了")
			break
		}
	}
	if idn == -1 {
		err = fmt.Errorf("没有找到")
	}
	return idn, err
}

// 修改

func modify(id string) {
	var is, scan string
	i, err := checkTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("===================")
	fmt.Printf("%v\n", task[i])

	for {
		fmt.Println("选择要修改的字段 1.name 2. status 3. detailed  4. finish_time ")
		fmt.Scanln(&is)
		switch {
		case is == "name":
			fmt.Printf("原名字:%#v\n请输入要修改的名字:", task[i].Name)
			fmt.Scanln(&scan)
			task[i].Name = scan
		case is == "status":
			fmt.Printf("原状态:%#v\n请输入要修改的状态:", task[i].Status)
			fmt.Scanln(&scan)
			task[i].Status = scan
		case is == "detailed":
			fmt.Printf("原说明:%#v\n请输入要修改的说明:", task[i].Detailed)
			fmt.Scanln(&scan)
			task[i].Detailed = scan
		case is == "finish_time":
			fmt.Printf("完成时间:%#v\n请输入要修改完成时间:", task[i].Finish_time)
			fmt.Scanln(&scan)
			task[i].Finish_time = scan
		case is == "exit":

			fmt.Println("===================")
			return
		default:
			fmt.Println("输入有误")
		}
	}

}

func delete(id string) (bool, error) {
	i, err := checkTask(id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	a := task[:i]
	b := task[i+1:]
	c := append(a, b...)
	copy(task, c)
	task = task[:len(task)-1]
	return true, nil
}

//检查密码文件
// 1. 判断文件是否存在
// 如果不存在则要创建并且让用户初始化密码，并把密码加密后的存入文件中
// 如果存在则要读取文件密码的内容，并且让用户有输入
// 修改密码

// 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func checkPasswdfile() {
	ok, err := PathExists(pathstr)
	if !ok && err == nil {
		err := os.MkdirAll(pathstr, os.ModeDir)
		if err != nil {
			fmt.Println(err)
		}
	}

	var b []byte = make([]byte, 1024)
	f, err := os.OpenFile(pwdfile, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("打开或创建文件报错")
		return
	}
	defer f.Close()

	// 读取文件的内容并判断是否又内容！
	l, err := f.Read(b)
	if l == 0 {
		err := initpasswd(f)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if err != nil || err != io.EOF {
		fmt.Println(err)
		return
	}
}

func initpasswd(f *os.File) error {
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

func readpwd() string {
	b := make([]byte, 1024)
	f, err := os.OpenFile(pwdfile, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)

	}

	defer f.Close()
	n, _ := f.Read(b)
	str := b[:n]

	return string(str)

}

func main() {
	// 创建一个task 任务管理界面
	// 需求
	// 1. 可以新建任务
	//    任务的字段有id(自增长)  任务名称   任务详情  任务创建时间
	// 任务结束时间   任务状态   任务创建人
	idInt := GetId()
	checkPasswdfile()

	for conut := 1; conut <= 3; conut++ {
		fmt.Println("请输入您的密码:")
		st, _ := gopass.GetPasswdMasked()
		ok := cryptomd5.Auto(string(st), readpwd())
		//ok := cryptomd5.Auto(string(st))
		if ok {
			sku = true
			break
		} else if conut == 3 {
			fmt.Println("输入密码3次全部失败退出")
			os.Exit(1)
		}
	}
	for sku {
		var scan string
		fmt.Println("欢迎使用task任务管理")
		fmt.Println("请输入你的操作")
		fmt.Printf("1.create 2. delete 3. revise 4. list 5. exit :")
		fmt.Scanln(&scan)
		switch scan {
		case "create":
			now := time.Now()
			var name, detailed, finish_time, founder string
			fmt.Print("输入任务名:")
			fmt.Scanln(&name)
			fmt.Print("任务详细内容:")
			fmt.Scanln(&detailed)
			fmt.Print("结束时间:")
			fmt.Scanln(&finish_time)
			fmt.Print("创建人:")
			fmt.Scanln(&founder)
			task = Addtask(idInt, name, detailed, finish_time, founder, now.Format("2006-01-02 15:04:05"))
			writetodolis(task)
		case "delete":
			var id string
			fmt.Printf("请输入删除ID:")
			fmt.Scanln(&id)
			ok, err := delete(id)
			if ok {
				fmt.Println("删除成功")
				writetodolis(task)
			} else {
				fmt.Println(err)
			}
		case "revise":
			var id string
			fmt.Printf("请输入修改ID:")
			fmt.Scanln(&id)

			modify(id)
			writetodolis(task)

		case "list":
			// 首先要变成二维数组
			// 1. 定义一个二维数组并且把map 中的数据存入到二维数组中
			sliceone := [][]string{}

			// for _, v := range task {
			// 	sliceone = append(sliceone, []string{v["id"], v["name"], v["detailed"], v["create_time"],
			// 		v["finish_time"], v["status"], v["founder"]})
			// }
			s, err := readtodolist()
			if err {
				for _, v := range s {
					sliceone = append(sliceone, []string{v.Id, v.Name, v.Detailed, v.Create_time, v.Finish_time, v.Status, v.Founder})
				}
				fromt.WriterTable(sliceone)
			} else {
				fmt.Println("请创建数据")
			}

		case "exit":
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Println("参数有错")
		}
	}
}

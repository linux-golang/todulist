package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	cryptomd5 "todoList/cryptoMd5"

	"github.com/howeyc/gopass"
)

// 创建task 函数
var (
	task = make([]map[string]string, 0)
	now  = time.Now()

	sku bool
)

const (
	statusRun    = "run"
	statusDelete = "delete"
	statusNone   = "none"
	statusFinish = "finish"
)

func GetId() func() int {

	var id int = 0

	return func() int {
		id++
		return id
	}
}

func createTask(idInt func() int, name, detailed, finish_time, founder string) map[string]string {
	idi := idInt()
	ids := strconv.Itoa(idi)
	return map[string]string{
		"id":          ids,
		"name":        name,
		"detailed":    detailed,
		"create_time": fmt.Sprintf(now.Format("2006-01-02 15:04:05")),
		"finish_time": finish_time,
		"status":      statusNone,
		"founder":     founder,
	}
}

func Addtask(idInt func() int, name, detailed, finish_time, founder string) []map[string]string {
	add := createTask(idInt, name, detailed, finish_time, founder)
	task = append(task, add)

	return task
}

// 验证任务是否存在
func checkTask(id string) (idn int, err error) {
	idn = -1
	//fmt.Println("====", task)
	for i, v := range task {
		c, _ := v["id"]
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
			fmt.Printf("原名字:%#v\n请输入要修改的名字:", task[i]["name"])
			fmt.Scanln(&scan)
			task[i]["name"] = scan
		case is == "status":
			fmt.Printf("原状态:%#v\n请输入要修改的状态:", task[i]["status"])
			fmt.Scanln(&scan)
			task[i]["status"] = scan
		case is == "detailed":
			fmt.Printf("原说明:%#v\n请输入要修改的说明:", task[i]["detailed"])
			fmt.Scanln(&scan)
			task[i]["detailed"] = scan
		case is == "finish_time":
			fmt.Printf("完成时间:%#v\n请输入要修改完成时间:", task[i]["finish_time"])
			fmt.Scanln(&scan)
			task[i]["finish_time"] = scan
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

func main() {
	// 创建一个task 任务管理界面
	// 需求
	// 1. 可以新建任务
	//    任务的字段有id(自增长)  任务名称   任务详情  任务创建时间
	// 任务结束时间   任务状态   任务创建人
	idInt := GetId()

	for conut := 1; conut <= 3; conut++ {
		fmt.Println("请输入您的密码:")
		st, _ := gopass.GetPasswdMasked()
		ok := cryptomd5.Auto(string(st))
		ok := cryptomd5.Auto(string(st))
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
		fmt.Println(strings.Repeat("-", 25))
		fmt.Println("欢迎使用task任务管理")
		fmt.Println("请输入你的操作")
		fmt.Printf("1.create 2. delete 3. revise 4. list 5. exit :")
		fmt.Scanln(&scan)
		switch scan {
		case "create":
			var name, detailed, finish_time, founder string
			fmt.Print("输入名字:")
			fmt.Scanln(&name)
			fmt.Print("任务内容:")
			fmt.Scanln(&detailed)
			fmt.Print("结束时间:")
			fmt.Scanln(&finish_time)
			fmt.Print("创建人:")
			fmt.Scanln(&founder)
			task = Addtask(idInt, name, detailed, finish_time, founder)

		case "delete":
			var id string
			fmt.Printf("请输入删除ID:")
			fmt.Scanln(&id)
			ok, err := delete(id)
			if ok {
				fmt.Println("删除成功")
			} else {
				fmt.Println(err)
			}
		case "revise":
			var id string
			fmt.Printf("请输入修改ID:")
			fmt.Scanln(&id)

			modify(id)

		case "list":
			fmt.Println("列出")
			for _, v := range task {
				fmt.Printf("ID:%v\n", v["id"])
				fmt.Printf("任务名称:%v\n", v["name"])
				fmt.Printf("任务详情:%v\n", v["detailed"])
				fmt.Printf("任务开始时间:%v\n", v["create_time"])
				fmt.Printf("任务结束时间:%v\n", v["finish_time"])
				fmt.Printf("状态:%v\n", v["status"])
				fmt.Printf("创建人:%v\n", v["founder"])
				fmt.Println()
			}

		case "exit":
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Println("参数有错")
		}
	}
}

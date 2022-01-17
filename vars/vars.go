package vars

import (
	"fmt"
	"strconv"
)

type Task struct {
	Id          string //ID
	Name        string // 任务名
	Detailed    string // 任务详情
	Create_time string // 创建时间
	Finish_time string // 完成时间
	Status      string // 状态
	Founder     string // 执行人
}


const (
	// 任务运行状态
	statusRun    = "run"
	statusDelete = "delete"
	statusNone   = "none"
	statusFinish = "finish"
)


var (
	Sku bool
	Pathdir string = "C:/todolist2/"
	Todolistfile string = Pathdir + "todo.txt"
	Pwdfile string = Pathdir + "pwd.pwd"
	Tasks = make([]*Task, 0)
)




func Addtask(idInt func() int, name, detailed, finish_time, founder, time string) []*Task {
	add := createTask(idInt, name, detailed, finish_time, founder, time)
	Tasks = append(Tasks, add)

	return Tasks
}


func createTask(idInt func() int, name, detailed, finish_time, founder, create_time string)  *Task {
	idi := idInt()
	ids := strconv.Itoa(idi)

	return &Task{
		Id:          ids,
		Name:        name,
		Detailed:    detailed,
		Create_time: create_time,
		Finish_time: finish_time,
		Status:      statusNone,
		Founder:     founder,
	}
}


func DeleteTask (id string) (bool, error) {
	i, err := CheckTask(id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	a := Tasks[:i]
	b := Tasks[i+1:]
	c := append(a, b...)
	copy(Tasks, c)
	Tasks = Tasks[:len(Tasks)-1]
	return true, nil
}



func CheckTask(id string) (idn int, err error) {
	idn = -1
	for i, v := range Tasks {
		c := v.Id
		if c == id {
			idn = i
			break
		}
	}
	if idn == -1 {
		err = fmt.Errorf("没有找到")
	}
	return idn, err
}



func Modify(id string) {
	var is, scan string
	i, err := CheckTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("===================")
	fmt.Printf("%v\n", Tasks[i])

	for {
		fmt.Println("选择要修改的字段 1.name 2. status 3. detailed  4. finish_time ")
		fmt.Scanln(&is)
		switch {
		case is == "name":
			fmt.Printf("原名字:%#v\n请输入要修改的名字:", Tasks[i].Name)
			fmt.Scanln(&scan)
			Tasks[i].Name = scan
		case is == "status":
			fmt.Printf("原状态:%#v\n请输入要修改的状态:", Tasks[i].Status)
			fmt.Scanln(&scan)
			Tasks[i].Status = scan
		case is == "detailed":
			fmt.Printf("原说明:%#v\n请输入要修改的说明:", Tasks[i].Detailed)
			fmt.Scanln(&scan)
			Tasks[i].Detailed = scan
		case is == "finish_time":
			fmt.Printf("完成时间:%#v\n请输入要修改完成时间:", Tasks[i].Finish_time)
			fmt.Scanln(&scan)
			Tasks[i].Finish_time = scan
		case is == "exit":

			fmt.Println("===================")
			return
		default:
			fmt.Println("输入有误")
		}
	}

}
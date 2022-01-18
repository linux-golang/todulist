package graphic

import (
	"fmt"
	"os"
	"time"
	"todoList/fileModel"
	"todoList/fromt"
	"todoList/vars"
)

func GraphicInterface(Idint func() int) {

	for vars.Sku {
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
			task := vars.CreateTask(Idint, name, detailed, finish_time, founder, now.Format("2006-01-02 15:04:05"))
			//fileModel.Writetodolis(task)
			fileModel.WritMarshal(task)
		case "delete":
			var id string
			fmt.Printf("请输入删除ID:")
			fmt.Scanln(&id)
			ok, err := vars.DeleteTask(id)
			if ok {
				fmt.Println("删除成功")
				//fileModel.Writetodolis(vars.Tasks)
			} else {
				fmt.Println(err)
			}
		case "revise":
			var id string
			fmt.Printf("请输入修改ID:")
			fmt.Scanln(&id)

			vars.Modify(id)

			// 在切片修改后，在更新到文件中

			fileModel.ModifyWritTodoList(vars.Tasks)

		case "list":
			// 首先要变成二维数组
			// 1. 定义一个二维数组并且把map 中的数据存入到二维数组中
			sliceone := [][]string{}

			task,b := fileModel.ReadFileTodolist()



			if b {
				for _, v := range task {
					sliceone = append(sliceone, []string{v.Id, v.Name, v.Detailed, v.Create_time, v.Finish_time, v.Status, v.Founder})
				}

				fromt.WriterTable(sliceone)
			}

		case "exit":
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Println("参数有错")
		}
	}
}

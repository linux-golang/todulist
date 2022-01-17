package main

import (
	"fmt"
	"todoList/fileModel"
	"todoList/graphic"
	"todoList/model"
	"todoList/vars"
)

// 创建task 函数

func main() {
	// 创建一个task 任务管理界面
	// 需求
	// 1. 可以新建任务
	//    任务的字段有id(自增长)  任务名称   任务详情  任务创建时间
	// 任务结束时间   任务状态   任务创建人
	IdInt := model.GetId()
	err := fileModel.ChechFileIsExit(vars.Pwdfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	model.CheckPasswd()
	graphic.GraphicInterface(IdInt)

}

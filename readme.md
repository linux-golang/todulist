

#### go module 模式

package  这个关键字是定义包
import 引入包
若要引入多个包可以放到()中
import (
     _ "xxxx"   // 这样表示这个包没有在程序中使用，但是会初始化这个包的一下函数
     . "xxxx"   // 这个表示直接把这个包的函数导入到当前程序中，放问的时候 可以不加包名来直接调用
     alisaName "xxxx" // 给这个包定义一个别名，调用alisaName.funcName() 即可
)

如何func 、type 、变量、常量位于不同的包下，则需要他们的首字母大写，表示是公开可以访问的，对于结构体下的字段，如果想要在包外访问，还需要将字段变量名首字母大写

Go Modules 
go modules 于1.11 版本初步引入，在1.12 中正式支持，
1. 如何使用Go modules 
   在系统环境开启GO111MODULE=on
2. 设置goproxy代理
    GOPROXY=https://goproxy.cn
3. 创建一个项目文件夹，进入文件夹后使用下面命令创建一个新的go.mod
   go mod init [project name]

4. 可以通过 go mod download  下载依赖
5. go mod tidy 命令来更新依赖关系

#### golang 标准包

+ flag   // 命令行参数解析
+ os     // 调用系统接口的
+ time   // 时间相关的包
+ sort   // 排序的包
+ log    // 日志包
+ exec   // 执行系统命令

>  用于查看各个包的api 文件请见官方文档https://pkg.go.dev/   or  https://golang.google.cn/pkg/





#### 自定义数据类型

​	在go中可以使用`type` 关键字来定义自定义数据类型

```go
package main

import "fmt"

// 定义一个名为con 的数据类型
type con int

// 定义一个map 类型的自定义数据类型
type todo map[string]string

// 也可以吧一个函数定义为类型
type ff func(s ...int)

func main() {

	var c con
	fmt.Println(c)

	var t todo = map[string]string{
		"jj": "ss",
		"22": "333",
	}
	fmt.Println(t)

	var f ff = func(s ...int) {
		fmt.Println(s)
	}
	f(1, 2, 3)

}

```



#### 结构体

​	Golang 也支持面向对象（OOP），但是和传统的面向对象编程有区别，并不是纯粹的面向对象语言，所以我们说Golang支持面向对象编程。

​	Golang没有类（Class）,Go语言的结构体（struct）和其它编程语言的类（Class）有同等的地位，你可以理解Golang是基于struct 来实现OOP

​	Golang面向对象编程非常简洁，仍然有面向对象编程的继承，封装，多态的特性，只是实现的方式和其它OOP语言不一样

##### struct 基本使用

```go
package main

import (
	"fmt"
	"time"
)

// task 结构体
type task struct {
	name      string
	status    int
	startTime time.Time
	auth      string
}

func main() {

	//创建一个task 实例
	var s task
	// 第一种方式赋值
	s = task{"name", 1, time.Now(), "guobb"}
	fmt.Printf("%#v\n", s)

	s = task{
		name:      "name",
		status:    2,
		startTime: time.Now(),
		auth:      "kk",
	}
	fmt.Printf("%#v\n", s)

	// 创建一个空的实例
	s = task{}
	fmt.Printf("%#v\n", s)

	// 单个赋值
	s.name = "这是一个测试"
	fmt.Printf("%#v\n", s)
    
    // 创建一个指针类型的实例
	var ptask *task    // 这里如果没有赋值 则是nil  

	ptask = &task{
		name:   "这是一个任务",
		status: 22,
	}
	fmt.Printf("%#v\n", ptask)
    
    // 创建一个空值的类型的struct 指针类型
	pst2 := new(task)
	fmt.Printf("%#v\n", pst2)

}

```



##### 方法

定义方法

```go

package main

import (
	"fmt"
	"time"
)

// 定义一个结构体
type Task struct {
	name      string
	id        int
	startTime *time.Time
	user      string
}

// 定义一个方法

func (task *Task) SetName(name string) {
	task.name = name
}

func (task *Task) GetName() string {
	return task.name
}

func main() {

	task := Task{name: "hhh"}
	task.SetName("123")
	fmt.Println(task.GetName())

}

```



```go

package main

import "fmt"

// 创建一个struct

type task struct {
	id      int
	name    string
	address string
}

// 创建一个GetName 函数获取name的值
// 如果只是查看某个值，则接受者是不是指针类型都可以
func (t task) GetName() {
	fmt.Println(t.name)
}

// 创建一个改变name值的一个方法
// 这里我们要改变结构体中的某个值，则必须要使用真正
func (t *task) SetName(name string) {
	t.name = name
	return
}

func main() {
	task1 := task{1, "gg", "北京"}  // 创一个task 实例e
	task2 := &task{2, "gg", "北京"} // 这是一个创建一个指针类型的

	//task1.SetName("bb")   //虽然task1 实例并不是指针类型，但是SetName 这个方法的接受者是指针
	// go 会自动传引用，所以这样写Name 的值也会变成bb
	(&task1).SetName("bb") // 这样也可以
	task1.GetName()
	fmt.Println("=======")
	task2.SetName("cc")
	task2.GetName()
}

```

```go
package main

import "fmt"

// 创建一个struct

type task struct {
	id      int
	name    string
	address string
}

// 创建一个GetName 函数获取name的值
// 如果只是查看某个值，则接受者是不是指针类型都可以
func (t task) GetName() {
	fmt.Println(t.name)
}

// 创建一个改变name值的一个方法
// 这里我们要改变结构体中的某个值，则必须要使用真正
func (t *task) SetName(name string) {
	t.name = name
	return
}

func main() {
	// 方法表达式  结构体.方法名
	// 对于值接受者，可以通过指针/值来获取方法表达式
	var t task

        method1 := task.GetName
	method2 := (*task).SetName // 这个指针类型的必须要指明。
	// method2 的类型 func(main.task,string)
	method1(t)
	method2(&t, "cc")
	fmt.Printf("%#v\n", t)

}

```




### go module 模式
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

flag   // 命令行参数解析
os     // 调用系统接口的
time   // 时间相关的包
sort   // 排序的包
log    // 日志包

package main

import (
	"fmt"
	"testing"
	"todoList/vars"
)

func TestCheckPasswdfile(t *testing.T) {
	checkPasswdfile()
}

func TestReadtodolist(t *testing.T) {
	f, _ := readtodolist()
	fmt.Printf("%#v\n", f)
}

func TestSplic(t *testing.T) {
	splic("2,2,2,2022-01-06 22:11:23,2022-01-06 22:11:23,none,2")
}

func TestWritetodolis(t *testing.T) {
	v := []vars.Task{vars.Task{Id: "1", Name: "1", Detailed: "1", Create_time: "2", Finish_time: "2", Status: "-", Founder: "1"}, vars.Task{Id: "2", Name: "2", Detailed: "2", Create_time: "2", Finish_time: "2", Status: "-", Founder: "1"}, vars.Task{Id: "3", Name: "3", Detailed: "3", Create_time: "2", Finish_time: "2", Status: "-", Founder: "1"}}
	writetodolis(v)

}

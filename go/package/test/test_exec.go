package test

import (
	"os/exec"
	"fmt"
)

func TestExec(){
	cmdo := exec.Command("dir")
	path := cmdo.Path
	fmt.Println("hello",path)
	if err := cmdo.Start(); err != nil {
		fmt.Println("Error: the Command No.0 can not be start up: %s\n" , err)
	}
}
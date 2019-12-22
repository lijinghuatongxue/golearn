package main

import (
	"os/exec"
)

func shell()  {
	cmd:=exec.Command("echo2")
	_,err:= cmd.Output()
	if err != nil{
		panic(err.Error())
	}
}
func main()  {
	shell()
}
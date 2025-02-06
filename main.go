package main

import (
	"os"
	"social_web_server/cmd"
)
func main(){
	if err:=cmd.Execute();err!=nil{
		os.Exit(1)
	}
}
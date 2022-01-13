package main

import (
	"../utils"
	"flag"
	"fmt"
	"strings"
)


func main(){
	userPtr := flag.String("user", "liuchengxiang", "input your username prefix")
	flag.Parse()

	msg, err := utils.RunGitCommand("git", "log", "--author="+*userPtr, "--pretty=tformat:", "--numstat")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	modifyFiles := strings.Split(msg, "\n")

	for _, item := range modifyFiles {
		fmt.Println(item)
	}
	fmt.Println(len(modifyFiles))
}

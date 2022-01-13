package main

import (
	"../utils"
	"flag"
	"fmt"
	"strconv"
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

	var addedSum = 0
	var removedSum = 0
	modifyFiles := strings.Split(msg, "\n")
	for _, item := range modifyFiles {
		changes := strings.Split(item, "\t")
		if len(changes) != 3 {
			continue
		}
		added, _ := strconv.Atoi(changes[0])
		removed, _ := strconv.Atoi(changes[1])

		addedSum += added
		removedSum += removed
	}

	fmt.Printf("In project, you have added %d lines, removed %d lines, total modified %d liens, good job! :)\n", addedSum, removedSum, addedSum+removedSum)
}

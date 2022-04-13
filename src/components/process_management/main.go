package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
)

func main() {
	listProcessInfo()
}

func listProcessInfo() {
	processes, _ := ps.Processes()
	var tasks []string
	for _, p := range processes {
		currenProccessInfo(p.Pid())
		isExist := false
		for _, t := range tasks {
			if t == p.Executable() {
				isExist = true
			}
		}
		if !isExist {
			tasks = append(tasks, p.Executable())
		}
	}
	for _, t := range tasks {
		fmt.Println(t)
	}
}

func currenProccessInfo(pid int) {
	// 全てのプロセスを取得
	pidInfo, _ := ps.FindProcess(pid)

	//fmt.Printf("> PID : %d", pidInfo.Pid())
	//fmt.Printf("> PPID : %d", pidInfo.PPid())
	//fmt.Printf("> Process name : %s", pidInfo.Executable())

	ppidInfo, err := ps.FindProcess(pidInfo.PPid())
	if err == nil && ppidInfo != nil {
		//fmt.Printf("> Parent process name : %s\n", ppidInfo.Executable())
	} else {
		//fmt.Printf("> Parent process name : noting\n")
	}
}

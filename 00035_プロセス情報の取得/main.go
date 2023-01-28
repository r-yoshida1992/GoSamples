package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
)

func main() {
	fetchProcessList()
}

// マシン上のプロセス情報を取得する
func fetchProcessList() {
	processes, _ := ps.Processes()
	for _, p := range processes {
		fmt.Printf("pid:%d,executable:%s,ppid:%d\n", p.Pid(), p.Executable(), p.PPid())
	}
}

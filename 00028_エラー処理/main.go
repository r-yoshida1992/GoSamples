package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When string // エラー発生時刻
	What string // エラー理由
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %v", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now().Format("2006/01/02 15:04:05"),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

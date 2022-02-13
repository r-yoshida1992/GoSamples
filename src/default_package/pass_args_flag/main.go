package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	_, err := fmt.Println(flag.Args()[0])
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var loadingText string
	for i := 0; i <= 100; i++ {
		loadingText = fmt.Sprintf("\r\x1b[36mLoading... %3d%%\x1b[0m", i)
		fmt.Print(loadingText)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Printf("\r%s", strings.Repeat(" ", len(loadingText)))
	fmt.Println("\rDone!")
}

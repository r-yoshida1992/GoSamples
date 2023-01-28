package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	writeA()
	fmt.Printf("elapsed time: %vms\n", time.Since(now).Milliseconds())

	fmt.Println()

	now = time.Now()
	writeB()
	fmt.Printf("elapsed time: %vms\n", time.Since(now).Milliseconds())
}

func writeA() {
	fmt.Println("executing writeA.")
	input, _ := os.Open("./input.csv")
	defer input.Close()
	output, _ := os.Create("./output.csv")
	defer output.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		output.Write(append([]byte(line), '\n'))
	}
	fmt.Println("Writing file to output.csv is completed.")
}

func writeB() {
	fmt.Println("executing writeB.")
	bytes, _ := os.ReadFile("./input.csv")
	os.WriteFile("./output.csv", bytes, 0666)
	fmt.Println("Writing file to output.csv is completed.")
}

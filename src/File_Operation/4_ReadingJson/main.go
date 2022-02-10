package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	raw, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	var persons []Person
	json.Unmarshal(raw, &persons)

	for _, person := range persons {
		fmt.Println(person.Name)
		fmt.Println(person.Age)
		fmt.Println(person.Gender)
	}

}

type Person struct {
	Name   string
	Age    int
	Gender string
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("<==>")

	newjson := &Input{}
	err := json.Unmarshal([]byte(input), &newjson)
	if err != nil {
		panic(err)
	}

	for k, v := range newjson.Machine.Beverages {
		fmt.Println(k, v)
	}

	fmt.Println("<==>")

	m := &Machine{}

	fmt.Println(m)

	m.Init(newjson.Machine)

	fmt.Println(m)

	m.TryServe(newjson.Machine.Beverages)
}

type inp struct {
	mp map[string]map[string]interface{}
}

type iMachine struct {
	Outlets int

	Items map[string]int
}

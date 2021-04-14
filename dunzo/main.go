package main

import (
	"encoding/json"
	"fmt"
	"lld/dunzo/models"
	"lld/dunzo/processing"
)

func main() {
	newjson := &models.JsonInput{}
	err := json.Unmarshal([]byte(input), &newjson)
	if err != nil {
		panic(err)
	}

	m := &models.Machine{}

	m.InitWithJson(newjson.InputMachine)

	fmt.Println(m)

	m.ServeJson(processing.Common, newjson.Beverages)

	fmt.Println(m)
}

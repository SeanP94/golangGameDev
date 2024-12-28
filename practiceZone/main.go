package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:age`
	Vibe int    `json:vibe`
}

func main() {
	json_data, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var person Person
	err = json.Unmarshal(json_data, &person)
	if err != nil {
		fmt.Println("Error reado JSON:", err)
		return
	}

	fmt.Printf(`%s is %d years old, with a vibe of %d out of 10.`, person.Name, person.Age, person.Vibe)
}

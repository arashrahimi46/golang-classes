package main

import (
	"fmt"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	persons := []Person{
		{
			FirstName: "arash",
			LastName:  "rahimi",
			Age:       20,
		}, {
			FirstName: "payam",
			LastName:  "payami",
			Age:       30,
		},
	}
	file, _ := os.Create("result.txt")
	defer file.Close()
	for _, person := range persons {
		file.WriteString("first name: " + person.FirstName + "\n")
		file.WriteString("last name: " + person.LastName + "\n")
		file.WriteString(fmt.Sprintf("%s %d \n" , "Age: " , person.Age))
		file.WriteString("------------------------ \n")
	}
}

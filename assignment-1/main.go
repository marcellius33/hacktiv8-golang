package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Name       string
	Address    string
	Occupation string
	Reason     string
}

func (s Student) Description() {
	fmt.Println("Description")
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Address: %s\n", s.Address)
	fmt.Printf("Occupation: %s\n", s.Occupation)
	fmt.Printf("Reason: %s\n", s.Reason)
}

func main() {
	id, _ := strconv.Atoi(os.Args[1])

	students := []Student{
		{
			Name:       "Varrel",
			Address:    "Jln. Angkasa",
			Occupation: "Chef",
			Reason:     "Love Golang",
		},
		{
			Name:       "Marcellius",
			Address:    "Jln. Bintang",
			Occupation: "Teacher",
			Reason:     "Eager to learn more",
		},
		{
			Name:       "William",
			Address:    "Jln. Kambing",
			Occupation: "Firefighter",
			Reason:     "Have an interest in programming",
		},
		{
			Name:       "James",
			Address:    "Jln. Bangka",
			Occupation: "Analyst",
			Reason:     "Love to learn new things",
		},
		{
			Name:       "Jason",
			Address:    "Jln. Kolonial",
			Occupation: "Sales",
			Reason:     "Want to use golang in his work",
		},
	}

	if id >= len(students) {
		fmt.Println("Id not found!!!")
		return
	}
	students[id].Description()
}

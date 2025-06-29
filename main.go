package main

import {
	"encoding/json"
	"fmt"
	"os"
	"strings"
}

type Contact struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var contacts []Contact

func main(){
	if len(os.Args)<2{
		fmt.Println("Usage: add/list/search")
		return
	}

	command := os.Args[1]

	switch command{
	case "add":
		if len(os.Args)<5{
			fmt.Println("Usage : add <name> <email> <phone>")
			return
		}
		name := os.Args[2]
		email := os.Args[3]
		phone := os.Args[4]
		addContact(name, email, phone)
	
	case "list":
		listContacts()
	
	case "search":
		if len(os.Args)<3 {
			fmt.Println("Usage: search <name>")
			return
		}
		searchContact(os.Args[2])

	default:
		fmt.Println("Unknown command:", command)

	}
}


package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var contacts []Contact

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: add/list/search")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 {
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
		if len(os.Args) < 3 {
			fmt.Println("Usage: search <name>")
			return
		}
		searchContact(os.Args[2])

	default:
		fmt.Println("Unknown command:", command)

	}
}

func addContact(name, email, phone string) {
	newContact := Contact{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	var file *os.File
	if _, err := os.Stat("contacts.json"); err == nil {
		fmt.Println("File exists.")
		content, err := os.ReadFile("contacts.json")
		if err != nil {
			fmt.Println("Error reading Existing file: ", err)
			return
		}

		err = json.Unmarshal(content, &contacts)
		if err != nil {
			fmt.Println("Error parsing existing contacts: ", err)
			return
		}

		file, err = os.Create("contacts.json")
		if err != nil {
			fmt.Println("Error opening file for writing: ", err)
			return
		}
	} else {
		fmt.Println("File does not exist, creating a new one.")
		file, err = os.Create("contacts.json")
		if err != nil {
			fmt.Println("Error creating file: ", err)
			return
		}
	}
	defer file.Close()

	contacts = append(contacts, newContact)

	saveContactsToFile()

	fmt.Println("Contact added sucsessfully.")
}

func saveContactsToFile() {

	file, err := os.OpenFile("contacts.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file for writing: ", err)
		return
	}
	defer file.Close()
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		fmt.Println("Error encoding contacts: ", err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}

}

func listContacts() {
	loadContactsFromFile()

	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	for i, contact := range contacts {
		fmt.Printf("%d. %s | %s | %s\n", i+1, contact.Name, contact.Email, contact.Phone)
	}
}

func loadContactsFromFile() {
	file, err := os.Open("contacts.json")
	if err != nil {
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	json.Unmarshal(data[:count], &contacts)
}

func searchContact(name string) {
	loadContactsFromFile()

	found := false
	for _, contact := range contacts {
		if strings.Contains(strings.ToLower(contact.Name), strings.ToLower(name)) {
			fmt.Printf("Found: %s | %s | %s\n", contact.Name, contact.Email, contact.Phone)
			found = true
		}
	}

	if !found {
		fmt.Println("No contact found with that name.")
	}
}

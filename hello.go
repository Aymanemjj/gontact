package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Contact struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Number   string `json:"number"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var PATH string = "contacts.json"

func main() {

	for {
		choice := menu()
		fmt.Println(choice)

		switch choice {
		case 1:
			contact := addContact()
			saveToJson(contact)
		case 2:
			listContacs()
		case 3:
			removeContact()
		case 4:
		case 5:
			fmt.Println("Good bye!")
			return
		}
	}
}

func menu() int {

	options := map[int]string{
		1: "Add contact",
		2: "List contacts",
		3: "Remove a contact",
		4: "Update a contact",
		5: "Exit",
	}

	fmt.Println("Pick an option")
	keys := make([]int, 0, len(options))

	for k := range options {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d. %s\n", k, options[k])
	}
	fmt.Print("\n")
	var choice int
	fmt.Scanln(&choice)
	fmt.Print("\n")
	return choice
}

func addContact() Contact {
	input := "static"
	c := Contact{}
	//Email
	fmt.Println("Input the contact's email")
	fmt.Scanln(&input)
	c.Email = input

	//Fullname
	fmt.Println("Input the contact's fullname")
	fmt.Scanln(&input)
	c.Fullname = input

	//Number
	fmt.Println("Input the contact's phone number")
	fmt.Scanln(&input)
	c.Number = input

	return c
}

func saveToJson(contact Contact) {
	var payload []Contact
	file, err := os.OpenFile(PATH, os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()
	content, err := os.ReadFile(PATH)
	check(err)
	if len(content) > 0 {
		err = json.Unmarshal(content, &payload)
		check(err)
	}
	payload = append(payload, contact)
	encodedContact, _ := json.MarshalIndent(payload, "", "  ")

	file.Write(encodedContact)
	check(err)
	file.Sync()
}

func listContacs() {
	content, err := os.ReadFile(PATH)
	if err != nil {
		fmt.Println("No contacts found yet.")
		return
	}

	var payload []Contact
	err = json.Unmarshal(content, &payload)
	check(err)
	for i, c := range payload{
		fmt.Printf("%d. FullName: %s, Email: %s, PhoneNumber: %s\n", i, c.Fullname, c.Email, c.Number)
	}
	fmt.Print("\n\n")
}

func removeContact(){
	var fname string
	fmt.Print("Input the fullname of the unwanted contact: ")
	fmt.Scanln(&fname)
	
}
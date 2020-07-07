package main

import "fmt"

type profile struct {
	name       string
	address    string
	email      string
	contactNum string
}

func (p profile) getName() string {
	return p.name
}

func (p profile) getAddress() string {
	return p.address
}

func (p profile) getContact() map[string]string {
	m := map[string]string{"email": p.email, "contactnum": p.contactNum}
	return m
}

func methodDemo() {
	s := profile{
		name:       "Sherlock",
		address:    "221B Baker Street",
		email:      "Sherlockhomes.com",
		contactNum: "18811904",
	}

	fmt.Println("Name", s.getName())
	fmt.Println("Address", s.getAddress())
	fmt.Println("Contact Details", s.getContact())
	fmt.Println("Email Extraction", s.getContact()["email"])
}

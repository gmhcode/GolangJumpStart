package main

import "fmt"

type userProfile interface {
	getName() string
	getAddress() string
	getContact() map[string]string
}

func getUserInfo(u userProfile) {
	fmt.Println("Name", u.getName())
	fmt.Println("Address", u.getAddress())
	fmt.Println("Contact Details", u.getContact())
	fmt.Println("Email Extraction", u.getContact()["email"])
}

func interfaceDemo() {
	s := profile{
		name:       "Sherlock",
		address:    "221B Baker Street",
		email:      "Sherlockhomes.com",
		contactNum: "18811904",
	}

	getUserInfo(s)

}

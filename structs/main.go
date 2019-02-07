package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main(){
	jose := person{
		firstName: "Jose",
		lastName: "Cordova",
		contact: contactInfo{
			email: "jfernancordova@gmail.com",
			zipCode: 1426,
		},
	}

	fmt.Printf("%+v", jose)
}
package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main(){
	jose := person{
		firstName: "Jose",
		lastName: "Cordova",
		contactInfo: contactInfo{
			email: "jfernancordova@gmail.com",
			zipCode: 1426,
		},
	}
	
	jose.updateName("fernando")
	jose.print()
}

func (p person) updateName(newFirstName string){
	p.firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}
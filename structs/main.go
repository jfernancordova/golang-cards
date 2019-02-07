package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main(){
	var jose person

	jose.firstName = "Jose"
	jose.lastName = "Cordova"

	fmt.Println(jose)
	fmt.Printf("%+v", jose)
}
package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main(){
	jose := person{ firstName: "José", lastName: "Cordova"}
	fmt.Println(jose)

}
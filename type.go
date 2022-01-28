package main

import "fmt"

func main() {
    type PersonData struct {
    	FirstName string
	LastName  string
    }

    Person := PersonData{
        FirstName: "Victor",
	LastName:  "Fernandez",
    }

    fmt.Println(Person)
    fmt.Println(Person.FirstName)
    fmt.Println(Person.LastName)
}

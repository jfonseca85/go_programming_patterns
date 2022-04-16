package main

import (
	"fmt"
)

//WithName - Struct embedding
type WithName struct {
	Name string
}

//Country - Structure that encapsulates the Country name
type Country struct {
	WithName
}

//City - Structure that encapsulates the City name
type City struct {
	WithName
}

//Printable - A Country and City interface Both implement PrintS() interface methods and produce themselves.
type Printable interface {
	PrintStr()
}

//PrintStr -Applying polymorphism in the PrintStr function to the WithName structure
func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

//main function that runs our scenario
func main() {
	c1 := Country{WithName{Name: "São Paulo"}} //messy startup
	c2 := City{WithName{Name: "Belém"}}        //messy startup
	c1.PrintStr()
	c2.PrintStr()
}

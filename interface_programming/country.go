package main

import "fmt"

//Country - Structure that encapsulates the Country name
type Country struct {
	Name string //duplicate code
}

//City - Structure that encapsulates the City name
type City struct {
	Name string //duplicate code
}

//Printable - A Country and City interface Both implement PrintS() interface methods and produce themselves.
type Printable interface {
	PrintStr()
}

//PrintStr -Applying polymorphism in the PrintStr function to the Country structure
func (c Country) PrintStr() {
	fmt.Println(c.Name) //duplicate code
}

//PrintStr -Applying polymorphism in the PrintStr function to the PrintStr structure
func (c City) PrintStr() {
	fmt.Println(c.Name) //duplicate code
}

//main function that runs our scenario
func main() {
	c1 := Country{"São Paulo"}
	c2 := City{"Belém"}
	c1.PrintStr()
	c2.PrintStr()
}

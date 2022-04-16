package main

import (
	"fmt"
)

//Country - Structure that encapsulates the Country name
type Country struct {
	Name string
}

//City - Structure that encapsulates the City name
type City struct {
	Name string
}

//Stringable - A Country and City interface Both implement PrintS() interface methods and produce themselves.
type Stringable interface {
	ToString() string
}

//ToString - Applying polymorphism in the ToString function to the Country structure
func (c Country) ToString() string {
	return "Country = " + c.Name
}

//ToString - Applying polymorphism in the ToString function to the City structure
func (c City) ToString() string {
	return "City = " + c.Name
}

//PrintStr - Function that receives the Stringable interface, which can be used by Country and City implementations
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

//main - function that runs our scenario
func main() {
	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}

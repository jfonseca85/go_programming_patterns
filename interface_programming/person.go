package main

import "fmt"

type Person struct {
	Name string
	Sex  string
	Age  int
}

//Print - Function that prints data as Person structure
func (p *Person) Print() {
	fmt.Printf("Name=%s, Sex=%s, Age=%d\n",
		p.Name, p.Sex, p.Age)
}

//PrintPerson - Member function ( Receiver ) that prints data as Person structure
func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sex=%s, Age=%d\n",
		p.Name, p.Sex, p.Age)
}

func main() {
	var p = Person{
		Name: "Rafael Oliveira",
		Sex:  "Male",
		Age:  35,
	}

	PrintPerson(&p)
	p.Print()
}

package main

import "fmt"

//var _ Shape = (*Square)(nil)

//Shape - Shape - Interface to abstract the polygon type
type Shape interface {
	Sides() int
	Area() int //Method not implemented
}

//Square - Structure that implements a Shape interface
type Square struct {
	len int
}

//Slices - Get the number of sides of the square
func (s *Square) Sides() int {
	return 4
}

//main - function that runs our scenario
func main() {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}

package main

import (
	"fmt"
)

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func NameVisitorFunc(info *Info, err error) error {
	fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
	fmt.Println("NameVisitor() after call function")
	return err
}

func OtherThingsVisitor(info *Info, err error) error {
	fmt.Println("OtherThingsVisitor() before call function")

	fmt.Printf("==> OtherThings=%s\n", info.OtherThings)

	fmt.Println("OtherThingsVisitor() after call function")
	return err

}

func main() {
	info := Info{}
	var v Visitor = &info

	loadFile := func(info *Info, err error) error {
		info.Name = "Jorge Luis"
		info.Namespace = "JorgeLuis"
		info.OtherThings = "We are running as remote team."
		return nil
	}

	v = NewDecoratedVisitor(v, NameVisitorFunc, OtherThingsVisitor)

	v.Visit(loadFile)

}

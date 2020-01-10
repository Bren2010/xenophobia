package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("plug/plug.so")
	if err != nil {
		fmt.Println(err)
	}
	b, err := p.Lookup("Been")
	if err != nil {
		fmt.Println(err)
	}
	been, ok := b.(func() error)
	if !ok {
		fmt.Println("Whoops!")
	}
	err = been()
	if err != nil {
		fmt.Println(err)
	}
	a, err := p.Lookup("Add")
	if err != nil {
		fmt.Println(err)
	}
	add, ok := a.(func())
	if !ok {
		fmt.Println("Whoops!")
	}
	add()
}

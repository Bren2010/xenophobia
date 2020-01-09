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
	x, err := p.Lookup("Xeno")
	if err != nil {
		fmt.Println(err)
	}
	xeno, ok := x.(func() error)
	if !ok {
		fmt.Println("Whoops!")
	}
	err = xeno()
	if err != nil {
		fmt.Println(err)
	}
}

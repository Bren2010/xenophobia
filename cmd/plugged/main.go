package main

import (
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("plug/plug.so")
	if err != nil {
		log.Fatal(err)
	}
	d, err := p.Lookup("Demo")
	if err != nil {
		log.Fatal(err)
	}
	demo := d.(func())
	demo()
}

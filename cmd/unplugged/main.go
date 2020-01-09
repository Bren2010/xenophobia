package main

import (
	"fmt"

	"gitlab.com/tymorl/xenophobia/pkg/been"
)

func main() {
	err := been.Xenophobic()
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"gitlab.com/tymorl/xenophobia/pkg/add"
	"gitlab.com/tymorl/xenophobia/pkg/been"
)

func Been() error {
	return been.Xenophobic()
}

func Add() {
	add.Xenophobic()
}

package model

import "fmt"

type Parent struct {
	Name   string
	Gender string
}

func NewParent(name string, gender string) *Parent {
	return &Parent{name, gender}
}

func (p *Parent) Walk() {
	fmt.Println("Parent Walk 被调用!")
}

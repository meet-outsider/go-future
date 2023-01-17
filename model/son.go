package model

import "fmt"

type Son struct {
	Parent
	action string
}

func NewSon(parent *Parent, action string) *Son {
	return &Son{*parent, action}
}

func (s Son) Walk() string {
	fmt.Println("action is:", s.action)
	return "Son is Walking"
}
func (s Son) Say() string {
	return "Son is Saying"
}
func (s Son) Run() string {
	return "Son is Running"
}

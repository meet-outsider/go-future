package model

type Action interface {
	Say() string
	Walk() string
	Run() string
}
type Walk interface {
	Walk() string
}

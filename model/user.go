package model

type User struct {
	Name   string
	age    uint32
	status int
}

func (u User) new() {

}
func (u User) Say() string {
	return "User is Saying"
}
func (u User) Walk() string {
	return "User is Walking"
}
func (u User) Run() string {
	return "User is Running"
}

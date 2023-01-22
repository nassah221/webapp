package model

type User struct {
	ID       string
	Username string
	Password []byte
	Feedback []Feedback
}

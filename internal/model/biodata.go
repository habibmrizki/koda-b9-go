package model

type User struct {
	Name        string
	Photo       string
	Email       string
	Age         uint8
	PhoneNumber string
	IsMarried   bool
	Education   []Educations
}

type Educations struct {
	Name  string
	Major string
}

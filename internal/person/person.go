package person

import (
	"fmt"
)

type Person struct {
	Name    string
	Address string
	Phone   string
}

func (p *Person) GetName() {
	fmt.Println(p.Name)
}

func (p *Person) GetAddress() string {
	return p.Address
}

func (p *Person) GetPhone() string {
	return p.Phone
}

// Method Print
func (p *Person) Print() string {
	return fmt.Sprintf("Nama saya adalah: %s, Alamat saya di: %s, Nomor Hp saya adalah: %s", p.Name, p.Address, p.Phone)
}

// Method Greet
func (p *Person) Greet() string {
	return "Hello, nama saya " + p.Name
}

// Method Setter
func (p *Person) SetName(newName string) {
	p.Name = newName
}

// Method Cunstructor
func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

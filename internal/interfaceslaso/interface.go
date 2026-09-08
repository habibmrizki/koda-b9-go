package interfaceslaso

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount []int) (string, error)
}

type Bank struct {
	Name string
}

type Online struct {
	Name string
}

type Fiktif struct {
	Lists []int
}

func (b *Bank) Pay(amount []int) (string, error) {
	total := 0
	for _, value := range amount {
		total += value
	}

	if total <= 0 {
		return "", errors.New("Total pembayaran harus lebih dari 0")
	}

	return fmt.Sprintf("Pembayaran %d berhasil via Bank %s\n", total, b.Name), nil
}

func (o *Online) Pay(amount []int) (string, error) {
	total := 0
	for _, value := range amount {
		total += value
	}

	if total <= 0 {
		return "", errors.New("Total pembayaran harus lebih dari 0")
	}

	return fmt.Sprintf("Pembayaran %d via Online %s\n", total, o.Name), nil
}

func (f *Fiktif) Pay(amount []int) (string, error) {
	total := 0
	for _, value := range amount {
		total += value
	}

	if total <= 0 {
		return "", errors.New("total pembayaran fiktif kurang dari atau sama dengan 0")
	}

	f.Lists = append(f.Lists, total)
	return "", nil
}

func (f *Fiktif) Total() int {
	totalFiktif := 0
	for _, v := range f.Lists {
		totalFiktif += v
	}
	return totalFiktif
}

func Checkout(method PaymentMethod, amount []int) (string, error) {
	return method.Pay(amount)
}

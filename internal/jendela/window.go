package jendela

import (
	"errors"
	"fmt"
)

func JendelaBingkai(ukuran uint8) error {
	if ukuran <= 0 {
		return errors.New("Ukuran jendela harus lebih besar dari 0")
	}

	tengah := ukuran / 2

	for i := uint8(0); i < ukuran; i++ {
		for j := uint8(0); j < ukuran; j++ {
			if i == 0 || i == ukuran-1 || j == 0 || j == ukuran-1 || i == tengah || j == tengah {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	// fmt.Println()
	return nil
}

package deferandpanic

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path []byte) {
	//  Defer untuk recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Continue...")
		}
	}()

	//  Buka file
	file, err := os.Open(string(path))
	if err != nil {
		fmt.Printf("gagal membuka file: %v\n", err)
		return
	}

	// Defer untuk menutup file
	defer func() {
		fmt.Println("Menutup file...")
		file.Close()
	}()

	// Baca konten file
	content, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("error membaca file: %v", err))
	}

	//  Jika sukses
	fmt.Println("File berhasil dibaca:")
	fmt.Println(string(content))

}

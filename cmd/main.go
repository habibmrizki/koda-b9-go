package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/habibmrizki/koda-b9-go/internal/jendela"
	"github.com/habibmrizki/koda-b9-go/internal/model"
	persegipanjang "github.com/habibmrizki/koda-b9-go/internal/persegiPanjang"
	"github.com/habibmrizki/koda-b9-go/internal/slice"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==========================================")
		fmt.Println("        MENU PROGRAM INTERAKTIF           ")
		fmt.Println("==========================================")
		fmt.Println("1. Hitung Luas & Keliling Persegi Panjang")
		fmt.Println("2. Buat Pola Jendela Bingkai")
		fmt.Println("3. Sisip Elemen pada Slice Angka")
		fmt.Println("4. Tampilkan Data Diri")
		fmt.Println("0. Keluar")
		fmt.Println("==========================================")
		fmt.Print("Masukkan pilihan menu (0-4): ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error saat membaca input:", err)
			}
			break
		}

		pilihanStr := strings.TrimSpace(scanner.Text())
		pilihan, err := strconv.Atoi(pilihanStr)
		if err != nil {
			fmt.Println("Input tidak valid! Harap masukkan angka.")
			continue
		}

		if pilihan == 0 {
			fmt.Println("\nApakah anda yakin ingin Keluar? (1:Ya , 0: Tidak)")
			if !scanner.Scan() {
				break
			}

			konfirmasiStr := strings.TrimSpace(scanner.Text())
			konfirmasi, err := strconv.Atoi(konfirmasiStr)

			if err != nil || (konfirmasi != 0 && konfirmasi != 1) {
				fmt.Println("Pilihan konfirmasi tidak valid! kembali ke menu utam")
				continue
			}

			if konfirmasi == 1 {
				fmt.Println("\nTerimkasih! Program selesai")
				break
			} else {
				fmt.Println("Batal Keluar! Kembali ke menu utama")
				continue
			}
		}

		switch pilihan {
		case 1:
			fmt.Println("\n--- Hitung Persegi Panjang ---")
			fmt.Print("Masukkan panjang: ")
			scanner.Scan()
			panjang, err1 := strconv.ParseUint(strings.TrimSpace(scanner.Text()), 10, 8)

			fmt.Print("Masukkan lebar: ")
			scanner.Scan()
			lebar, err2 := strconv.ParseUint(strings.TrimSpace(scanner.Text()), 10, 8)

			if err1 != nil || err2 != nil {
				fmt.Println("Input panjang atau lebar tidak valid!")
				continue
			}

			luas, keliling := persegipanjang.LuasDanKelilingPersegiPanjang(uint8(panjang), uint8(lebar))
			fmt.Printf("Hasil -> Luas: %d, Keliling: %d\n", luas, keliling)

		case 2:
			fmt.Println("\n--- Pola Jendela Bingkai ---")
			fmt.Print("Masukkan ukuran jendela: ")
			scanner.Scan()
			ukuran, err := strconv.ParseUint(strings.TrimSpace(scanner.Text()), 10, 8)
			if err != nil {
				fmt.Println("Ukuran jendela harus berupa angka!")
				continue
			}

			fmt.Println("Hasil Pola:")
			err = jendela.JendelaBingkai(uint8(ukuran))
			if err != nil {
				fmt.Println("Error:", err.Error())
			}

		case 3:
			fmt.Println("\n--- Pensisipan Slice Angka ---")
			numbers := []int{50, 75, 66, 20, 32, 90}
			fmt.Println("Slice awal :", numbers)
			result := slice.SlicesNumbers(numbers)
			fmt.Println("Setelah disisipkan 88 di tengah:", result)

		case 4:
			fmt.Println("\n--- Data Diri ---")
			user := model.Person{
				Name:        "Habib Muhammad Rizki",
				Photo:       "habib.jpeg",
				Email:       "habib@gmail.com",
				Age:         24,
				PhoneNumber: "086786573645",
				IsMarried:   false,
				Education: []model.Educations{
					{
						Name:  "SMA 1 N Cilacap",
						Major: "IPS",
					},
				},
			}

			fmt.Println("Nama             :", user.Name)
			fmt.Println("Foto             :", user.Photo)
			fmt.Println("Email            :", user.Email)
			fmt.Println("Umur             :", user.Age)
			fmt.Println("Nomor Telepon    :", user.PhoneNumber)
			fmt.Println("Status Pernikahan:", user.IsMarried)
			for i, edu := range user.Education {
				fmt.Printf("Pendidikan %d     : %s - %s\n", i+1, edu.Name, edu.Major)
			}

		default:
			fmt.Println("Pilihan tidak tersedia. Silakan pilih 0-4.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error saat membaca input:", err)
	}
}

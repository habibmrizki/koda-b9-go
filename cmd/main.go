package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/habibmrizki/koda-b9-go/internal/filereader"
	"github.com/habibmrizki/koda-b9-go/internal/jendela"
	"github.com/habibmrizki/koda-b9-go/internal/messaging"
	"github.com/habibmrizki/koda-b9-go/internal/model"
	"github.com/habibmrizki/koda-b9-go/internal/payment"
	"github.com/habibmrizki/koda-b9-go/internal/persegipanjang"
	"github.com/habibmrizki/koda-b9-go/internal/person"
	"github.com/habibmrizki/koda-b9-go/internal/routines"
	"github.com/habibmrizki/koda-b9-go/internal/slice"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fiktif := &payment.Fiktif{}

	for {
		fmt.Println("\n==========================================")
		fmt.Println("        MENU PROGRAM INTERAKTIF           ")
		fmt.Println("==========================================")
		fmt.Println("1. Hitung Luas & Keliling Persegi Panjang")
		fmt.Println("2. Buat Pola Jendela Bingkai")
		fmt.Println("3. Sisip Elemen pada Slice Angka")
		fmt.Println("4. Tampilkan Data Diri")
		fmt.Println("5. Buka & Baca File")
		fmt.Println("6. Method Person")
		fmt.Println("7. Sistem Checkout Pembayaran (Interface)")
		fmt.Println("8. Jalankan Rutinitas Pagi (Goroutines)")
		fmt.Println("9. Kirim & Terima Pesan (Channel)")
		fmt.Println("0. Keluar")
		fmt.Println("==========================================")
		fmt.Print("Masukkan pilihan menu (0-7): ")

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
				fmt.Println("Terimakasih! program selesai")
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
			user := model.User{
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

		case 5:
			fmt.Println("\n--- Buka & Baca File ---")
			fmt.Print("Masukkan path/nama file (misal: go.mod): ")
			scanner.Scan()
			path := strings.TrimSpace(scanner.Text())
			if path == "" {
				fmt.Println("Path file tidak boleh kosong!")
				continue
			}

			filereader.ReadFile([]byte(path))

		case 6:
			fmt.Println("\n--- Person Method (Input User) ---")

			// Input data person dari user
			fmt.Print("Masukkan Nama   : ")
			scanner.Scan()
			nama := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukkan Alamat : ")
			scanner.Scan()
			alamat := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukkan No HP  : ")
			scanner.Scan()
			hp := strings.TrimSpace(scanner.Text())

			// Buat objek Person baru
			p := person.NewPerson(nama, alamat, hp)

			// Panggil method Print & Greet
			fmt.Println("\n>> Hasil Method Print():")
			fmt.Println(p.Print())

			fmt.Println("\n>> Hasil Method Greet():")
			fmt.Println(p.Greet())

			//  method SetName dari input user
			fmt.Print("\nMasukkan Nama Baru (untuk test SetName): ")
			scanner.Scan()
			namaBaru := strings.TrimSpace(scanner.Text())

			p.SetName(namaBaru)
			fmt.Println(">> Setelah SetName diubah:")
			fmt.Println(p.Greet())

		case 7:
			fmt.Println("\n--- Sistem Checkout Pembayaran  ---")
			fmt.Println("Pilih Metode Pembayaran:")
			fmt.Println("1. Transfer Bank")
			fmt.Println("2. Pembayaran Online / E-Wallet")
			fmt.Println("3. Pembayaran Fiktif")
			fmt.Print("Pilih metode (1-3): ")
			scanner.Scan()
			metodePilihan := strings.TrimSpace(scanner.Text())

			var method payment.PaymentMethod
			switch metodePilihan {
			case "1":
				fmt.Print("Masukkan Nama Bank (misal: BCA, Mandiri, BRI): ")
				scanner.Scan()
				namaBank := strings.TrimSpace(scanner.Text())
				if namaBank == "" {
					namaBank = "BCA"
				}
				method = &payment.Bank{Name: namaBank}

			case "2":
				fmt.Print("Masukkan Nama Online / E-Wallet (misal: GoPay, OVO, Dana): ")
				scanner.Scan()
				namaOnline := strings.TrimSpace(scanner.Text())
				if namaOnline == "" {
					namaOnline = "GoPay"
				}
				method = &payment.Online{Name: namaOnline}

			case "3":
				method = fiktif

			default:
				fmt.Println("Metode pembayaran tidak valid!")
				continue
			}

			fmt.Print("Masukkan daftar harga barang (pisahkan dengan koma, misal: 25000,50000): ")
			scanner.Scan()
			inputHarga := strings.TrimSpace(scanner.Text())
			daftarStr := strings.Split(inputHarga, ",")
			var amounts []int

			for _, s := range daftarStr {
				val, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					continue
				}
				amounts = append(amounts, val)
			}

			if len(amounts) == 0 {
				fmt.Println("Daftar harga tidak valid!")
				continue
			}

			// Panggil fungsi Checkout interface
			msg, err := payment.Checkout(method, amounts)
			if err != nil {
				fmt.Println("Error Pembayaran:", err)
			} else if msg != "" {
				fmt.Print(msg)
			} else {
				fmt.Println(">> Sukses! Pembayaran fiktif berhasil dicatat ke dalam slice.")
			}

			// total pembayaran fiktif
			fmt.Printf("\n[Laporan Sistem Fiktif] Total Tersimpan: Rp %d | Riwayat Transaksi: %v\n", fiktif.Total(), fiktif.Lists)

		case 8:
			fmt.Println("\n--- Memulai Rutinitas Pagi secara Bersamaan (Concurrent) ---")

			var wg sync.WaitGroup

			wg.Add(4)

			go routines.Mandi(&wg)
			go routines.BuatKopi(&wg)
			go routines.MenyiapkanSarapan(&wg)
			go routines.MerapikanKamar(&wg)

			wg.Wait()

			fmt.Println(">> Semua rutinitas pagi telah selesai!")

		case 9:
			fmt.Println("\n--- Fitur Perpesanan (Goroutine & Channel) ---")
			msgChannel := make(chan messaging.Message)

			go messaging.ReceivedMessage(msgChannel)

			messaging.SendMessage(msgChannel, "Habib", "Halo, selamat pagi!")
			messaging.SendMessage(msgChannel, "Budi", "Pagi juga Habib!")
			messaging.SendMessage(msgChannel, "Sistem", "Server akan dimatikan dalam 5 menit.")

			time.Sleep(500 * time.Millisecond)

			close(msgChannel)

		default:
			fmt.Println("Pilihan tidak tersedia. Silakan pilih 0-7.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error saat membaca input:", err)
	}
}

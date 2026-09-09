package routines

import (
	"fmt"
	"sync"
	"time"
)

func Mandi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Mandi] Mulai Mandi...")
	time.Sleep(2 * time.Second)
	fmt.Println("[Mandi] Selesai Mandi")
}

func BuatKopi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Buat Kopi] Mulai Menyeduh Kopi")
	time.Sleep(1 * time.Second)
	fmt.Println("[Buat Kopi] Selesai Buat kopi")
}
func MenyiapkanSarapan(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Sarapan] Mulai menyiapkan sarapan...")
	time.Sleep(3 * time.Second)
	fmt.Println("[Sarapan] Selesai menyiapkan sarapan.")
}

func MerapikanKamar(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Rapi Kamar] Mulai merapikan kamar...")
	time.Sleep(2 * time.Second)
	fmt.Println("[Rapi Kamar] Selesai merapikan kamar.")
}

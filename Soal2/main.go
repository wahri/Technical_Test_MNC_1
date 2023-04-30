package main

import (
	"fmt"
	"math"
)

func calculateChange(totalBelanja, jumlahUang float64) (map[float64]int, float64, bool) {

	pecahan := [...]float64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	kembalian := jumlahUang - totalBelanja

	// Return false jika kurang bayar
	if kembalian < 0 {
		return nil, 0, false
	}

	// Bulatkan kebawah
	kembalian = math.Floor(kembalian/100) * 100
	// Buat variabel ulang untuk ditampilkan di luar
	pembulatanKembalian := kembalian

	jumlahPecahan := make(map[float64]int)

	// Looping kembalian
	for _, p := range pecahan {
		if kembalian >= p {
			jumlah := int(math.Floor(kembalian / p))
			jumlahPecahan[p] = jumlah
			kembalian -= float64(jumlah) * p
		}
	}

	return jumlahPecahan, pembulatanKembalian, true
}

func main() {

	var totalBelanja, jumlahUang float64

	// Input total belanja dan jumlah uang
	fmt.Printf("Total belanja seorang customer: ")
	fmt.Scan(&totalBelanja)
	fmt.Printf("Pembeli membayar: ")
	fmt.Scan(&jumlahUang)

	pecahan, kembalian, ok := calculateChange(totalBelanja, jumlahUang)
	if ok {
		fmt.Printf("Kembalian yang harus diberikan kasir: Rp.%.0f\n", (totalBelanja-jumlahUang)*-1)
		fmt.Printf("Dibulatkan menjadi %.0f\n", kembalian)
		fmt.Println("Pecahan uang: ")

		// Print kembalian
		for rupiah, jumlah := range pecahan {
			if rupiah < 1000 {
				fmt.Printf("%d koin %.0f \n", jumlah, rupiah)
			} else {
				fmt.Printf("%d lembar %.0f \n", jumlah, rupiah)
			}
		}
	} else {
		fmt.Println("False, kurang bayar")
	}
}

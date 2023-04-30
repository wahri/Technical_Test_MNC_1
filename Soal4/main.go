package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var tglMasukstr, tglCutiStr string
	var jumlahCuti, cutiBersama int

	// Input jumlah cuti bersama pada tahun
	fmt.Print("Jumlah cuti bersama: ")
	fmt.Scan(&cutiBersama)

	// input tanggal join karyawan
	fmt.Print("Tanggal join karyawan (format: yyyy-mm-dd): ")
	fmt.Scan(&tglMasukstr)
	tglMasuk, err := time.Parse("2006-01-02", tglMasukstr)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	// input tanggal rencana cuti
	fmt.Print("Tanggal rencana cuti (format: yyyy-mm-dd): ")
	fmt.Scan(&tglCutiStr)
	tglCuti, err := time.Parse("2006-01-02", tglCutiStr)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	// input durasi cuti
	fmt.Print("Durasi cuti (hari): ")
	fmt.Scan(&jumlahCuti)

	validasiCuti, alasan := valdasiCuti(tglMasuk, cutiBersama, tglCuti, jumlahCuti)
	fmt.Println(validasiCuti)
	if !validasiCuti {
		fmt.Println(alasan)
	}
}

func valdasiCuti(tglMasuk time.Time, cutiBersama int, tglCuti time.Time, jumlahCuti int) (bool, string) {
	jmlCutiKantor := 14
	maksJumlahCutiPribadi := jmlCutiKantor - cutiBersama
	var jmlCutiPribadi float64

	durasiKerja := int(tglCuti.Sub(tglMasuk).Hours() / 24)
	if durasiKerja < 180 {
		return false, "Alasan: Karena belum 180 hari sejak tanggal join karyawan"
	} else if durasiKerja < 365 {
		akhirTahun := time.Date(tglCuti.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
		jmlDurasiKerjaSetelah180 := math.Floor(akhirTahun.Sub(tglMasuk.AddDate(0, 0, 180)).Hours() / 24)
		jmlCutiPribadi = math.Floor(jmlDurasiKerjaSetelah180 / 365.0 * float64(maksJumlahCutiPribadi))
	} else {
		jmlCutiPribadi = float64(maksJumlahCutiPribadi)
	}

	if jumlahCuti > int(jmlCutiPribadi) {
		return false, fmt.Sprintf("Alasan: Karena hanya boleh mengambil %d hari cuti", int(jmlCutiPribadi))
	} else if jumlahCuti > 3 {
		return false, "Alasan: Maksimal cuti berturut-turut adalah 3 hari"
	} else {
		return true, "True"
	}
}

package main

import "fmt"

func cetakBintang(jumlah int) {
	if jumlah > 0 {
		cetakBintang(jumlah - 1)
		fmt.Print("*")
	}
}

func polaBintang(baris int, saatIni int) {
	if saatIni > baris {
		return
	}
	cetakBintang(saatIni)
	fmt.Println()
	polaBintang(baris, saatIni+1)
}

func main() {
	var totalBaris int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&totalBaris)
	fmt.Println("Pola bintang:")
	polaBintang(totalBaris, 1)
}

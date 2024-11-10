package main

import "fmt"

func cetakPola(nilai, saatIni int) {
	fmt.Print(saatIni, " ")

	if saatIni == 1 {
		return
	}

	cetakPola(nilai, saatIni-1)

	fmt.Print(saatIni, " ")
}

func main() {
	var angka int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&angka)

	cetakPola(angka, angka)
	fmt.Println()
}

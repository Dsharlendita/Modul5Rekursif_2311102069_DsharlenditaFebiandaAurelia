package main

import "fmt"

func cetakBilanganGanjil(batas, angka int) {
	if angka > batas {
		return
	}

	fmt.Print(angka, " ")

	cetakBilanganGanjil(batas, angka+2)
}

func main() {
	var nilai int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&nilai)

	cetakBilanganGanjil(nilai, 1)
	fmt.Println()
}

package main

import "fmt"

func cetakFaktor(bilangan, pembagi int) {
	if pembagi > bilangan {
		return
	}
	if bilangan%pembagi == 0 {
		fmt.Printf("%d ", pembagi)
	}
	cetakFaktor(bilangan, pembagi+1)
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)
	fmt.Printf("Faktor dari %d adalah: ", angka)
	cetakFaktor(angka, 1)
	fmt.Println()
}

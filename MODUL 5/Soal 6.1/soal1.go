package main

import "fmt"

func hitungFibonacci(posisi int) int {
	if posisi <= 1 {
		return posisi
	}
	return hitungFibonacci(posisi-1) + hitungFibonacci(posisi-2)
}

func main() {
	var jumlahSuku int = 10
	fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", jumlahSuku)
	for indeks := 0; indeks < jumlahSuku; indeks++ {
		fmt.Printf("%d ", hitungFibonacci(indeks))
	}
	fmt.Println()
}

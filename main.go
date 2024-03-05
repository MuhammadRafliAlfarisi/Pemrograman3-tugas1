package main

import "fmt"

func main() {
	fmt.Println("Hello, World!!")
	var a, b int

	// Input angka pertama
	fmt.Println("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	// Input angka kedua
	fmt.Println("Masukkan angka kedua: ")
	fmt.Scanln(&b)

	// Penjumlahan
	fmt.Println("Penjumlahan: Hasil =", Penjumlahan(a, b))
	fmt.Println()
	fmt.Println("Pengurangan: Hasil =", Pengurangan(a, b))
	fmt.Println()
	fmt.Println("Perkalian: Hasil =", Perkalian(a, b))
	fmt.Println()
	fmt.Println("Pembagian: Hasil =", Pembagian(a, b))
	fmt.Println()
}

func Penjumlahan(a, b int) int {
	return a + b
}

func Pengurangan(a, b int) int {
	return a - b
}

func Perkalian(a, b int) int {
	return a * b
}

func Pembagian(a, b int) float64 {
	if b == 0 {
		fmt.Println("Pembagian dengan nol tidak diperbolehkan.")
		return 0
	}
	return float64(a) / float64(b)
}

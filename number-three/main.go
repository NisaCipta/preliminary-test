package main

import "fmt"

func printNumberParts(num int) {
	divisor := 1

	// Menemukan divisor terbesar yang dapat membagi angka
	for divisor <= num {
		divisor = divisor * 10
	}

	// Mengubah divisor menjadi 10 kali lebih kecil untuk mulai dari posisi yang benar
	divisor = divisor / 10
	// 1000.000
	// 100.000
	// 10.000
	// 1.000

	// Perulangan untuk membagi angka dan mencetak hasilnya
	for divisor > 0 {
		result := num / divisor
		fmt.Println(result * divisor)
		num %= divisor
		divisor /= 10
	}
}

func main() {
	intInput := 1345679
	printNumberParts(intInput)
}

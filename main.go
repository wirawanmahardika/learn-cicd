package main

import "fmt"

// Tambah adalah fungsi sederhana yang menerima dua integer dan mengembalikan hasil penjumlahannya.
func Tambah(a int, b int) int {
	return a + b
}

func main() {
	name := "wirawan"
	hasil := Tambah(5, 5)
	fmt.Printf("Aplikasi berjalan. Hasil 5 + 5 = %d\n", hasil)
	fmt.Printf("Hello %s", name)
}

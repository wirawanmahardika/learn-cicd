package main

import "testing"

// TestTambah menguji apakah fungsi Tambah menghasilkan nilai yang benar.
func TestTambah(t *testing.T) {
	ekspektasi := 10
	hasil := Tambah(5, 5)

	if hasil != ekspektasi {
		t.Errorf("Fungsi Tambah() gagal: ekspektasi %d, tetapi mendapatkan %d", ekspektasi, hasil)
	}
}

// TestTambahSalah sengaja dibuat untuk skenario pengujian dengan input lain.
func TestTambahLain(t *testing.T) {
	ekspektasi := 0
	hasil := Tambah(5, -5)

	if hasil != ekspektasi {
		t.Errorf("Fungsi Tambah() gagal: ekspektasi %d, tetapi mendapatkan %d", ekspektasi, hasil)
	}
}

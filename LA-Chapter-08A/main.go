package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x // p adalah pointer ke x

	fmt.Println("Nilai x:", x)
	fmt.Println("Alamat memori x:", &x)
	fmt.Println("Nilai p (alamat memori x):", p)
	fmt.Println("Nilai yang ditunjuk p:", *p) // dereference untuk mendapatkan nilai x

	// Mendeklarasikan pointer baru untuk integer
	y := new(int)

	// Menggunakan pointer untuk mengubah nilai
	*y = 42

	fmt.Println("Nilai yang ditunjuk oleh pointer:", *y)
	fmt.Println("Alamat memori yang ditunjuk oleh pointer:", y)

}

package main

import "fmt"

// Deklarasi struct
type Person struct {
	Name string
	Age  int
}

type Rectangle struct {
	width  float64
	height float64
}

// Method untuk menghitung luas persegi panjang
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method untuk mengubah nilai lebar persegi panjang
func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}

func main() {
	// Inisialisasi struct dengan nilai-nilai awal
	person1 := Person{Name: "Andi", Age: 30}

	// Mengakses dan mencetak nilai bidang struct
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)

	// Inisialisasi objek struct dan pointer
	var personPtr *Person
	person := Person{Name: "Siska", Age: 26}
	personPtr = &person // Inisialisasi pointer ke objek struct
	// Akses dan modifikasi nilai objek struct melalui pointer
	fmt.Println("Name (before):", personPtr.Name)
	fmt.Println("Age (before):", personPtr.Age)

	fmt.Println("Name (before):", person.Name)
	fmt.Println("Age (before):", person.Age)

	personPtr.Name = "Deni"
	personPtr.Age = 25

	// Cetak nilai objek struct melalui pointer setelah dimodifikasi
	fmt.Println("Name (after):", personPtr.Name)
	fmt.Println("Age (after):", personPtr.Age)

	fmt.Println("Name (after):", person.Name)
	fmt.Println("Age (after):", person.Age)

	// inisialisasi struct rectangle
	rectangle := Rectangle{width: 5, height: 3}

	// Memanggil method pada struct
	fmt.Println("Luas Persegi Panjang:", rectangle.Area())

	// call value rectangle
	fmt.Println("Lebar sebelum diubah:", rectangle.width)

	// Memanggil method untuk mengubah nilai lebar
	rectangle.SetWidth(10)

	fmt.Println("Lebar setelah diubah:", rectangle.width)
}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person   // Embedding struct Person
	Position string
}

func main() {
	emp := Employee{
		Person:   Person{Name: "Wendi", Age: 17},
		Position: "Developer",
	}

	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Position:", emp.Position)

	// Deklarasi dan inisialisasi anonymous struct
	person := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// Mendeklarasikan slice dari struct
	var people []Person

	// Menambahkan instansi struct ke dalam slice
	people = append(people, Person{Name: "Siska", Age: 20})
	people = append(people, Person{Name: "Danur", Age: 25})

	// Iterasi melalui slice dan mencetak nilai struct
	for _, p := range people {
		fmt.Println("Name:", p.Name)
		fmt.Println("Age:", p.Age)
	}

}

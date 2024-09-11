package main

import "fmt"

// Definisikan interface Animal
type Animal interface {
	Speak() string
}

// Definisikan struct Dog
type Dog struct{}

// Implementasi metode Speak untuk Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Definisikan struct Cat
type Cat struct{}

// Implementasi metode Speak untuk Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Fungsi yang menerima parameter bertipe Animal
func makeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	// inisialiasasi struct
	dog := Dog{}
	cat := Cat{}

	// implementation interface
	makeAnimalSpeak(dog)
	makeAnimalSpeak(cat)
}

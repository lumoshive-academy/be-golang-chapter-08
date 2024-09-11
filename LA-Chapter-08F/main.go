package main

import (
	"fmt"
	"math"
)

// Interface untuk bentuk dua dimensi
type Shape interface {
	Area() float64
}

// Interface untuk bentuk tiga dimensi
type Volume interface {
	Volume() float64
}

// Interface untuk bentuk tiga dimensi yang juga memiliki area
type Solid interface {
	Shape
	Volume
}

// Definisikan struct Cylinder
type Cylinder struct {
	Radius, Height float64
}

// Implementasi metode Area untuk Cylinder
func (c Cylinder) Area() float64 {
	return 2*math.Pi*c.Radius*c.Height + 2*math.Pi*c.Radius*c.Radius
}

// Implementasi metode Volume untuk Cylinder
func (c Cylinder) Volume() float64 {
	return math.Pi * c.Radius * c.Radius * c.Height
}

func printSolidDetails(s Solid) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Volume: %.2f\n", s.Volume())
}

// Fungsi yang menerima parameter dengan tipe interface kosong
func PrintValue(val interface{}) {
	fmt.Println("Value:", val)
}

func main() {
	c := Cylinder{Radius: 3, Height: 5}
	printSolidDetails(c)

	PrintValue(42)             // Integer
	PrintValue("Hello")        // String
	PrintValue(3.14)           // Float
	PrintValue([]int{1, 2, 3}) // Slice

}

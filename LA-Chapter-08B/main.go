package main

import "fmt"

// pointer in parameter
func changeValue(val *int) {
	*val = 30
}

func main() {
	// change value pointer
	var x int = 10
	var p *int = &x

	fmt.Println("Nilai x sebelum diubah:", x)
	*p = 20 // mengubah nilai x melalui pointer p
	fmt.Println("Nilai x setelah diubah:", x)

	// call funtions with parameter pointer
	var y int = 10
	fmt.Println("Nilai x sebelum fungsi changeValue:", y)

	changeValue(&y)
	fmt.Println("Nilai x setelah fungsi changeValue:", y)

}

package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func _main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// x = x + 1
	// y = "bye"

	fmt.Println(x)
	fmt.Println(y)

	const x = 10
	var z float64 = x
	var d byte = x
	fmt.Println(y, z, d)

	const typedX int = 10
}

func _main2() {
	x := 10
	x = 20
	fmt.Println(x)
	x = 30
}

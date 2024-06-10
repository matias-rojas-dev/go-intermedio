package main

import "fmt"

func sum(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

func names(values ...string) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(sum(1, 2, 3))
	names("Matías", "Ignaci", "ROjas")
}

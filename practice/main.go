package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 21
	numbers[0] = 25
	numbers[0] = 41
	numbers[0] = 51
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}

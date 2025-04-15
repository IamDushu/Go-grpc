package main

import (
	"fmt"
	"sort"
)

func sort_main() {
	numbers := []int{10, 50, 3, 2, 48}
	names := []string{"Dhana", "Dushu", "Chaitu"}

	sort.Ints(numbers)
	sort.Strings(names)

	fmt.Println(numbers)
	fmt.Println(names)
}

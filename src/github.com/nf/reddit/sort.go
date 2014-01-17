package main

import (
	"fmt"
	"sort"
)

func main() {
	av := []int{1, 4, 5, 6, 2}
	av = append(av, 9, 78, 9, 7, 78, 8, 89, 9)

	av2 := [5]int{1, 2, 3, 2, 1}

	fmt.Println(av)
	sort.Ints(av)
	fmt.Println(av)
	fmt.Println(av2)

	for i := 0; i < len(av); i++ {
		fmt.Print(av[i])
		fmt.Println()
	}
	return
}

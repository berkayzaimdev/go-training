package main

import "fmt"

// comparable, hazır interface olarak karşımıza çıkar. java'daki comparable mantığı tamamen
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	sb := []bool{true, true, true} 
	fmt.Println(Index(sb, false))
}

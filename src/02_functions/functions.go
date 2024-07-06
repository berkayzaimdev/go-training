package main

import "fmt"

func add(x int, y int) int { // diğer dillerden farklı olarak önce değişken ismi sonra type
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

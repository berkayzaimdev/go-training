package main

import "fmt"

func add(x int, y int) int { // diğer dillerden farklı olarak önce değişken ismi sonra type
	return x + y
}

func split(sum int) (x, y int) { // (int, int) olarak da tanımlayıp "return x, y" diyebilirdik
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(split(17))
}
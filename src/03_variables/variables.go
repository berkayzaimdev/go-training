package main

import "fmt"

var c, python, java bool // fonksiyon dışında "var" keywordü ile liste olarak tanımlanabilir

func main() {
	var i, j int = 1, 2 //
	k := 3
	c, python, java := true, false, "no!" // := ile otomatik type belirleme (C# var gibi)

	fmt.Println(i, j, k, c, python, java)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // type belirlediğimiz değişkenler için parantez yoluyla conversion yaparız
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

/*

   - bool, string

   - int, int8, int16, int32, int64
   - uint, '', '' ..

   - byte => uint8
   - rune => int32

   - float32, float64
   - complex64, complex128 (karmaşık sayılar için)

*/

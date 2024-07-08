package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) AbsMethod() float64 { // go'da class yapısı yoktur. fakat biz "v Vertex'e ait bir metot olan AbsMethod()"u tanımlayabiliriz. bu şekilde değişkene özgü metot setleri oluşturup class benzeri bir yapı kurabiliriz.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 { // "v Vertex'i parametre kabul eden AbsFunc()". metottan farkı bu şekilde
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main() {
	v := Vertex{3, 4}
	fmt.Println(v.AbsMethod())
	fmt.Println(AbsFunc(v))
}

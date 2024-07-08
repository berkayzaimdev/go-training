package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { // metotları pointer olarak kullanmanın sağlayacağı en büyük fayda, metodun etkilediği değişkende kalıcı değişimler yapabilmektir
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	// ScaleFunc(v, 10) // derleme hatası verir. Go'da fonksiyon ile metodun temel farkı budur. Metot kullanımı bize '&' olmasa dahi ScaleFunc((&)v, 10) kullanımını otomatik olarak sağlayacaktı. Fonksiyonda ise dönüşüm için & kullanmamız gerekir.

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)


	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	fmt.Println(p.Abs()) // dikkat edersek Abs metodu bir 'değişken' kabul ediyor ama verdiğimiz parametre bir adres. Metot kullanımı sayesinde otomatik dönüşüm yapıyor ve hata olmuyor
	fmt.Println(AbsFunc(*p))
}

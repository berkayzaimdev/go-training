package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  
	v2 = Vertex{X: 1}  // C#'daki gibi parametre seçme mantığı burada da var. (X,Y) = (1,0)
	v3 = Vertex{}      // propertylere değer ataması yapılmazsa default değeri alırlar (int için 0)
	p  = &Vertex{1, 2} 
)

func main() {
	fmt.Println(v1, p, v2, v3) // pointer'ı yazdırırken bellek adresi yerine syntaxtaki & ifadesini yazdırır

	v4 := Vertex{1, 2}
	p2 := &v4
	p2.X = 1e9 // pointer'ın işaret ettiği struct ı değiştirdik
	fmt.Println(v4)
}

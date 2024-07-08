package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

// Bu metot, f değişkeni ile F tipinin I interface'ini implement ettiğini ifade eder
func (f F) M() {
	fmt.Println(f)
}

func (f F) M2() {
	fmt.Println(f)
}


func main() {
	/*
	var i I
	// panic döner. çünkü NULL interface'in ne herhangi bir değeri, ne de concrete tipi vardır. runtime hatasıdır
	describe(i)
	i.M()
	*/

	var i I // diğer OOP dillerden farklı olarak interface tanımlamasında ek bir tip belirtmeden nesne oluşturabiliyoruz
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t // eğer interface'in concrete değeri NULL ise, metot çağrısı sonucu <nil> döner. 
	describe(i)
	i.M()

	var ie interface{} // boş bir interface, özel bir duruma teşkil eder. Bu interface, 0 metodu implemente eder. 
	describeForEmpty(ie)

	ie = 42 // boş interface, herhangi bir tipe eşitlenebilir. çünkü her bir tip, en azından 0 metodu uygular. (primitive tipler özellikle)
	describeForEmpty(ie)

	ie = "hello"
	describeForEmpty(ie)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i) // interface (değeri, tipi) tuple şeklinde düşünebiliriz
}

func describeForEmpty(ie interface{}) {
	fmt.Printf("(%v, %T)\n", ie, ie)
}
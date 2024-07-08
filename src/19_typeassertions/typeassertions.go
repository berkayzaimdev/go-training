package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var i interface{} = "hello"
	// kısaca t := i.(T); T, interface'in değerini tanımlayan concrete type ismidir. interface değerini döndürür. verdiğimiz type ile eşleşmiyorsa ve ok kullanıldıysa, default değer döndürür (0,false,"")
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*
	f = i.(float64) // panic fırlatır. "ok" kullanımı bu noktada hata yönetimi için önemlidir.
	fmt.Println(f)
	*/

	do(21)
	do("hello")
	do(true)
}
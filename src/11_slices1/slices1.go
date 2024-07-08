package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // python'daki dizi bölümleme mantığı ile tamamen aynı
	fmt.Println(s)

	fmt.Println();
	fmt.Println();
	fmt.Println();

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // bu şekilde slice'ta indisler üzerinden değer değiştirebiliriz.
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println();
	fmt.Println();
	fmt.Println();



	q := []int{2, 3, 5, 7, 11, 13} // dizi tanımlamasını boyut belirlemeden yapmak, slice tanımlamasına karşılık gelir

	q = q[1:4] // 1'den başla 4'e kadar al
	fmt.Println(q)

	q = q[:2]
	fmt.Println(q)

	q = q[1:]
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct { // struct initialize
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)
}

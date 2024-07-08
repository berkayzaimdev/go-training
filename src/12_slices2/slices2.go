package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// slice işleminden sonra 0 uzunluğa sahip yeni bir slice
	s = s[:0]
	printSlice(s)

	// slice'ın uzunluğunu arttırdık
	s = s[:4]
	printSlice(s)

	// slice'ta var olan elemanın sayısını azaltarak, kapasitesini azaltmış olduk
	s = s[2:]
	printSlice(s)

	var s2 []int // boş bir şekilde slice tanımladığımızda, bu slice NULL'ı ifade edecek şekilde "nil" durumundadır
	fmt.Println(s2, len(s2), cap(s2)) // dizi boş şekilde yazdırılır. ayrıca hem uzunluğu hem de kapasitesi 0'a eşittir.
	if s2 == nil {
		fmt.Println("I'm a NULL slice")
	}

	s2 = append(s2, 0) // hazır "append" fonksiyonu ile slice'a eleman ekleyip "nil" durumundan çıkarabiliriz
	printSlice(s2)


	s2 = append(s2, 1)
	printSlice(s2)

	s2 = append(s2, 2, 3, 4) // burada len=5, cap=6 olur. Bunun sebebi işletim sistemi farklılığı.
	printSlice(s2)

	printSlice(s2[0:6]) //  bir slice'ın kapasitesi miktarınca eleman seçebiliriz, fakat default değeri elde ederiz.
	// printSlice(s2[0:7]) //  kapasiteyi aşarsak panic fırlatır
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

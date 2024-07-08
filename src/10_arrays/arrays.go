package main

import "fmt"

func main() {
	var a [2]string // type isimleri yine en sağda yer alıyor, oluşturma aşamasında ise dizi boyutu sola yazılıyor.
	a[0] = "Hello" // dizi elemanı ataması bildiğimiz gibi, bir farklılık yok
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a) // elemanları köşeli parantezler arasında boşluklu bir şekilde yazdırır

	primes := [6]int{2, 3, 5, 7, 11, 13} // new'lemek yerine bu şekilde initialize edebiliriz
	fmt.Println(primes) 
}
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // init; condition; post
		sum += i
	}

	for ; sum < 1000; { // init ve post isteğe bağlıdır
		sum += sum
	}

	for sum < 1000 { // sadece condition yazarsak while döngüsü gibi kullanabiliriz
		sum += sum
	}

	/*
	for { // sonsuz döngü

	}
	*/

	fmt.Println(sum)
}

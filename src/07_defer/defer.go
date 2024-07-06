package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ { 
		defer fmt.Println(i) // defer; uygulandığı fonksiyonu, defer kullanmadan çağrılana kadar bekletir
	}

	fmt.Println("done") // defer, stack mantığıyla çalışır. örnek uygulamada 9,8,...,0 sırası ile yazar
}

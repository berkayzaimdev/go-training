package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // c kanalına, toplam değerini gönder
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 2) // kanal oluştururken limit belirleyebiliriz
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	
	x, y := <-c, <-c // c kanalından değer çek (queue mantığı var, stack değil)
	fmt.Println(x, y, x+y)
}

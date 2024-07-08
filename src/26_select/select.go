package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // select'teki anahtar kelime "bekleme"
		// select, hazır olan case'teki işlemleri derhal gerçekleştirir. birden fazla işlem hazırsa seçimi rastgele yapae
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {	
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			time.Sleep(1*time.Second)
		}
		time.Sleep(3*time.Second)
		quit <- 0	
	}()
	fibonacci(c, quit)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: // tespit edilen kanal olayı yoksa default çalışır. Bu bağlamda sürekli çalışacağı için Sleep kullanımı önemlidir. Waitpid, örnek olarak düşünülebilir
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

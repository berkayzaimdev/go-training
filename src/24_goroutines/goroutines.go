package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*

C'deki thread olayına benzer bir olay Go'da karşımıza çıkıyor
go işaretli çağırdığımız fonksiyon, main thread ile aynı anda ilerleyecek şekilde ayarlanır. Sonuç olarak "hello" ve "world", her adımda aynı anda yazdırılır
stringlerin yazılma sırası ise "race condition" a bağlıdır

*/
package main

import (
	"fmt"

	"rsc.io/quote"
)

func Quote() {
    fmt.Println(quote.Go())
}

/*

bir paket ile çalışmak için 
   - importa ilgili paketler eklenir
   - go mod init -p
   - go mod tidy

main'i tekrar etme hatasına karşı en iyi çözüm çalışacak her go dosyası için ayrı klasör açmak

*/
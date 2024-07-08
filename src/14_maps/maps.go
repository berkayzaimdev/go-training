package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m2 = map[string]Vertex{ // var ile map initializing
	"Bell Labs": Vertex{ 
		40.68433, -74.39967,
	},
	"Google": { // Vertex diye tip belirtmemize gerek yok
		37.42202, -122.08408,
	},
}


func main() {
	m = make(map[string]Vertex) // indisi string, değeri struct olan bir map oluşturduk. Dictionary mantığı ile düşünebiliriz
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])


	fmt.Println(m2)


	m3 := make(map[string]int)

	m3["Answer"] = 42 // eleman ekleme
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48 // elemanı değiştirme, (duplicate yok)
	fmt.Println("The value:", m3["Answer"])

	delete(m3, "Answer") // eleman silme
	fmt.Println("The value:", m3["Answer"])

	v, ok := m3["Answer"] // eleman(key), map'te bulunuyorsa true
	fmt.Println("The value:", v, "Present?", ok)
}

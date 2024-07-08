package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         
	fmt.Println(*p) 
	*p = 21         
	fmt.Println(i) 

	p = &j        
	*p = *p / 37 
	fmt.Println(j)
}

/*

C'nin pointer olayından çok bir farkı yok. Göze çarpan bir fark "pointer aritmetiği"nin olmayışı

*/

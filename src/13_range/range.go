package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { // range ile baştan sonra diziyi gezebiliyoruz
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()


	pow2 := make([]int, 10)
	for i := range pow2 { // indis gezme
		fmt.Println("i = ", i)
		pow2[i] = 1 << uint(i) // == 2**i; byte kaydırma işlemi yaptık.
	}

	
	for _, value := range pow2 { // değer gezme (underscore mantığı C# ile aynı)
		fmt.Printf("%d\n", value)
	}
}

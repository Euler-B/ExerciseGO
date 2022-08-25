package main 

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n - 1)
}

func main () {
	var n uint

	fmt.Println("Introduzca un numero:")
	fmt.Scanf("%d", &n)
	
	if (n < 0) {
		fmt.Printf("La propiedad del factorial no existe para los numeros negativos")
	} else {
		fmt.Println("El factorial del numero es:")
		fmt.Println(factorial(n))
	}

}
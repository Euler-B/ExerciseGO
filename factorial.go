package main 

import "fmt"

func main () {
	var i int
	var n int
	var f = 1

	fmt.Println("Introduzca un numero:")
	fmt.Scanf("%d", &n)
	
	for i = n; i > 0; i-- {
		f = f * i
	}

	fmt.Printf("El factorial del numero %d es: %d \n", n, f)

}



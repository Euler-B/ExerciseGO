package main 

import "fmt"

func main () {
	var nombre string
	var numero int

	fmt.Println("Saludos...")

	fmt.Println("Introduzca su nombre:")
	fmt.Scanf("%s", &nombre)

	fmt.Println("Introduzca un numero entre uno y dos:")
	fmt.Scanf("%d", &numero)

	switch {
	case (numero == 1):
		fmt.Printf("Hola %s desde la ciudad de Merida \n", nombre)
	case (numero == 2):
		fmt.Printf("Hola %s desde el sector Belen \n", nombre)
	case (numero >= 3 || numero <= 0):
		fmt.Printf("Hola %s, igual le saludamos aunque introdujo una opcion incorrecta \n", nombre)
	}
}
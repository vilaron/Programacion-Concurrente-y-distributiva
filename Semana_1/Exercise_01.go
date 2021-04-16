package main

import "fmt"

//main function
func main() {
	//variable
	var largo float64 = 12.58
	//short way
	ancho := 19

	fmt.Println("curso de Programación concurrente !!!")

	//conversion cast
	fmt.Println("El area del rectangulo es: ", largo*float64(ancho))
	fmt.Println("Es el largo del rectangulo mayor a a ancho?", largo > float64(ancho))
}

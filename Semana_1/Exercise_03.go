package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	menu := `
	           Bienvenido
			   [1] Pizza
			   [2] Empanadas
			   Selecciona la comida de tu preferencia
	`
	fmt.Println(menu)

	//Lectura de datos por consola
	bufferIn := bufio.NewReader(os.Stdin)
	ingreso, _ := bufferIn.ReadString('\n')
	opcionstr := strings.TrimRight(ingreso, "\r\n")

	opcion, _ := strconv.Atoi(opcionstr)

	switch opcion {
	case 1:
		fmt.Println("Elegiste Pizza")
	case 2:
		fmt.Println("Elegiste Empanada")
	default:
		fmt.Println("Opcion invalida")
	}
	if opcion == 1 || opcion == 2 {
		fmt.Println("Ingresa la cantidad")
		var cantidad int
		fmt.Scanf("%d", &cantidad)

		fmt.Println("El valor a pagar es : ", cantidad*5)
	}
}

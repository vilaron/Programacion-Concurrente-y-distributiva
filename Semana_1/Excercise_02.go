import "fmt"

func main() {
	//arreglo
	arreglo := [7]int{0, 5, 2, 6, 6, 10, 8, 5}
	//estructura repetitiva del FOR

	//1.- forma accesso a colecciones
	for i, v := range arreglo {
		fmt.Printf("El valor de v es %d en la posicion %d\n", v, i)
	}

	for _, v := range arreglo {
		fmt.Printf("El valor de v es %d\n", v)
	}
	//2.- While
	i := 0
	for {
		if i == 5 {
			break
		}
		Printf("Numero de intentos:", i)
		i++
	}
	//3 .- Forma tradicional
	j := 0
	for i > 0 {
		if j == 10 {
			break
		}
		j++
	}

	fmt.Println("Numero de intentos de j", j)

}
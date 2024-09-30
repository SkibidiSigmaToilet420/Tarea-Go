package main

import (
	"fmt"
)

func main(){

	var num1, num2 int64 = 0, 0
	var resultado float64 = 0
	var resp int64 = 1

	fmt.Println("Calculadora")
	
	fmt.Println("Ingrese el primer numero:")
	fmt.Scan(&num1)

	fmt.Println("Ingrese el segundo numero:")
	fmt.Scan(&num2)

	fmt.Println("Escoja su operacion 1. suma 2. resta 3. Multiplicación 4. División")
	fmt.Scan(&resp)

	if resp == 1{
		resultado=float64(num1+num2)
	} else if resp == 2 {
		resultado=float64(num1-num2)
	} else if resp == 3 {
		resultado=float64(num1*num2)
	} else {
		if num2==0{
			fmt.Println("No se puede dividir entre cero")
			
		}else{
			resultado=float64(num1/num2)
		}

		
	}

	fmt.Print("El resultado es: ", resultado)
	

}


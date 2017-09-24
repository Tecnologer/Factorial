package main

import "fmt"

func main(){
	var number int
	fmt.Print("Escribe un numero entero: ")
	_, err := fmt.Scan(&number)
	
	for err != nil{		
		fmt.Print("Se esperaba un numero entero. Intenta de nuevo: ")
		_, err = fmt.Scan(&number)
	}

	fmt.Println(factorial(number))
}

func factorial(number int) uint64{
	if(number == 0){
		return 1
	}
	
	return uint64(number) * factorial(number -1)
}
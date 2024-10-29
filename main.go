package main

import "fmt"

func ScanNumber () int {
	var i int
	Printf("escreva um numero inteiro:",0)
	fmt.Scan(&i)
	return i
}

func Printf(str string, num int) int {
	if num != 0 {
		fmt.Printf("%s %d\n", str, num)
	} else {
		fmt.Println(str)
	}
	return num
}


func soma(a int, b int) int {
	if a != 0 && b != 0 {
		return a + b
	}else{
		Printf("nenhum parametro foi passado",0)
		return 0
	}
	/*
	poderia adicionar apenar o return a+b mas fiz uma palhacada :)
	*/
}

func main(){

	/*
	Printf("hello amigos da 42 vou dominar esse porra de go aqui !!!",0)
	Printf("hello Word",0) // ex00
	*/
	
	/* exercicio 01
	var result = ScanNumber()
	Printf("o resultado do numero solicitado foi:", result)
	*/
	
	/*
	//Faça um Programa que peça dois números e imprima a soma.
	var result = soma(4,5)
	Printf("o resultado da soma dos inteiro é ", result)
	*/
	
	


}
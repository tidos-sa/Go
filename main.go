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

func media(a , b, c, d int) float64{
	media := a+b+c+d
	return float64(media) / 4.0
}

func QuestInt() float64{
	var a,b,c,d int
	fmt.Println("digite a primeira nota")
	fmt.Scan(&a)
	fmt.Println("digite a segunda nota")
	fmt.Scan(&b)
	fmt.Println("digite a terceira nota")
	fmt.Scan(&c)
	fmt.Println("digite a quarta nota")
	fmt.Scan(&d)

	return media(a,b,c,d)
	
}

func conversorFarehin(num float64) float64{
	var C float64
	C = 5*((num-32)/9)
	return (C)
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

	/*
	//Faça um Programa que peça as 4 notas bimestrais e mostre a média.
	result := media(5,5,5,5)
	fmt.Println("a media das notas é: " ,result)
	*/

	/*
	media da escola
	result := QuestInt()
	fmt.Println("a media dos numeros é: ", result)
	*/
	

	var result float64
	result = conversorFarehin(100)
	fmt.Println("temperatura em farenhein seila oque: ",result)
}
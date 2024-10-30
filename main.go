package main

import "fmt"
// estrutura sequencial -- lista de logica

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
// aquicimento acabou hora dos boss :)

/*
Faça um Programa que pergunte quanto você ganha por hora e o número de horas trabalhadas no mês. Calcule e mostre o total do seu salário no referido mês, sabendo-se que são descontados 11% para o Imposto de Renda, 8% para o INSS e 5% para o sindicato, faça um programa que nos dê:
salário bruto.
quanto pagou ao INSS.
quanto pagou ao sindicato.
o salário líquido.
calcule os descontos e o salário líquido, conforme a tabela abaixo:
+ Salário Bruto : R$
- IR (11%) : R$
- INSS (8%) : R$
- Sindicato ( 5%) : R$
= Salário Liquido : R$
*/

func Anlsalario(){
	var dias int
	var valorhPorHoras int
	var horas int

	fmt.Println("O quanto vocë recebe por *HORA* ?")
	fmt.Scan(&valorhPorHoras)
	fmt.Println("Quantos dias voce trabalha por mes ?")
	// calculo convertendo dias trabalhados em horas considerando que se trabalha 9h por dia.
	fmt.Scan(&dias) // 23 x (9 horas) por dia = totalhoras trabalhadas 
	fmt.Println("quantas horas por dia vc trabalha?")
	fmt.Scan(&horas)

	var salarioTotal  = valorhPorHoras*(horas*dias)
	var IR  = (salarioTotal * 11) / 100
	var INSS = (salarioTotal * 8) / 100
	var Sindicato = (salarioTotal * 5) / 100
	var liquido	= salarioTotal - (IR+INSS+Sindicato)


	fmt.Println("o salario total: ", salarioTotal,"\nIR: -",IR,"\nINSS : -",INSS,"\nSINDICADO: -",Sindicato,"\nSÁLARIO LIQUIDO :",liquido)


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
	
	/*
	var result float64
	result = conversorFarehin(100)
	fmt.Println("temperatura em farenhein seila oque: ",result)
	*/
	Anlsalario()
}
package main

import "fmt"

func main()  {
	//Exepciones de tipado
	const pi float64 = 3.14;
	const pi2 = 3.1415;

	fmt.Println("variable de pi declara con float: ", pi , " Variable no declarada: ", pi2 );
	
	//Declaracion de variables
	base := 14; // inicializar variable sin que se declare
	var altura  int = 14; //declara e inicializar
	var area int; //declaracion de variable sin incilizar

	fmt.Println("base: ", base , " altura: ", altura, " area: ", area);
	
	//Zero Value
	var numero int;
	var decimal float64;
	var caracter string;
	var binario bool;

	fmt.Println("Zero Numero: ", numero , " Zero decimal: ", decimal , " zero String: ", caracter, " Zero Boolean: ", binario);
}
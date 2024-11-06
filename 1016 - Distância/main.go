package main

import "fmt"

func main(){
	var kmpm int = 2
	var km int
	fmt.Scanf("%d",&km)
	distancia := kmpm * km
	fmt.Printf("%d minutos\n", distancia)
}
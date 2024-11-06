package main

import (
	"fmt"
	"math"
)

func main()  {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a,&b,&c)
	triangulo := (a * c) / 2
	circulo := 3.14159 * math.Pow(c, 2)
	trapezio := (a+b) * (c/2)
	quadrado := b * b
	retangulo := a*b
	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3F\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n", triangulo, circulo, trapezio, quadrado,retangulo)
}
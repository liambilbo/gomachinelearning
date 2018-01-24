package main

import (
	"github.com/gonum/floats"
	"fmt"
)

func main() {
	vectorA:=[]float64{11.0,5.2,-1.3}
	vectorB:=[]float64{-7.2,4.2,5.1}

	//Compute the dot product of to vectors
	dotProduct:=floats.Dot(vectorA,vectorB)

	fmt.Printf("The product of the dot product is %0.2f\n",dotProduct)

	//Scale each element of Vector A by 1.5
	floats.Scale(1.5,vectorA)

	fmt.Println("Scaling vectorA by 1.5 give %v \n",vectorA)

	//Compute the norm length
	norm:=floats.Norm(vectorB,2)
	fmt.Printf("The norm/length of vectoB by 2 is %0.2f",norm)



}

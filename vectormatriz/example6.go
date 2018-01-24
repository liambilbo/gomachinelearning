package main

import (
	"gonum.org/v1/gonum/mat"
	"fmt"
	"log"
)

func main(){
	a:=mat.NewDense(3,3,[]float64{1,2,3,0,4,5,0,0,6})

	//Compute the transpose of the Matrix
	ft:=mat.Formatted(a.T(),mat.Prefix("        "))
	fmt.Printf("A ^ T:  %v\n\n",ft)

	//Compute and output de determinant of a
	fmt.Printf("Det(A):  %.2f\n\n",mat.Det(a))


	//Compute Inversa

	d:=mat.NewDense(0,0,nil)

	if err:=d.Inverse(a); err!=nil{
		log.Fatal(err)
	}
	df:=mat.Formatted(d,mat.Prefix("         "))

	fmt.Printf("A ^ -1 : %0.2v \n\n",df )

}

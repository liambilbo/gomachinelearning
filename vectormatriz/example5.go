package main

import (
	"gonum.org/v1/gonum/mat"
	"fmt"
	"math"
)

func main(){
	a:=mat.NewDense(3,3,[]float64{1,2,3,0,4,5,0,0,6})
	b:=mat.NewDense(3,3,[]float64{8,9,10,1,4,2,9,0,2})

	c:=mat.NewDense(3,2,[]float64{3,2,1,4,0,8})

	d:=mat.NewDense(0,0,nil)

	//Add a and b
	d.Add(a,b)

	fd:=mat.Formatted(d,mat.Prefix("        "))

	fmt.Printf("A + B = %0.4v \n\n",fd)

	//Multipliy A an C
	d=mat.NewDense(0,0,nil)

	d.Mul(a,c)

	fd=mat.Formatted(d,mat.Prefix("        "))

	fmt.Printf("A * C = %0.4v\n\n",fd)

	//Raising a Matrix to a Power

	d=mat.NewDense(0,0,nil)

	d.Pow(a,5)

	fd=mat.Formatted(d,mat.Prefix("        "))

	fmt.Printf("A ^ 5 = %0.4v\n\n",fd)

	//Apply a funcion to each of the elements of a.

	d=mat.NewDense(0,0,nil)

	sqrt:=func (_,_ int, v float64) float64 {
		return math.Sqrt(v)
	}

	d.Apply(sqrt,a)

	fd=mat.Formatted(d,mat.Prefix("         "))

	fmt.Printf("A Sqrt = %0.4v\n\n",fd)


}

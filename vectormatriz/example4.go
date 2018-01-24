package main

import (
	"gonum.org/v1/gonum/mat"
	"fmt"
)

func main(){

	components:=[]float64{1.2,-5.7,-2.4,7.3}

	m:=mat.NewDense(2,2,components)

	fmt.Printf("Matriz %+v \n", m)

	fm:=mat.Formatted(m,mat.Prefix(" "))

	fmt.Printf(" %v \n\n", fm)


	val:=m.At(0,0)

	fmt.Printf("Element at 0,0 %0.2f",val)


	//Get the valus of a column
	col:=mat.Col(nil,0,m)

	fmt.Printf("Value in first column is %v \n\n",col)

	//Gt values of a row

	row:=mat.Row(nil,0,m)

	fmt.Printf("Value in the first row is %v \n\n",row)

	//Set a value
	m.Set(0,0,3.3)

	m.SetRow(0,[]float64{14.3,-4.2})

	m.SetCol(0,[]float64{1.7,-0.3})

	fm=mat.Formatted(m,mat.Prefix(" "))

	fmt.Printf(" %v \n\n", fm)
}

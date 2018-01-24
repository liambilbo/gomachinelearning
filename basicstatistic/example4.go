package main

import (
	"github.com/gonum/stat"
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
)

func main(){

	//Previous data
	//0.60 % No Regular Exercise
	//0.25 % Exporadic Exercise
	//0.15 % Regular Exercise


	//New survey
	//260 No Regular Exrcise
	//135 Exporadic Exercise
	//105 Regular Exercise

	observed:=[]float64{260,135,105}

	totalObserved:=500.0

	expected:=[]float64{0.60 * totalObserved,0.25 * totalObserved,0.15 * totalObserved}

	chiSquare:=stat.ChiSquare(observed,expected)

	fmt.Printf("ChiSquare : %0.2f \n",chiSquare)

	//Create a Chi-Squared distribution with K degrees of freedom
	distuvC:=distuv.ChiSquared{
		K:2,
		Src:nil,
	}

	pValue:=distuvC.Prob(chiSquare)

	fmt.Printf("P-Value : %0.4f", pValue)
}

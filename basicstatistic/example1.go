package main

import (
	"github.com/kniren/gota/dataframe"
	"os"
	"log"
	"github.com/gonum/stat"
	"github.com/montanaflynn/stats"
	"fmt"
	"github.com/gonum/floats"
)

func main(){

	irisCsv,err:=os.Open("basicstatistic/data/iris.csv")

	if err!=nil{
		log.Fatal(err)
	}

	defer irisCsv.Close()

	//Create a dataframe from the file
	irisDF:=dataframe.ReadCSV(irisCsv)

	sepalLength:=irisDF.Col("sepal_length").Float()

	meanVal:=stat.Mean(sepalLength,nil)

	modeVal,modeCount:=stat.Mode(sepalLength,nil)

	medianVal,err:=stats.Median(sepalLength)

	if err!=nil{
		log.Fatal(err)
	}

	sepalLengthMax:=floats.Max(sepalLength)
	sepalLengthMin:=floats.Min(sepalLength)

	sepalLengthRange:=sepalLengthMax-sepalLengthMin

	sepalLengthVariance:=stat.Variance(sepalLength,nil)
	sepalLengthStandardDev:=stat.StdDev(sepalLength,nil)

	inds:=make([]int,len(sepalLength))
	floats.Argsort(sepalLength,inds)

	quant25:=stat.Quantile(0.25,stat.Empirical,sepalLength,nil)
	quant50:=stat.Quantile(0.50,stat.Empirical,sepalLength,nil)
	quant75:=stat.Quantile(0.75,stat.Empirical,sepalLength,nil)

	fmt.Printf("Mean value : %0.2f \n",meanVal)
	fmt.Printf("Mode value : %0.2f , Number %f \n",modeVal,modeCount)
	fmt.Printf("Median value : %0.2f \n",medianVal)
	fmt.Printf("Max value : %0.2f \n",sepalLengthMax)
	fmt.Printf("Min value : %0.2f \n",sepalLengthMin)
	fmt.Printf("Range value : %0.2f \n",sepalLengthRange)
	fmt.Printf("Variance value : %0.2f \n",sepalLengthVariance)
	fmt.Printf("StdDev value : %0.2f \n",sepalLengthStandardDev)
	fmt.Printf("Quantile 0.25 value : %0.2f \n",quant25)
	fmt.Printf("Quantile 0.50 value : %0.2f \n",quant50)
	fmt.Printf("Quantile 0.75 value : %0.2f \n",quant75)
}

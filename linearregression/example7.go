package main

import (
	"os"
	"log"
	"encoding/csv"
	"strconv"
	"github.com/sajari/regression"
	"fmt"
	"math"
)

func main(){

	adCsv,err:=os.Open("linearregression/data/Advertising.csv")
	if err!=nil{
		log.Fatal(err)
	}
	defer adCsv.Close()


	reader:=csv.NewReader(adCsv)

	reader.FieldsPerRecord=4

	data,err:=reader.ReadAll()

	if err!=nil{
		log.Fatal(err)
	}

	var r regression.Regression

	r.SetObserved("Sales")
	r.SetVar(0,"Tv")
	r.SetVar(1,"Radio")

	for i,record :=range data {

		if i==0{
			continue
		}

		tvVal,err:=strconv.ParseFloat(record[0],64)
		if err!=nil{
			log.Fatal(err)
		}
		radioVal,err:=strconv.ParseFloat(record[1],64)
		if err!=nil{
			log.Fatal(err)
		}

		yVal,err:=strconv.ParseFloat(record[3],64)
		if err!=nil{
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(yVal,[]float64{tvVal,radioVal}))

	}

	r.Run()

	fmt.Printf("The formula %v \n\n",r.Formula)


	//Calculate MAE
	var mae float64

	for i,record :=range data {

		if i==0{
			continue
		}

		tvVal,err:=strconv.ParseFloat(record[0],64)
		if err!=nil{
			log.Fatal(err)
		}
		radioVal,err:=strconv.ParseFloat(record[1],64)
		if err!=nil{
			log.Fatal(err)
		}

		yVal,err:=strconv.ParseFloat(record[3],64)
		if err!=nil{
			log.Fatal(err)
		}

		predictVal,err:=r.Predict([]float64{tvVal,radioVal})
		if err!=nil{
			log.Fatal(err)
		}

		mae +=math.Abs(yVal-predictVal) / float64(len(data))


	}

	fmt.Printf("MAE %0.02f \n\n",mae)






}

package main

import (
	"os"
	"github.com/sajari/regression"
	"log"
	"encoding/csv"
	"strconv"
	"fmt"
	"math"
)

func main(){
	trainingCsv,err:=os.Open("linearregression/data/training.csv")
	if err!=nil{
		log.Fatal(err)
	}
	defer trainingCsv.Close()


	reader:=csv.NewReader(trainingCsv)

	reader.FieldsPerRecord=4

	trainingData,err:=reader.ReadAll()

	var r regression.Regression

	r.SetObserved("Sales")
	r.SetVar(0,"TV")

	for i,record :=range trainingData {

		if i==0 {
			continue
		}

		yVal,err:=strconv.ParseFloat(record[3],64)
		if err!=nil{
			log.Fatal(err)
		}
		salesVal,err:=strconv.ParseFloat(record[0],64)
		if err!=nil{
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(yVal,[]float64{salesVal}))

	}

	r.Run()

	fmt.Printf("Regression formula : \n Predicted = %v \n\n",r.Formula)

	testingCsv,err:=os.Open("linearregression/data/testing.csv")

	if err!=nil{
		log.Fatal(err)
	}
	defer testingCsv.Close()


	testingreader:=csv.NewReader(testingCsv)

	testingreader.FieldsPerRecord=4

	testingData,err:=testingreader.ReadAll()

	if err!=nil{
		log.Fatal(err)
	}

	var mAE float64

	for i,record := range testingData{

		if i==0 {
			continue
		}
		v,err:=strconv.ParseFloat(record[0],64)
		if err!=nil{
			log.Fatal(err)
		}

		yObserved,err:=strconv.ParseFloat(record[3],64)
		if err!=nil{
			log.Fatal(err)
		}

		yPredicted,err:=r.Predict([]float64{v})
		if err!=nil{
			log.Fatal(err)
		}

		mAE+=math.Abs(yPredicted-yObserved) / float64(len(testingData))


	}

	fmt.Printf("MAE %0.2f \n\n",mAE)

}

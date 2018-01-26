package main

import (
	"os"
	"log"
	"encoding/csv"
	"strconv"
	"github.com/gonum/matrix/mat64"
	"github.com/berkmancenter/ridge"
	"fmt"
)

func main(){

	trainingCsv,err:=os.Open("linearregression/data/training.csv")
	if err!=nil{
		log.Fatal(err)
	}
	defer trainingCsv.Close()

	reader:=csv.NewReader(trainingCsv)
	reader.FieldsPerRecord=4

	data,err:=reader.ReadAll()
	if err!=nil{
		log.Fatal(err)
	}



	featureData:=make([]float64,4 * len(data))
	yData:=make([]float64,len(data))

	var idxFeature int
	var idY int

	for irr,record :=range data {
		if irr==0{
			continue
		}
		for idf, value:= range record {

			floatValue,err:=strconv.ParseFloat(value,64)
			if err!=nil{
				log.Fatal(err)
			}

			if idf==0 {
				featureData[idxFeature]=1
				idxFeature++
			}


			if idf==3 {
				yData[idY]=floatValue
				idY++
			} else {
				featureData[idxFeature]=floatValue
				idxFeature++
			}
		}
	}

	featureMx:=mat64.NewDense(len(data),4,featureData)

	yVec:=mat64.NewVector(len(data),yData)

	r:=ridge.New(featureMx,yVec,1.0)
	r.Regress()

	w1:=r.Coefficients.At(0,0)
	w2:=r.Coefficients.At(1,0)
	w3:=r.Coefficients.At(2,0)
	w4:=r.Coefficients.At(3,0)



	fmt.Printf("Formula : y = %0.03f +  %0.03f TV  +  %0.03f Radio  +  %0.03f Newspaper ",w1,w2,w3,w4)

}

package main

import (
	"os"
	"log"
	"github.com/kniren/gota/dataframe"
	"bufio"
)

func main(){
	advertisingCSV,err:=os.Open("linearregression/data/Advertising.csv")
	if err!=nil{
		log.Fatal(err)
	}
	defer  advertisingCSV.Close()

	advertisingDF:=dataframe.ReadCSV(advertisingCSV)

	numTraining:=(4 * advertisingDF.Nrow()) / 5
	numTesting:=advertisingDF.Nrow() / 5
	
	if numTraining+numTesting<advertisingDF.Nrow() {
		numTraining++
	}

	trainingIdx:=make([]int,numTraining)
	testingIdx:=make([]int,numTesting)

	for i:=0 ; i<numTraining ; i++  {
		trainingIdx[i]=i
	}

	for i:=0 ; i<numTesting ; i++  {
		testingIdx[i]=numTraining+i
	}

	trainingDF:=advertisingDF.Subset(trainingIdx)
	testingDF:=advertisingDF.Subset(testingIdx)

	setMap:=map[int]dataframe.DataFrame{
		0:trainingDF,
		1:testingDF,
	}

	for i,v:=range []string{"training.csv","testing.csv"} {

		subFileCsv,err:=os.Create("linearregression/data/"+v)

		if err!=nil{
			log.Fatal(err)
		}

		buff:=bufio.NewWriter(subFileCsv)

		if err:=setMap[i].WriteCSV(buff);err!=nil{
			log.Fatal(err)
		}

	}

}

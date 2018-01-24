package main

import (
	"os"
	"log"
	"github.com/kniren/gota/dataframe"
	"fmt"
)

func main() {
	irishFile,err:=os.Open("readingcsv/data/iris_labeled.csv")

	if err!=nil{
		log.Fatal(err)
	}

	defer irishFile.Close()

	irisDF:=dataframe.ReadCSV(irishFile)

	fmt.Println(irisDF)

	filter:=dataframe.F{
		Colname:"species",
		Comparator:"==",
		Comparando:"Iris-versicolor",
	}

	versicolorDF:=irisDF.Filter(filter)

	if versicolorDF.Err!=nil{
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println(versicolorDF)

	versicolorDF=irisDF.Filter(filter).Select([]string{"sepal_width","species"})

	if versicolorDF.Err!=nil{
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println(versicolorDF)

	versicolorDF=irisDF.Filter(filter).Select([]string{"sepal_width","species"}).Subset([]int{0,1,2})

	if versicolorDF.Err!=nil{
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println(versicolorDF)




}

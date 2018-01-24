package main

import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
)

func main() {

	irisCsv,err:=os.Open("readingcsv/data/iris.csv")

	if err!=nil{
		log.Fatal(err)

	}

	reader:=csv.NewReader(irisCsv)

	reader.FieldsPerRecord=-1

	var irisData [][]string

	for {
		record,err:=reader.Read()
		if err!=nil {
			if err==io.EOF {
				break
			}
			log.Fatal(err)
		}

		irisData=append(irisData,record)
	}

	fmt.Println(irisData)

}

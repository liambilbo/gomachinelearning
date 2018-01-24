package main

import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
)

func main(){

	irisFile,err:=os.Open("readingcsv/data/iris_unexpected_fields.csv")

	if err!=nil{
		log.Fatal(err)
	}


	var irisData [][] string

	reader:=csv.NewReader(irisFile)

	reader.FieldsPerRecord=5

	for {

		record,err:=reader.Read()

		if err!=nil{
			if err==io.EOF{
				break
			}
			log.Println(err)
			continue
		}
		irisData = append(irisData,record)
	}

	fmt.Printf("Parsed %d lines successfully \n",len(irisData))
}

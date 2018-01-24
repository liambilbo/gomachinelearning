package main

import (
	"os"
	"log"

	"encoding/csv"
	"fmt"
)

func main() {

	irisFile,err:=os.Open("readingcsv/data/iris.csv")
	if err!=nil {
		log.Fatal(err)
	}

	reader:=csv.NewReader(irisFile)

	reader.FieldsPerRecord=-1

	record,err:=reader.ReadAll()

	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println(record)

}




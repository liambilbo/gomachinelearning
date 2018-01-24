package main

import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
	"strconv"
)

type CsvRecord struct{
	SepalLength float64
	SepalWidth float64
	PetalLength float64
	PetalWidth float64
	Spices string
	ParseError error
}

func main() {

	var csvData []CsvRecord

	irisFile,err:=os.Open("readingcsv/data/iris_mixed_types.csv")
	if err!=nil{
		log.Fatal(err)
	}
	defer irisFile.Close()

	reader:=csv.NewReader(irisFile)
	reader.FieldsPerRecord=5

	line:=1

	for {
		record,err:=reader.Read()

		if err!=nil{
			if err==io.EOF{
				log.Println(err)
				break
			}
			continue
		}

		var csvRecord CsvRecord

		for i,v:=range record{
			if i==4 {
				if v=="" {
					log.Printf("Error line %d - Unexpected type field \n",i)
					csvRecord.ParseError=fmt.Errorf("Empty string value")
					break
				}
				csvRecord.Spices=v
				continue
			}

			var floatValue float64

			if floatValue,err=strconv.ParseFloat(v,64); err!=nil {
				log.Println(err)
				csvRecord.ParseError=fmt.Errorf("Not a float value")
				break
			}

			switch i {
				case 0:
					csvRecord.SepalLength=floatValue
				case 1:
					csvRecord.SepalWidth=floatValue
				case 2:
					csvRecord.PetalLength=floatValue
				case 3:
					csvRecord.PetalWidth=floatValue
			}

			if csvRecord.ParseError==nil{
				csvData=append(csvData,csvRecord)
			}
		}

		line++


	}

	fmt.Printf("Total Number %d / Correct %d",line,len(csvData))

}
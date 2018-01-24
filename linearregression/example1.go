package main

import (
	"os"
	"log"
	"github.com/kniren/gota/dataframe"
	"fmt"
)

func main() {

	advertFile , err := os.Open("linearregression/data/Advertising.csv")
	//advertFile , err := os.Open("../Advertising.csv")


	if err!=nil{
		log.Fatalf("Error open file %v",err)
	}

	defer advertFile.Close()

	advertFrame:=dataframe.ReadCSV(advertFile)

	advertSummary:=advertFrame.Describe()

	fmt.Println(advertSummary)



}

package main

import (
	"os"
	"log"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot"
	"fmt"
	"gonum.org/v1/plot/vg"
)

func main() {

	advertFile,err:=os.Open("linearregression/data/Advertising.csv")

	defer advertFile.Close()

	if err!=nil {
		log.Fatal(err)
	}


	advertDF:=dataframe.ReadCSV(advertFile)


	//Create an Histogram for each column
	for _,colName := range advertDF.Names() {

		plotVals:=make(plotter.Values,advertDF.Nrow())

		for i,value:= range advertDF.Col(colName).Float() {
			plotVals[i]=value
		}

		p,err := plot.New()

		if err!=nil {
			log.Fatal(err)
		}

		p.Title.Text = fmt.Sprintf("Histogram of %s " , colName)

		h,err:=plotter.NewHist(plotVals,16)

		if err!=nil {
			log.Fatal(err)
		}
		h.Normalize(1)

		p.Add(h)

		if err:=p.Save(4 * vg.Inch,4 * vg.Inch, "linearregression/data/"+colName+".png");err!=nil{
			log.Fatal(err)
		}

	}



}
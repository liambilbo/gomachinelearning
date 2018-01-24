package main

import (
	"fmt"
	"log"
	"os"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/kniren/gota/dataframe"
)

func main(){

	irisCsv,err:=os.Open("basicstatistic/data/iris.csv")
	if err!=nil{
		log.Fatal(err)
	}

	defer irisCsv.Close()

	irisDF:=dataframe.ReadCSV(irisCsv)

	for _,colName:=range irisDF.Names() {

		if colName!="species" {

			v:=make(plotter.Values,irisDF.Nrow())

			for i,value:=range irisDF.Col(colName).Float() {
				v[i]=value
			}


			p,err:=plot.New()

			if err!=nil{
				log.Fatal(err)
			}

			p.Title.Text =fmt.Sprintf("Histogram of a %s",colName)
			h,err:=plotter.NewHist(v,16)

			if err!=nil{
				log.Fatal(err)
			}

			h.Normalize(1)

			p.Add(h)

			if err:=p.Save(4*vg.Inch,4*vg.Inch,"basicstatistic/data/"+colName+".png");err!=nil{
				log.Fatal(err)
			}


		}




	}

}

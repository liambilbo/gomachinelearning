package main

import (
	"os"
	"log"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/plotter"
)

func main(){

	irisFile,err:=os.Open("basicstatistic/data/iris.csv")

	if err!=nil{
		log.Fatal(err)
	}


	defer irisFile.Close()

	irisDF:=dataframe.ReadCSV(irisFile)

	p,err:=plot.New()

	if err!=nil{
		log.Fatal(err)
	}

	p.Title.Text = "Box Plots"
	p.Y.Label.Text = "Values"

	w:=vg.Points(50)

	for idx,colName:=range irisDF.Names() {

		if colName!="species"{
			v:=make(plotter.Values,irisDF.Nrow())
			v=irisDF.Col(colName).Float()

			b,err:=plotter.NewBoxPlot(w,float64(idx),v)

			if err!=nil{
				log.Fatal(err)
			}

			p.Add(b)
		}

	}

	p.NominalX("sepal_length","sepal_width","petal_length","petal_width")

	if err:=p.Save(6 * vg.Inch,8 * vg.Inch,"basicstatistic/data/boxplots.png");err!=nil{
		log.Fatal(err)
	}




}

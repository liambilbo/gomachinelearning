package main

import (
	"os"
	"log"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	advertisingCsv,err:=os.Open("linearregression/data/Advertising.csv")

	if err!=nil{
		log.Fatal(err)
	}

	defer advertisingCsv.Close()

	advertisingDF:=dataframe.ReadCSV(advertisingCsv)


	yVal:=advertisingDF.Col("Sales").Float()


	for _,colName:= range advertisingDF.Names() {

		pts:=make(plotter.XYs,advertisingDF.Nrow())

		for i,floatVal:=range advertisingDF.Col(colName).Float(){
			pts[i].Y=yVal[i]
			pts[i].X=floatVal
		}

		plot,err :=plot.New()
		if err!=nil{
			log.Fatal(err)
		}

		graph,err:=plotter.NewScatter(pts)

		if err!=nil{
			log.Fatal(err)
		}

		plot.Y.Label.Text="Sales"
		plot.X.Label.Text=colName
		plot.Add(plotter.NewGrid())

		graph.GlyphStyle.Radius = vg.Points(3)

		plot.Add(graph)

		if err=plot.Save(4 * vg.Inch,4 * vg.Inch,"linearregression/data/"+colName+"_scatter.png");err!=nil{
			log.Fatal(err)
		}

	}



}

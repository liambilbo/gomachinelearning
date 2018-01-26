package main

import (
	"os"
	"log"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func predict(tv float64) float64 {
	return 7.07+tv * 0.05
}

func main(){

		adCsv,err:=os.Open("linearregression/data/Advertising.csv")
		if err!=nil{
			log.Fatal(err)
		}
		defer adCsv.Close()

		adDF:=dataframe.ReadCSV(adCsv)

		plotg,err:=plot.New()
		if err!=nil{
			log.Fatal(err)
		}

		plotg.Y.Label.Text="Sales"
	    plotg.X.Label.Text="Tv"

	    pts:=make(plotter.XYs,adDF.Nrow())
	    ptspredict:=make(plotter.XYs,adDF.Nrow())

	    yVal:=adDF.Col("Sales").Float()

		for i,v := range adDF.Col("TV").Float(){
			pts[i].X = v
			ptspredict[i].X = v
			pts[i].Y = yVal[i]
			ptspredict[i].Y = predict(v)
		}

		plotg.Add(plotter.NewGrid())

		scatter,err:=plotter.NewScatter(pts)
		if err!=nil{
			log.Fatal(err)
		}
		plotg.Add(scatter)

	    scatter.GlyphStyle.Radius=3

	    line , err:=plotter.NewLine(ptspredict)
		if err!=nil{
			log.Fatal(err)
		}
		line.LineStyle.Width=vg.Points(1)
		line.LineStyle.Dashes=[]vg.Length{vg.Points(5),vg.Points(5)}
		plotg.Add(line)

		if err:=plotg.Save(4 * vg.Inch,4* vg.Inch,"linearregression/data/regression_line.png");err!=nil{
			log.Fatal(err)
		}


}

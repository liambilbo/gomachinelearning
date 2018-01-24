package main

import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"strconv"
	"math"
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main(){

  dataCvs,err:=os.Open("evaluation/data/continuous_data.csv")
  if err!=nil{
  	log.Fatal(err)
  }

  defer dataCvs.Close()

  var observed []float64
  var predicted []float64

  reader:=csv.NewReader(dataCvs)

  var line int

  for {
	  row,err:=reader.Read()
	  line++

	  if line==1 {
	  	continue
	  }

	  if err!=nil{
	  	if err==io.EOF {
	  		break
		}
		log.Fatal(err)

	  }


	  observedVal,err:=strconv.ParseFloat(row[0],64)
	  if err!=nil{
	  	log.Printf("Parsing line %d failed, unexpected type\n",line)
	  	continue
	  }

	  predictedVal,err:=strconv.ParseFloat(row[1],64)

	  if err!=nil{
		  log.Printf("Parsing line %d failed, unexpected type\n",line)
		  continue
	  }
	  observed=append(observed,observedVal)
	  predicted=append(predicted,predictedVal)

  	}
	//Calculate MSE MAE
	var sumSE float64
	var sumAE float64

	for idx,value:=range observed{
		sumSE=sumSE+math.Pow(value-predicted[idx],2)
		sumAE=sumAE+math.Abs(value-predicted[idx])
	}

	mae:=sumAE/float64(len(observed))
	mse:=sumSE/float64(len(observed))

	fmt.Printf("MAE : %0.2f \n",mae)
	fmt.Printf("MSE : %0.2f \n",mse)

	//Calculate the R^2

	rSquared:=stat.RSquaredFrom(predicted,observed,nil)

	fmt.Printf("R^2 : %0.2f \n",rSquared)

	rSquared=stat.RSquaredFrom(observed,predicted,nil)

	fmt.Printf("R^2 : %0.2f \n",rSquared)


}
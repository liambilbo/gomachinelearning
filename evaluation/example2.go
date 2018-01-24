package main

import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"strconv"
	"fmt"
)

func main(){

	dataCsv,err:=os.Open("evaluation/data/labeled.csv")

	if err!=nil{
		log.Fatal(err)
	}
	defer dataCsv.Close()

	reader:=csv.NewReader(dataCsv)

	var line int
	var observed []int
	var predicted []int

	for {

		row,err:=reader.Read()
		line++


		if err!=nil{
			if err==io.EOF{
				break
			}

			log.Fatal(err)
		}

		if line== 1 {
			continue
		}

		observedV,err:=strconv.Atoi(row[0])
		if err!=nil{
			log.Printf("Error paring line %f",line)
		}

		predictedV,err:=strconv.Atoi(row[1])
		if err!=nil{
			log.Printf("Error paring line %f",line)
		}

		observed=append(observed,observedV)
		predicted=append(predicted,predictedV)



	}

	//Calculate accurate (TP+TN)/(TP+TN+FP+FN)
	var t float64
	for idx,_:=range observed{
		if observed[idx]==predicted[idx] {
			t++
		}
	}
	accuracy:=t/float64(len(observed))

	fmt.Printf("Accuracy : %0.2f \n",accuracy)


	//Calculata Precision & Recall

	class:=[]int{0,1,2}

	for _,valueClass:=range class {

		var tp,tn,fp,fn int

		for  idx,_ :=range observed{
			if observed[idx]==predicted[idx] {
				if predicted[idx]==valueClass {
					tp++
				}else {
					tn++
				}
			}else {
				if predicted[idx]==valueClass {
					fp++
				}else {
					fn++
				}
			}
		}


		precision:= float64(tp)/float64(tp+fp)
		recall:=float64(tp)/float64(tp+fn)

		fmt.Printf("Class %d Precision : %0.2f \n",valueClass,precision)
		fmt.Printf("Class %d Recall : %0.2f \n",valueClass,recall)
	}


}

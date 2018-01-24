package main

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL int `json:"ttl"`
	Data struct {
		Stations []station `json:"stations"`
	} `json:"data"`

}

type station struct{
	StationId string `json:"station_id"`
	NumBikesAvaible int `json:"num_bikes_avaible"`
	NumBikesDisabled int `json:"num_bikes_disabled"`
	NumDocksAvaible int `json:"num_docks_avaible"`
	NumDocksDisabled int `json:"num_docks_disabled"`
	IsInstalled int `json:"is_installed"`
	IsRenting int `json:"is_renting"`
	IsReturning int `json:"is_returning"`
	LastReported int `json:"last_reported"`
	EightdHasAvailableKeys bool `json:"eightd_has_available_keys"`


}

func main(){

	//Get File From Url
	response,err:=http.Get(citiBikeURL)

	if err!=nil{
		log.Fatal(err)
	}

	defer response.Body.Close()

	//Read the body of the response in []bytes
	data,err:=ioutil.ReadAll(response.Body)

	if err!=nil{
		log.Fatal(err)
	}

	var sD stationData

	err=json.Unmarshal(data,&sD)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n",sD.Data.Stations[0])


	data,err=json.Marshal(sD)

	if err!=nil{
		log.Fatal(err)
	}

	ioutil.WriteFile("readingjson/output/station.json",data,0644)

}

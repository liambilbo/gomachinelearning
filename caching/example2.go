package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main(){

	db,err:=bolt.Open("caching/data/embedded.db",0600,nil)

	if err!=nil{
		log.Fatal(err)
	}

	defer db.Close()

	//Create Bucket inside db
	if err=db.Update(func(tx *bolt.Tx) error {

		_,err:=tx.CreateBucket([]byte("MyBucket") )
		if err!=nil{
			return fmt.Errorf("Error creating Bucket %s \n",err)
		}
		return nil
	});err!=nil{
		log.Fatal(err)
	}


	if err=db.Update(func(tx *bolt.Tx) error {
		myBucket:=tx.Bucket([]byte("MyBucket") )
		err=myBucket.Put([]byte("mykey"),[]byte("myvalue"))
		return err
	});err!=nil{
		log.Fatal(err)
	}


	if err=db.View(func(tx *bolt.Tx) error {
		myBucket:=tx.Bucket([]byte("MyBucket") )
		cursor:=myBucket.Cursor()

		for  k,v:=cursor.First() ; k!=nil ; k,v =cursor.Next(){
			fmt.Printf("Key : %s , Value %s \n", k,v)
		}
		return nil
	});err!=nil{
		log.Fatal(err)
	}


}

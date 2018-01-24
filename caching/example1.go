package main

import (
	"github.com/patrickmn/go-cache"
	"time"
	"fmt"
)

func main(){

	c:=cache.New(5 * time.Minute,30 * time.Second)

	c.Set("key","example value",cache.DefaultExpiration)

	if v,found:=c.Get("key") ; found {
		fmt.Printf("Value of key % v", v)
	}

}

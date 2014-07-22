package main

import (
	"fmt"
	"log"
	"github.com/nurikevenoglu/reddit"
)

func main(){
	items, err := reddit.Get("golang")

	if err != nil {
		log.Fatal(err)
	}

	for i,item := range items {
		fmt.Println(i, item)
	}
}

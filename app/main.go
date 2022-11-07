package main

import (
	"fmt"
	"app/routecheck"
)

func main() {
	coche, err := routecheck.NewCoche(1,2,1)
	if err != nil{
		print(err)
	}else{
		fmt.Println(coche)
	}
}
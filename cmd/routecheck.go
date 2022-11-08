package main

import (
	"fmt"
	"github.com/dmonjasm/RouteCheck/internal"
)

func main() {
	coche, err := internal.NewCoche(internal.Ligero, 3.5, 8)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Coche: ", *coche)
	}
	inicio, err := internal.NewPosicion(23, -100)
	if err != nil{
		fmt.Println(err)
	}
	fin, err := internal.NewPosicion(-90, 150)
	if err != nil{
		fmt.Println(err)
	}
	peaje, err := internal.NewPeaje(map[internal.Peso]float32{internal.Ligero: 2.3, internal.Pesado: 5.1})
	if err != nil{
		fmt.Println(err)
	}

	tramo, err := internal.NewTramo(inicio, fin, peaje, 120)
	if err != nil{
		fmt.Println(err)
	}
	ruta, err := internal.NewRuta(tramo, nil, 200)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Ruta: ", *ruta)
	}
	
}
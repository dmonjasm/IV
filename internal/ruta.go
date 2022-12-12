package internal

import "errors"

type ConsumoVehiculo float32
type Peaje float32

type Ruta struct {
	tramos []Tramo
	peaje  Peaje
}

func NewRuta(tramos []Tramo, peaje Peaje) (*Ruta, error) {
	if len(tramos) == 0 {
		return nil, errors.New("Una ruta debe contener mínimo un tramo.")
	}
	if peaje < 0 {
		return nil, errors.New("Un peaje no puede ser negativo.")
	}
	return &Ruta{tramos, peaje}, nil
}

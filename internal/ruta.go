package internal
import "errors"

type ConsumoVehiculo float32

type Ruta struct{
	precio float32
	tramos [] Tramo
}

func NewRuta(tramos [] Tramo) (*Ruta, error){
	if len(tramos) == 0 {
		return nil, errors.New("Una ruta debe contener mínimo un tramo.")
	}
	return &Ruta{0, tramos}, nil
}

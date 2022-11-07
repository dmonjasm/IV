package routecheck
import "errors"

type Ruta struct{
	tramo Tramo
	ruta *Ruta
	puntuacion float32
}

func NewRuta(tramo Tramo, ruta *Ruta, puntuacion float32) (*Ruta, error){
	err := PuntuacionError(puntuacion)
	if err != nil{
		return nil, err
	}
	return &Ruta{tramo, ruta, puntuacion}, nil
}

func PuntuacionError(puntuacion float32) error{
	if puntuacion < 0 {
		return errors.New("La puntuación de una ruta no puede ser negativa.")
	}
	return nil
}
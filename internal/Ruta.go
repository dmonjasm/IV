package internal
import "errors"

type Ruta struct{
	tramo Tramo
	ruta *Ruta
	precio float32
}

func NewRuta(tramo *Tramo, ruta *Ruta, precio float32) (*Ruta, error){
	err := PuntuacionError(precio)
	if err != nil{
		return nil, err
	}
	return &Ruta{*tramo, ruta, precio}, nil
}

func PuntuacionError(precio float32) error{
	if precio < 0 {
		return errors.New("El precio no puede ser negativo.")
	}
	return nil
}
package internal
import "errors"

type Tramo struct{
    distancia uint
    limite uint8
}

func NewTramo(distancia uint, limite uint8)(*Tramo, error){
	err := NewTramoError(distancia, limite)
	if err != nil{
		return nil, err
	}
    return &Tramo{distancia, limite}, nil
}

func NewTramoError(distancia uint,limite uint8) error{
    if limite < 10 || limite > 120{
        return errors.New("El límite de velocidad no es válido.")
    }
    if distancia <= 0{
        return errors.New("La distancia no puede ser negativa.")
    }
	return nil
}
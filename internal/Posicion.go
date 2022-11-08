package internal
import "errors"

type Posicion struct{
    X float32
    Y float32
}

func NewPosicion(x float32, y float32) (*Posicion, error){
    err := PosicionError(x, y)
    if err != nil{
        return nil, err
    }
    return &Posicion{x, y}, nil
}

func PosicionError(x float32, y float32) error{
    if x > 90 || x < -90 || y < -180 || y > 180{
        return errors.New("Fuera de rango([-90,90],[-180,180]).")
    }
    return nil
}
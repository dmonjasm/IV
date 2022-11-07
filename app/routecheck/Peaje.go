package routecheck
import "errors"

type Peaje struct{
    precios map[Peso]float32
}

func NewPeaje(precios map[Peso]float32) (*Peaje, error){

    err := NewPeajeError(precios)
    if err != nil{
        return nil, err
    }
    return &Peaje{precios}, nil
}

func (peaje Peaje) GetPeaje(peso Peso) (float32, error){
    err := PesoError(peso)
    if err != nil{
        return 0, err   
    }
    return peaje.precios[peso], nil
}


func NewPeajeError(precios map[Peso]float32) error{
    for k, v := range precios{
        err := PesoError(k)
        if err != nil{
            return err
        }
        if v < 0{
            return errors.New("El precio de un peaje no puede ser negativo.")
        }
    }
    return nil
}
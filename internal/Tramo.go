package internal
import "errors"

type Tramo struct{
    inicio Posicion
    fin Posicion
    distancia float32
    peaje Peaje
    limite uint8
}

func NewTramo(inicio *Posicion, fin *Posicion, peaje *Peaje, limite uint8)(*Tramo, error){
	err := LimiteError(limite)
	if err != nil{
		return nil, err
	}
    return &Tramo{*inicio, *fin, 0, *peaje, limite}, nil	//implementar el calculo entre dos coordenadas para la distancia
}

func LimiteError(limite uint8) error{
    if limite < 10 || limite > 120{
        return errors.New("El límite de velocidad no es válido.")
    }
	return nil
}
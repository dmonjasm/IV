package internal

import "errors"

type Tramo struct {
	distancia       uint
	velocidadLimite uint8
}

func NewTramo(distancia uint, velocidadLimite uint8) (*Tramo, error) {
	err := NewTramoError(distancia, velocidadLimite)
	if err != nil {
		return nil, err
	}
	return &Tramo{distancia, velocidadLimite}, nil
}

func NewTramoError(distancia uint, velocidadLimite uint8) error {
	if velocidadLimite < 10 || velocidadLimite > 120 {
		return errors.New("El límite de velocidad no es válido.")
	}
	if distancia <= 0 {
		return errors.New("La distancia no puede ser negativa.")
	}
	return nil
}

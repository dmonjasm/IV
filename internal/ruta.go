package internal

import "errors"

type ConsumoVehiculo float32
type Peaje float32

type Ruta struct {
	tramos []Tramo
	peaje  Peaje
}

func NewRuta(tramos []Tramo, peaje Peaje) (Ruta, error) {
	if len(tramos) == 0 {
		return Ruta{}, errors.New("Una ruta debe contener mínimo un tramo.")
	}
	if peaje < 0 {
		return Ruta{}, errors.New("Un peaje no puede ser negativo.")
	}
	return Ruta{tramos, peaje}, nil
}

func EstimatedDistance(ruta Ruta) (uint, error) {
	if len(ruta.tramos) == 0 {
		return 0, errors.New("Una ruta debe contener mínimo un tramo.")
	}

	var distancia_total uint
	distancia_total = 0

	for i := 0; i < len(ruta.tramos); i++ {
		distancia_total += ruta.tramos[i].distancia
	}

	return distancia_total, nil
}

func EstimatedTime(ruta Ruta) (float32, error) {
	if len(ruta.tramos) == 0 {
		return -1.0, errors.New("Una ruta debe contener mínimo un tramo.")
	}

	var tiempo_total float32
	tiempo_total = 0.0

	for i := 0; i < len(ruta.tramos); i++ {
		tiempo_total += float32(ruta.tramos[i].distancia) / float32(ruta.tramos[i].velocidadLimite)
	}

	return tiempo_total, nil
}

func FastestRoute(rutas []Ruta) (Ruta, error) {
	if len(rutas) == 0 {
		return Ruta{}, errors.New("Se debe proporcionar mínimo una ruta.")
	}

	var fastest_route Ruta
	var mejor_tiempo float32
	var tiempo_actual float32

	mejor_tiempo = EstimatedTime(rutas[0])
	fastest_route = rutas[0]
	

	for i := 1; i < len(rutas); i++ {
		tiempo_actual = EstimatedTime(rutas[i])

		if tiempo_actual < tiempo_total{
			mejor_tiempo = tiempo_actual
			fastest_route = rutas[i]
		}

		else if tiempo_actual == tiempo_total {
			if rutas[i].peaje < fastest_route.peaje{
				fastest_route = rutas[i]
			}
		}
	}

	return fastest_route, nil
}

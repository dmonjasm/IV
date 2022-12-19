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

	mejor_tiempo, _ = EstimatedTime(rutas[0])
	fastest_route = rutas[0]
	

	for i := 1; i < len(rutas); i++ {
		tiempo_actual, _ = EstimatedTime(rutas[i])

		if tiempo_actual < mejor_tiempo{
			mejor_tiempo = tiempo_actual
			fastest_route = rutas[i]
		} else if tiempo_actual == mejor_tiempo {
			if rutas[i].peaje < fastest_route.peaje{
				fastest_route = rutas[i]
			}
		}
	}

	return fastest_route, nil
}

func CheapestRoute(rutas []Ruta) (Ruta, error) {
	if len(rutas) == 0 {
		return Ruta{}, errors.New("Se debe proporcionar mínimo una ruta.")
	}

	var cheapest_route Ruta
	var mejor_precio Peaje
	var precio_actual Peaje

	mejor_precio = rutas[0].peaje
	cheapest_route = rutas[0]
	

	for i := 1; i < len(rutas); i++ {
		precio_actual = rutas[i].peaje

		if precio_actual < mejor_precio{
			mejor_precio = precio_actual
			cheapest_route = rutas[i]
		} else if precio_actual == mejor_precio {
			tiempo_actual, _ := EstimatedTime(rutas[i])
			tiempo_mejor_ruta, _ := EstimatedTime(cheapest_route)
			if tiempo_actual < tiempo_mejor_ruta{
				cheapest_route = rutas[i]
			}
		}
	}

	return cheapest_route, nil
}

func CostTimeRoute(rutas []Ruta) (Ruta, error) {
	if len(rutas) == 0 {
		return Ruta{}, errors.New("Se debe proporcionar mínimo una ruta.")
	}

	var best_route Ruta
	var best_cost_time_relation float32
	var cost_time_actual float32
	var tiempo_actual float32
	var distancia_actual uint

	//Se calcula la velocidad media de la ruta pero se le quita a la misma el precio del via ya que
	//se entiende que una ruta cara no debería sobrepasar los 50€ en peajes y que la velocidad media
	//debería ser alta.
	tiempo_actual, _ = EstimatedTime(rutas[0])
	distancia_actual, _ = EstimatedDistance(rutas[0])
	best_cost_time_relation = (float32(distancia_actual) / tiempo_actual) - float32(rutas[0].peaje)
	best_route = rutas[0]

	for i := 1; i < len(rutas); i++ {
		tiempo_actual, _ = EstimatedTime(rutas[i])
		distancia_actual, _ = EstimatedDistance(rutas[i])
		cost_time_actual = (float32(distancia_actual) / tiempo_actual) - float32(rutas[i].peaje)

		if cost_time_actual > best_cost_time_relation {
			best_cost_time_relation = cost_time_actual
			best_route = rutas[i] 
		} else if cost_time_actual > best_cost_time_relation{
			//Como no se le ha dado un gran peso al precio consideramos que dos rutas con una relación similar
			//mantendrán una velocidad media similar, luego optaremos por la más barata
			if rutas[i].peaje < best_route.peaje{
				best_route = rutas[i]
			}
		}
	}

	return best_route, nil
}

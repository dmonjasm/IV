package internal
import "errors"

type Peso uint8
const(
    Ligero Peso = iota
    Pesado
)
const PesoTotal uint8 = 2

type Vehiculo struct{
    peso Peso
    consumoCiudad float32
    cosumoCarretera float32
}

func NewVehiculo(peso Peso, cCiudad float32, cCarretera float32) (*Vehiculo, error){
    err := NewVehiculoError(peso, cCiudad, cCarretera)
    if err != nil {
        return nil, err
    }
    return &Vehiculo{peso, cCiudad, cCarretera}, nil
}

func NewVehiculoError(peso Peso, cCiudad float32, cCarretera float32) error{
    if cCiudad <= 0{
        return errors.New("El consumo en ciudad no puede ser menor o igual que 0.")
    }
    if cCarretera <= 0{
        return errors.New("El consumo en carretera no puede ser menor o igual que 0.")
    }
    return PesoError(peso)
}

func PesoError(peso Peso) error{
    if uint8(peso) >= PesoTotal{
        return errors.New("El peso no está definido.")
    }
    return nil
}
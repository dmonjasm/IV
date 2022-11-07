package routecheck
import "errors"

type Peso uint8
const(
    Ligero Peso = iota
    Pesado
)
const PesoTotal uint8 = 2

type Coche struct{
    peso Peso
    consumoCiudad float32
    cosumoCarretera float32
}

func NewCoche(peso Peso, cCiudad float32, cCarretera float32) (*Coche, error){
    err := NewCocheError(peso, cCiudad, cCarretera)
    if err != nil {
        return nil, err
    }
    return &Coche{peso, cCiudad, cCarretera}, nil
}

func NewCocheError(peso Peso, cCiudad float32, cCarretera float32) error{
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
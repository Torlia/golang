package medidas //creamos el paquete llamado medidas

//Definimos tres constantes
const KmPorMilla = 1.60934
const MetrosPorPie = 0.3048
const CentimetrosPorPulagada = 2.54

//Creamos tres funciones para la conversión de unidades del SI y SD
func MillasEnKms(s float64) float64 {
	return s * KmPorMilla
}

func PiesEnMetros(s float64) float64 {
	return s * MetrosPorPie
}

func PulgadasEnCentimetros(s float64) float64 {
	return s * CentimetrosPorPulagada
}

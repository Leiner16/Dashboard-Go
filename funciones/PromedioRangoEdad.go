package funcs

import (
	"math"

	"github.com/derly/TallerGoWeb/modelos"
)

func PromedioTodosRangosEdad(data []modelos.Estudiante) []modelos.RangoPromedio {
	var promedios []modelos.RangoPromedio

	promedios = append(promedios, PromedioEstudiantesRangoEdad(data, 20, 29))
	promedios = append(promedios, PromedioEstudiantesRangoEdad(data, 30, 39))
	promedios = append(promedios, PromedioEstudiantesRangoEdad(data, 40, 49))
	promedios = append(promedios, PromedioEstudiantesRangoEdad(data, 50, 59))
	promedios = append(promedios, PromedioEstudiantesRangoEdad(data, 60, 69))
	promedios = append(promedios, PromedioEstudiantesRangoEdad(data, 70, 79))
	promedios = append(promedios, PromedioEstudiantesRangoEdad(data, 80, 89))

	return promedios
}

func PromedioEstudiantesRangoEdad(data []modelos.Estudiante, rangoMenor int, rangoMayor int) modelos.RangoPromedio {

	suma := 0.0
	contador := 0

	for _, estudiante := range data {

		if estudiante.Edad >= rangoMenor && estudiante.Edad <= rangoMayor {
			//calcular promedio de notas
			for _, curso := range estudiante.Cursos {
				suma += curso.Nota
				contador++
			}

		}
	}
	if contador == 0 {
		rangoPromedio := modelos.RangoPromedio{
			RangoMayor: rangoMayor,
			RangoMenor: rangoMenor,
			Promedio:   0,
		}
		return rangoPromedio
	}
	promedio := suma / float64(contador)
	promedioRedondeado := math.Round(promedio*100) / 100

	rangoPromedio := modelos.RangoPromedio{
		RangoMayor: rangoMayor,
		RangoMenor: rangoMenor,
		Promedio:   promedioRedondeado,
	}
	return rangoPromedio

}

package funcs

import (
	"github.com/derly/TallerGoWeb/modelos"
)

func MayorEdadHombres(data []modelos.Estudiante) []modelos.Estudiante {

	MayorEdadHombres := []modelos.Estudiante{}

	for _, estudiante := range data {

		if estudiante.Gender == "male" {
			if len(MayorEdadHombres) == 0 || (estudiante.Edad > MayorEdadHombres[0].Edad) {
				MayorEdadHombres = []modelos.Estudiante{estudiante}
			} else if estudiante.Edad == MayorEdadHombres[0].Edad {
				MayorEdadHombres = append(MayorEdadHombres, estudiante)
			}
		}

	}

	return MayorEdadHombres
}

func MayorEdadMujeres(data []modelos.Estudiante) []modelos.Estudiante {

	MayorEdadMujeres := []modelos.Estudiante{}

	for _, estudiante := range data {

		if estudiante.Gender == "female" {
			if len(MayorEdadMujeres) == 0 || estudiante.Edad > MayorEdadMujeres[0].Edad {
				MayorEdadMujeres = []modelos.Estudiante{estudiante}
			} else if estudiante.Edad == MayorEdadMujeres[0].Edad {
				MayorEdadMujeres = append(MayorEdadMujeres, estudiante)
			}
		}

	}

	return MayorEdadMujeres
}

func MenorEdadHombres(data []modelos.Estudiante) []modelos.Estudiante {

	MenorEdadHombres := []modelos.Estudiante{}

	for _, estudiante := range data {

		if estudiante.Gender == "male" {
			if len(MenorEdadHombres) == 0 || estudiante.Edad < MenorEdadHombres[0].Edad {
				MenorEdadHombres = []modelos.Estudiante{estudiante}
			} else if estudiante.Edad == MenorEdadHombres[0].Edad {
				MenorEdadHombres = append(MenorEdadHombres, estudiante)
			}
		}

	}
	return MenorEdadHombres
}

func MenorEdadMujeres(data []modelos.Estudiante) []modelos.Estudiante {

	MenorEdadMujeres := []modelos.Estudiante{}

	for _, estudiante := range data {
		if estudiante.Gender == "female" {
			if len(MenorEdadMujeres) == 0 || estudiante.Edad < MenorEdadMujeres[0].Edad {
				MenorEdadMujeres = []modelos.Estudiante{estudiante}
			} else if estudiante.Edad == MenorEdadMujeres[0].Edad {
				MenorEdadMujeres = append(MenorEdadMujeres, estudiante)
			}
		}
	}

	return MenorEdadMujeres

}

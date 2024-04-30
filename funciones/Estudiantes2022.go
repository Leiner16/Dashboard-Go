package funcs

import (
	"github.com/derly/TallerGoWeb/modelos"
)

func Estudiantes2022(data []modelos.Estudiante) []modelos.Estudiante {

	Matricula2022 := []modelos.Estudiante{}

	for _, estudiante := range data {

		if estudiante.Matriculado[:4] == "2022" {
			Matricula2022 = append(Matricula2022, estudiante)
		}

	}

	return Matricula2022
}

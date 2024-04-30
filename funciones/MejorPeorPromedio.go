package funcs

import (
	"github.com/derly/TallerGoWeb/modelos"
)

func MejorPromedio(data []modelos.Estudiante) modelos.Estudiante {

	//estudiante con mejor promedio
	suma := 0.0
	estudianteMejorPromedio := data[0]
	promedioMejorEstudiante := 0.0
	for _, estudiante := range data {
		suma = 0.0
		for _, curso := range estudiante.Cursos {
			suma += curso.Nota
		}

		if suma/float64(len(estudiante.Cursos)) > promedioMejorEstudiante {
			estudianteMejorPromedio = estudiante
			promedioMejorEstudiante = suma / float64(len(estudiante.Cursos))
		}
	}

	return estudianteMejorPromedio
}

func PeorPromedio(data []modelos.Estudiante) modelos.Estudiante {

	//estudiante con mejor promedio
	suma := 0.0
	estudiantePeorPromedio := data[0]
	promedioPeorEstudiante := 100.0
	for _, estudiante := range data {

		suma = 0.0
		for _, curso := range estudiante.Cursos {
			suma += curso.Nota
		}

		if suma/float64(len(estudiante.Cursos)) < promedioPeorEstudiante {
			estudiantePeorPromedio = estudiante
			promedioPeorEstudiante = suma / float64(len(estudiante.Cursos))
		}

	}

	return estudiantePeorPromedio
}

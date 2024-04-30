package funcs

import (
	"math"

	"github.com/derly/TallerGoWeb/modelos"
)

func CursosPromedioVarianzaDesviacionRango(data []modelos.Estudiante) []modelos.EstadisticasCurso {

	var Estadisticas []modelos.EstadisticasCurso

	Estadisticas = append(Estadisticas, EstadisticasCurso(data, "Algebra lineal"))
	Estadisticas = append(Estadisticas, EstadisticasCurso(data, "Calculo diferencial"))
	Estadisticas = append(Estadisticas, EstadisticasCurso(data, "POO"))
	Estadisticas = append(Estadisticas, EstadisticasCurso(data, "CTD"))

	return Estadisticas
}

func EstadisticasCurso(data []modelos.Estudiante, curso string) modelos.EstadisticasCurso {
	//calculo de promedio
	suma := 0.0
	for _, estudiante := range data {
		for _, cursoEstudiante := range estudiante.Cursos {
			if cursoEstudiante.Curso == curso {
				suma += cursoEstudiante.Nota
			}
		}
	}
	promedio := suma / float64(len(data))

	//calculo de nota
	notaMayor := 0.0
	notaMenor := 100.0
	for _, estudiante := range data {
		for _, cursoEstudiante := range estudiante.Cursos {
			if cursoEstudiante.Curso == curso {
				if cursoEstudiante.Nota > notaMayor {
					notaMayor = cursoEstudiante.Nota
				}
				if cursoEstudiante.Nota < notaMenor {
					notaMenor = cursoEstudiante.Nota
				}
			}
		}
	}

	//calculo de varianza y desviacion estandar
	suma = 0.0
	for _, estudiante := range data {
		for _, cursoEstudiante := range estudiante.Cursos {
			if cursoEstudiante.Curso == curso {
				suma += (cursoEstudiante.Nota - promedio) * (cursoEstudiante.Nota - promedio)
			}
		}
	}
	varianza := suma / float64(len(data))
	desviacionEstandar := math.Sqrt(varianza)

	EstadisticasCurso := modelos.EstadisticasCurso{
		Cursos:             curso,
		Promedio:           promedio,
		NotaMenor:          notaMenor,
		NotaMayor:          notaMayor,
		Varianza:           varianza,
		DesviacionEstandar: desviacionEstandar,
	}

	return EstadisticasCurso
}

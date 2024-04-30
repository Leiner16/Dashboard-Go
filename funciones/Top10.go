package funcs

import (
	"sort"

	"github.com/derly/TallerGoWeb/modelos"
)

func Top10Mejor(data []modelos.Estudiante) modelos.TopCurso {

	algebraEstudiantes := make([]modelos.Estudiante, 0)
	calculoEstudiantes := make([]modelos.Estudiante, 0)
	programacionEstudiantes := make([]modelos.Estudiante, 0)
	ctdEstudiantes := make([]modelos.Estudiante, 0)

	for _, estudiante := range data {
		for _, curso := range estudiante.Cursos {
			if curso.Curso == "Algebra lineal" {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:    estudiante.Index,
					Nombre:   estudiante.Nombre,
					Apellido: estudiante.Apellido,
					Edad:     estudiante.Edad,
					Gender:   estudiante.Gender,
					Email:    estudiante.Email,
					Cursos:   []modelos.Curso{curso},
				}

				algebraEstudiantes = append(algebraEstudiantes, nuevoEstudiante)
			}
			if curso.Curso == "Calculo diferencial" {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:    estudiante.Index,
					Nombre:   estudiante.Nombre,
					Apellido: estudiante.Apellido,
					Edad:     estudiante.Edad,
					Gender:   estudiante.Gender,
					Email:    estudiante.Email,
					Cursos:   []modelos.Curso{curso},
				}

				calculoEstudiantes = append(calculoEstudiantes, nuevoEstudiante)
			}
			if curso.Curso == "POO" {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:    estudiante.Index,
					Nombre:   estudiante.Nombre,
					Apellido: estudiante.Apellido,
					Edad:     estudiante.Edad,
					Gender:   estudiante.Gender,
					Email:    estudiante.Email,
					Cursos:   []modelos.Curso{curso},
				}

				programacionEstudiantes = append(programacionEstudiantes, nuevoEstudiante)
			}
			if curso.Curso == "CTD" {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:    estudiante.Index,
					Nombre:   estudiante.Nombre,
					Apellido: estudiante.Apellido,
					Edad:     estudiante.Edad,
					Gender:   estudiante.Gender,
					Email:    estudiante.Email,
					Cursos:   []modelos.Curso{curso},
				}

				ctdEstudiantes = append(ctdEstudiantes, nuevoEstudiante)
			}
		}
	}

	sort.Slice(algebraEstudiantes, func(i, j int) bool {
		return algebraEstudiantes[i].Cursos[0].Nota > algebraEstudiantes[j].Cursos[0].Nota
	})
	sort.Slice(calculoEstudiantes, func(i, j int) bool {
		return calculoEstudiantes[i].Cursos[0].Nota > calculoEstudiantes[j].Cursos[0].Nota
	})
	sort.Slice(programacionEstudiantes, func(i, j int) bool {
		return programacionEstudiantes[i].Cursos[0].Nota > programacionEstudiantes[j].Cursos[0].Nota
	})
	sort.Slice(ctdEstudiantes, func(i, j int) bool {
		return ctdEstudiantes[i].Cursos[0].Nota > ctdEstudiantes[j].Cursos[0].Nota
	})

	var Algebra []modelos.Estudiante
	var Calculo []modelos.Estudiante
	var POO []modelos.Estudiante
	var CTD []modelos.Estudiante

	for i := 0; i < 10; i++ {
		Algebra = append(Algebra, algebraEstudiantes[i])
	}

	for i := 0; i < 10; i++ {
		Calculo = append(Calculo, calculoEstudiantes[i])
	}

	for i := 0; i < 10; i++ {
		POO = append(POO, programacionEstudiantes[i])
	}

	for i := 0; i < 10; i++ {
		CTD = append(CTD, ctdEstudiantes[i])
	}

	var TopMejor = modelos.TopCurso{
		Algebra: Algebra,
		Calculo: Calculo,
		POO:     POO,
		CTD:     CTD,
	}

	return TopMejor
}

func Top10Peor(data []modelos.Estudiante) modelos.TopCurso {

	algebraEstudiantes := make([]modelos.Estudiante, 0)
	calculoEstudiantes := make([]modelos.Estudiante, 0)
	programacionEstudiantes := make([]modelos.Estudiante, 0)
	ctdEstudiantes := make([]modelos.Estudiante, 0)

	for _, estudiante := range data {
		for _, curso := range estudiante.Cursos {
			if curso.Curso == "Algebra lineal" {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:    estudiante.Index,
					Nombre:   estudiante.Nombre,
					Apellido: estudiante.Apellido,
					Edad:     estudiante.Edad,
					Gender:   estudiante.Gender,
					Email:    estudiante.Email,
					Cursos:   []modelos.Curso{curso},
				}

				algebraEstudiantes = append(algebraEstudiantes, nuevoEstudiante)
			}
			if curso.Curso == "Calculo diferencial" {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:    estudiante.Index,
					Nombre:   estudiante.Nombre,
					Apellido: estudiante.Apellido,
					Edad:     estudiante.Edad,
					Gender:   estudiante.Gender,
					Email:    estudiante.Email,
					Cursos:   []modelos.Curso{curso},
				}

				calculoEstudiantes = append(calculoEstudiantes, nuevoEstudiante)
			}
			if curso.Curso == "POO" {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:    estudiante.Index,
					Nombre:   estudiante.Nombre,
					Apellido: estudiante.Apellido,
					Edad:     estudiante.Edad,
					Gender:   estudiante.Gender,
					Email:    estudiante.Email,
					Cursos:   []modelos.Curso{curso},
				}

				programacionEstudiantes = append(programacionEstudiantes, nuevoEstudiante)
			}
			if curso.Curso == "CTD" {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:    estudiante.Index,
					Nombre:   estudiante.Nombre,
					Apellido: estudiante.Apellido,
					Edad:     estudiante.Edad,
					Gender:   estudiante.Gender,
					Email:    estudiante.Email,
					Cursos:   []modelos.Curso{curso},
				}

				ctdEstudiantes = append(ctdEstudiantes, nuevoEstudiante)
			}
		}
	}

	var Algebra []modelos.Estudiante
	var Calculo []modelos.Estudiante
	var POO []modelos.Estudiante
	var CTD []modelos.Estudiante

	sort.Slice(algebraEstudiantes, func(i, j int) bool {
		return algebraEstudiantes[i].Cursos[0].Nota < algebraEstudiantes[j].Cursos[0].Nota
	})
	sort.Slice(calculoEstudiantes, func(i, j int) bool {
		return calculoEstudiantes[i].Cursos[0].Nota < calculoEstudiantes[j].Cursos[0].Nota
	})
	sort.Slice(programacionEstudiantes, func(i, j int) bool {
		return programacionEstudiantes[i].Cursos[0].Nota < programacionEstudiantes[j].Cursos[0].Nota
	})
	sort.Slice(ctdEstudiantes, func(i, j int) bool {
		return ctdEstudiantes[i].Cursos[0].Nota < ctdEstudiantes[j].Cursos[0].Nota
	})

	for i := 0; i < 10; i++ {
		Algebra = append(Algebra, algebraEstudiantes[i])
	}

	for i := 0; i < 10; i++ {
		Calculo = append(Calculo, calculoEstudiantes[i])
	}

	for i := 0; i < 10; i++ {
		POO = append(POO, programacionEstudiantes[i])
	}

	for i := 0; i < 10; i++ {
		CTD = append(CTD, ctdEstudiantes[i])
	}

	var TopMejor = modelos.TopCurso{
		Algebra: Algebra,
		Calculo: Calculo,
		POO:     POO,
		CTD:     CTD,
	}

	return TopMejor
}

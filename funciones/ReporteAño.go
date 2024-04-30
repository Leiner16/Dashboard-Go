package funcs

import "github.com/derly/TallerGoWeb/modelos"

func ReporteMateriaEstudiantesOutlider(data []modelos.Estudiante, numeroPagina int) []modelos.ReporteAño {

	EstAlgebra := EstadisticasCurso(data, "Algebra lineal")
	var EstPOO = EstadisticasCurso(data, "POO")
	var EstCalculo = EstadisticasCurso(data, "Calculo diferencial")
	var EstCTD = EstadisticasCurso(data, "CTD")

	var desvAlgebra = EstAlgebra.DesviacionEstandar
	var desvPOO = EstPOO.DesviacionEstandar
	var desvCalculo = EstCalculo.DesviacionEstandar
	var desvCTD = EstCTD.DesviacionEstandar

	algebraEstudiantes := make([]modelos.Estudiante, 0)
	calculoEstudiantes := make([]modelos.Estudiante, 0)
	programacionEstudiantes := make([]modelos.Estudiante, 0)
	ctdEstudiantes := make([]modelos.Estudiante, 0)

	algebra2014 := make([]modelos.Estudiante, 0)
	calculo2014 := make([]modelos.Estudiante, 0)
	programacion2014 := make([]modelos.Estudiante, 0)
	ctd2014 := make([]modelos.Estudiante, 0)

	algebra2015 := make([]modelos.Estudiante, 0)
	calculo2015 := make([]modelos.Estudiante, 0)
	programacion2015 := make([]modelos.Estudiante, 0)
	ctd2015 := make([]modelos.Estudiante, 0)

	algebra2016 := make([]modelos.Estudiante, 0)
	calculo2016 := make([]modelos.Estudiante, 0)
	programacion2016 := make([]modelos.Estudiante, 0)
	ctd2016 := make([]modelos.Estudiante, 0)

	algebra2017 := make([]modelos.Estudiante, 0)
	calculo2017 := make([]modelos.Estudiante, 0)
	programacion2017 := make([]modelos.Estudiante, 0)
	ctd2017 := make([]modelos.Estudiante, 0)

	algebra2018 := make([]modelos.Estudiante, 0)
	calculo2018 := make([]modelos.Estudiante, 0)
	programacion2018 := make([]modelos.Estudiante, 0)
	ctd2018 := make([]modelos.Estudiante, 0)

	algebra2019 := make([]modelos.Estudiante, 0)
	calculo2019 := make([]modelos.Estudiante, 0)
	programacion2019 := make([]modelos.Estudiante, 0)
	ctd2019 := make([]modelos.Estudiante, 0)

	algebra2020 := make([]modelos.Estudiante, 0)
	calculo2020 := make([]modelos.Estudiante, 0)
	programacion2020 := make([]modelos.Estudiante, 0)
	ctd2020 := make([]modelos.Estudiante, 0)

	algebra2021 := make([]modelos.Estudiante, 0)
	calculo2021 := make([]modelos.Estudiante, 0)
	programacion2021 := make([]modelos.Estudiante, 0)
	ctd2021 := make([]modelos.Estudiante, 0)

	algebra2022 := make([]modelos.Estudiante, 0)
	calculo2022 := make([]modelos.Estudiante, 0)
	programacion2022 := make([]modelos.Estudiante, 0)
	ctd2022 := make([]modelos.Estudiante, 0)

	for _, estudiante := range data {
		for _, curso := range estudiante.Cursos {
			if curso.Curso == "Algebra lineal" && (curso.Nota >= desvAlgebra*2 || curso.Nota <= desvAlgebra*-2) {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:       estudiante.Index,
					Nombre:      estudiante.Nombre,
					Apellido:    estudiante.Apellido,
					Edad:        estudiante.Edad,
					Gender:      estudiante.Gender,
					Email:       estudiante.Email,
					Matriculado: estudiante.Matriculado,
					Cursos:      []modelos.Curso{curso},
				}

				algebraEstudiantes = append(algebraEstudiantes, nuevoEstudiante)

				//clasificar matriz año
				if estudiante.Matriculado[:4] == "2014" {
					algebra2014 = append(algebra2014, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2015" {
					algebra2015 = append(algebra2015, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2016" {
					algebra2016 = append(algebra2016, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2017" {
					algebra2017 = append(algebra2017, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2018" {
					algebra2018 = append(algebra2018, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2019" {
					algebra2019 = append(algebra2019, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2020" {
					algebra2020 = append(algebra2020, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2021" {
					algebra2021 = append(algebra2021, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2022" {
					algebra2022 = append(algebra2022, nuevoEstudiante)
				}

			}
			if curso.Curso == "Calculo diferencial" && (curso.Nota >= desvCalculo*2 || curso.Nota <= desvCalculo*-2) {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:       estudiante.Index,
					Nombre:      estudiante.Nombre,
					Apellido:    estudiante.Apellido,
					Edad:        estudiante.Edad,
					Gender:      estudiante.Gender,
					Email:       estudiante.Email,
					Matriculado: estudiante.Matriculado,
					Cursos:      []modelos.Curso{curso},
				}

				calculoEstudiantes = append(calculoEstudiantes, nuevoEstudiante)

				if estudiante.Matriculado[:4] == "2014" {
					calculo2014 = append(calculo2014, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2015" {
					calculo2015 = append(calculo2015, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2016" {
					calculo2016 = append(calculo2016, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2017" {
					calculo2017 = append(calculo2017, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2018" {
					calculo2018 = append(calculo2018, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2019" {
					calculo2019 = append(calculo2019, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2020" {
					calculo2020 = append(calculo2020, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2021" {
					calculo2021 = append(calculo2021, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2022" {
					calculo2022 = append(calculo2022, nuevoEstudiante)
				}
			}
			if curso.Curso == "POO" && (curso.Nota >= desvPOO*2 || curso.Nota <= desvPOO*-2) {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:       estudiante.Index,
					Nombre:      estudiante.Nombre,
					Apellido:    estudiante.Apellido,
					Edad:        estudiante.Edad,
					Gender:      estudiante.Gender,
					Email:       estudiante.Email,
					Matriculado: estudiante.Matriculado,
					Cursos:      []modelos.Curso{curso},
				}

				programacionEstudiantes = append(programacionEstudiantes, nuevoEstudiante)

				if estudiante.Matriculado[:4] == "2014" {
					programacion2014 = append(programacion2014, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2015" {
					programacion2015 = append(programacion2015, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2016" {
					programacion2016 = append(programacion2016, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2017" {
					programacion2017 = append(programacion2017, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2018" {
					programacion2018 = append(programacion2018, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2019" {
					programacion2019 = append(programacion2019, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2020" {
					programacion2020 = append(programacion2020, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2021" {
					programacion2021 = append(programacion2021, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2022" {
					programacion2022 = append(programacion2022, nuevoEstudiante)
				}

			}
			if curso.Curso == "CTD" && (curso.Nota >= desvCTD*2 || curso.Nota <= desvCTD*-2) {
				//nuevo estudiante con el curso
				nuevoEstudiante := modelos.Estudiante{
					Index:       estudiante.Index,
					Nombre:      estudiante.Nombre,
					Apellido:    estudiante.Apellido,
					Edad:        estudiante.Edad,
					Gender:      estudiante.Gender,
					Email:       estudiante.Email,
					Matriculado: estudiante.Matriculado,
					Cursos:      []modelos.Curso{curso},
				}

				ctdEstudiantes = append(ctdEstudiantes, nuevoEstudiante)

				if estudiante.Matriculado[:4] == "2014" {
					ctd2014 = append(ctd2014, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2015" {
					ctd2015 = append(ctd2015, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2016" {
					ctd2016 = append(ctd2016, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2017" {
					ctd2017 = append(ctd2017, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2018" {
					ctd2018 = append(ctd2018, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2019" {
					ctd2019 = append(ctd2019, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2020" {
					ctd2020 = append(ctd2020, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2021" {
					ctd2021 = append(ctd2021, nuevoEstudiante)
				}
				if estudiante.Matriculado[:4] == "2022" {
					ctd2022 = append(ctd2022, nuevoEstudiante)
				}

			}
		}
	}

	var cursos2014 = modelos.TopCurso{
		Algebra: algebra2014,
		Calculo: calculo2014,
		POO:     programacion2014,
		CTD:     ctd2014,
	}

	var ReporteAño []modelos.ReporteAño
	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2014", cursos2014})

	var cursos2015 = modelos.TopCurso{
		Algebra: algebra2015,
		Calculo: calculo2015,
		POO:     programacion2015,
		CTD:     ctd2015,
	}

	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2015", cursos2015})

	var cursos2016 = modelos.TopCurso{
		Algebra: algebra2016,
		Calculo: calculo2016,
		POO:     programacion2016,
		CTD:     ctd2016,
	}

	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2016", cursos2016})

	var cursos2017 = modelos.TopCurso{
		Algebra: algebra2017,
		Calculo: calculo2017,
		POO:     programacion2017,
		CTD:     ctd2017,
	}

	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2017", cursos2017})

	var cursos2018 = modelos.TopCurso{
		Algebra: algebra2018,
		Calculo: calculo2018,
		POO:     programacion2018,
		CTD:     ctd2018,
	}

	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2018", cursos2018})

	var cursos2019 = modelos.TopCurso{
		Algebra: algebra2019,
		Calculo: calculo2019,
		POO:     programacion2019,
		CTD:     ctd2019,
	}

	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2019", cursos2019})

	var cursos2020 = modelos.TopCurso{
		Algebra: algebra2020,
		Calculo: calculo2020,
		POO:     programacion2020,
		CTD:     ctd2020,
	}

	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2020", cursos2020})

	var cursos2021 = modelos.TopCurso{
		Algebra: algebra2021,
		Calculo: calculo2021,
		POO:     programacion2021,
		CTD:     ctd2021,
	}

	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2021", cursos2021})

	var cursos2022 = modelos.TopCurso{
		Algebra: algebra2022,
		Calculo: calculo2022,
		POO:     programacion2022,
		CTD:     ctd2022,
	}

	ReporteAño = append(ReporteAño, modelos.ReporteAño{"2022", cursos2022})

	return ReporteAño

}

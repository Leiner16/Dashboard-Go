package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"sort"

	funcs "github.com/derly/TallerGoWeb/funciones"
	"github.com/derly/TallerGoWeb/modelos"
)

func Index(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprint(w, "<h1> Pagina no encontrada <h1>")
		panic(err)
	} else {
		template.Execute(w, "personas")
	}
}

func MejorEstudiante(w http.ResponseWriter, r *http.Request) {
	var data []modelos.Estudiante = CargaDatos()
	var estudiante modelos.Estudiante = funcs.MejorPromedio(data)

	template, err := template.ParseFiles("templates/mejorPromedio.html")
	if err != nil {
		fmt.Fprint(w, "<h1> Pagina no encontrada <h1>")
		panic(err)
	} else {
		template.Execute(w, estudiante)
	}
}
func PeorEstudiante(w http.ResponseWriter, r *http.Request) {
	var data []modelos.Estudiante = CargaDatos()
	var estudiante modelos.Estudiante = funcs.PeorPromedio(data)

	template, err := template.ParseFiles("templates/peorPromedio.html")
	if err != nil {
		fmt.Fprint(w, "<h1> Pagina no encontrada <h1>")
		panic(err)
	} else {
		template.Execute(w, estudiante)
	}
}

func Top10MejoresPromedios(w http.ResponseWriter, r *http.Request) {
    curso := r.URL.Query().Get("curso") // Obtiene el curso desde el parámetro de la URL

	if curso == "" {
        curso = "Algebra lineal" // Valor por defecto
    }

    // Carga de datos de estudiantes
    var data []modelos.Estudiante = CargaDatos()
    var estudiantesFiltrados []modelos.Estudiante

    // Filtrar estudiantes que están matriculados en el curso seleccionado
    for _, estudiante := range data {
        for _, c := range estudiante.Cursos {
            if c.Curso == curso {
                estudiante.Cursos = []modelos.Curso{c} // Reten solo el curso relevante
                estudiantesFiltrados = append(estudiantesFiltrados, estudiante)
                break
            }
        }
    }

    // Ordenar los estudiantes filtrados por nota, asegurando que se usa la nota del curso relevante
    sort.Slice(estudiantesFiltrados, func(i, j int) bool {
        return estudiantesFiltrados[i].Cursos[0].Nota > estudiantesFiltrados[j].Cursos[0].Nota
    })

    // Si hay más de 10 estudiantes, recortar el slice a los 10 primeros
    if len(estudiantesFiltrados) > 10 {
        estudiantesFiltrados = estudiantesFiltrados[:10]
    }

    // Preparar y ejecutar el template con los estudiantes filtrados
    template, err := template.ParseFiles("templates/top10Mejor.html")
    if err != nil {
        http.Error(w, "Página no encontrada", http.StatusNotFound)
        return
    }
    template.Execute(w, estudiantesFiltrados)
}

func Top10PeoresPromedios(w http.ResponseWriter, r *http.Request) {
    curso := r.URL.Query().Get("curso") // Obtiene el curso desde el parámetro de la URL

	if curso == "" {
        curso = "Algebra lineal" // Valor por defecto
    }

    var data []modelos.Estudiante = CargaDatos()
    var estudiantesFiltrados []modelos.Estudiante

    // Filtrar estudiantes que están matriculados en el curso seleccionado
    for _, estudiante := range data {
        for _, c := range estudiante.Cursos {
            if c.Curso == curso {
                estudiante.Cursos = []modelos.Curso{c} // Reten solo el curso relevante
                estudiantesFiltrados = append(estudiantesFiltrados, estudiante)
                break
            }
        }
    }

    // Ordenar los estudiantes filtrados por nota, de menor a mayor
    sort.Slice(estudiantesFiltrados, func(i, j int) bool {
        return estudiantesFiltrados[i].Cursos[0].Nota < estudiantesFiltrados[j].Cursos[0].Nota
    })

    // Si hay más de 10 estudiantes, recortar el slice a los 10 primeros
    if len(estudiantesFiltrados) > 10 {
        estudiantesFiltrados = estudiantesFiltrados[:10]
    }

    // Preparar y ejecutar el template con los estudiantes filtrados
    template, err := template.ParseFiles("templates/top10Peor.html")
    if err != nil {
        http.Error(w, "Página no encontrada", http.StatusNotFound)
        return
    }
    template.Execute(w, estudiantesFiltrados)
}


func EstudiantesMatriculados2022(w http.ResponseWriter, r *http.Request) {
    var data []modelos.Estudiante = CargaDatos()
    
    // Obtener los parámetros de query
    nombre := r.URL.Query().Get("nombre")
    apellido := r.URL.Query().Get("apellido")

    // Filtrar estudiantes por nombre y apellido si se proporcionan
    var estudiantesFiltrados []modelos.Estudiante
    for _, estudiante := range data {
        if estudiante.Matriculado[:4] == "2022" &&
           (nombre == "" || strings.Contains(strings.ToLower(estudiante.Nombre), strings.ToLower(nombre))) &&
           (apellido == "" || strings.Contains(strings.ToLower(estudiante.Apellido), strings.ToLower(apellido))) {
            estudiantesFiltrados = append(estudiantesFiltrados, estudiante)
        }
    }

    template, err := template.ParseFiles("templates/estudiantes2022.html")
    if err != nil {
        fmt.Fprint(w, "<h1> Página no encontrada <h1>")
        panic(err)
    } else {
        template.Execute(w, estudiantesFiltrados)
    }
}


func EstudianteMasculinoMayor(w http.ResponseWriter, r *http.Request) {
	var data []modelos.Estudiante = CargaDatos()
	var estudiante []modelos.Estudiante = funcs.MayorEdadHombres(data)

	template, err := template.ParseFiles("templates/EstudianteMasculinoMayorEdad.html")
	if err != nil {
		fmt.Fprint(w, "<h1> Pagina no encontrada <h1>")
		panic(err)
	} else {
		template.Execute(w, estudiante)
	}
}
func EstudianteFemeninoMayor(w http.ResponseWriter, r *http.Request) {
	var data []modelos.Estudiante = CargaDatos()
	var estudiante []modelos.Estudiante = funcs.MayorEdadMujeres(data)

	template, err := template.ParseFiles("templates/EstudianteFemeninoMayorEdad.html")
	if err != nil {
		fmt.Fprint(w, "<h1> Pagina no encontrada <h1>")
		panic(err)
	} else {
		template.Execute(w, estudiante)
	}
}

func CalculosPorCurso(w http.ResponseWriter, r *http.Request) {
	var data []modelos.Estudiante = CargaDatos()
	var estadisticas []modelos.EstadisticasCurso = funcs.CursosPromedioVarianzaDesviacionRango(data)

	template, err := template.ParseFiles("templates/EstadisticasCurso.html")
	if err != nil {
		fmt.Fprint(w, "<h1> Pagina no encontrada <h1>")
		panic(err)
	} else {
		template.Execute(w, estadisticas)
	}
}
func PromedioRangoEdad(w http.ResponseWriter, r *http.Request) {
	var data []modelos.Estudiante = CargaDatos()
	var promedios []modelos.RangoPromedio = funcs.PromedioTodosRangosEdad(data)

	template, err := template.ParseFiles("templates/PromedioRangoEdad.html")
	if err != nil {
		fmt.Fprint(w, "<h1> Pagina no encontrada <h1>")
		panic(err)
	} else {
		template.Execute(w, promedios)
	}
}

func ReporteAño(w http.ResponseWriter, r *http.Request) {
	var data []modelos.Estudiante = CargaDatos()
	var reporte []modelos.ReporteAño = funcs.ReporteMateriaEstudiantesOutlider(data, 2)

	template, err := template.ParseFiles("templates/ReporteAños.html")
	if err != nil {
		fmt.Fprint(w, "<h1> Pagina no encontrada <h1>")
		panic(err)
	} else {
		template.Execute(w, reporte)
	}
}

func CargaDatos() []modelos.Estudiante {
	file, err := os.Open("generated.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var data []modelos.Estudiante
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return data
}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/mejor-promedio", MejorEstudiante)
	http.HandleFunc("/peor-promedio", PeorEstudiante)

	http.HandleFunc("/top10-mejor", Top10MejoresPromedios)
	http.HandleFunc("/top10-peores", Top10PeoresPromedios)

	http.HandleFunc("/estudiante-masculino-mayor", EstudianteMasculinoMayor)
	http.HandleFunc("/estudiante-femenino-mayor", EstudianteFemeninoMayor)

	http.HandleFunc("/estadisticas-cursos", CalculosPorCurso)

	http.HandleFunc("/promedio-rango-edad", PromedioRangoEdad)

	http.HandleFunc("/matriculados/estudiantes-2022", EstudiantesMatriculados2022)

	http.HandleFunc("reporte", ReporteAño)

	fmt.Println("Escuchando en http://localhost:8080..")
	http.ListenAndServe("localhost:8080", nil)

}

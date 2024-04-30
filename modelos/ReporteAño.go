package modelos

type ReporteAño struct {
	Año    string   `json:"año"`
	Cursos TopCurso `json:"cursos"`
}

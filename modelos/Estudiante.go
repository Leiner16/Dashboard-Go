package modelos

type Estudiante struct {
	Index       int     `json:"index"`
	Nombre      string  `json:"nombre"`
	Apellido    string  `json:"apellido"`
	Edad        int     `json:"edad"`
	Gender      string  `json:"gender"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	About       string  `json:"about"`
	Matriculado string  `json:"matriculado"`
	Cursos      []Curso `json:"cursos"`
}

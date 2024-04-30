package modelos

type Curso struct {
	ID    int     `json:"id"`
	Curso string  `json:"curso"`
	Nota  float64 `json:"nota"`
}

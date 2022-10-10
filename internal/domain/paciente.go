package domain

type Paciente struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Domicilio string `json:"domicilio"`
	Dni       int    `json:"dni"`
	FechaAlta string `json:"fecha_de_alta"`
}

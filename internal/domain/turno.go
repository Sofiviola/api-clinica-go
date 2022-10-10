package domain

type Turno struct {
	Id          int    `json:"id"`
	DentistaId  int    `json:"dentista_id"`
	PacienteId  int    `json:"paciente_id"`
	Fecha       string `json:"fecha"`
	Hora        string `json:"hora"`
	Descripcion string `json:"descripcion"`
}

type TurnoDTO struct {
	Id          int      `json:"id"`
	Dentista    Dentista `json:"dentista"`
	Paciente    Paciente `json:"paciente"`
	Fecha       string   `json:"fecha"`
	Hora        string   `json:"hora"`
	Descripcion string   `json:"descripcion"`
}

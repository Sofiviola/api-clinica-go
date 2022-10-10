package turnoStore

import "github.com/Sofiviola/Examen.git/internal/domain"

type StoreInterface interface {
	Read(id int) (domain.TurnoDTO, error)
	Create(t domain.TurnoDTO) (  error)
	Update(id int,turno domain.Turno) error
	Delete(id int) error
	GetAll() ([]domain.Turno, error)
	ExistTurno(paciente_id int, dentista_id int) bool
	ReadDniPaciente(dni int) (domain.TurnoDTO, error)
}

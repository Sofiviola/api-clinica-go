package pacienteStore

import "github.com/Sofiviola/Examen.git/internal/domain"

type StoreInterface interface {
	Read(id int) (domain.Paciente, error)
	Create(product domain.Paciente) error
	Update(id int,product domain.Paciente) error
	Delete(id int) error
	GetAll() ([]domain.Paciente, error)
	Exists(dni int) bool
	ReadDni(dni int) (domain.Paciente, error)
}

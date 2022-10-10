package turno

import (
	"errors"

	"github.com/Sofiviola/Examen.git/internal/domain"
	"github.com/Sofiviola/Examen.git/pkg/store/turnoStore"
)

type Repository interface {
	GetByID(id int) (domain.TurnoDTO, error)
	Create(p domain.TurnoDTO) (domain.TurnoDTO, error)
	Update(id int, u domain.Turno) (domain.Turno, error)
	Delete(id int) error
	GetByDniPaciente(dni int) (domain.TurnoDTO, error) 
}

type repository struct {
	storage turnoStore.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage turnoStore.StoreInterface) Repository {
	return &repository{
		storage: storage,
	}
}


// GetByID busca un Turno por su id
func (r *repository) GetByID(id int) (domain.TurnoDTO, error) {
	turno, err := r.storage.Read(id)
	if err != nil {
		return domain.TurnoDTO{}, errors.New("Turno not found")
	}
	return turno, nil

}

// GetByID busca un Turno por su id
func (r *repository) GetByDniPaciente(dni int) (domain.TurnoDTO, error) {
	turno, err := r.storage.ReadDniPaciente(dni)
	if err != nil {
		return domain.TurnoDTO{}, errors.New("Turno not found")
	}
	return turno, nil

}


// Create agrega un nuevo Turno
func (r *repository) Create(d domain.TurnoDTO) (domain.TurnoDTO, error) {
	if r.storage.ExistTurno(d.Paciente.Id,d.Dentista.Id) {
		return domain.TurnoDTO{}, errors.New("Turno ya existe, solo se puede tener un turno por persona")
	}
	err := r.storage.Create(d)
	if err != nil {
		return domain.TurnoDTO{}, errors.New("Error al crear un turno")
	}
	return d, nil
	
}

// Delete elimina un Turno
func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Turno
func (r *repository) Update(id int, u domain.Turno) (domain.Turno, error) {
	u.Id = id
	err := r.storage.Update(id,u)
	if err != nil {
		return domain.Turno{}, errors.New("error updating Turno")
	}
	return u, nil
}


package paciente

import (
	"errors"

	"github.com/Sofiviola/Examen.git/internal/domain"
	"github.com/Sofiviola/Examen.git/pkg/store/pacienteStore"
)

type Repository interface {
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
	GetByDni(dni int) (domain.Paciente, error)
}

type repository struct {
	storage pacienteStore.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage pacienteStore.StoreInterface) Repository {
	return &repository{
		storage: storage,
	}
}


// GetByID busca un Paciente por su id
func (r *repository) GetByID(id int) (domain.Paciente, error) {
	paciente, err := r.storage.Read(id)
	if err != nil {
		return domain.Paciente{}, errors.New("Paciente not found")
	}
	return paciente, nil

}

// GetByDni busca un Paciente por su dni
func (r *repository) GetByDni(dni int) (domain.Paciente, error) {
	paciente, err := r.storage.ReadDni(dni)
	if err != nil {
		return domain.Paciente{}, errors.New("Paciente not found repository")
	}
	return paciente, nil

}

// Create agrega un nuevo Paciente
func (r *repository) Create(d domain.Paciente) (domain.Paciente, error) {
	if r.storage.Exists(d.Dni) {
		return domain.Paciente{}, errors.New("Paciente already exists")
	}
	err := r.storage.Create(d)
	if err != nil {
		return domain.Paciente{}, errors.New("error creating Paciente")
	}
	return d, nil
}

// Delete elimina un Paciente
func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Paciente
func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	p.Id = id
	err := r.storage.Update(id,p)
	if err != nil {
		return domain.Paciente{}, errors.New("error updating Paciente")
	}
	return p, nil
}

// validateCodeValue valida que la dni de un paciente no se repita
func (r *repository) validateCodeValue(dni int) bool {
	list, err := r.storage.GetAll()
	if err != nil {
		return false
	}
	for _, paciente := range list {
		if paciente.Dni == dni {
			return false
		}
	}
	return true
}
package dentista

import (
	"errors"

	"github.com/Sofiviola/Examen.git/internal/domain"
	"github.com/Sofiviola/Examen.git/pkg/store/dentistaStore"
)

type Repository interface {
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Update(id int,p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	GetByMatricula(matricula string) (domain.Dentista, error)
}

type repository struct {
	storage dentistaStore.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage dentistaStore.StoreInterface) Repository {
	return &repository{
		storage: storage,
	}
}


// GetByID busca un Dentista por su id
func (r *repository) GetByID(id int) (domain.Dentista, error) {
	dentista, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, errors.New("Dentista not found")
	}
	return dentista, nil

}

// GetByMatricula busca un Dentista por su Matricula
func (r *repository) GetByMatricula(matricula string) (domain.Dentista, error) {
	dentista, err := r.storage.ReadMatricula(matricula)
	if err != nil {
		return domain.Dentista{}, errors.New("Dentista not found")
	}
	return dentista, nil

}


// Create agrega un nuevo Dentista
func (r *repository) Create(d domain.Dentista) (domain.Dentista, error) {
	if r.storage.Exists(d.Matricula) {
		return domain.Dentista{}, errors.New("Matricula already exists")
	}
	err := r.storage.Create(d)
	if err != nil {
		return domain.Dentista{}, errors.New("error creating Dentista")
	}
	return d, nil
}

// Delete elimina un Dentista
func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Dentista
func (r *repository) Update(id int,p domain.Dentista) (domain.Dentista, error) {
	err := r.storage.Update(id,p)
	if err != nil {
		return domain.Dentista{}, errors.New("error updating Dentista")
	}
	return p, nil
}

// validateCodeValue valida que la matricula de un dentista no se repita
func (r *repository) validateCodeValue(matricula string) bool {
	list, err := r.storage.GetAll()
	if err != nil {
		return false
	}
	for _, dentista := range list {
		if dentista.Matricula == matricula {
			return false
		}
	}
	return true
}
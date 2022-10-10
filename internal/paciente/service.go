package paciente

import (
	"errors"
	"github.com/Sofiviola/Examen.git/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
	GetByDni(dni int) (domain.Paciente, error)
	Patch(id int, u domain.Paciente) (domain.Paciente, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetByID busca un Paciente por su id
func (s *service) GetByID(id int) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

// GetByDni busca un Paciente por su dni
func (s *service) GetByDni(dni int) (domain.Paciente, error) {
	p, err := s.r.GetByDni(dni)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

// Create agrega un nuevo Paciente
func (s *service) Create(p domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

// Delete elimina un Paciente
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Paciente
func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Update(id, u)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

// Actualizacion parcial paciente
func (s *service) Patch(id int, u domain.Paciente) (domain.Paciente, error) {
	 p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{},  errors.New("No se encontro el Paciente")
	}

	update := domain.Paciente{
		    Id: p.Id,
		    Nombre:        p.Nombre,
			Apellido:    p.Apellido,
			Domicilio:  p.Domicilio,
			Dni:   p.Dni,
			FechaAlta:  p.FechaAlta,
	}

       if update.Id == id {
		switch {
		case   u.Nombre != "":
			update.Nombre = u.Nombre
		
		case u.Apellido != "":
			update.Apellido = u.Apellido

		case  u.Domicilio != "":
			update.Domicilio = u.Domicilio
        
		case u.Dni > 0:
			update.Dni = u.Dni
		
		case u.FechaAlta != "":
			update.FechaAlta = u.FechaAlta
		}
    }

	a, err := s.r.Update(id, update)
	if err != nil {
		return domain.Paciente{}, errors.New("No se pudo aplicar la actualizacion parcial")
	}
	return a, nil
}

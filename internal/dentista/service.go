package dentista

import (
	"github.com/Sofiviola/Examen.git/internal/domain"
    "errors"
)

type Service interface {
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	GetByMatricula(matricula string) (domain.Dentista, error)
	Patch(id int, u domain.Dentista) (domain.Dentista, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetByID busca un Dentista por su id
func (s *service) GetByID(id int) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

// GetByMatricula busca un Dentista por su id
func (s *service) GetByMatricula(matricula string) (domain.Dentista, error) {
	p, err := s.r.GetByMatricula(matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}


// Create agrega un nuevo Dentista
func (s *service) Create(p domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

// Delete elimina un Dentista
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Dentista
func (s *service) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.Update(id, u)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

// Actualizacion parcial paciente
func (s *service) Patch(id int, u domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
   if err != nil {
	   return domain.Dentista{},  errors.New("No se encontro el Dentista")
   }

   update := domain.Dentista{
		   Id: p.Id,
		   Nombre:        p.Nombre,
		   Apellido:    p.Apellido,
		   Matricula:  p.Matricula,
   }

	  if update.Id == id {
	   switch {
	   case   u.Nombre != "":
		   update.Nombre = u.Nombre
	   
	   case u.Apellido != "":
		   update.Apellido = u.Apellido

	   case  u.Matricula != "":
		   update.Matricula = u.Matricula
	   
	   }
   }

   a, err := s.r.Update(id, update)
   if err != nil {
	   return domain.Dentista{}, errors.New("No se pudo aplicar la actualizacion parcial")
   }
   return a, nil
}


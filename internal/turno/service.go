package turno

import (
	"errors"
	"github.com/Sofiviola/Examen.git/internal/domain"
	"github.com/Sofiviola/Examen.git/internal/dentista"
	"github.com/Sofiviola/Examen.git/internal/paciente"
)

type Service interface {
	GetByID(id int) (domain.TurnoDTO, error)
	Create(p domain.TurnoDTO) (domain.TurnoDTO,error)
	Update(id int, u domain.Turno) (domain.Turno, error)
	Delete(id int) error
	GetByDniPaciente(dni int) (domain.TurnoDTO, error) 
	Patch(id int, u domain.Turno) (domain.Turno, error)
}

type service struct {
	r Repository
	repositoryPaciente paciente.Repository
	repositoryDentista dentista.Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository, repositoryPaciente paciente.Repository,repositoryDentista dentista.Repository ) Service {
	return &service{r, repositoryPaciente, repositoryDentista}
}

// GetByID busca un Turno por su id
func (s *service) GetByID(id int) (domain.TurnoDTO, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.TurnoDTO{}, err
	}
	return p, nil
}

//busca un Turno por su dni paciente  
func (s *service) GetByDniPaciente(dni int) (domain.TurnoDTO, error) {
p, err := s.r.GetByDniPaciente(dni)
if err != nil {
	return domain.TurnoDTO{}, err
}
return p, nil
}

// Create agrega un nuevo Turno
func (s *service) Create(p domain.TurnoDTO) ( domain.TurnoDTO,error) {
	dentista, err := s.repositoryDentista.GetByMatricula(p.Dentista.Matricula)
	if err != nil {
		return domain.TurnoDTO{}, errors.New("El dentista no existe")
	}

	paciente, err := s.repositoryPaciente.GetByDni(p.Paciente.Dni)
	if err != nil {
		return domain.TurnoDTO{}, errors.New("El paciente no existe")
	}

	//Obtengo el id del dentista y paciente
	var turnoDTO domain.TurnoDTO
	turnoDTO.Dentista = dentista
	turnoDTO.Paciente = paciente
	turnoDTO.Fecha = p.Fecha
	turnoDTO.Hora = p.Hora
	turnoDTO.Descripcion = p.Descripcion

	a,err := s.r.Create(turnoDTO)
	if err != nil {
		return   domain.TurnoDTO{},err
	}
	return a,nil
}

// Delete elimina un Turno
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Turno
func (s *service) Update(id int, u domain.Turno) (domain.Turno, error) {
	p, err := s.r.Update(id, u)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
	
}

// Patch actualiza un Turno
func (s *service) Patch(id int, u domain.Turno) (domain.Turno, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{},  errors.New("No se encontro el turno")
	}

	update := domain.Turno{
		Id: p.Id,
		DentistaId  :p.Dentista.Id,
		PacienteId  :p.Paciente.Id,
		Fecha       :p.Fecha,
		Hora       :p.Hora,
		Descripcion :p.Descripcion,
	}

       if update.Id == id {
		switch {
		case   u.DentistaId > 0:
			update.DentistaId = u.DentistaId
		
		case u.PacienteId > 0:
			update.PacienteId = u.PacienteId

		case  u.Fecha != "":
			update.Fecha = u.Fecha
        
		case u.Hora != "":
			update.Hora = u.Hora
		
		case u.Descripcion != "":
			update.Descripcion = u.Descripcion
		}
    }
	a, err := s.r.Update(id, update)
	if err != nil {
		return domain.Turno{}, errors.New("No se pudo aplicar la actualizacion parcial")
	}
	return a, nil
}
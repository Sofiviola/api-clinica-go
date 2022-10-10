package handler

import (
	"errors"
	"strconv"
	_ "strings"

	"github.com/Sofiviola/Examen.git/internal/domain"
	"github.com/Sofiviola/Examen.git/internal/turno"
	"github.com/Sofiviola/Examen.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.Service
}

// NewTurnoHandler crea un nuevo controller de turnos
func NewTurnoHandler(s turno.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// GetByID obtiene un turno por su id
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		turno, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turno not found"))
			return
		}
		web.Success(c, 200, turno)
	}
}

// GetByDniPaciente obtiene un turno por su id
func (h *turnoHandler) GetByDniPaciente() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid dni"))
			return
		}
		turno, err := h.s.GetByDniPaciente(dni)
		if err != nil {
			web.Failure(c, 404, errors.New("turno not found"))
			return
		}
		web.Success(c, 200, turno)
	}
}

func (h *turnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.TurnoDTO
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		valid, err := validateEmptysTurno(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un Turno
func (h *turnoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un Turno
func (h *turnoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var turno domain.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		valid, err := validateEmptysTurnoUpdate(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *turnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		DentistaId  int    `json:"dentista_id,omitempty"`
		PacienteId  int    `json:"paciente_id,omitempty"`
		Fecha       string `json:"fecha,omitempty"`
		Hora        string `json:"hora,omitempty"`
		Descripcion string `json:"descripcion,omitempty"`
	}

	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}

		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Turno{
			DentistaId:  r.DentistaId,
			PacienteId:  r.PacienteId,
			Fecha:       r.Fecha,
			Hora:        r.Hora,
			Descripcion: r.Descripcion,
		}

		p, err := h.s.Patch(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// validateEmptysTurno valida que los campos no esten vacios para el post
func validateEmptysTurno(turno *domain.TurnoDTO) (bool, error) {
	switch {
	case turno.Fecha == "" || turno.Hora == "" || turno.Descripcion == "" || turno.Dentista.Matricula == "" || turno.Paciente.Dni <= 0:
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// validateEmptysTurnoUpdate valida que los campos no esten vacios para el update
func validateEmptysTurnoUpdate(turno *domain.Turno) (bool, error) {
	switch {
	case turno.DentistaId <= 0 || turno.PacienteId <= 0 || turno.Fecha == "" || turno.Hora == "" || turno.Descripcion == "":
		return false, errors.New("De completar la Fecha, hora, descricion e id de paciente y dentista para actualizar. ")
	}
	return true, nil
}

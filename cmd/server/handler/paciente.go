package handler

import (
	"errors"
	"strconv"
	_"strings"

	"github.com/Sofiviola/Examen.git/internal/domain"
	"github.com/Sofiviola/Examen.git/internal/paciente"
	"github.com/Sofiviola/Examen.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

// NewPacienteHandler crea un nuevo controller de pacientes
func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// GetByID obtiene un paciente por su id
func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("paciente not found"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

// GetByID obtiene un paciente por su id
func (h *pacienteHandler) GetByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		paciente, err := h.s.GetByDni(dni)
		if err != nil {
			web.Failure(c, 404, errors.New("Error: paciente not found"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente
		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysPaciente(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un Paciente
func (h *pacienteHandler) Delete() gin.HandlerFunc {
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

// Put actualiza un Paciente
func (h *pacienteHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysPaciente(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nombre       string  `json:"nombre,omitempty"`
		Apellido    string     `json:"apellido,omitempty"`
		Domicilio   string  `json:"domicilio,omitempty"`
		Dni          int  `json:"dni,omitempty"`
		FechaAlta   string  `json:"fechaAlta,omitempty"`
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
		update := domain.Paciente{
			Nombre:        r.Nombre,
			Apellido:    r.Apellido,
			Domicilio:   r.Domicilio,
			Dni:   r.Dni,
			FechaAlta:   r.FechaAlta,
		}
		p, err := h.s.Patch(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// validateEmptysaciente valida que los campos no esten vacios
func validateEmptysPaciente(paciente *domain.Paciente) (bool, error) {
	switch {
	case paciente.Nombre == "" || paciente.Apellido == "" || paciente.Domicilio == "" || paciente.Dni  < 0 || paciente.FechaAlta == "" :
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}


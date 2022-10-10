package handler

import (
	"errors"
	"strconv"

	"github.com/Sofiviola/Examen.git/internal/domain"
	"github.com/Sofiviola/Examen.git/internal/dentista"
	"github.com/Sofiviola/Examen.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	s dentista.Service
}

// NewDentistaHandler crea un nuevo controller de dentistas
func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}

// GetByID obtiene un dentista por su id
func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentista, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}
		web.Success(c, 200, dentista)
	}
}

// GetByID obtiene un dentista por su id
func (h *dentistaHandler) GetByMatricula() gin.HandlerFunc {
	return func(c *gin.Context) {
		matriculaParam := c.Param("matricula")
		dentista, err := h.s.GetByMatricula(matriculaParam)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}
		web.Success(c, 200, dentista)
	}
}

func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentista domain.Dentista
		err := c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysDentista(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(dentista)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un Dentista
func (h *dentistaHandler) Delete() gin.HandlerFunc {
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

// Put actualiza un Dentista
func (h *dentistaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var dentista domain.Dentista
		err = c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysDentista(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, dentista)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *dentistaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nombre       string  `json:"nombre,omitempty"`
		Apellido    string     `json:"apellido,omitempty"`
		Matricula   string  `json:"matricula,omitempty"`
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
		update := domain.Dentista{
			Nombre:        r.Nombre,
			Apellido:    r.Apellido,
			Matricula:   r.Matricula,
		}
		p, err := h.s.Patch(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// validateEmptysDentista valida que los campos no esten vacios
func validateEmptysDentista(dentista *domain.Dentista) (bool, error) {
	switch {
	case dentista.Nombre == " " || dentista.Apellido == " " || dentista.Matricula == " ":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}


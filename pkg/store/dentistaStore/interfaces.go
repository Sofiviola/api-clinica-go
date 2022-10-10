package dentistaStore

import "github.com/Sofiviola/Examen.git/internal/domain"

type StoreInterface interface {
	Read(id int) (domain.Dentista, error)
	Create(product domain.Dentista) error
	Update(id int,product domain.Dentista) error
	Delete(id int) error
	GetAll() ([]domain.Dentista, error)
	Exists(matricula string) bool
	ReadMatricula(matricula string) (domain.Dentista, error)
}

package dentistaStore

import (
	"database/sql"
	"github.com/Sofiviola/Examen.git/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

// Read devuelve un Dentista por su id
func (s *sqlStore) Read(id int) (domain.Dentista, error) {
	var dentista domain.Dentista
	row := s.db.QueryRow("select * from dentista where id = ?", id)
	err := row.Scan(&dentista.Id, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dentista, nil
}

// Read devuelve un Dentista por su matricula
func (s *sqlStore) ReadMatricula(matricula string) (domain.Dentista, error) {
	var dentista domain.Dentista
	row := s.db.QueryRow("select * from dentista where Matricula = ?", matricula)
	err := row.Scan(&dentista.Id, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dentista, nil
}



// Create agrega un nuevo Dentista
func (s *sqlStore) Create(dentista domain.Dentista) error {
	query := "insert into dentista (nombre, apellido, matricula) values (?, ?, ?)"
	st, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(dentista.Nombre, dentista.Apellido, dentista.Matricula)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Dentista
func (s *sqlStore) Update(id int,dentista domain.Dentista) error {
	query := "UPDATE dentista SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&dentista.Nombre, &dentista.Apellido, &dentista.Matricula,id)
	if err != nil {
		return err
	}
	return nil
}

// Delete elimina un Dentista
func (s *sqlStore) Delete(id int) error {
	stmt := "delete from dentista where id = ?"
	_, err := s.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// Trae todos los Dentista
func (s *sqlStore) GetAll() ([]domain.Dentista, error) {
	list := []domain.Dentista{}
	rows, err := s.db.Query("select * from dentista")
	if err != nil {
		return list, err
	}

	if rows.Err() != nil {
		return list, rows.Err()
	}

	for rows.Next() {
		var dentista domain.Dentista
		err := rows.Scan(&dentista.Id, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
		if err != nil {
			return nil, err
		}
		list = append(list, dentista)
	}

	return list, nil
}

// Exists verifica si un Dentista existe
func (s *sqlStore) Exists(matricula string) bool {
	var id int
	row := s.db.QueryRow("select id from Dentista where matricula = ?", matricula)
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}

	return false
}
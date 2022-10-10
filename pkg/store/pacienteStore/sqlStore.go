package pacienteStore

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

// Read devuelve un Paciente por su id
func (s *sqlStore) Read(id int) (domain.Paciente, error) {
	var paciente domain.Paciente
	row := s.db.QueryRow("select * from paciente where id = ?", id)
	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio,&paciente.Dni,&paciente.FechaAlta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

// Read devuelve un Paciente por su dni
func (s *sqlStore) ReadDni(dni int) (domain.Paciente, error) {
	var paciente domain.Paciente
	row := s.db.QueryRow("select * from paciente where dni = ?", dni)
	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio,&paciente.Dni,&paciente.FechaAlta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

// Create agrega un nuevo Paciente
func (s *sqlStore) Create(paciente domain.Paciente) error {
	query := "insert into paciente (nombre, apellido, domicilio,dni,fecha_de_alta) values (?,?,?,?,?)"
	st, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio,&paciente.Dni,paciente.FechaAlta)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Paciente
func (s *sqlStore) Update(id int,paciente domain.Paciente) error {
	stmt, err := s.db.Prepare("UPDATE paciente SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_de_alta=? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&paciente.Nombre, &paciente.Apellido, &paciente.Domicilio,&paciente.Dni,&paciente.FechaAlta,id)
	if err != nil {
		return err
	}
	return nil
}

// Delete elimina un Paciente
func (s *sqlStore) Delete(id int) error {
	stmt := "delete from paciente where id = ?"
	_, err := s.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// Trae todos los Paciente
func (s *sqlStore) GetAll() ([]domain.Paciente, error) {
	list := []domain.Paciente{}
	rows, err := s.db.Query("select * from paciente")
	if err != nil {
		return list, err
	}

	if rows.Err() != nil {
		return list, rows.Err()
	}

	for rows.Next() {
		var paciente domain.Paciente
		err := rows.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio,&paciente.Dni,&paciente.FechaAlta )
		if err != nil {
			return nil, err
		}
		list = append(list, paciente)
	}

	return list, nil
}

// Exists verifica si un Paciente existe
func (s *sqlStore) Exists(dni int) bool {
	var id int
	row := s.db.QueryRow("select id from paciente where dni = ?", dni)
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}

	return false
}
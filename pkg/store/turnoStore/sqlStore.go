package turnoStore

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

// Read devuelve un Turno por su id
func (s *sqlStore) Read(id int) (domain.TurnoDTO, error) {
	var turno domain.TurnoDTO
	row := s.db.QueryRow("select t.id, t.fecha, t.hora, t.descripcion,  " + 
	"d.id, d.nombre, d.apellido, d.matricula, " +
	"p.id, p.nombre, p.apellido, p.domicilio, p.dni, p.fecha_de_alta FROM turno t " +
	"JOIN dentista d ON d.id = t.dentista_id  " +
	"JOIN paciente p ON p.id = t.paciente_id " +
	" WHERE t.id= ?;", id)
	err := row.Scan(&turno.Id, &turno.Fecha, &turno.Hora, &turno.Descripcion,
		&turno.Dentista.Id, &turno.Dentista.Nombre, &turno.Dentista.Apellido, &turno.Dentista.Matricula,
		&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.Dni, &turno.Paciente.FechaAlta)
	if err != nil {
		return domain.TurnoDTO{}, err
	}
	return turno, nil
}

// Read devuelve un Turno por su dni paciente
func (s *sqlStore) ReadDniPaciente(dni int) (domain.TurnoDTO, error) {
	var turno domain.TurnoDTO
	row := s.db.QueryRow("select t.id, t.fecha, t.hora, t.descripcion,  " + 
	"d.id, d.nombre, d.apellido, d.matricula, " +
	"p.id, p.nombre, p.apellido, p.domicilio, p.dni, p.fecha_de_alta FROM turno t " +
	"JOIN dentista d ON d.id = t.dentista_id  " +
	"JOIN paciente p ON p.id = t.paciente_id " +
	" WHERE p.dni= ?;", dni)
	err := row.Scan(&turno.Id, &turno.Fecha, &turno.Hora, &turno.Descripcion,
		&turno.Dentista.Id, &turno.Dentista.Nombre, &turno.Dentista.Apellido, &turno.Dentista.Matricula,
		&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.Dni, &turno.Paciente.FechaAlta)
	if err != nil {
		return domain.TurnoDTO{}, err
	}
	return turno, nil
}



// Create agrega un nuevo Turno
func (s *sqlStore) Create(turno domain.TurnoDTO) ( error) {
	query := "insert into turno (dentista_id, paciente_id, fecha,hora,descripcion) values (?, ?, ?,?,?)"
	st, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(turno.Dentista.Id, turno.Paciente.Id, turno.Fecha, turno.Hora ,turno.Descripcion)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un Turno
func (s *sqlStore) Update(id int,turno domain.Turno) error {
	stmt, err := s.db.Prepare("UPDATE turno SET dentista_id = ?, paciente_id = ?, fecha = ?,hora = ?, descripcion = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&turno.DentistaId, &turno.PacienteId,&turno.Fecha,&turno.Hora,&turno.Descripcion,id )
	if err != nil {
		return err
	}
	return nil
}

// Delete elimina un Turno
func (s *sqlStore) Delete(id int) error {
	stmt := "delete from turno where id = ?"
	_, err := s.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}


// Trae todos los Turno
func (s *sqlStore) GetAll() ([]domain.Turno, error) {
	list := []domain.Turno{}
	rows, err := s.db.Query("select * from turno")
	if err != nil {
		return list, err
	}

	if rows.Err() != nil {
		return list, rows.Err()
	}

	for rows.Next() {
		var turno domain.Turno
		err := rows.Scan(&turno.DentistaId, &turno.PacienteId, &turno.Fecha,&turno.Hora,&turno.Descripcion )
		if err != nil {
			return nil, err
		}
		list = append(list, turno)
	}

	return list, nil
}

// Exists verifica si un turno existe
func (s *sqlStore) ExistTurno(paciente_id int, dentista_id int) bool {
	var id int
	row := s.db.QueryRow("select id from turno where dentista_id = ? and paciente_id = ?", dentista_id,paciente_id)
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}

	return false
}

package formularios

import (
	"log"

	"github.com/test/pkg/db"
	"github.com/test/pkg/res"
)

func CreateTables() error {
	log.Println("Create tables formularios")
	sql := `
    CREATE TABLE IF NOT EXISTS formularios (
      numIdentificacion INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      nombre TEXT,
      apellidos TEXT,
      tipoIdentificacion TEXT,
      estadoCivil INTEGER NOT NULL,
      fechaNacimiento TEXT,
      numBeneficiarios INTEGER NOT NULL,
      fechaIngreso TEXT
    );
  `

	_, err := db.Db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func InsertFormulario(nombre, apellidos, tipoIdentificacion string, estadoCivil int, fechaNacimiento string, numBeneficiarios int, fechaIngreso string) (res.Json, error) {
	sql := `
	INSERT INTO formularios (nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso)
	VALUES (?,?,?,?,?,?,?);`

	_, err := db.Db.Exec(sql, nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso)
	if err != nil {
		return nil, err
	}

	return res.Json{
		"mensaje": "Formulario creado",
	}, nil
}

func SelectAllFormularios() (res.Json, error) {
	sql := `
	SELECT * FROM formularios;
	`

	rows, err := db.Db.Query(sql)
	if err != nil {
		return nil, err
	}

	var formularios []res.Json
	for rows.Next() {
		var numIdentificacion int
		var nombre string
		var apellidos string
		var tipoIdentificacion string
		var estadoCivil int
		var fechaNacimiento string
		var numBeneficiarios int
		var fechaIngreso string

		err := rows.Scan(&numIdentificacion, &nombre, &apellidos, &tipoIdentificacion, &estadoCivil, &fechaNacimiento, &numBeneficiarios, &fechaIngreso)
		if err != nil {
			return nil, err
		}

		formulario := res.Json{
			"numIdentificacion":  numIdentificacion,
			"nombre":             nombre,
			"apellidos":          apellidos,
			"tipoIdentificacion": tipoIdentificacion,
			"estadoCivil":        estadoCivil,
			"fechaNacimiento":    fechaNacimiento,
			"numBeneficiarios":   numBeneficiarios,
			"fechaIngreso":       fechaIngreso,
		}

		formularios = append(formularios, formulario)
	}

	return res.Json{
		"formularios": formularios,
	}, nil

}

func UpdateFormulario(numIdentificacion int, nombre, apellidos, tipoIdentificacion string, estadoCivil int, fechaNacimiento string, numBeneficiarios int, fechaIngreso string) (res.Json, error) {
	sql := `
	UPDATE formularios SET 
	nombre=?,
	apellidos=?,
	tipoIdentificacion=?,
	estadoCivil=?,
	fechaNacimiento=?,
	numBeneficiarios=?,
	fechaIngreso=?
	WHERE numIdentificacion=?;
	`

	_, err := db.Db.Exec(sql, nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso, numIdentificacion)

	if err != nil {
		return nil, err
	}

	return res.Json{
		"message": "Formulario actualizado",
	}, nil
}

func DeleteFormulario(numIdentificacion int) (res.Json, error) {
	sql := `
	DELETE FROM formularios WHERE numIdentificacion=?;
	`

	_, err := db.Db.Exec(sql, numIdentificacion)
	if err != nil {
		return nil, err
	}

	return res.Json{
		"message": "Formulario eliminado",
	}, nil
}

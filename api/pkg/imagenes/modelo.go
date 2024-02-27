package imagenes

import (
	"log"

	"github.com/test/pkg/db"
	"github.com/test/pkg/res"
)

func CreateTables() error {
	log.Println("Create tables imagenes")

	sql := `CREATE TABLE IF NOT EXISTS imagenes (
		numIdentificacion INTEGER not null,
		nombreImagen Text,
		fecha TEXT,
		PRIMARY KEY(numIdentificacion, nombreimagen),
		CONSTRAINT imagenes_fk1 FOREIGN KEY(numIdentificacion) REFERENCES imagenes(numIdentificacionj) ON UPDATE CASCADE ON DELETE NO ACTION
	);
	`

	_, err := db.Db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func InsertImagen(numIdentificacion int, nombreImagen string, fecha string) (res.Json, error) {
	sql := `
	INSERT INTO imagenes (numIdentificacion, nombreImagen, fecha)
	VALUES (?,?,?)`

	_, err := db.Db.Exec(sql, numIdentificacion, nombreImagen, fecha)
	if err != nil {
		return nil, err
	}

	return res.Json{
		"mensaje": "Imsagen creada",
	}, nil
}

func SelectAllImagen() (res.Json, error) {
	sql := `
	SELECT * FROM imagenes;
	`

	rows, err := db.Db.Query(sql)
	if err != nil {
		return nil, err
	}

	var imagenes []res.Json
	for rows.Next() {
		var numIdentificacion int
		var nombreImagen string
		var fecha string

		err := rows.Scan(&numIdentificacion, &nombreImagen, &fecha)
		if err != nil {
			return nil, err
		}

		imagen := res.Json{
			"numIdentificacion": numIdentificacion,
			"nombreImagen":      nombreImagen,
			"fecha":             fecha,
		}

		imagenes = append(imagenes, imagen)
	}

	return res.Json{
		"imagenes": imagenes,
	}, nil

}

func UpdateImagen(numIdentificacion int, nombreImagen, fecha string) (res.Json, error) {
	sql := `
	UPDATE imagenes SET 
	nombreImagen=?,
	fecha=?
	WHERE numIdentificacion=?;
	`
	_, err := db.Db.Exec(sql, nombreImagen, fecha, numIdentificacion)

	if err != nil {
		return nil, err
	}

	return res.Json{
		"message": "Imagen actualizado",
	}, nil
}

func DeleteImagen(numIdentificacion int, nombreImagen string) (res.Json, error) {
	sql := `
	DELETE FROM imagenes WHERE numIdentificacion=? AND nombreImagen=?;
	`

	_, err := db.Db.Exec(sql, numIdentificacion, nombreImagen)
	if err != nil {
		return nil, err
	}

	return res.Json{
		"message": "Imagen eliminado",
	}, nil
}

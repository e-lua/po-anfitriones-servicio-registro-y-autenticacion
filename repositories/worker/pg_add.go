package repositories

import (
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Add(anfitrion_pg models.Pg_BusinessWorker) (int, error) {

	db := models.Conectar_Pg_DB()

	//Agregamos el Anfitrion
	id_inserted := 0
	err_add_business := db.QueryRow("INSERT INTO BusinessWorker(idcountry,phone,name,lastname,password,updateddate,idrol) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING idworker", anfitrion_pg.IdCountry, anfitrion_pg.Phone, anfitrion_pg.Name, anfitrion_pg.LastName, anfitrion_pg.Password, time.Now(), 1).Scan(&id_inserted)
	if err_add_business != nil {
		return id_inserted, err_add_business
	}

	//Actualizamos el negocio con el password
	q_2 := "UPDATE BusinessWorker SET idbusiness=$1 WHERE idworker=$2"
	update_idworker, error_update := db.Prepare(q_2)
	if error_update != nil {
		defer db.Close()
		return id_inserted, error_update
	}
	//Scaneamos l resultado y lo asignamos a la variable instanciada
	update_idworker.Exec(id_inserted, id_inserted)

	defer db.Close()

	return id_inserted, nil
}

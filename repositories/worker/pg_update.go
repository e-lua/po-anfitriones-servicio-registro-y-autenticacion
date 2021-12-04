package repositories

import (
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Update_Password(anfitrion models.Pg_BusinessWorker) (int, error) {

	idbusiness := 0

	db := models.Conectar_Pg_DB()

	//Buscamos el Id del negocio
	q := "SELECT idbusiness FROM BusinessWorker WHERE phone=$1"
	error_showidbusiness := db.QueryRow(q, anfitrion.Phone).Scan(&idbusiness)
	if error_showidbusiness != nil {
		defer db.Close()
		return idbusiness, error_showidbusiness
	}

	//Actualizamos el negocio con el password
	q_2 := "UPDATE business SET password=$1,updateddate=$2 WHERE idbusiness=$3"
	updateBusiness_Name, error_update := db.Prepare(q_2)
	if error_update != nil {
		defer db.Close()
		return idbusiness, error_update
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	updateBusiness_Name.Exec(anfitrion.Password, anfitrion.UpdatedDate, idbusiness)

	defer db.Close()
	return idbusiness, nil
}

func Pg_Update_NameLastName(anfitrion models.Pg_BusinessWorker) (int, error) {

	idbusiness := 0

	db := models.Conectar_Pg_DB()

	//Buscamos el Id del negocio
	q := "SELECT idbusiness FROM BusinessWorker WHERE phone=$1"
	error_showidbusiness := db.QueryRow(q, anfitrion.Phone).Scan(&idbusiness)
	if error_showidbusiness != nil {
		defer db.Close()
		return idbusiness, error_showidbusiness
	}

	//Actualizacion de nombre y apellido
	q_2 := "UPDATE business SET name=$1,lastname=$2, updateddate=$3 WHERE idbusiness=$4"
	updateBusiness_Name, error_update := db.Prepare(q_2)

	//Instanciamos una variable del modelo R_Country
	if error_update != nil {
		defer db.Close()
		return idbusiness, error_update
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	updateBusiness_Name.Exec(anfitrion.Name, anfitrion.LastName, anfitrion.UpdatedDate, idbusiness)

	defer db.Close()
	return idbusiness, nil
}

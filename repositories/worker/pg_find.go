package repositories

import (
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Find_ById(idbusiness int) (int, error) {

	var idbusiness_q int

	db := models.Conectar_Pg_DB()
	q := "SELECT idbusiness FROM BusinessWorker WHERE idbusiness=$1 LIMIT 1"
	error_query := db.QueryRow(q, idbusiness).Scan(&idbusiness_q)

	if error_query != nil {
		return idbusiness_q, error_query
	}

	defer db.Close()

	return idbusiness_q, nil

}

func Pg_FindByPhone(phone int) (models.Pg_BusinessWorker, error) {

	var anfitrion models.Pg_BusinessWorker

	db := models.Conectar_Pg_DB()
	q := "SELECT idbusiness,idworker,idcountry,idrol,phone,password FROM BusinessWorker WHERE phone=$1 LIMIT 1"
	error_query := db.QueryRow(q, phone).Scan(&anfitrion.IdBusiness, &anfitrion.IdWorker, &anfitrion.IdCountry, &anfitrion.IdRol, &anfitrion.Phone, &anfitrion.Password)

	if error_query != nil {
		return anfitrion, error_query
	}

	defer db.Close()

	return anfitrion, nil

}

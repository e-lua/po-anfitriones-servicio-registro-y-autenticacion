package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Find_ById(idbusiness int, idcountry int) (int, error) {

	var idbusiness_q int

	db := models.Conectar_Pg_DB()
	q := `SELECT idbusiness FROM BusinessWorker WHERE idbusiness=$1 AND idcountry=$2 LIMIT 1`
	error_show := db.QueryRow(context.Background(), q, idbusiness, idcountry).Scan(&idbusiness_q)

	if error_show != nil {

		return idbusiness_q, error_show
	}
	return idbusiness_q, nil

}

func Pg_FindByPhone(phone int, idcountry int) (models.Pg_BusinessWorker, error) {

	var anfitrion models.Pg_BusinessWorker

	db := models.Conectar_Pg_DB()
	q := `SELECT idbusiness,idworker,idcountry,idrol,phone,password,name,lastname FROM BusinessWorker WHERE phone=$1 AND idcountry=$2 LIMIT 1`
	error_show := db.QueryRow(context.Background(), q, phone, idcountry).Scan(&anfitrion.IdBusiness, &anfitrion.IdWorker, &anfitrion.IdCountry, &anfitrion.IdRol, &anfitrion.Phone, &anfitrion.Password, &anfitrion.Name, &anfitrion.LastName)

	if error_show != nil {
		return anfitrion, error_show
	}

	return anfitrion, nil
}

func Pg_FindPassword_ById(idbusiness int) (string, error) {

	var pass string

	db := models.Conectar_Pg_DB()
	q := "SELECT password FROM BusinessWorker WHERE idbusiness=$1"
	error_showname := db.QueryRow(context.Background(), q, idbusiness).Scan(&pass)

	if error_showname != nil {
		return pass, error_showname
	}

	//Si todo esta bien
	return pass, nil

}

func Pg_Find_QtyCodesRegistered(idbusiness int, idcountry int) (int, error) {

	var codesregistered_pg int

	db := models.Conectar_Pg_DB()
	q := "SELECT codesregistered FROM BusinessWorker WHERE idcomensal=$1 AND idcountry=$2 LIMIT 1"
	error_query := db.QueryRow(context.Background(), q, idbusiness, idcountry).Scan(&codesregistered_pg)

	if error_query != nil {
		return codesregistered_pg, error_query
	}

	return codesregistered_pg, nil
}

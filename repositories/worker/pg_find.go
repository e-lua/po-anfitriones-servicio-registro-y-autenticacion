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
		defer db.Close()
		return idbusiness_q, error_show
	}
	db.Close()
	return idbusiness_q, nil

}

func Pg_FindByPhone(phone int, idcountry int) (models.Pg_BusinessWorker, error) {

	var anfitrion models.Pg_BusinessWorker

	db := models.Conectar_Pg_DB()
	q := `SELECT idbusiness,idworker,idcountry,idrol,phone,password FROM BusinessWorker WHERE phone=$1 AND idcountry=$2 LIMIT 1`
	error_show := db.QueryRow(context.Background(), q, phone, idcountry).Scan(&anfitrion.IdBusiness, &anfitrion.IdWorker, &anfitrion.IdCountry, &anfitrion.IdRol, &anfitrion.Phone, &anfitrion.Password)

	if error_show != nil {
		defer db.Close()
		return anfitrion, error_show
	}
	defer db.Close()
	return anfitrion, nil
}

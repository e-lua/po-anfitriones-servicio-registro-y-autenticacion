package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Find_ById(idbusiness int, idcountry int) (int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var idbusiness_q int

	db := models.Conectar_Pg_DB()
	q := `SELECT idbusiness FROM BusinessWorker WHERE idbusiness=$1 AND idcountry=$2 LIMIT 1`
	error_show := db.QueryRow(ctx, q, idbusiness, idcountry).Scan(&idbusiness_q)

	if error_show != nil {

		return idbusiness_q, error_show
	}
	return idbusiness_q, nil

}

func Pg_FindByPhone(phone int, idcountry int) (models.Pg_BusinessWorker, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var anfitrion models.Pg_BusinessWorker

	db := models.Conectar_Pg_DB()
	q := `SELECT idbusiness,idworker,idcountry,idrol,phone,password,name,lastname,isbanned,sessioncode,isdeleted FROM BusinessWorker WHERE phone=$1 AND idcountry=$2 LIMIT 1`
	error_show := db.QueryRow(ctx, q, phone, idcountry).Scan(&anfitrion.IdBusiness, &anfitrion.IdWorker, &anfitrion.IdCountry, &anfitrion.IdRol, &anfitrion.Phone, &anfitrion.Password, &anfitrion.Name, &anfitrion.LastName, &anfitrion.Isbanned, &anfitrion.SessionCode, &anfitrion.IsDeleted)

	if error_show != nil {
		return anfitrion, error_show
	}

	return anfitrion, nil
}

func Pg_FindPassword_ById(idbusiness int) (string, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var pass string

	db := models.Conectar_Pg_DB()
	q := "SELECT password FROM BusinessWorker WHERE idbusiness=$1"
	error_showname := db.QueryRow(ctx, q, idbusiness).Scan(&pass)

	if error_showname != nil {
		return pass, error_showname
	}

	//Si todo esta bien
	return pass, nil

}

func Pg_Find_QtyCodesRegistered(idbusiness int, idcountry int) (int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var codesregistered_pg int

	db := models.Conectar_Pg_DB()
	q := "SELECT codesregistered FROM BusinessWorker WHERE idcomensal=$1 AND idcountry=$2 LIMIT 1"
	error_query := db.QueryRow(ctx, q, idbusiness, idcountry).Scan(&codesregistered_pg)

	if error_query != nil {
		return codesregistered_pg, error_query
	}

	return codesregistered_pg, nil
}

func Pg_Find_SubWorkers(idbusiness int) ([]models.Pg_SubWorker, int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Contador
	quantity := 0

	db := models.Conectar_Pg_DB()
	q := "SELECT idworker,idbusiness,name,lastname,idcountry,phone,TO_CHAR(createddate, 'dd Mon, yyyy HH12:MI AM') FROM businessworker WHERE isdeleted=false AND idrol=2 AND idbusiness=$1"
	rows, error_query := db.Query(ctx, q, idbusiness)

	//Instanciamos una variable del modelo Pg_SubWorker
	var oListSubWorker []models.Pg_SubWorker

	if error_query != nil {
		return oListSubWorker, quantity, error_query
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		subworker_pg := models.Pg_SubWorker{}
		rows.Scan(&subworker_pg.IdWorker, &subworker_pg.IdBusiness, &subworker_pg.Name, &subworker_pg.LastName, &subworker_pg.IdCountry, &subworker_pg.Phone, &subworker_pg.DateRegistered)
		oListSubWorker = append(oListSubWorker, subworker_pg)
		quantity = quantity + 1
	}

	return oListSubWorker, quantity, nil
}

package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Update_Password(password string, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE businessworker SET password=$1,updateddate=$2 WHERE idbusiness=$3`
	if _, err_update := db.Exec(ctx, query, password, time.Now(), idbusiness); err_update != nil {
		return err_update
	}

	return nil
}

func Pg_Update_NameLastName(name string, lastname string, idworker int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE businessworker SET name=$1,lastname=$2, updateddate=$3 WHERE idworker=$4`
	if _, err_update := db.Exec(ctx, query, name, lastname, time.Now(), idworker); err_update != nil {
		return err_update
	}

	return nil
}

func Pg_Update_Email(email string, idworker int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE businessworker SET email=$1,isexported=false,updateddate=$2 WHERE idworker=$3`
	if _, err_update := db.Exec(ctx, query, email, time.Now(), idworker); err_update != nil {
		return err_update
	}

	return nil
}

func Pg_Update_IdBusiness(idworker int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE BusinessWorker SET idbusiness=$1 WHERE idworker=$2`
	if _, err_update := db.Exec(ctx, query, idworker, idworker); err_update != nil {
		return err_update
	}

	return nil
}

func Pg_Update_QtyCodesRegistered(phonenumber int, country int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE BusinessWorker SET codesregistered=codesregistered+1,updateddate=$1 WHERE phone=$2 AND idcountry=$3`
	if _, err_update := db.Exec(ctx, query, time.Now(), phonenumber, country); err_update != nil {
		return err_update
	}

	return nil
}

func Pg_Update_Password_Recovery(password string, phone int, code int, sessioncode int) error {

	//Tiempo limite al contextos
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE BusinessWorker SET password=$1,updateddate=$2,updatedPassword=$3,sessioncode=$4 WHERE phone=$5 AND idcountry=$6`
	if _, err_update := db.Exec(ctx, query, password, time.Now(), time.Now(), sessioncode, phone, code); err_update != nil {
		return err_update
	}

	return nil
}

func Pg_Update_IsDeleted(idworker int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE BusinessWorker SET isexported=$1,isdeleted=$2,sessioncode=$3 WHERE idworker=$4`
	if _, err_update := db.Exec(ctx, query, false, true, 123456, idworker); err_update != nil {
		return err_update
	}

	return nil
}

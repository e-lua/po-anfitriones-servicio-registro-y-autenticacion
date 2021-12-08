package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Update_Password(anfitrion models.Pg_BusinessWorker) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE business SET password=$1,updateddate=$2 WHERE phone=$3`
	if _, err_update := db.Exec(context.Background(), query, anfitrion.Password, anfitrion.UpdatedDate, anfitrion.Phone); err_update != nil {
		return err_update
	}

	defer db.Close()
	return nil
}

func Pg_Update_NameLastName(anfitrion models.Pg_BusinessWorker) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE business SET name=$1,lastname=$2, updateddate=$3 WHERE phone=$4`
	if _, err_update := db.Exec(context.Background(), query, anfitrion.Name, anfitrion.LastName, anfitrion.UpdatedDate, anfitrion.Phone); err_update != nil {
		return err_update
	}

	return nil
}

func Pg_Update_IdBusiness(idworker int) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE BusinessWorker SET idbusiness=$1 WHERE idworker=$2`
	if _, err_update := db.Exec(context.Background(), query, idworker, idworker); err_update != nil {
		return err_update
	}

	return nil
}

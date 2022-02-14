package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Delete_SubWorker(idworker int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `DELETE FROM businessworker WHERE idworker=$1`
	if _, err_update := db.Exec(ctx, query, idworker); err_update != nil {
		return err_update
	}

	return nil
}

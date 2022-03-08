package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Update_IDDevice(idworker int, id_device string) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	//Id del worker insertado
	var id_inserted int

	query := `UPDATE BusinessWorker SET iddevice=$1 WHERE idworker=$2`
	inserted := db.QueryRow(ctx, query, id_device, idworker)

	inserted.Scan(&id_inserted)

	return nil
}

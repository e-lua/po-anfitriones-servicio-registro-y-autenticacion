package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Find_IfIsAvailable() (bool, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var available bool

	db := models.Conectar_Pg_DB()

	q := "SELECT availableregister FROM comensal WHERE idcomensal=1"
	error_query := db.QueryRow(ctx, q).Scan(&available)

	if error_query != nil {
		return available, error_query
	}

	return available, nil
}

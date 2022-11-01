package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Find_Qty_Subsidiary(idbusiness int) (int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var qty int

	db := models.Conectar_Pg_DB()

	q := `SELECT COUNT(idworker) FROM businessworker WHERE issubsidiary=true AND idbusiness=$1`
	error_show := db.QueryRow(ctx, q, idbusiness).Scan(&qty)
	if error_show != nil {
		return qty, error_show
	}

	return qty, nil
}

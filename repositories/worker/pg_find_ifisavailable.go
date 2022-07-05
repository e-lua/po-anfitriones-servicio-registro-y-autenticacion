package repositories

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find_IfIsAvailable() (bool, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var available bool
	var db *pgxpool.Pool

	random := rand.Intn(4)
	if random%2 == 0 {
		db = models.Conectar_Pg_DB()
	} else {
		db = models.Conectar_Pg_DB_Slave()
	}

	q := "SELECT availableregister FROM comensal WHERE idcomensal=1"
	error_query := db.QueryRow(ctx, q).Scan(&available)

	if error_query != nil {
		return available, error_query
	}

	return available, nil
}

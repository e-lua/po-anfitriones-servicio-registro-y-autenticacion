package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Add(anfitrion_pg models.Pg_BusinessWorker) (int, error) {

	db := models.Conectar_Pg_DB()

	//Id del worker insertado
	var id_inserted int

	query := `INSERT INTO BusinessWorker(idcountry,phone,name,lastname,password,updateddate,idrol) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING idworker`
	inserted := db.QueryRow(context.Background(), query, anfitrion_pg.IdCountry, anfitrion_pg.Phone, anfitrion_pg.Name, anfitrion_pg.LastName, anfitrion_pg.Password, time.Now(), 1)

	inserted.Scan(&id_inserted)

	return id_inserted, nil
}

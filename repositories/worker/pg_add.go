package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Pg_Add(anfitrion_pg models.Pg_BusinessWorker) (int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	//Id del worker insertado
	var id_inserted int

	query := `INSERT INTO BusinessWorker(idcountry,phone,name,lastname,password,updateddate,idrol) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING idworker`
	inserted := db.QueryRow(ctx, query, anfitrion_pg.IdCountry, anfitrion_pg.Phone, anfitrion_pg.Name, anfitrion_pg.LastName, anfitrion_pg.Password, time.Now(), 1)

	inserted.Scan(&id_inserted)

	return id_inserted, nil
}

func Pg_Add_Subworker(anfitrion_pg models.Pg_BusinessWorker) (int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Id del worker insertado
	var id_inserted int

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO BusinessWorker(idcountry,phone,name,lastname,password,updateddate,idrol,idbusiness) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING idworker`
	inserted := db.QueryRow(ctx, query, anfitrion_pg.IdCountry, anfitrion_pg.Phone, anfitrion_pg.Name, anfitrion_pg.LastName, anfitrion_pg.Password, time.Now(), 2, anfitrion_pg.IdBusiness)

	inserted.Scan(&id_inserted)

	return id_inserted, nil
}

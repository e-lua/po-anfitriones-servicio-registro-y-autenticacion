package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Re_Set_Id(idbusiness int, idcountry int) (int, error) {

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(idbusiness)+strconv.Itoa(idcountry), strconv.Itoa(idbusiness)+strconv.Itoa(idcountry), "NX")
	if err_do != nil {
		models.RedisCN.Get().Close()
		return 0, err_do
	}

	models.RedisCN.Get().Close()
	return idbusiness, nil
}

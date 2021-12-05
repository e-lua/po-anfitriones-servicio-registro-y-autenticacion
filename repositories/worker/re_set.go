package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Re_Set_Id(idbusiness int, idcountry int) (int, error) {

	_, err_do := models.RedisCN.Do("SET", strconv.Itoa(idbusiness)+strconv.Itoa(idcountry), strconv.Itoa(idbusiness)+strconv.Itoa(idcountry), "NX")
	if err_do != nil {
		return 0, err_do
	}

	return idbusiness, nil
}

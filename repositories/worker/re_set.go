package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Re_Set_Id(idbusiness int, idcountry int, sessioncode int, phone int) (int, error) {

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(idbusiness)+strconv.Itoa(idcountry)+strconv.Itoa(phone), strconv.Itoa(idbusiness)+strconv.Itoa(sessioncode)+strconv.Itoa(idcountry)+strconv.Itoa(phone), "EX", 7776000)
	if err_do != nil {
		return 0, err_do
	}

	return idbusiness, nil
}

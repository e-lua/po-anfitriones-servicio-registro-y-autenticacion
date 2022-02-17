package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Re_Set_ID(idworker int, idcountry int, sessioncode int, idbusiness int) (int, error) {

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(idworker)+strconv.Itoa(idcountry)+strconv.Itoa(idbusiness), strconv.Itoa(idworker)+strconv.Itoa(sessioncode)+strconv.Itoa(idbusiness), "EX", 7776000)
	if err_do != nil {
		return 0, err_do
	}

	return idbusiness, nil
}

func Re_Set_Email(idworker int, sessioncode int, idrol int) error {

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(idworker)+strconv.Itoa(idrol), strconv.Itoa(idworker)+strconv.Itoa(sessioncode)+strconv.Itoa(idrol), "EX", 7776000)
	if err_do != nil {
		return err_do
	}

	return nil
}

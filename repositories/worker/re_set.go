package repositories

import (
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Re_Set_Id(idbusiness int) (int, error) {

	_, err_do := models.RedisCN.Do("SET", idbusiness, idbusiness, "NX")
	if err_do != nil {
		return 0, err_do
	}

	return idbusiness, nil
}

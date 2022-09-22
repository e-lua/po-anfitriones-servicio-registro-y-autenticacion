package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Re_Set_Request(phone int, country int, quantity int) error {

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(phone)+strconv.Itoa(country)+"REQUEST", strconv.Itoa(quantity), "EX", 86400)
	if err_do != nil {
		return err_do
	}

	return nil
}

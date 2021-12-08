package repositories

import (
	"encoding/json"
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

func Re_Set_Phone(code models.Re_SetGetCode) (int, error) {

	uJson, err_marshal := json.Marshal(code)
	if err_marshal != nil {
		defer models.RedisCN.Get().Close()
		return 0, err_marshal
	}

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(code.PhoneRegister_Key)+strconv.Itoa(code.Country), uJson, "EX", 300)
	if err_do != nil {
		defer models.RedisCN.Get().Close()
		return 0, err_do
	}

	defer models.RedisCN.Get().Close()
	return code.PhoneRegister_Key, nil
}

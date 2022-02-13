package registro

import (
	"strconv"

	"github.com/labstack/echo/v4"

	//MDOELS
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

var RegisterRouter *registerRouter

type registerRouter struct {
}

func (rr *registerRouter) AvailableRegister(c echo.Context) error {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AvailableRegister_Service()
	results := Response_Available{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (rr *registerRouter) SignUpNumber(c echo.Context) error {

	//Instanciamos una variable del modelo Code
	var code models.Re_SetGetCode

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&code)
	if err != nil {
		results := Response_WithInt{Error: true, DataError: "Se debe enviar el numero y el pais, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if code.PhoneRegister_Key < 999999 {
		results := Response_WithInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SignUpNumber_Service(code)
	results := Response_SignInFirstStep{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *registerRouter) UpdateCodeWithCode(c echo.Context) error {

	//Recibimos el id del Business Owner
	phoneregister := c.Param("phoneRegister")
	country := c.Param("country")

	//Instanciamos una variable del modelo Code
	var code models.Re_SetGetCode

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&code)
	if err != nil {
		results := Response_WithInt{Error: true, DataError: "Se debe enviar el código, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(phoneregister) < 9 {
		results := Response_WithInt{Error: true, DataError: "Los valores ingresados no cumplen con la regla de negocio", Data: 0}
		return c.JSON(400, results)
	}

	//Convertimos texto a numero
	numero_registro, _ := strconv.Atoi(phoneregister)
	country_registro, _ := strconv.Atoi(country)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateWithCode_Service(numero_registro, code, country_registro)
	results := Response_WithPhoneCountryCode{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *registerRouter) UpdateWithCodeRecovery(c echo.Context) error {

	//Recibimos el id del Business Owner
	phoneregister := c.Param("phoneRegister")
	country := c.Param("country")

	//Instanciamos una variable del modelo Code
	var code models.Re_SetGetCode

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&code)
	if err != nil {
		results := Response_WithInt{Error: true, DataError: "Se debe enviar el código, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(phoneregister) < 9 {
		results := Response_WithInt{Error: true, DataError: "Los valores ingresados no cumplen con la regla de negocio", Data: 0}
		return c.JSON(400, results)
	}

	//Convertimos texto a numero
	numero_registro, _ := strconv.Atoi(phoneregister)
	country_registro, _ := strconv.Atoi(country)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateWithCodeRecovery_Service(numero_registro, code, country_registro)
	results := Response_WithPhoneCountryCode{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *registerRouter) RegisterAnfitrion(c echo.Context) error {

	//Instanciamos una variable del modelo Code
	var anfitrion models.Pg_BusinessWorker

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&anfitrion)
	if err != nil {
		results := Response_WithString{Error: true, DataError: "Se debe enviar los datos del pais, nombre, apellido y contraseña del anfitrion, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if anfitrion.Phone < 999999 || len(anfitrion.Password) < 8 || len(anfitrion.Name) < 1 || len(anfitrion.LastName) < 1 || anfitrion.IdCountry != 51 && anfitrion.IdCountry != 52 {
		results := Response_WithString{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := RegisterAnfitrion_Service(anfitrion)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *registerRouter) UpdatePassword_Recover(c echo.Context) error {

	//Instanciamos una variable del modelo Code
	var entrydata EntryData_Password

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&entrydata)
	if err != nil {
		results := Response_WithString{Error: true, DataError: "Se debe enviar los datos necesarios para actualizar la contraseña, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if entrydata.Phone < 999999 || entrydata.Country != 51 && entrydata.Country != 52 || len(entrydata.NewPassword) < 8 {
		results := Response_WithString{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdatePassword_Recover_Service(entrydata)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *registerRouter) RegisterColaborador(c echo.Context) error {

	//Instanciamos una variable del modelo Code
	var anfitrion models.Pg_BusinessWorker

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&anfitrion)
	if err != nil {
		results := Response_WithString{Error: true, DataError: "Se debe enviar los datos del pais, nombre, apellido y contraseña del anfitrion, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if anfitrion.Phone < 999999 || len(anfitrion.Password) < 8 || len(anfitrion.Name) < 1 || len(anfitrion.LastName) < 1 || anfitrion.IdCountry != 51 && anfitrion.IdCountry != 52 {
		results := Response_WithString{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := RegisterColaborador_Service(anfitrion)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

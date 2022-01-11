package profile

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	//MDOELS
)

var ProfileRouter *profileRouter

type profileRouter struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://localhost:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

/*----------------------INICIO DEL ROUTER----------------------*/

func (pr *profileRouter) UpdateNameLastName(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response_WithString{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response_WithString{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Code
	var anfitrion Entry_Profile

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&anfitrion)
	if err != nil {
		results := Response_WithString{Error: true, DataError: "Se debe enviar los datos del pais, nombre, apellido y contraseña del anfitrion, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(anfitrion.Name) < 1 && len(anfitrion.Name) > 20 || len(anfitrion.LastName) < 1 && len(anfitrion.LastName) > 30 {
		results := Response_WithString{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateNameLastName_Service(anfitrion, data_idbusiness)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (pr *profileRouter) UpdatePassword(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response_WithString{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response_WithString{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Code
	var entrydata EntryData_Password

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&entrydata)
	if err != nil {
		results := Response_WithString{Error: true, DataError: "Se debe enviar los datos necesarios para actualizar la contraseña, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if entrydata.Phone < 999999 || entrydata.Country != 51 && entrydata.Country != 52 || len(entrydata.OldPassword) < 8 || len(entrydata.NewPassword) < 8 {
		results := Response_WithString{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdatePassword_Service(entrydata, data_idbusiness)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

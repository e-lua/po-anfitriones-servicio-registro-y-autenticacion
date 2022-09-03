package solicitud_plan

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

var AnfitrionRouter_pg *anfitrionRouter_pg

type anfitrionRouter_pg struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT_Anfitrion(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta_anfitrion, _ := http.Get("http://localhost:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT_Anfitrion
	error_decode_respuesta := json.NewDecoder(respuesta_anfitrion.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

/*----------------------FUNCIONES DEL SERVICIO----------------------*/

func (ar *anfitrionRouter_pg) Anfitriones_SendRequest(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Anfitriones_SendRequest_Service(data_idbusiness, "-5")
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

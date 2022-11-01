package subsidiary

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	//MDOELS
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

var SubsidiaryRouter *subsidiaryRouter

type subsidiaryRouter struct {
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

/*----------------------COMIENZA EL ROUTER----------------------*/

func (sr *subsidiaryRouter) AddSubsidiary(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Code
	var subsidiary models.Pg_BusinessWorker

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&subsidiary)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos del pais, nombre, apellido y contraseña del anfitrion, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	/*if len(subsidiary.Name) < 1 && len(subsidiary.Name) > 20 || len(subsidiary.LastName) < 1 && len(subsidiary.LastName) > 30 || subsidiary.Phone=0 ||subsidiary.IdCountry=0{
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}*/

	if len(subsidiary.Password) < 1 || len(subsidiary.Name) < 1 && len(subsidiary.Name) > 20 || len(subsidiary.LastName) < 1 && len(subsidiary.LastName) > 30 || subsidiary.Phone == 0 || subsidiary.IdCountry == 0 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}

	if len(subsidiary.Email) == 0 {
		subsidiary.Email = "na"
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddSubsidiary_Service(data_idbusiness, subsidiary)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

package profile

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func GetJWTCountry(jwt string) (int, bool, string, int, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://localhost:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0, 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness, get_respuesta.Data.IdCountry
}

func GetJWTRol(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://localhost:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdRol
}

func GetJWTRol_Country_Business(jwt string) (int, bool, string, int, int, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://localhost:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0, 0, 0
	}
	return 200, false, "", get_respuesta.Data.IdRol, get_respuesta.Data.IdCountry, get_respuesta.Data.IdBusiness
}

func GetJWTSubWorker(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://localhost:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdWorker
}

/*----------------------INICIO DEL ROUTER----------------------*/

func (pr *profileRouter) UpdateNameLastNameEmail(c echo.Context) error {

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

	if len(anfitrion.Email) == 0 {
		anfitrion.Email = "na"
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateNameLastNameEmail_Service(anfitrion, data_idbusiness)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (pr *profileRouter) UpdatePassword(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness, data_idcountry := GetJWTCountry(c.Request().Header.Get("Authorization"))
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
	status, boolerror, dataerror, data := UpdatePassword_Service(entrydata, data_idbusiness, data_idcountry)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (pr *profileRouter) DeleteAnfitrion(c echo.Context) error {

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

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := DeleteAnfitrion_Service(data_idbusiness)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (pr *profileRouter) DeleteColaborador(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idrol, data_idcountry, data_idbusines := GetJWTRol_Country_Business(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response_WithString{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idrol <= 0 {
		results := Response_WithString{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if data_idrol != 1 {
		results := Response_WithString{Error: true, DataError: "Este rol no puede eliminar colaboradores", Data: ""}
		return c.JSON(403, results)
	}

	idsubworker := c.Param("idsubworker")
	idsubworker_int, _ := strconv.Atoi(idsubworker)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := DeleteColaborador_Service(idsubworker_int, data_idrol, data_idcountry, data_idbusines)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (pr *profileRouter) GetColaborador(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idrol := GetJWTRol(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response_WithString{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idrol <= 0 {
		results := Response_WithString{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if data_idrol != 1 {
		results := Response_WithString{Error: true, DataError: "Este rol no puede listar colaboradores", Data: ""}
		return c.JSON(403, results)
	}

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

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetColaborador_Service(data_idbusiness)
	results := Response_SubWorkers{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (pr *profileRouter) GetEmail(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().URL.Query().Get("Authorization"))
	if dataerror != "" {
		results := Response_WithString{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response_WithString{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetEmail_Service(data_idbusiness)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (pr *profileRouter) GetColaboradorToExport(c echo.Context) error {

	idsubworker := c.Param("idsubworker")
	idsubworker_int, _ := strconv.Atoi(idsubworker)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := V2_GetColaboradorToExport_Service(idsubworker_int)
	results := Response_SubWorker_ToExport{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (pr *profileRouter) UpdatIDDevice(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idworker := GetJWTSubWorker(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response_WithString{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idworker <= 0 {
		results := Response_WithString{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	var id_device Input_IDDevice

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&id_device)
	if err != nil {
		results := Response_WithString{Error: true, DataError: "Se debe enviar los datos necesarios para actualizar la contraseña, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(id_device.IDDevice) < 5 {
		results := Response_WithString{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateIDDevice_Service(data_idworker, id_device.IDDevice)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*=======================================*/
/*===============VERSION 2===============*/
/*=======================================*/

func (pr *profileRouter) V2_GetColaborador(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idrol := GetJWTRol(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response_WithString{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idrol <= 0 {
		results := Response_WithString{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if data_idrol != 1 {
		results := Response_WithString{Error: true, DataError: "Este rol no puede listar colaboradores", Data: ""}
		return c.JSON(403, results)
	}

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

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := V2_GetColaborador_Service(data_idbusiness)
	results := Response_SubWorkers_V2{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

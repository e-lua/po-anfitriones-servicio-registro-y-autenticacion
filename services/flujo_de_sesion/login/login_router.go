package login

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	//MDOELS
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
)

var Loginrouter *loginRouter

type loginRouter struct {
}

func (lr *loginRouter) Login(c echo.Context) error {

	//Instanciamos una variable del modelo Business Worker
	var anfitrion models.Pg_BusinessWorker

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&anfitrion)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar todos los datos solicitados, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(strconv.Itoa(anfitrion.Phone)) < 8 || len(anfitrion.Password) < 8 {
		results := Response{Error: true, DataError: "Los valores ingresados no cumplen con las reglas de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Login_Service(anfitrion)

	/* --Grabar una cookie desde el back--
	Primero crearemos un campo fecha para saber la expiración
	*/
	expirationTime := time.Now().Add(72 * time.Hour)
	http.SetCookie(c.Response().Writer, &http.Cookie{
		Name:    "token",
		Value:   data.JWT,
		Expires: expirationTime,
	})

	results := Responde_JWTAndRol{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (lr *loginRouter) TryingLogin(c echo.Context) error {

	//Recibimos el JWT
	jwt := c.Request().URL.Query().Get("jwt")

	//Recibimos el JWT
	service := c.Request().URL.Query().Get("service")

	//Recibimos el JWT
	module := c.Request().URL.Query().Get("module")

	//Recibimos el JWT
	epic := c.Request().URL.Query().Get("epic")

	//Recibimos el JWT
	endpoint := c.Request().URL.Query().Get("endpoint")

	//Validamos los valores enviados
	if len(jwt) < 8 {
		results := Response{Error: true, DataError: "Los valores ingresados no cumplen con las reglas de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	anfitrionjwt, boolerror, _, error := TryingLogin_Service(jwt, service, module, epic, endpoint)

	results := ResponseJWT{Error: boolerror, DataError: "Null" + error, Data: anfitrionjwt}
	return c.JSON(200, results)

}

/*=======================================*/
/*===============VERSION 2===============*/
/*=======================================*/

func (lr *loginRouter) V2_Login(c echo.Context) error {

	//Instanciamos una variable del modelo Business Worker
	var login_data Input_BusinessWorker_login

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&login_data)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar todos los datos solicitados, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	if login_data.IsAnfitrion {
		//Validamos los valores enviados
		if len(strconv.Itoa(login_data.Phone)) < 8 || len(login_data.Password) < 8 {
			results := Response{Error: true, DataError: "Los valores ingresados no cumplen con las reglas de negocio", Data: ""}
			return c.JSON(400, results)
		}
	} else {
		//Validamos los valores enviados
		if len(login_data.Email) < 3 && len(login_data.Email) > 100 || len(login_data.Password) < 8 {
			results := Response{Error: true, DataError: "Los valores ingresados no cumplen con las reglas de negocio", Data: ""}
			return c.JSON(400, results)
		}
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := V2_Login_Service(login_data)

	expirationTime := time.Now().Add(72 * time.Hour)
	http.SetCookie(c.Response().Writer, &http.Cookie{
		Name:    "token",
		Value:   data.JWT,
		Expires: expirationTime,
	})

	results := Responde_JWTAndRol{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

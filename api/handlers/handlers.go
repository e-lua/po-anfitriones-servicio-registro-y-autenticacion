package api

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	profile "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/services/flujo_de_perfil/profile"
	login "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/services/flujo_de_sesion/login"
	register "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/services/flujo_de_sesion/registro"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	//V1 TO AVAILABLE REGISTER
	router_available_v1 := version_1.Group("/available")
	router_available_v1.GET("", register.RegisterRouter.AvailableRegister)

	//V1 TO LOGIN
	router_login_v1 := version_1.Group("/login")
	router_login_v1.POST("", login.Loginrouter.Login)

	//V1 TO ENTITY-CODE
	router_code_v1 := version_1.Group("/codes")
	router_code_v1.POST("", register.RegisterRouter.SignUpNumber)
	router_code_v1.PUT("/:phoneRegister/:country", register.RegisterRouter.UpdateCodeWithCode)

	//V1 TO RECOVER
	router_recover_v1 := version_1.Group("/recover")
	router_recover_v1.PUT("/code/:phoneRegister/:country", register.RegisterRouter.UpdateWithCodeRecovery)
	router_recover_v1.PUT("/password", register.RegisterRouter.UpdatePassword_Recover)

	//V1 TO ANFITRION
	router_anfitrion_v1 := version_1.Group("/worker")
	router_anfitrion_v1.POST("", register.RegisterRouter.RegisterAnfitrion)
	router_anfitrion_v1.PUT("/password", profile.ProfileRouter.UpdatePassword)
	router_anfitrion_v1.PUT("/profile", profile.ProfileRouter.UpdateNameLastName)

	//V1 TO TRYLOGIN
	router_login := version_1.Group("/trylogin")
	router_login.GET("", login.Loginrouter.TryingLogin)

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "5000"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Acceso no autorizado")
}

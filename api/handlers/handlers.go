package api

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	export "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/services/flujo_de_perfil/export"
	profile "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/services/flujo_de_perfil/profile"
	login "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/services/flujo_de_sesion/login"
	register "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/services/flujo_de_sesion/registro"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)

	/*=======================================*/
	/*===============VERSION 1===============*/
	/*=======================================*/

	//VERSION
	version_1 := e.Group("/v1")

	//V1 TO AVAILABLE REGISTER
	router_available_v1 := version_1.Group("/available")
	router_available_v1.GET("", register.RegisterRouter.AvailableRegister)

	//V1 TO LOGIN
	router_login_v1 := version_1.Group("/login")
	router_login_v1.POST("", login.Loginrouter.Login)

	//V1 TO DEVICE
	router_device_v1 := version_1.Group("/device")
	router_device_v1.PUT("", profile.ProfileRouter.UpdatIDDevice)

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
	router_anfitrion_v1.PUT("/email", profile.ProfileRouter.UpdateEmail)
	router_anfitrion_v1.GET("/email", profile.ProfileRouter.GetEmail)
	router_anfitrion_v1.DELETE("", profile.ProfileRouter.DeleteAnfitrion)

	//V1 TO COLABORADOR
	router_colaborador_v1 := version_1.Group("/subworker")
	router_colaborador_v1.POST("", register.RegisterRouter.RegisterColaborador)
	router_colaborador_v1.GET("", profile.ProfileRouter.GetColaborador)
	router_colaborador_v1.DELETE("/:idsubworker", profile.ProfileRouter.DeleteColaborador)

	//V1 TO TRYLOGIN
	router_login := version_1.Group("/trylogin")
	router_login.GET("", login.Loginrouter.TryingLogin)

	//V1 TO COLABORADOR A EXPORTAR - SE UTILIZA EN PEDIDOS, PERO SE VA A BORRAR
	router_colaborador_to_export_v1 := version_1.Group("/subworkertoexport")
	router_colaborador_to_export_v1.GET("/:idsubworker", profile.ProfileRouter.GetColaboradorToExport)

	/*=======================================*/
	/*===============VERSION 2===============*/
	/*=======================================*/

	//VERSION
	version_2 := e.Group("/v2")

	//V2 TO LOGIN
	router_login_v2 := version_2.Group("/login")
	router_login_v2.POST("", login.Loginrouter.V2_Login)

	//V2 TO COLABORADOR
	router_colaborador_v2 := version_2.Group("/subworker")
	router_colaborador_v2.POST("", register.RegisterRouter.V2_RegisterColaborador)
	router_colaborador_v2.GET("", profile.ProfileRouter.V2_GetColaborador)
	router_colaborador_v2.DELETE("/:idsubworker", profile.ProfileRouter.DeleteColaborador)

	/*=======================================*/
	/*=======ADDITONAL CONFIGURATIONS========*/
	/*=======================================*/

	//V1 TO EXPORT
	router_export := version_1.Group("/export")
	router_export.POST("", export.ExportRouter.ExportIDDevice)

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

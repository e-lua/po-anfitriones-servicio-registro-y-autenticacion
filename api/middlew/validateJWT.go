package api

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataerror"`
	Data      string `json:"data"`
}

/*
func ValidoJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, idbusines, error_token_process := login.TryingLogin_Service(c.Request().Header.Get("Authorization"))
		if error_token_process != nil {

			results := Response{Error: true, DataError: "Token erroneo, detalle: " + error_token_process.Error(), Data: idbusines}
			return c.JSON(406, results)
		}
		return next(c)
	}
}*/

package login

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	//MDOELS
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	worker_reposiroty "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/repositories/worker"
)

//FUNCIONES PRIVADAS
func compareToken(inputpassword_worker_founded string, inputpassword string) error {

	//brypt trabaja con slices de bytes - Esta password no esta encriptada
	passwordBytes := []byte(inputpassword)

	//Password se encripta
	passwordBD := []byte(inputpassword_worker_founded)

	//Comparamos si el hash encriptado es el password que se escribio en el Login
	error_compare_hash := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if error_compare_hash != nil {
		return error_compare_hash
	}

	return nil
}

func generateJWT(anfitrion models.Pg_BusinessWorker) (string, error) {
	miClave := []byte("TokenGeneradorRestoner")

	payload := jwt.MapClaims{
		"business":    anfitrion.IdBusiness,
		"worker":      anfitrion.IdWorker,
		"rol":         anfitrion.IdRol,
		"country":     anfitrion.IdCountry,
		"sessioncode": anfitrion.SessionCode,
		"exp":         time.Now().Add(time.Hour * 1460).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	//Se añade el string de firma para completar los 3 campos que se pide en http...
	tokenStr, err_signedString := token.SignedString(miClave)
	if err_signedString != nil {
		return tokenStr, err_signedString
	}

	return tokenStr, nil
}

//FUNCIONES PUBLICAS

func TryingLogin_Service(inpuToken string, inputService string, inputModule string, inputEpic string, inputEndpoint string) (JWTRequest, bool, string, string) {

	//Generador de token
	miClave := []byte("TokenGeneradorRestoner")
	claims := &models.Claim{}

	var anfitrionjwt JWTRequest

	/*splitToken para quitar el Bearer del token, por lo tanto
	Bearer sera 0 y el token sera 1*/
	/*splitToken := strings.Split(inpuToken, "Bearer")

	var anfitrion models.Pg_BusinessWorker

	if len(splitToken) != 2 {
		return anfitrion, false, string(""), errors.New("formato de token inválido")
	}

	//Quitar espacios
	inpuToken = strings.TrimSpace(splitToken[1])*/

	token, error_parse := jwt.ParseWithClaims(inpuToken, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if error_parse == nil {

		if claims.IDRol == 1 {
			//Buscamos la existencia del registro en Pg - Redis
			idbusiness_and_sessioncode, error_get_re := worker_reposiroty.Re_Get_Id(claims.Worker, claims.Country, claims.Business)
			if idbusiness_and_sessioncode != strconv.Itoa(claims.Worker)+strconv.Itoa(claims.SessionCode)+strconv.Itoa(claims.Business) {
				return anfitrionjwt, true, "N", "sesion inválida"
			}

			if error_get_re != nil {
				idworker, error_findworker := worker_reposiroty.Pg_Find_ById_TryLogin(claims.Worker, claims.Country)
				if error_findworker != nil {
					return anfitrionjwt, true, "N", error_findworker.Error()
				}
				//Registramos en Redis
				_, err_add_re := worker_reposiroty.Re_Set_ID(idworker, claims.Country, claims.SessionCode, claims.Business)
				if err_add_re != nil {
					return anfitrionjwt, true, "N", err_add_re.Error()
				}
			}
		}

		if claims.IDRol == 2 {
			//Buscamos la existencia del registro en Pg - Redis
			idbusiness_and_sessioncode, error_get_re := worker_reposiroty.Re_Get_Email(claims.Worker, claims.SessionCode, claims.IDRol)
			if idbusiness_and_sessioncode != strconv.Itoa(claims.Worker)+strconv.Itoa(claims.SessionCode)+strconv.Itoa(2) {
				return anfitrionjwt, true, "N", "sesion inválida"
			}

			if error_get_re != nil {
				_, error_findworker := worker_reposiroty.Pg_Find_ById(claims.Business, claims.Country)
				if error_findworker != nil {
					return anfitrionjwt, true, "N", error_findworker.Error()
				}
				//Registramos en Redis
				_, err_add_re := worker_reposiroty.Re_Set_ID(claims.Worker, claims.Country, claims.SessionCode, claims.Business)
				if err_add_re != nil {
					return anfitrionjwt, true, "N", err_add_re.Error()
				}
			}
		}

		anfitrionjwt.IdBusiness = claims.Business
		anfitrionjwt.IdWorker = claims.Worker
		anfitrionjwt.IdCountry = claims.Country
		anfitrionjwt.IdRol = claims.IDRol
		anfitrionjwt.Service = inputService
		anfitrionjwt.Module = inputModule
		anfitrionjwt.Epic = inputEpic
		anfitrionjwt.Endpoint = inputEndpoint

		return anfitrionjwt, false, "N", ""
	}

	//Si el token es inválido
	if !token.Valid {
		return anfitrionjwt, true, "N", "token invalido"
	}

	return anfitrionjwt, true, "N", error_parse.Error()
}

func Login_Service(inputanfitrion models.Pg_BusinessWorker) (int, bool, string, JWTAndRol) {

	//Variable
	var jwt_and_rol JWTAndRol

	//Buscamos la existencia del registro en Pg
	worker_found, error_findworker := worker_reposiroty.Pg_FindByPhone(inputanfitrion.Phone, inputanfitrion.IdCountry)
	if error_findworker != nil {
		return 500, true, "Error en el servidor interno al intentar buscar el anfitrión, detalle: " + error_findworker.Error(), jwt_and_rol
	}
	if strconv.Itoa(worker_found.Phone) == "" || worker_found.IsDeleted {
		return 404, true, "666" + "Este anfitrion no se encuentra registrado", jwt_and_rol
	}
	if worker_found.Isbanned {
		return 404, true, "777" + "Este anfitrión se encuentra baneado", jwt_and_rol
	}

	//Intentamos el login
	error_compareToken := compareToken(worker_found.Password, inputanfitrion.Password)
	if error_compareToken != nil {
		return 403, true, "888" + "Telefono y/o Contraseña incorrectos, detalle: " + error_compareToken.Error(), jwt_and_rol
	}

	jwtKey, error_generatingJWT := generateJWT(worker_found)
	if error_generatingJWT != nil {
		return 500, true, "Error en el servidor interno al intentar generar el token, detalle: " + error_generatingJWT.Error(), jwt_and_rol
	}

	//Registramos en Redis
	_, err_add_re := worker_reposiroty.Re_Set_ID(worker_found.IdWorker, worker_found.IdCountry, worker_found.SessionCode, worker_found.IdBusiness)
	if err_add_re != nil {
		return 500, true, "Error en el servidor interno al intentar registrar el código en cache, detalle: " + err_add_re.Error(), jwt_and_rol
	}

	jwt_and_rol.JWT = jwtKey
	jwt_and_rol.Rol = worker_found.IdRol
	jwt_and_rol.Phone = inputanfitrion.Phone
	jwt_and_rol.Country = worker_found.IdCountry
	jwt_and_rol.Name = worker_found.Name
	jwt_and_rol.Lastname = worker_found.LastName
	jwt_and_rol.ID = worker_found.IdBusiness

	return 201, false, "", jwt_and_rol
}

/*=======================================*/
/*===============VERSION 2===============*/
/*=======================================*/

func V2_Login_Service(input_login Input_BusinessWorker_login) (int, bool, string, JWTAndRol) {

	//Variable
	var jwt_and_rol JWTAndRol

	if input_login.IsAnfitrion {
		//SI ES ANFITRION
		worker_found, error_findworker := worker_reposiroty.Pg_FindByPhone(input_login.Phone, input_login.IdCountry)
		if error_findworker != nil {
			return 500, true, "Error en el servidor interno al intentar buscar el anfitrión, detalle: " + error_findworker.Error(), jwt_and_rol
		}
		if strconv.Itoa(worker_found.Phone) == "" || worker_found.IsDeleted {
			return 404, true, "666" + "Este anfitrion no se encuentra registrado", jwt_and_rol
		}
		if worker_found.Isbanned {
			return 404, true, "777" + "Este anfitrión se encuentra baneado", jwt_and_rol
		}

		//Intentamos el login
		error_compareToken := compareToken(worker_found.Password, input_login.Password)
		if error_compareToken != nil {
			return 403, true, "888" + "Telefono y/o Contraseña incorrectos, detalle: " + error_compareToken.Error(), jwt_and_rol
		}

		jwtKey, error_generatingJWT := generateJWT(worker_found)
		if error_generatingJWT != nil {
			return 500, true, "Error en el servidor interno al intentar generar el token, detalle: " + error_generatingJWT.Error(), jwt_and_rol
		}

		//Registramos en Redis
		_, err_add_re := worker_reposiroty.Re_Set_ID(worker_found.IdWorker, worker_found.IdCountry, worker_found.SessionCode, worker_found.IdBusiness)
		if err_add_re != nil {
			return 500, true, "Error en el servidor interno al intentar registrar el código en cache, detalle: " + err_add_re.Error(), jwt_and_rol
		}

		jwt_and_rol.JWT = jwtKey
		jwt_and_rol.Rol = worker_found.IdRol
		jwt_and_rol.Phone = input_login.Phone
		jwt_and_rol.Country = worker_found.IdCountry
		jwt_and_rol.Name = worker_found.Name
		jwt_and_rol.Lastname = worker_found.LastName
		jwt_and_rol.Email = "anfitrion - sin email"
		jwt_and_rol.ID = worker_found.IdBusiness

	} else {
		//SI NO ES ANFITRION
		subworker_found, error_findsubworker := worker_reposiroty.Pg_FindByEmail(input_login.Email)
		if error_findsubworker != nil {
			return 500, true, "Error en el servidor interno al intentar buscar el colaborador, detalle: " + error_findsubworker.Error(), jwt_and_rol
		}
		if len(subworker_found.Email) < 2 || subworker_found.IsDeleted {
			return 404, true, "666" + "Este anfitrion no se encuentra registrado", jwt_and_rol
		}
		if subworker_found.Isbanned {
			return 404, true, "777" + "Este anfitrión se encuentra baneado", jwt_and_rol
		}

		//Intentamos el login
		error_compareToken := compareToken(subworker_found.Password, input_login.Password)
		if error_compareToken != nil {
			return 403, true, "888" + "Telefono y/o Contraseña incorrectos, detalle: " + error_compareToken.Error(), jwt_and_rol
		}

		jwtKey, error_generatingJWT := generateJWT(subworker_found)
		if error_generatingJWT != nil {
			return 500, true, "Error en el servidor interno al intentar generar el token, detalle: " + error_generatingJWT.Error(), jwt_and_rol
		}

		//Registramos en Redis
		err_add_re := worker_reposiroty.Re_Set_Email(subworker_found.IdWorker, subworker_found.SessionCode, subworker_found.IdRol)
		if err_add_re != nil {
			return 500, true, "Error en el servidor interno al intentar registrar el código en cache, detalle: " + err_add_re.Error(), jwt_and_rol
		}

		jwt_and_rol.JWT = jwtKey
		jwt_and_rol.Rol = subworker_found.IdRol
		jwt_and_rol.Phone = subworker_found.Phone
		jwt_and_rol.Country = subworker_found.IdCountry
		jwt_and_rol.Name = subworker_found.Name
		jwt_and_rol.Lastname = subworker_found.LastName
		jwt_and_rol.Email = input_login.Email
		jwt_and_rol.ID = subworker_found.IdBusiness
	}

	return 201, false, "", jwt_and_rol
}

package solicitud_plan

import (
	"log"

	//TWILIO
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Anfitriones_SendRequest_Service(idbusiness int, timezone string) (int, bool, string, string) {

	//Enviamos el codigo al anfitrion
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: "ACeb643456bb0e06813948315b95c3aa98",
		Password: "b6febb18bf85369763c4a303937137d9",
	})
	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:+51938488229")
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("Hello from Golang!")

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		log.Println("Error Twilio---->", err.Error())
		return 200, false, "", "Solicitud enviada correctamente"
	}

	return 200, false, "", "Solicitud enviada correctamente"
}

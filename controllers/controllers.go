package controllers

import (
	"log"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/gofiber/fiber/v2"
)

// TwilioIncoming represents the structure of incoming Twilio webhook data
type TwilioIncoming struct {
	ChannelPrefix     string
	ApiVersion        string
	MessageStatus     string
	SmsSid            string
	SmsStatus         string
	ChannelInstallSid string
	To                string
	From              string
	MessageSid        string
	StructuredMessage bool
	AccountSid        string
	ChannelToAddress  string
}

// Incoming handles incoming Twilio webhook requests
func Incoming(c *fiber.Ctx) error {
	log.Printf(">> %s: %s\n", c.Method(), string(c.Request().RequestURI()))

	var body TwilioIncoming

	err := c.BodyParser(&body)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return err
	}

	log.Printf(">> BODY: \n%-v\n", body)

	err = getMessagePrice(body.MessageSid)

	return err
}

func getMessagePrice(messageSid string) error {
	client := twilio.NewRestClient()

	params := &api.FetchMessageParams{}

	resp, err := client.Api.FetchMessage(messageSid, params)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return err
	}

	log.Printf("Price: %d %v \n\n", resp.Price, *resp.PriceUnit)

	return nil
}

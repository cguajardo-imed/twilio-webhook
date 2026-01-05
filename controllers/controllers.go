package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/gofiber/fiber/v2"
)

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

func Incoming(c *fiber.Ctx) error {
	var body TwilioIncoming
	err := c.BodyParser(&body)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Incoming: \nBODY: %-v\n", body)

	client := twilio.NewRestClient()

	params := &api.FetchMessageParams{}

	resp, err := client.Api.FetchMessage(body.MessageSid,
		params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		log.Printf("Price: %d %v", resp.Price, *resp.PriceUnit)
	}

	return nil
}

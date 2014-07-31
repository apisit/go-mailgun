package mailgun

import (
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {
	var domain = "socialprintstudio.com"
	var apiKey = "key"
	mailgun := Init(domain, apiKey)
	var message Message
	message.From = "apisit@sps.io"
	message.To = "apisit@socialprintstudio.com"
	message.Subject = "test from golang"
	message.Html = "<b>test html</b>"
	res, err := mailgun.Send(message)
	fmt.Printf("%s", res)
	if err != nil {
		t.Error(fmt.Sprintf("%s", err))
	}

}

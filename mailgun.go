package mailgun

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Message struct {
	From    string
	To      string
	Cc      string
	Bcc     string
	Subject string
	Text    string
	Html    string
}

var baseUri = "https://api.mailgun.net/v2"

type Client struct {
	ApiKey string
	Domain string
}

func Init(domain, apiKey string) *Client {
	return &Client{apiKey, domain}
}

func (m Message) ResourceName() string {
	return "messages"
}

func (m Message) Body() io.Reader {
	values := make(url.Values)
	values.Set("from", m.From)
	values.Set("to", m.To)
	values.Set("subject", m.Subject)
	values.Set("text", m.Text)
	values.Set("html", m.Html)
	return strings.NewReader(values.Encode())
}

func (m Message) Isvalid() bool {
	if m.From == "" || m.To == "" || m.Subject == "" || (m.Text == "" && m.Html == "") {
		return false
	}
	return true
}
func (m Message) EndPoint(c Client) string {
	return fmt.Sprintf("%s/%s", c.EndPoint(), m.ResourceName())
}

func (client Client) EndPoint() string {
	return fmt.Sprintf("%s/%s", baseUri, client.Domain)
}

func (client Client) Send(message Message) (response string, err error) {
	if !message.Isvalid() {
		var errorInvalidMessage = errors.New("Invalid message")
		return "", errorInvalidMessage
	}
	req, _ := http.NewRequest("POST", message.EndPoint(client), message.Body())
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("api", client.ApiKey)
	httpClient := &http.Client{}
	res, _ := httpClient.Do(req)
	defer res.Body.Close()
	responseBytes, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		return "", readErr
	}
	return string(responseBytes), nil
}

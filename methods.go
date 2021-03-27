package telebbb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	Store all telegram methods list to call to

	Excample of request:
	https://api.telegram.org/bot<token>/METHOD_NAME
	https://api.telegram.org/bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11/getMe
*/

const (
	// URL contains main telegram url to place our calls
	URL = "https://api.telegram.org/bot%s/%s"
	// GetUpdates Returns new messages and incoming updates to bot
	GetUpdates = "getUpdates"
	// LogOut Use this method to log out from the cloud Bot API server before launching the bot locally
	LogOut = "logOut"
	// Close Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart.
	Close = "close"
	// SendMessage Use this method to send text messages. On success, the sent Message is returned. Read about Message formats you can use on: https://core.telegram.org/bots/api#formatting-options
	SendMessage = "sendMessage"
	// ForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned.
	ForwardMessage = "forwardMessage"
	// CopyMessage same as forward message but without a link to message source. Posts as new message
	CopyMessage = "copyMessage"
)

// GetMe returns User infor about our bot
func (t *TbBot) GetMe() (u *User, e error) {
	type responce struct {
		IsOk bool `json:"ok,omitempty"`
		Type User `json:"result,omitempty"`
	}
	// Recreate all in one local struct
	var m responce
	req, e := http.NewRequest("GET", fmt.Sprintf(URL, t.token, "getMe"), nil)
	if e != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	r, e := t.client.Do(req)
	if e != nil {
		return
	}
	defer r.Body.Close()
	d, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return
	}
	if e = json.Unmarshal(d, &m); e != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		e = fmt.Errorf("we got invalid status code responce, code responce is %d", r.StatusCode)
		return
	}
	if m.IsOk {
		u = &m.Type
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", m)
	}
	return
}

// SendMessage Sends message, we take interface{} you can send any struct as message but we recomend to use SendMessageType type as a message to avoid error responce from Telegram API will return Message type
func (t *TbBot) SendMessage(message interface{}) (m Message, e error) {
	type responce struct {
		IsOk bool    `json:"ok,omitempty"`
		Type Message `json:"result,omitempty"`
	}
	fmt.Println(message)
	mrsh, e := json.Marshal(message)
	if e != nil {
		return
	}
	b := bytes.NewReader(mrsh)
	req, e := http.NewRequest("POST", fmt.Sprintf(URL, t.token, "sendMessage"), b)
	if e != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	hresp, e := t.client.Do(req)
	if e != nil {
		return
	}
	defer hresp.Body.Close()
	if hresp.StatusCode != http.StatusOK {
		e = fmt.Errorf("we got invalid status code responce, code responce is %d", hresp.StatusCode)
		return
	}
	d, e := ioutil.ReadAll(hresp.Body)
	if e != nil {
		return
	}
	// Working with ressponce
	var r responce
	if e = json.Unmarshal(d, &r); e != nil {
		return
	}
	if r.IsOk {
		m = r.Type
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", m)
	}
	return
}

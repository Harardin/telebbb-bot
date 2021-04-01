package telebbb

import "fmt"

// PostMessage Sends message by specifyed telegram method in mehod param
// This method can be used if don't want to chose from different method and just want to send message as it is
func (t *TbBot) PostMessage(message interface{}, method string) (a interface{}, e error) {
	switch method {
	case "sendMessage":
		r, e := t.SendMessage(message)
		if e != nil {
			return nil, e
		}
		return r, nil

	// TODO
	// Other methods
	default:
		e = fmt.Errorf("invalid method type")
		return nil, e
	}
}

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
	case "forwardMessage":
		r, e := t.ForwardMessage(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "copyMessage":
		r, e := t.CopyMessage(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "sendLocation":
		r, e := t.SendLocation(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "editMessageLiveLocation":
		r, e := t.EditMessageLiveLocation(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "stopMessageLiveLocation":
		r, e := t.StopMessageLiveLocation(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "sendVenue":
		r, e := t.SendVenue(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "sendContact":
		r, e := t.SendContact(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "sendPoll":
		r, e := t.SendPoll(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "sendDice":
		r, e := t.SendDice(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "sendChatAction":
		r, e := t.SendChatAction(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "getUserProfilePhotos":
		r, e := t.GetUserProfilePhotos(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "getFile":
		r, e := t.GetFile(message)
		if e != nil {
			return nil, e
		}
		return r, nil
	case "kickChatMember":
		r, e := t.KickChatMember(message)
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

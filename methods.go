package telebbb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

type responce struct {
	IsOk bool    `json:"ok,omitempty"`
	Type Message `json:"result,omitempty"`
}

// GetMe returns User infor about our bot
func (t *TbBot) GetMe() (u *User, e error) {
	type responce struct {
		IsOk bool `json:"ok,omitempty"`
		Type User `json:"result,omitempty"`
	}
	resp, e := t.sendGet("getMe")
	if e != nil {
		return
	}
	// Recreate all in one local struct
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		u = &r.Type
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// SendMessage Sends message, we take interface{} you can send any struct as message but we recomend to use SendMessageType type as a message to avoid error responce from Telegram API will return Message type
func (t *TbBot) SendMessage(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "forwardMessage")
	if e != nil {
		return
	}
	// Working with ressponce
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = &r.Type
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// ForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned. Accepts type ForwardMessageType call it from tbb library
func (t *TbBot) ForwardMessage(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "forwardMessage")
	if e != nil {
		return
	}
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = &r.Type
		return
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// CopyMessage Use this method to copy messages of any kind. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success. Accept CopyMessageType struct, but can also take any interface.
func (t *TbBot) CopyMessage(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "copyMessage")
	if e != nil {
		return
	}
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = &r.Type
		return
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// SendPhoto Use this method to send photos. On success, the sent Message is returned. Send nil if using ID of a file or file not using direct file to send
func (t *TbBot) SendPhoto(message interface{}, file *os.File) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	var resp []byte
	if file != nil {
		resp, e = t.uploadFile(file, "sendPhoto", "photo", message)
		if e != nil {
			return nil, e
		}
	} else {
		resp, e = t.sendPost(message, "sendPhoto")
		if e != nil {
			return
		}
	}
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = &r.Type
		return
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// SendAudio Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future. For sending voice messages, use the sendVoice method instead. Accepts SendAudioType type as a struct, but can accept interface if needed
func (t *TbBot) SendAudio(message interface{}, file *os.File) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	var resp []byte
	if file != nil {
		resp, e = t.uploadFile(file, "sendAudio", "audio", message)
		if e != nil {
			return nil, e
		}
	} else {
		resp, e = t.sendPost(message, "sendAudio")
		if e != nil {
			return
		}
	}
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = &r.Type
		return
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// SendDocument Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future. Accepts SendDocumentType as message. but can accept interface if needed
func (t *TbBot) SendDocument(message interface{}, file *os.File) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	var resp []byte
	if file != nil {
		resp, e = t.uploadFile(file, "sendDocument", "document", message)
		if e != nil {
			return nil, e
		}
	} else {
		resp, e = t.sendPost(message, "sendDocument")
		if e != nil {
			return
		}
	}
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = &r.Type
		return
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// SendVideo Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future. Accepts SendVideoType as a struct. but can accept interface if needed
func (t *TbBot) SendVideo(message interface{}, file *os.File) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	var resp []byte
	if file != nil {
		resp, e = t.uploadFile(file, "sendVideo", "video", message)
		if e != nil {
			return nil, e
		}
	} else {
		resp, e = t.sendPost(message, "sendVideo")
		if e != nil {
			return
		}
	}
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = &r.Type
		return
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// SendAnimation Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future. Accepts SendAnimationType struct, but can accept interface if needed
func (t *TbBot) SendAnimation(message interface{}, file *os.File) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	var resp []byte
	if file != nil {
		resp, e = t.uploadFile(file, "sendAnimation", "animation", message)
		if e != nil {
			return nil, e
		}
	} else {
		resp, e = t.sendPost(message, "sendAnimation")
		if e != nil {
			return
		}
	}
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = &r.Type
		return
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// ----------------------------------

// Additional functions -------------

func (t *TbBot) sendGet(method string) ([]byte, error) {
	req, e := http.NewRequest("GET", fmt.Sprintf(URL, t.token, method), nil)
	if e != nil {
		return nil, e
	}
	req.Header.Set("Content-Type", "application/json")
	r, e := t.client.Do(req)
	if e != nil {
		return nil, e
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		e = fmt.Errorf("we got invalid status code responce, code responce is %d", r.StatusCode)
		return nil, e
	}
	d, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func (t *TbBot) sendPost(data interface{}, method string) ([]byte, error) {
	mrsh, e := json.Marshal(data)
	if e != nil {
		return nil, e
	}
	b := bytes.NewReader(mrsh)
	req, e := http.NewRequest("POST", fmt.Sprintf(URL, t.token, method), b)
	if e != nil {
		return nil, e
	}
	req.Header.Set("Content-Type", "application/json")
	resp, e := t.client.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		e = fmt.Errorf("we got invalid status code responce, code responce is %d", resp.StatusCode)
		return nil, e
	}
	d, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func (t *TbBot) uploadFile(file *os.File, method string, name string, message interface{}) ([]byte, error) {
	if file == nil {
		return nil, fmt.Errorf("can't upload file that don't exists")
	}
	defer file.Close()
	buff := &bytes.Buffer{}
	writer := multipart.NewWriter(buff)

	// Adding message params
	msg, e := json.Marshal(message)
	if e != nil {
		return nil, e
	}
	var params map[string]interface{}
	if e = json.Unmarshal(msg, &params); e != nil {
		return nil, e
	}
	for k, v := range params {
		var field string
		switch d := v.(type) {
		case float64, float32:
			field = fmt.Sprintf("%f", d)
		default:
			field = fmt.Sprint(d)
		}
		if e = writer.WriteField(k, field); e != nil {
			return nil, e
		}
	}

	boundary := writer.Boundary()
	contentType := "multipart/form-data; boundary=" + boundary
	closeBoundary := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	closeBuffer := bytes.NewBufferString(closeBoundary)
	fi, e := os.Stat(filepath.Base(file.Name()))
	if e != nil {
		return nil, e
	}
	size := fi.Size()
	if _, e := writer.CreateFormFile(name, filepath.Base(file.Name())); e != nil {
		return nil, e
	}
	// Setup request
	req, e := http.NewRequest("POST", fmt.Sprintf(URL, t.token, method), nil)
	if e != nil {
		return nil, e
	}
	req.Body = ioutil.NopCloser(io.MultiReader(buff, file, closeBuffer))
	req.Header.Add("Content-Type", contentType)
	req.ContentLength = size + int64(buff.Len()) + int64(closeBuffer.Len())
	// Making request
	resp, e := t.client.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	data, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	return data, nil
}

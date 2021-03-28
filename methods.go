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

// URL contains main telegram url to place our calls
const URL = "https://api.telegram.org/bot%s/%s"

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

// SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future. Accepts SendVoiceType struct, but can accept interface if needed
func (t *TbBot) SendVoice(message interface{}, file *os.File) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	var resp []byte
	if file != nil {
		resp, e = t.uploadFile(file, "sendVoice", "voice", message)
		if e != nil {
			return nil, e
		}
	} else {
		resp, e = t.sendPost(message, "sendVoice")
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

// SendVideoNote As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned. Accepts SendVideoNoteType struct, but can accept interface if needed
func (t *TbBot) SendVideoNote(message interface{}, file *os.File) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	var resp []byte
	if file != nil {
		resp, e = t.uploadFile(file, "sendVideoNote", "videonote", message)
		if e != nil {
			return nil, e
		}
	} else {
		resp, e = t.sendPost(message, "sendVideoNote")
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

// SendMediaGroup Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned. Accepts SendMediaGroupType struct, but can accept interface if needed
func (t *TbBot) SendMediaGroup(message interface{}, file *os.File) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	var resp []byte
	if file != nil {
		resp, e = t.uploadFile(file, "sendMediaGroup", "mediagroup", message)
		if e != nil {
			return nil, e
		}
	} else {
		resp, e = t.sendPost(message, "sendMediaGroup")
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

// SendLocation Use this method to send point on the map. On success, the sent Message is returned. Accepts SendLocationType struct, but can accept interface if needed
func (t *TbBot) SendLocation(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "sendLocation")
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

// EditMessageLiveLocation Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Accepts EditMessageLiveLocationType struct, but can accept interface if needed.
func (t *TbBot) EditMessageLiveLocation(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "editMessageLiveLocation")
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

// StopMessageLiveLocation Use this method to stop updating a live location message before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned. Accepts StopMessageLiveLocationType struct, but can accept interface if needed.
func (t *TbBot) StopMessageLiveLocation(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "stopMessageLiveLocation")
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

// SendVenue Use this method to send information about a venue. On success, the sent Message is returned. Accepts SendVenue struct, but can accept interface if needed.
func (t *TbBot) SendVenue(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "sendVenue")
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

// SendContact Use this method to send phone contacts. On success, the sent Message is returned. Accepts SendContactType struct, but can accept interface if needed.
func (t *TbBot) SendContact(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "sendContact")
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

// SendPoll Use this method to send a native poll. On success, the sent Message is returned. Accepts SendPollType struct, but can accept interface if needed.
func (t *TbBot) SendPoll(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "sendPoll")
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

// SendDice Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned. Accepts SendDiceType struct, but can accept interface if needed.
func (t *TbBot) SendDice(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "sendDice")
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

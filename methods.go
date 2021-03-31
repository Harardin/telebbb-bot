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
	// Working with responce
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
	// Working with responce
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
	// Working with responce
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
	// Working with responce
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
	// Working with responce
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
	// Working with responce
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
	// Working with responce
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
	// Working with responce
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

// SendChatAction Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success. Accepts SendChatActionType struct, but can accept interface if needed.
func (t *TbBot) SendChatAction(message interface{}) (m *Message, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "sendChatAction")
	if e != nil {
		return
	}
	// Working with responce
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

// GetUserProfilePhotos Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (t *TbBot) GetUserProfilePhotos(message interface{}) (m *UserProfilePhotos, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "getUserProfilePhotos")
	if e != nil {
		return
	}
	// Working with responce
	type userProf struct {
		IsOk bool              `json:"ok,omitempty"`
		Type UserProfilePhotos `json:"result,omitempty"`
	}
	var r userProf
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

// GetFile Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
func (t *TbBot) GetFile(message interface{}) (m *File, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "getFile")
	if e != nil {
		return
	}
	// Working with responce
	type file struct {
		IsOk bool `json:"ok,omitempty"`
		Type File `json:"result,omitempty"`
	}
	var r file
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

// KickChatMember Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success. Accepts KickChatMemberType or any interface
func (t *TbBot) KickChatMember(message interface{}) (m bool, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "kickChatMember")
	if e != nil {
		return
	}
	// Working with responce
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	m = r.IsOk
	return
}

// UnbanChatMember Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
/*
	Parameter 		Type 				Required 	Description
	chat_id 		Integer or String 	Yes 		Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
	user_id 		Integer 			Yes 		Unique identifier of the target user
	only_if_banned 	Boolean 			Optional 	Do nothing if the user is not banned
*/
func (t *TbBot) UnbanChatMember(message interface{}) (m bool, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "unbanChatMember")
	if e != nil {
		return
	}
	// Working with responce
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	m = r.IsOk
	return
}

// RestrictChatMember Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
/*
	Parameter 		Type 				Required 	Description
	chat_id 		Integer or String 	Yes 		Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	user_id 		Integer 			Yes 		Unique identifier of the target user
	permissions 	ChatPermissions 	Yes 		A JSON-serialized object for new user permissions
	until_date 		Integer 			Optional 	Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
*/
func (t *TbBot) RestrictChatMember(message interface{}) (m bool, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "restrictChatMember")
	if e != nil {
		return
	}
	// Working with responce
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	m = r.IsOk
	return
}

// PromoteChatMember Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to demote a user. Returns True on success. Accepts PromoteChatMemberType struct but also can take interface
func (t *TbBot) PromoteChatMember(message interface{}) (m bool, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "promoteChatMember")
	if e != nil {
		return
	}
	// Working with responce
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	m = r.IsOk
	return
}

// SetChatAdministratorCustomTitle Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
/*
	Parameter 		Type 				Required 	Description
	chat_id 		Integer or String 	Yes 		Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	user_id 		Integer 			Yes 		Unique identifier of the target user
	custom_title 	String 				Yes 		New custom title for the administrator; 0-16 characters, emoji are not allowed
*/
func (t *TbBot) SetChatAdministratorCustomTitle(message interface{}) (m bool, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "setChatAdministratorCustomTitle")
	if e != nil {
		return
	}
	// Working with responce
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	m = r.IsOk
	return
}

// SetChatPermissions Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members admin rights. Returns True on success. Accepts SetChatPermissionsType struct, but also can take an interface.
func (t *TbBot) SetChatPermissions(message interface{}) (m bool, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "setChatPermissions")
	if e != nil {
		return
	}
	// Working with responce
	var r responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	m = r.IsOk
	return
}

// ExportChatInviteLink Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite link as String on success.
/*
	Parameter 	Type 				Required 	Description
	chat_id 	Integer or String 	Yes 		Unique identifier for the target chat or username of the target channel (in the format @channelusername)
*/
func (t *TbBot) ExportChatInviteLink(message interface{}) (m interface{}, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "exportChatInviteLink")
	if e != nil {
		return
	}
	type link struct {
		IsOk bool        `json:"ok,omitempty"`
		Type interface{} `json:"result,omitempty"`
	}
	var r link
	// Working with responce
	if e = json.Unmarshal(resp, &r); e != nil {
		return
	}
	if r.IsOk {
		m = r.Type
	} else {
		e = fmt.Errorf("we got 200 responce but have false in status returned struct %+v", r)
	}
	return
}

// CreateChatInviteLink Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object. Returns ChatInviteLinkType type
/*
	Parameter 		Type 				Required 	Description
	chat_id 		Integer or String 	Yes 		Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	expire_date 	Integer 			Optional 	Point in time (Unix timestamp) when the link will expire
	member_limit 	Integer 			Optional 	Maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
*/
func (t *TbBot) CreateChatInviteLink(message interface{}) (m *ChatInviteLinkType, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "createChatInviteLink")
	if e != nil {
		return
	}
	type link struct {
		IsOk bool               `json:"ok,omitempty"`
		Type ChatInviteLinkType `json:"result,omitempty"`
	}
	var r link
	// Working with responce
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

// EditChatInviteLink Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the edited invite link as a ChatInviteLink object.
/*
	Parameter 		Type 				Required 	Description
	chat_id 		Integer or String 	Yes 		Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	invite_link 	String 				Yes 		The invite link to edit
	expire_date 	Integer 			Optional 	Point in time (Unix timestamp) when the link will expire
	member_limit 	Integer 			Optional 	Maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
*/
func (t *TbBot) EditChatInviteLink(message interface{}) (m *ChatInviteLinkType, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "editChatInviteLink")
	if e != nil {
		return
	}
	type link struct {
		IsOk bool               `json:"ok,omitempty"`
		Type ChatInviteLinkType `json:"result,omitempty"`
	}
	var r link
	// Working with responce
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

// RevokeChatInviteLink Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the revoked invite link as ChatInviteLink object.
/*
	Parameter 		Type 				Required 	Description
	chat_id 		Integer or String 	Yes 		Unique identifier of the target chat or username of the target channel (in the format @channelusername)
	invite_link 	String 				Yes 		The invite link to revoke
*/
func (t *TbBot) RevokeChatInviteLink(message interface{}) (m *ChatInviteLinkType, e error) {
	if message == nil {
		e = fmt.Errorf("message can't be nil")
		return
	}
	resp, e := t.sendPost(message, "revokeChatInviteLink")
	if e != nil {
		return
	}
	type link struct {
		IsOk bool               `json:"ok,omitempty"`
		Type ChatInviteLinkType `json:"result,omitempty"`
	}
	var r link
	// Working with responce
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

// TODO
// Other functions

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

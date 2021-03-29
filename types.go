package telebbb

import "net/http"

// BotConfig Main bot configuration
type BotConfig struct {
	Type string // Type - what bot type to use
	/*
		Options:
		- webhook
		- local
	*/
	Token string // Token - insert your bot token string
	Port  string // Port - port to listen webhook default is :8000
}

// TbBot Main Bot struct to stor all data, and call bot functions
type TbBot struct {
	client   *http.Client
	token    string
	Incoming chan interface{}
	Errors   chan error // Will return error from deep routines to process
}

// ------------------------------
// Telegram types

// --------------------------
// User Object types

// User This object represents a Telegram user or bot.
type User struct {
	ID                int    `json:"id,omitempty"`                          // Unique identifier for this user or bot
	IsBot             bool   `json:"is_bot,omitempty"`                      // True, if this user is a bot
	FirstName         string `json:"first_name,omitempty"`                  // User's or bot's first name
	LastName          string `json:"last_name,omitempty"`                   // Optional. User's or bot's last name
	UserName          string `json:"username,omitempty"`                    // Optional. User's or bot's username
	LangCode          string `json:"language_code,omitempty"`               // Optional. IETF language tag of the user's language
	CanJoinGroups     bool   `json:"can_join_groups,omitempty"`             // Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanReadGroupsMsgs bool   `json:"can_read_all_group_messages,omitempty"` // Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	IsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`     // Optional. True, if the bot supports inline queries. Returned only in getMe.
}

// UserProfilePhotos This object represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int         `json:"total_count,omitempty"` // Total number of profile pictures the target user has
	Photos     []PhotoSize `json:"photos,omitempty"`      // Requested profile pictures (in up to 4 sizes each)
}

// --------------------------
// All Chat objects

// Chat This object represents a chat.
type Chat struct {
	ID                int              `json:"id,omitempty"`                       // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type              string           `json:"type,omitempty"`                     // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title             string           `json:"title,omitempty"`                    // Optional. Title, for supergroups, channels and group chats
	UserName          string           `json:"username,omitempty"`                 // Optional. Username, for private chats, supergroups and channels if available
	FirstName         string           `json:"first_name,omitempty"`               // Optional. First name of the other party in a private chat
	LastName          string           `json:"last_name,omitempty"`                // Optional. Last name of the other party in a private chat
	Photo             *ChatPhoto       `json:"photo,omitempty"`                    // Optional. Chat photo. Returned only in getChat.
	Bio               string           `json:"bio,omitempty"`                      // Optional. Bio of the other party in a private chat. Returned only in getChat.
	Description       string           `json:"description,omitempty"`              // Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink        string           `json:"invite_link,omitempty"`              // Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	PinnedMsg         *Message         `json:"pinned_message,omitempty"`           // Optional. The most recent pinned message (by sending date). Returned only in getChat.
	Permission        *ChatPermissions `json:"permissions,omitempty"`              // Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModDelay      int              `json:"slow_mode_delay,omitempty"`          // Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user. Returned only in getChat.
	MsgAutoDeleteTime int              `json:"message_auto_delete_time,omitempty"` // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
	StickerSetName    string           `json:"sticker_set_name,omitempty"`         // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetSticker     bool             `json:"can_set_sticker_set,omitempty"`      // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	LinkedChatID      int              `json:"linked_chat_id,omitempty"`           // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
	Location          *ChatLocation    `json:"location,omitempty"`                 // Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
}

// ChatPhoto This object represents a chat photo.
type ChatPhoto struct {
	SmallID     string `json:"small_file_id,omitempty"`        // File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallIDUniq string `json:"small_file_unique_id,omitempty"` // Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigID       string `json:"big_file_id,omitempty"`          // File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigIDUniq   string `json:"big_file_unique_id,omitempty"`   // Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}

// ChatPermissions Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMsg           bool `json:"can_send_messages,omitempty"`         // Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMediaMsg      bool `json:"can_send_media_messages,omitempty"`   // Optional. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendPolls         bool `json:"can_send_polls,omitempty"`            // Optional. True, if the user is allowed to send polls, implies can_send_messages
	CanSendOtherMsgs     bool `json:"can_send_other_messages,omitempty"`   // Optional. True, if the user is allowed to send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePrewiew bool `json:"can_add_web_page_previews,omitempty"` // Optional. True, if the user is allowed to add web page previews to their messages, implies can_send_media_messages
	CanChangeInfo        bool `json:"can_change_info,omitempty"`           // Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanInviteUsers       bool `json:"can_invite_users,omitempty"`          // Optional. True, if the user is allowed to invite new users to the chat
	CanPinMsgs           bool `json:"can_pin_messages,omitempty"`          // Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
}

// ChatLocation Represents a location to which a chat is connected.
type ChatLocation struct {
	Loc  Location `json:"location,omitempty"` // The location to which the supergroup is connected. Can't be a live location.
	Addr string   `json:"address,omitempty"`  // Location address; 1-64 characters, as defined by the chat owner
}

// ChatMember This object contains information about one member of a chat.
type ChatMember struct {
	Usr               *User  `json:"user,omitempty"`                      // Information about the user
	Status            string `json:"status,omitempty"`                    // The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
	CustomTitle       string `json:"custom_title,omitempty"`              // Optional. Owner and administrators only. Custom title for this user
	IsAnonymous       bool   `json:"is_anonymous,omitempty"`              // Optional. Owner and administrators only. True, if the user's presence in the chat is hidden
	CanBeEdited       bool   `json:"can_be_edited,omitempty"`             // Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanPostMsg        bool   `json:"can_post_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can post in the channel; channels only
	CanEditMsg        bool   `json:"can_edit_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages; channels only
	CanDeleteMsg      bool   `json:"can_delete_messages,omitempty"`       // Optional. Administrators only. True, if the administrator can delete messages of other users
	CanRestrictMsg    bool   `json:"can_restrict_members,omitempty"`      // Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanPromoteMembers bool   `json:"can_promote_members,omitempty"`       // Optional. Administrators only. True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo     bool   `json:"can_change_info,omitempty"`           // Optional. Administrators and restricted only. True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers    bool   `json:"can_invite_users,omitempty"`          // Optional. Administrators and restricted only. True, if the user is allowed to invite new users to the chat
	CanPinMsg         bool   `json:"can_pin_messages,omitempty"`          // Optional. Administrators and restricted only. True, if the user is allowed to pin messages; groups and supergroups only
	IsMember          bool   `json:"is_member,omitempty"`                 // Optional. Restricted only. True, if the user is a member of the chat at the moment of the request
	CanSendMsg        bool   `json:"can_send_messages,omitempty"`         // Optional. Restricted only. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMedia      bool   `json:"can_send_media_messages,omitempty"`   // Optional. Restricted only. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes
	CanSendPolls      bool   `json:"can_send_polls,omitempty"`            // Optional. Restricted only. True, if the user is allowed to send polls
	CanSendOtherMsg   bool   `json:"can_send_other_messages,omitempty"`   // Optional. Restricted only. True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePrw  bool   `json:"can_add_web_page_previews,omitempty"` // Optional. Restricted only. True, if the user is allowed to add web page previews to their messages
	UntilDate         int    `json:"until_date,omitempty"`                // Optional. Restricted and kicked only. Date when restrictions will be lifted for this user; unix time
}

// --------------------------

// --------------------------
// All Message objects

// Message This object represents a message.
type Message struct {
	MessageID              int              `json:"message_id,omitempty"`              // Unique message identifier inside this chat
	From                   *User            `json:"from,omitempty"`                    // Optional. Sender, empty for messages sent to channels
	SenderChat             *Chat            `json:"sender_chat,omitempty"`             // Optional. Sender of the message, sent on behalf of a chat. The channel itself for channel messages. The supergroup itself for messages from anonymous group administrators. The linked channel for messages automatically forwarded to the discussion group
	Date                   int              `json:"date,omitempty"`                    // Date the message was sent in Unix time
	Chat                   *Chat            `json:"chat,omitempty"`                    // Conversation the message belongs to
	ForwardedFrom          *User            `json:"forward_from,omitempty"`            // Optional. For forwarded messages, sender of the original message
	ForwardedFromChat      *Chat            `json:"forward_from_chat,omitempty"`       // Optional. For messages forwarded from channels or from anonymous administrators, information about the original sender chat
	ForwardedFromMessageID int              `json:"forward_from_message_id,omitempty"` // Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature       string           `json:"forward_signature,omitempty"`       // Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSenderName      string           `json:"forward_sender_name,omitempty"`     // Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
	ForwardDate            int              `json:"forward_date,omitempty"`            // Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage         *Message         `json:"reply_to_message,omitempty"`        // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ViaBot                 *User            `json:"via_bot,omitempty"`                 // Optional. Bot through which the message was sent
	EditDate               int              `json:"edit_date,omitempty"`               // Optional. Date the message was last edited in Unix time
	MediaGroupID           string           `json:"media_group_id,omitempty"`          // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature        string           `json:"author_signature,omitempty"`        // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                   string           `json:"text,omitempty"`                    // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Entities               []*MessageEntity `json:"entities,omitempty"`                // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Animation              *Animation       `json:"animation,omitempty"`               // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                  *Audio           `json:"audio,omitempty"`                   // Optional. Message is an audio file, information about the file
	Document               *Document        `json:"document,omitempty"`                // Optional. Message is a general file, information about the file
	Photo                  []*PhotoSize     `json:"photo,omitempty"`                   // Optional. Message is a photo, available sizes of the photo
	Sticker                *StickerType     `json:"sticker,omitempty"`                 // Optional. Message is a sticker, information about the sticker
	Video                  *Video           `json:"video,omitempty"`                   // Optional. Message is a video, information about the video
	VideoNote              *VideoNote       `json:"video_note,omitempty"`              // Optional. Message is a video note, information about the video message
	Voice                  *Voice           `json:"voice,omitempty"`                   // Optional. Message is a voice message, information about the file
	Caption                string           `json:"caption,omitempty"`                 // Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	CaptionEntities        []*MessageEntity `json:"caption_entities,omitempty"`        // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	Contact                *Contact         `json:"contact,omitempty"`                 // Optional. Message is a shared contact, information about the contact
	Dice                   *Dice            `json:"dice,omitempty"`                    // Optional. Message is a dice with random value

	// TODO
	// Add all Message fields
}

// MessageID This object represents a unique message identifier.
type MessageID struct {
	ID int `json:"message_id,omitempty"` // Unique message identifier
}

// MessageEntity This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   string `json:"type,omitempty"`     // Type of the entity. Can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames)
	Offset int    `json:"offset,omitempty"`   // Offset in UTF-16 code units to the start of the entity
	Length int    `json:"length,omitempty"`   // Length of the entity in UTF-16 code units
	URL    string `json:"url,omitempty"`      // Optional. For “text_link” only, url that will be opened after user taps on the text
	Usr    *User  `json:"user,omitempty"`     // Optional. For “text_mention” only, the mentioned user
	Lang   string `json:"language,omitempty"` // Optional. For “pre” only, the programming language of the entity text
}

// PhotoSize This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID     string `json:"file_id,omitempty"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqID string `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width      int    `json:"width,omitempty"`          // Photo width
	Height     int    `json:"height,omitempty"`         // Photo height
	FileSize   int    `json:"file_size,omitempty"`      // Optional. File size
}

// Animation This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileID     string     `json:"file_id,omitempty"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqID string     `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width      int        `json:"width,omitempty"`          // Video width as defined by sender
	Height     int        `json:"height,omitempty"`         // Video height as defined by sender
	Duration   int        `json:"duration,omitempty"`       // Duration of the video in seconds as defined by sender
	Thumb      *PhotoSize `json:"thumb,omitempty"`          // Optional. Animation thumbnail as defined by sender
	FileName   string     `json:"file_name,omitempty"`      // Optional. Original animation filename as defined by sender
	MIMEType   string     `json:"mime_type,omitempty"`      // Optional. MIME type of the file as defined by sender
	FileSize   int        `json:"file_size,omitempty"`      // Optional. File size
}

// Audio This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID     string     `json:"file_id,omitempty"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqID string     `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration   int        `json:"duration,omitempty"`       // Duration of the audio in seconds as defined by sender
	Performer  string     `json:"performer,omitempty"`      // Optional. Performer of the audio as defined by sender or by audio tags
	Title      string     `json:"title,omitempty"`          // Optional. Title of the audio as defined by sender or by audio tags
	FileName   string     `json:"file_name,omitempty"`      // Optional. Original filename as defined by sender
	MIMEType   string     `json:"mime_type,omitempty"`      // Optional. MIME type of the file as defined by sender
	FileSize   int        `json:"file_size,omitempty"`      // Optional. File size
	Thumb      *PhotoSize `json:"thumb,omitempty"`          // Optional. Thumbnail of the album cover to which the music file belongs
}

// Document This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID     string     `json:"file_id,omitempty"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqID string     `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumb      *PhotoSize `json:"thumb,omitempty"`          // Optional. Document thumbnail as defined by sender
	FileName   string     `json:"file_name,omitempty"`      // Optional. Original filename as defined by sender
	MIMEType   string     `json:"mime_type,omitempty"`      // Optional. MIME type of the file as defined by sender
	FileSize   int        `json:"file_size,omitempty"`      // Optional. File size
}

// Video This object represents a video file.
type Video struct {
	FileID     string     `json:"file_id,omitempty"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqID string     `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width      int        `json:"width,omitempty"`          // Video width as defined by sender
	Height     int        `json:"height,omitempty"`         // Video height as defined by sender
	Duration   int        `json:"duration,omitempty"`       // Duration of the video in seconds as defined by sender
	Thumb      *PhotoSize `json:"thumb,omitempty"`          // Optional. Video thumbnail
	FileName   string     `json:"file_name,omitempty"`      // Optional. Original filename as defined by sender
	MIMEType   string     `json:"mime_type,omitempty"`      // Optional. Mime type of a file as defined by sender
	FileSize   int        `json:"file_size,omitempty"`      // Optional. File size
}

// VideoNote This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileID     string     `json:"file_id,omitempty"`        //  	Identifier for this file, which can be used to download or reuse the file
	FileUniqID string     `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length     int        `json:"length,omitempty"`         // Video width and height (diameter of the video message) as defined by sender
	Duration   int        `json:"duration,omitempty"`       // Duration of the video in seconds as defined by sender
	Thumb      *PhotoSize `json:"thumb,omitempty"`          // Optional. Video thumbnail
	FileSize   int        `json:"file_size,omitempty"`      // Optional. File size
}

// Voice This object represents a voice note.
type Voice struct {
	FileID     string `json:"file_id,omitempty"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqID string `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration   int    `json:"duration,omitempty"`       // Duration of the audio in seconds as defined by sender
	MIMEType   string `json:"mime_type,omitempty"`      // Optional. MIME type of the file as defined by sender
	FileSize   int    `json:"file_size,omitempty"`      // Optional. File size
}

// Contact This object represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number,omitempty"` // Contact's phone number
	FirstName   string `json:"first_name,omitempty"`   // Contact's first name
	LastName    string `json:"last_name,omitempty"`    // Optional. Contact's last name
	UserID      int    `json:"user_id,omitempty"`      // Optional. Contact's user identifier in Telegram
	VCard       string `json:"vcard,omitempty"`        // Optional. Additional data about the contact in the form of a vCard
}

// Dice This object represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji,omitempty"` // Emoji on which the dice throw animation is based
	Value int    `json:"value,omitempty"` // Value of the dice, 1-6 for “🎲” and “🎯” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
}

// PollOption This object contains information about one answer option in a poll.
type PollOption struct {
	Text       string `json:"text,omitempty"`        // Option text, 1-100 characters
	VoterCount int    `json:"voter_count,omitempty"` // Number of users that voted for this option
}

// PollAnswer This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string `json:"poll_id,omitempty"`    // Unique poll identifier
	Usr       *User  `json:"user,omitempty"`       // The user, who changed the answer to the poll
	OptionIDs []int  `json:"option_ids,omitempty"` // 0-based identifiers of answer options, chosen by the user. May be empty if the user retracted their vote.
}

// Poll This object contains information about a poll.
type Poll struct {
	ID            string           `json:"id,omitempty"`                      // Unique poll identifier
	Question      string           `json:"question,omitempty"`                // Poll question, 1-300 characters
	Options       []*PollOption    `json:"options,omitempty"`                 // List of poll options
	VoterCount    int              `json:"total_voter_count,omitempty"`       // Total number of users that voted in the poll
	IsClosed      bool             `json:"is_closed,omitempty"`               // True, if the poll is closed
	IsAnon        bool             `json:"is_anonymous,omitempty"`            // True, if the poll is anonymous
	Type          string           `json:"type,omitempty"`                    // Poll type, currently can be “regular” or “quiz”
	IsMultAnswers bool             `json:"allows_multiple_answers,omitempty"` // True, if the poll allows multiple answers
	CorrecOptID   int              `json:"correct_option_id,omitempty"`       // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	Explanation   string           `json:"explanation,omitempty"`             // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	ExplEntities  []*MessageEntity `json:"explanation_entities,omitempty"`    // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	OpenPeriod    int              `json:"open_period,omitempty"`             // Optional. Amount of time in seconds the poll will be active after creation
	CloseDate     int              `json:"close_date,omitempty"`              // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
}

// --------------------------

// --------------------------
// All location objects

// Location This object represents a point on the map
type Location struct {
	Longitude     float64 `json:"longitude,omitempty"`              // Longitude as defined by sender
	Latitude      float64 `json:"latitude,omitempty"`               // Latitude as defined by sender
	HorizontalAcc float64 `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod    int     `json:"live_period,omitempty"`            // Optional. Time relative to the message sending date, during which the location can be updated, in seconds. For active live locations only.
	Heading       int     `json:"heading,omitempty"`                // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	AlertRadius   int     `json:"proximity_alert_radius,omitempty"` // Optional. Maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}

// Venue This object represents a venue.
type Venue struct {
	Loc             *Location `json:"location,omitempty"`          // Venue location. Can't be a live location
	Title           string    `json:"title,omitempty"`             // Name of the venue
	Address         string    `json:"address,omitempty"`           // Address of the venue
	FoursquareID    string    `json:"foursquare_id,omitempty"`     // Optional. Foursquare identifier of the venue
	FoursquareType  string    `json:"foursquare_type,omitempty"`   // Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceID   string    `json:"google_place_id,omitempty"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string    `json:"google_place_type,omitempty"` // Optional. Google Places type of the venue. (See: https://developers.google.com/places/web-service/supported_types)
}

// ProximityAlertTriggered This object represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler,omitempty"` // User that triggered the alert
	Watcher  *User `json:"watcher,omitempty"`  // User that set the alert
	Distance int   `json:"distance,omitempty"` // The distance between the users
}

// --------------------------

// --------------------------

// File This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
// Maximum file size to download is 20 MB
type File struct {
	FileID     string `json:"file_id,omitempty"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqID string `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize   int    `json:"file_size,omitempty"`      // Optional. File size, if known
	FilePath   string `json:"file_path,omitempty"`      // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

// --------------------------

// --------------------------
// Keyboard button types

// ReplyKeyboardMarkup This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard        []*KeyboardButton `json:"keyboard,omitempty"`          // Array of button rows, each represented by an Array of KeyboardButton objects
	ResizeKeyboard  bool              `json:"resize_keyboard,omitempty"`   // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard bool              `json:"one_time_keyboard,omitempty"` // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	Selective       bool              `json:"selective,omitempty"`         // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

// KeyboardButton This object represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields request_contact, request_location, and request_poll are mutually exclusive.
type KeyboardButton struct {
	Text            string                  `json:"text,omitempty"`             // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestContact  bool                    `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestLocation bool                    `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
	RequestPool     *KeyboardButtonPollType `json:"request_poll,omitempty"`     // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only
}

// KeyboardButtonPollType This object represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"` // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

// ReplyKeyboardRemove Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"` // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective,omitempty"`       // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

// LoginURL This object represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
type LoginURL struct {
	URL            string `json:"url,omitempty"`                  //  	An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.
	ForwardText    string `json:"forward_text,omitempty"`         // Optional. New text of the button in forwarded messages.
	BotUserName    string `json:"bot_username,omitempty"`         // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	ReqWriteAccess bool   `json:"request_write_access,omitempty"` // Optional. Pass True to request the permission for your bot to send messages to the user.
}

// CallbackGame No info for now
type CallbackGame struct{} // A placeholder, currently holds no information. Use BotFather to set up your game.

// InlineKeyboardButton This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text              string    `json:"text,omitempty"`                // Label text on the button
	URL               string    `json:"url,omitempty"`                 // Optional. HTTP or tg:// url to be opened when button is pressed
	LoginURL          *LoginURL `json:"login_url,omitempty"`           // Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	CallbackData      string    `json:"callback_data,omitempty"`       // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery string    `json:"switch_inline_query,omitempty"` // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. Can be empty, in which case just the bot's username will be inserted.
	// Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineChat string `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot's username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"` // Optional. Description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
	Pay bool `json:"pay,omitempty"` // Optional. Specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row.
}

// InlineKeyboardMarkup This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard []*InlineKeyboardButton `json:"inline_keyboard,omitempty"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

// CallbackQuery This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ID            string   `json:"id,omitempty"`                // Unique identifier for this query
	From          *User    `json:"from,omitempty"`              // Sender
	Msg           *Message `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMsg     string   `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance  string   `json:"chat_instance,omitempty"`     // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data          string   `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	GameShortName string   `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

// ForceReply Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool `json:"force_reply,omitempty"` // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	Selective  bool `json:"selective,omitempty"`   // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// --------------------------
// Bot Commands

// BotCommand This object represents a bot command.
type BotCommand struct {
	Command     string `json:"command,omitempty"`     // Text of the command, 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Description string `json:"description,omitempty"` // Description of the command, 3-256 characters.
}

// ResponseParameters Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id,omitempty"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      int `json:"retry_after,omitempty"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

// --------------------------

// --------------------------
// InputMedia Structs

// InputMediaPhoto Represents a photo to be sent.
type InputMediaPhoto struct {
	Type            string           `json:"type,omitempty"`             // Type of the result, must be photo
	Media           string           `json:"media,omitempty"`            // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. https://core.telegram.org/bots/api#sending-files
	Caption         string           `json:"caption,omitempty"`          // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode       string           `json:"parse_mode,omitempty"`       // Optional. Mode for parsing entities in the photo caption. See formatting options for more details. (https://core.telegram.org/bots/api#formatting-options)
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
}

// InputMediaVideo Represents a video to be sent.
type InputMediaVideo struct {
	Type             string           `json:"type,omitempty"`               // Type of the result, must be video
	Media            string           `json:"media,omitempty"`              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. (https://core.telegram.org/bots/api#sending-files)
	Thumb            interface{}      `json:"thumb,omitempty"`              // InputFile(type) or String Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Caption          string           `json:"caption,omitempty"`            // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode        string           `json:"parse_mode,omitempty"`         // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities  []*MessageEntity `json:"caption_entities,omitempty"`   // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Width            int              `json:"width,omitempty"`              // Optional. Video width
	Height           int              `json:"height,omitempty"`             // Optional. Video height
	Duration         int              `json:"duration,omitempty"`           // Optional. Video duration
	SupportStreaming bool             `json:"supports_streaming,omitempty"` // Optional. Pass True, if the uploaded video is suitable for streaming
}

// InputMediaAnimation Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Type            string           `json:"type,omitempty"`             // Type of the result, must be animation
	Media           string           `json:"media,omitempty"`            // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	Thumb           interface{}      `json:"thumb,omitempty"`            // InputFile(type) or String Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Caption         string           `json:"caption,omitempty"`          // Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	ParseMode       string           `json:"parse_mode,omitempty"`       // Optional. Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Width           int              `json:"width,omitempty"`            // Optional. Animation width
	Height          int              `json:"height,omitempty"`           // Optional. Animation height
	Duration        int              `json:"duration,omitempty"`         // Optional. Animation duration
}

// InputMediaAudio Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	Type            string           `json:"type,omitempty"`             // Type of the result, must be audio
	Media           string           `json:"media,omitempty"`            // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	Thumb           interface{}      `json:"thumb,omitempty"`            // InputFile(type) or String Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Caption         string           `json:"caption,omitempty"`          // Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	ParseMode       string           `json:"parse_mode,omitempty"`       // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration        int              `json:"duration,omitempty"`         // Optional. Duration of the audio in seconds
	Performer       string           `json:"performer,omitempty"`        // Optional. Performer of the audio
	Title           string           `json:"title,omitempty"`            // Optional. Title of the audio
}

// InputMediaDocument Represents a general file to be sent.
type InputMediaDocument struct {
	Type                        string           `json:"type,omitempty"`                           // Type of the result, must be document
	Media                       string           `json:"media,omitempty"`                          // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name
	Thumb                       interface{}      `json:"thumb,omitempty"`                          // InputFile(type) or String Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Caption                     string           `json:"caption,omitempty"`                        // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode                   string           `json:"parse_mode,omitempty"`                     // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`               // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always true, if the document is sent as part of an album.
}

// --------------------------

// ------ Types to work with telegram methods

// SendMessageType SendMessageType to send message to telegram
type SendMessageType struct {
	// ChatID string or int
	ChatID                interface{}    `json:"chat_id,omitempty"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                  string         `json:"text,omitempty"`                        // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode             string         `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the message text. See https://core.telegram.org/bots/api#formatting-options for more details.
	Entities              *MessageEntity `json:"entities,omitempty"`                    // Optional.  	List of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePrewiew bool           `json:"disable_web_page_preview,omitempty"`    // Optional. Disables link previews for links in this message
	DisableNotification   bool           `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMsg            int            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowWithoutReply     bool           `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup           interface{}    `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// ForwardMessageType ForwardMessageType Use this method to forward messages of any kind. On success, the sent Message is returned.
type ForwardMessageType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// ForwardChatID string or int
	ForwardChatID       interface{} `json:"from_chat_id,omitempty"`         // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool        `json:"disable_notification,omitempty"` // Optional. Sends the message silently. Users will receive a notification with no sound.
	MessageID           int         `json:"message_id,omitempty"`           // Message identifier in the chat specified in from_chat_id
}

// CopyMessageType Use this method to copy messages of any kind. The method is analogous to the method forwardMessages, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
type CopyMessageType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// FromChatID string or int
	FromChatID               interface{}      `json:"from_chat_id,omitempty"`                // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageID                int              `json:"message_id,omitempty"`                  // Message identifier in the chat specified in from_chat_id
	Caption                  string           `json:"caption,omitempty"`                     // Optional. New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	ParseMode                string           `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. List of special entities that appear in the new caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendPhotoType Use this method to send photos. On success, the sent Message is returned.
type SendPhotoType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// Photo string(file_id) InputFile type
	Photo                    interface{}      `json:"photo,omitempty"`                       // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendAudioType Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
type SendAudioType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// Audio string(fileID) or InputFile type
	Audio           interface{}      `json:"audio,omitempty"`            // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data.
	Caption         string           `json:"caption,omitempty"`          // Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode       string           `json:"parse_mode,omitempty"`       // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration        int              `json:"duration,omitempty"`         // Optional. Duration of the audio in seconds
	Performer       string           `json:"performer,omitempty"`        // Optional. Performer
	Title           string           `json:"title,omitempty"`            // Optional. Title
	// Thumb string(imgID) InputFile type
	Thumb                    interface{} `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int         `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendDocumentType Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocumentType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// Document string(docID) InputFile type
	Document interface{} `json:"document,omitempty"` // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	// Thumb string(fileID) InputFile type
	Thumb                       interface{}      `json:"thumb,omitempty"`                          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Caption                     string           `json:"caption,omitempty"`                        // Optional. Document caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode                   string           `json:"parse_mode,omitempty"`                     // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`               // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableNotification         bool             `json:"disable_notification,omitempty"`           // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID            int              `json:"reply_to_message_id,omitempty"`            // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply    bool             `json:"allow_sending_without_reply,omitempty"`    // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup                 interface{}      `json:"reply_markup,omitempty"`                   // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendVideoType Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideoType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// Video string(file_id) or InputFile type
	Video    interface{} `json:"video,omitempty"`    // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data.
	Duration int         `json:"duration,omitempty"` // Optional. Duration of sent video in seconds
	Width    int         `json:"width,omitempty"`    // Optional. Video width
	Height   int         `json:"height,omitempty"`   // Optional. Video height
	// Thumb string(fileID) InputFile type
	Thumb                    interface{}      `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	SupportsStreaming        bool             `json:"supports_streaming,omitempty"`          // Optional. Pass True, if the uploaded video is suitable for streaming
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendAnimationType Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimationType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// Animation string(file_id) or InputFile type
	Animation interface{} `json:"animation,omitempty"` // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data.
	Duration  int         `json:"duration,omitempty"`  // Optional. Duration of sent video in seconds
	Width     int         `json:"width,omitempty"`     // Optional. Video width
	Height    int         `json:"height,omitempty"`    // Optional. Video height
	// Thumb string(fileID) InputFile type
	Thumb                    interface{}      `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Animation caption (may also be used when resending animations by file_id), 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendVoiceType Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoiceType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// Voice string(file_id) or InputFile type
	Voice                    interface{}      `json:"voice,omitempty"`                       // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Voice message caption, 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration                 int              `json:"duration,omitempty"`                    // Optional. Duration of sent video in seconds
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendVideoNoteType As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNoteType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// VideoNote string(file_id) or InputFile type
	VideoNote interface{} `json:"video_note,omitempty"` // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
	Duration  int         `json:"duration,omitempty"`   // Optional. Duration of sent video in seconds
	Length    int         `json:"length,omitempty"`     // Optional. Video width and height, i.e. diameter of the video message
	// Thumb string(fileID) InputFile type
	Thumb                    interface{} `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int         `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendMediaGroupType Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
type SendMediaGroupType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	// Array of InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo
	Media                    interface{} `json:"media,omitempty"`                       // A JSON-serialized array describing messages to be sent, must include 2-10 items
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int         `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
}

// SendLocationType Use this method to send point on the map. On success, the sent Message is returned.
type SendLocationType struct {
	// ChatID string or int
	ChatID                   interface{} `json:"chat_id,omitempty"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude                 float64     `json:"latitude,omitempty"`                    // Latitude of the location
	Longitude                float64     `json:"longitude,omitempty"`                   // Longitude of the location
	HorizontalAcc            float64     `json:"horizontal_accuracy,omitempty"`         // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod               int         `json:"live_period,omitempty"`                 // Optional. Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	Heading                  int         `json:"heading,omitempty"`                     // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius     int         `json:"proximity_alert_radius,omitempty"`      // For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int         `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// EditMessageLiveLocationType Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocationType struct {
	// ChatID string or int
	ChatID               interface{} `json:"chat_id,omitempty"`                // Optional. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int         `json:"message_id,omitempty"`             // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      string      `json:"inline_message_id,omitempty"`      // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude             float64     `json:"latitude,omitempty"`               // Latitude of the location
	Longitude            float64     `json:"longitude,omitempty"`              // Longitude of the location
	HorizontalAcc        float64     `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	Heading              int         `json:"heading,omitempty"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int         `json:"proximity_alert_radius,omitempty"` // For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          interface{} `json:"reply_markup,omitempty"`           // Optional. A JSON-serialized object for a new inline keyboard.
}

// StopMessageLiveLocationType Use this method to stop updating a live location message before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
type StopMessageLiveLocationType struct {
	// ChatID string or int
	ChatID          interface{} `json:"chat_id,omitempty"`           // Optional. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID       int         `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageID string      `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     interface{} `json:"reply_markup,omitempty"`      // A JSON-serialized object for a new inline keyboard.
}

// SendVenue Use this method to send information about a venue. On success, the sent Message is returned.
type SendVenue struct {
	// ChatID string or int
	ChatID                   interface{} `json:"chat_id,omitempty"`                     // Optional. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude                 float64     `json:"latitude,omitempty"`                    // Latitude of the location
	Longitude                float64     `json:"longitude,omitempty"`                   // Longitude of the location
	Title                    string      `json:"title,omitempty"`                       // Name of the venue
	Address                  string      `json:"address,omitempty"`                     // Address of the venue
	FoursquareID             string      `json:"foursquare_id,omitempty"`               // Optional. Foursquare identifier of the venue
	FoursquareType           string      `json:"foursquare_type,omitempty"`             // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceID            string      `json:"google_place_id,omitempty"`             // Optional. Google Places identifier of the venue
	GooglePlaceType          string      `json:"google_place_type,omitempty"`           // Google Places type of the venue.
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int         `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendContactType Use this method to send phone contacts. On success, the sent Message is returned.
type SendContactType struct {
	// ChatID string or int
	ChatID                   interface{} `json:"chat_id,omitempty"`                     // Optional. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	PhoneNumber              string      `json:"phone_number,omitempty"`                // Contact's phone number
	FirstName                string      `json:"first_name,omitempty"`                  // Contact's first name
	LastName                 string      `json:"last_name,omitempty"`                   // Optional. Contact's last name
	VCard                    string      `json:"vcard,omitempty"`                       // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int         `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendPollType Use this method to send a native poll. On success, the sent Message is returned.
type SendPollType struct {
	// ChatID string or int
	ChatID                   interface{}      `json:"chat_id,omitempty"`                     // Optional. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Question                 string           `json:"question,omitempty"`                    // Poll question, 1-300 characters
	Options                  []string         `json:"options,omitempty"`                     // A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	IsAnonymous              bool             `json:"is_anonymous,omitempty"`                // Optional 	True, if the poll needs to be anonymous, defaults to True
	Type                     string           `json:"type,omitempty"`                        // Optional 	Poll type, “quiz” or “regular”, defaults to “regular”
	AllowMultipleAnswers     bool             `json:"allows_multiple_answers,omitempty"`     // Optional 	True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrecOptID              int              `json:"correct_option_id,omitempty"`           // Optional 	0-based identifier of the correct answer option, required for polls in quiz mode
	Explanation              string           `json:"explanation,omitempty"`                 // Optional 	Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	ExplanationParseMode     string           `json:"explanation_parse_mode,omitempty"`      // Optional 	Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationEntities      []*MessageEntity `json:"explanation_entities,omitempty"`        // Optional 	List of special entities that appear in the poll explanation, which can be specified instead of parse_mode
	OpenPeriod               int              `json:"open_period,omitempty"`                 // Optional 	Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	CloseDate                int              `json:"close_date,omitempty"`                  // Optional 	Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	IsClosed                 bool             `json:"is_closed,omitempty"`                   // Optional 	Pass True, if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendDiceType Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
type SendDiceType struct {
	// ChatID string or int
	ChatID                   interface{} `json:"chat_id,omitempty"`                     // Optional. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Emoji                    string      `json:"emoji,omitempty"`                       // Optional 	Emoji on which the dice throw animation is based. Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, or “🎰”. Dice can have values 1-6 for “🎲” and “🎯”, values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”. Defaults to “🎲”
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID         int         `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply types can be used
}

// SendChatActionType Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive
type SendChatActionType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action string      `json:"action,omitempty"`  // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
}

// GetUserProfilePhotos Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object
type GetUserProfilePhotos struct {
	UserID int `json:"user_id,omitempty"` // Unique identifier of the target user
	Offset int `json:"offset,omitempty"`  // Optional 	Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int `json:"limit,omitempty"`   // Optional 	Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

// GetFileType Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
type GetFileType struct {
	FileID string `json:"file_id,omitempty"` // File identifier to get info about
}

// KickChatMemberType Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type KickChatMemberType struct {
	// ChatID string or int
	ChatID         interface{} `json:"chat_id,omitempty"`         // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID         int         `json:"user_id,omitempty"`         // Unique identifier of the target user
	UntilDate      int         `json:"until_date,omitempty"`      // Optional 	Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	RevokeMessages bool        `json:"revoke_messages,omitempty"` // Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
}

// UnbanChatMemberType Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
type UnbanChatMemberType struct {
	// ChatID string or int
	ChatID   interface{} `json:"chat_id,omitempty"`        // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID   int         `json:"user_id,omitempty"`        // Unique identifier of the target user
	IfBanned bool        `json:"only_if_banned,omitempty"` // Optional 	Do nothing if the user is not banned
}

// RestrictChatMemberType Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
type RestrictChatMemberType struct {
	// ChatID string or int
	ChatID      interface{}      `json:"chat_id,omitempty"`     // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID      int              `json:"user_id,omitempty"`     // Unique identifier of the target user
	Permissions *ChatPermissions `json:"permissions,omitempty"` // A JSON-serialized object for new user permissions
	UntilDate   int              `json:"until_date,omitempty"`  //  	Optional 	Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

// PromoteChatMemberType Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to demote a user. Returns True on success.
type PromoteChatMemberType struct {
	// ChatID string or int
	ChatID             interface{} `json:"chat_id,omitempty"`              // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID             int         `json:"user_id,omitempty"`              // Unique identifier of the target user
	IsAnonymous        bool        `json:"is_anonymous,omitempty"`         // Optional. Pass True, if the administrator's presence in the chat is hidden
	CanChangeInfo      bool        `json:"can_change_info,omitempty"`      // Optional. Pass True, if the administrator can change chat title, photo and other settings
	CanPostMsg         bool        `json:"can_post_messages,omitempty"`    // Optional 	Pass True, if the administrator can create channel posts, channels only
	CanEditMsg         bool        `json:"can_edit_messages,omitempty"`    // Optional 	Pass True, if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMsg       bool        `json:"can_delete_messages,omitempty"`  // Optional 	Pass True, if the administrator can delete messages of other users
	CanInviteUsers     bool        `json:"can_invite_users,omitempty"`     // Optional 	Pass True, if the administrator can invite new users to the chat
	CanRestrictMembers bool        `json:"can_restrict_members,omitempty"` // Optional 	Pass True, if the administrator can restrict, ban or unban chat members
	CanPinMsg          bool        `json:"can_pin_messages,omitempty"`     // Optional 	Pass True, if the administrator can pin messages, supergroups only
	CanPromoteMembers  bool        `json:"can_promote_members,omitempty"`  // Optional 	Pass True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him)
}

// SetChatAdministratorCustomTitleType Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
type SetChatAdministratorCustomTitleType struct {
	// ChatID string or int
	ChatID      interface{} `json:"chat_id,omitempty"`      // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID      int         `json:"user_id,omitempty"`      // Unique identifier of the target user
	CustomTitle string      `json:"custom_title,omitempty"` // New custom title for the administrator; 0-16 characters, emoji are not allowed
}

// SetChatPermissionsType Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members admin rights. Returns True on success.
type SetChatPermissionsType struct {
	// ChatID string or int
	ChatID     interface{}     `json:"chat_id,omitempty"`     // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	Permission ChatPermissions `json:"permissions,omitempty"` // New default chat permissions
}

// ExportChatInviteLinkType Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite link as String on success.
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink — after this the link will become available to the bot via the getChat method. If your bot needs to generate a new invite link replacing its previous one, use exportChatInviteLink again.
type ExportChatInviteLinkType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
}

// SetChatPhotoType Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatPhotoType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	Photo  interface{} `json:"photo,omitempty"`   // New chat photo, uploaded using multipart/form-data
}

// DeleteChatPhotoType Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type DeleteChatPhotoType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
}

// SetChatTitleType Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatTitleType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	Title  string      `json:"title,omitempty"`   // New chat title, 1-255 characters
}

// SetChatDescriptionType Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatDescriptionType struct {
	// ChatID string or int
	ChatID      interface{} `json:"chat_id,omitempty"`     // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	Description string      `json:"description,omitempty"` // Optional. New chat description, 0-255 characters
}

// PinChatMessageType Use this method to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel. Returns True on success.
type PinChatMessageType struct {
	// ChatID string or int
	ChatID              interface{} `json:"chat_id,omitempty"`              // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	MessageID           int         `json:"message_id,omitempty"`           // Identifier of a message to pin
	DisableNotification bool        `json:"disable_notification,omitempty"` // Optional. Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
}

// UnpinChatMessageType Use this method to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel. Returns True on success.
type UnpinChatMessageType struct {
	// ChatID string or int
	ChatID    interface{} `json:"chat_id,omitempty"`    // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	MessageID int         `json:"message_id,omitempty"` // Optional. Identifier of a message to unpin. If not specified, the most recent pinned message (by sending date) will be unpinned.
}

// UnpinAllChatMessagesType Use this method to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel. Returns True on success.
type UnpinAllChatMessagesType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
}

// LeaveChatType Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChatType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatType Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
type GetChatType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatAdministratorsType Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
type GetChatAdministratorsType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatMembersCountType Use this method to get the number of members in a chat. Returns Int on success.
type GetChatMembersCountType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatMemberType Use this method to get information about a member of a chat. Returns a ChatMember object on success.
type GetChatMemberType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID int         `json:"user_id,omitempty"` // Unique identifier of the target user
}

// SetChatStickerSetType Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type SetChatStickerSetType struct {
	// ChatID string or int
	ChatID         interface{} `json:"chat_id,omitempty"`          // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	StickerSetName string      `json:"sticker_set_name,omitempty"` // Name of the sticker set to be set as the group sticker set
}

// DeleteChatStickerSetType Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type DeleteChatStickerSetType struct {
	// ChatID string or int
	ChatID interface{} `json:"chat_id,omitempty"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
}

// AnswerCallbackQueryType Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @Botfather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
type AnswerCallbackQueryType struct {
	CallbackQuery string `json:"callback_query_id,omitempty"` // Unique identifier for the query to be answered
	Text          string `json:"text,omitempty"`              // Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert     bool   `json:"show_alert,omitempty"`        // Optional. If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	URL           string `json:"url,omitempty"`               // Optional. URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game — note that this will only work if the query comes from a callback_game button. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime     int    `json:"cache_time,omitempty"`        // Optional. The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

// SetMyCommandsType Use this method to change the list of the bot's commands. Returns True on success.
type SetMyCommandsType struct {
	Commands []*BotCommand `json:"commands,omitempty"` // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
}

// Updating Messages types

// EditMessageTextType Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageTextType struct {
	// ChatID string or int
	ChatID                interface{}           `json:"chat_id,omitempty"`                  // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID             int                   `json:"message_id,omitempty"`               // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID       string                `json:"inline_message_id,omitempty"`        // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                  string                `json:"text,omitempty"`                     // New text of the message, 1-4096 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities              []*MessageEntity      `json:"entities,omitempty"`                 // Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePrewiew bool                  `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in this message
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. A JSON-serialized object for an inline keyboard.
}

// EditMessageCaptionType Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageCaptionType struct {
	// ChatID string or int
	ChatID          interface{}           `json:"chat_id,omitempty"`           // Optional.Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID       int                   `json:"message_id,omitempty"`        // Optional.Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption         string                `json:"caption,omitempty"`           // Optional. New caption of the message, 0-1024 characters after entities parsing
	ParseMode       string                `json:"parse_mode,omitempty"`        // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	CaptionEntities []*MessageEntity      `json:",omitempty"`                  // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for an inline keyboard.
}

// -----------------------------------------------
// Stickers types Structs

// StickerType This object represents a sticker.
type StickerType struct {
	FileID       string            `json:"file_id,omitempty"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqID   string            `json:"file_unique_id,omitempty"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file
	Width        int               `json:"width,omitempty"`          // Sticker width
	Height       int               `json:"height,omitempty"`         // Sticker height
	IsAnimated   bool              `json:"is_animated,omitempty"`    // True, if the sticker is animated
	Thumb        *PhotoSize        `json:"thumb,omitempty"`          // Optional. Sticker thumbnail in the .WEBP or .JPG format
	Emoji        string            `json:"emoji,omitempty"`          // Optional. Emoji associated with the sticker
	SetName      string            `json:"set_name,omitempty"`       // Optional. Name of the sticker set to which the sticker belongs
	MaskPosition *MaskPositionType `json:"mask_position,omitempty"`  // Optional. For mask stickers, the position where the mask should be placed
	FileSize     int               `json:"file_size,omitempty"`      // Optional. File size
}

// MaskPositionType This object describes the position on faces where a mask should be placed by default.
type MaskPositionType struct {
	Point  string  `json:"point,omitempty"`   // The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	XShift float64 `json:"x_shift,omitempty"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float64 `json:"y_shift,omitempty"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float64 `json:"scale,omitempty"`   // Mask scaling coefficient. For example, 2.0 means double size.
}

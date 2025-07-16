package gotwtr

type DMEventField string

const (
	DirectMessageFieldID               DMEventField = "id"
	DirectMessageFieldText             DMEventField = "text"
	DirectMessageFieldEventType        DMEventField = "event_type"
	DirectMessageFieldCreatedAt        DMEventField = "created_at"
	DirectMessageFieldDMConversationID DMEventField = "dm_conversation_id"
	DirectMessageFieldSenderID         DMEventField = "sender_id"
	DirectMessageFieldParticipantIDs   DMEventField = "participant_ids"
	DirectMessageFieldReferencedTweets DMEventField = "referenced_tweets"
	DirectMessageFieldAttachments      DMEventField = "attachments"
	DirectMessageFieldMediaKeys        DMEventField = "media_keys"
	DirectMessageFieldPublicMetrics    DMEventField = "public_metrics"
)

type EventTypes string

const (
	EventTypesFieldMessageCreate     EventTypes = "MessageCreate"
	EventTypesFieldParticipantsJoin  EventTypes = "ParticipantsJoin"
	EventTypesFieldParticipantsLeave EventTypes = "ParticipantsLeave"
)

type DirectMessage struct {
	Attachments      []DirectMessageAttachment `json:"attachments,omitempty"`
	CreatedAt        string                    `json:"created_at"`
	DMConversationID string                    `json:"dm_conversation_id"`
	EventType        string                    `json:"event_type"`
	ID               string                    `json:"id"`
	SenderID         string                    `json:"sender_id"`
	Text             string                    `json:"text,omitempty"`
	MediaKeys        []string                  `json:"media_keys,omitempty"`
	PublicMetrics    *DirectMessageMetrics     `json:"public_metrics,omitempty"`
}

type DirectMessageMeta struct {
	ResultCount   int    `json:"result_count"`
	PreviousToken string `json:"previous_token,omitempty"`
	NextToken     string `json:"next_token,omitempty"`
}

type DirectMessageAttachment struct {
	MediaID string `json:"media_id"`
}

type DirectMessageMetrics struct {
	ReadCount int `json:"read_count"`
}

type CreateOneToOneDMBody struct {
	Text        string                    `json:"text,omitempty"`
	Attachments []DirectMessageAttachment `json:"attachments,omitempty"`
}

type CreateOneToOneDMResponse struct {
	DMConversationID string              `json:"dm_conversation_id,omitempty"`
	DMEventFieldID   string              `json:"dm_event_id,omitempty"`
	Errors           []*APIResponseError `json:"errors,omitempty"`
	Title            string              `json:"title,omitempty"`
	Detail           string              `json:"detail,omitempty"`
	Type             string              `json:"type,omitempty"`
}

type CreateNewGroupDMBody struct {
	Text        string                    `json:"text,omitempty"`
	Attachments []DirectMessageAttachment `json:"attachments,omitempty"`
}

type CreateNewGroupDMResponse struct {
	DMConversationID string              `json:"dm_conversation_id,omitempty"`
	DMEventFieldID   string              `json:"dm_event_id,omitempty"`
	Errors           []*APIResponseError `json:"errors,omitempty"`
	Title            string              `json:"title,omitempty"`
	Detail           string              `json:"detail,omitempty"`
	Type             string              `json:"type,omitempty"`
}

type PostDMBody struct {
	ConversationType string         `json:"conversation_type"`
	Message          *DirectMessage `json:"message"`
	ParticipantIDs   []string       `json:"participant_id"`
}

type PostDMResponse struct {
	DMConversationID string              `json:"dm_conversation_id,omitempty"`
	DMEventFieldID   string              `json:"dm_event_id,omitempty"`
	Errors           []*APIResponseError `json:"errors,omitempty"`
	Title            string              `json:"title,omitempty"`
	Detail           string              `json:"detail,omitempty"`
	Type             string              `json:"type,omitempty"`
}

type LookUpAllOneToOneDMResponse struct {
	Message []*DirectMessage    `json:"data"`
	Errors  []*APIResponseError `json:"errors,omitempty"`
	Meta    *DirectMessageMeta  `json:"meta,omitempty"`
	Title   string              `json:"title,omitempty"`
	Detail  string              `json:"detail,omitempty"`
	Type    string              `json:"type,omitempty"`
}

type LookUpDMResponse struct {
	Message []*DirectMessage    `json:"data"`
	Errors  []*APIResponseError `json:"errors,omitempty"`
	Meta    *DirectMessageMeta  `json:"meta,omitempty"`
	Title   string              `json:"title,omitempty"`
	Detail  string              `json:"detail,omitempty"`
	Type    string              `json:"type,omitempty"`
}

type LookUpAllDMResponse struct {
	Message []*DirectMessage    `json:"data"`
	Errors  []*APIResponseError `json:"errors,omitempty"`
	Meta    *DirectMessageMeta  `json:"meta,omitempty"`
	Title   string              `json:"title,omitempty"`
	Detail  string              `json:"detail,omitempty"`
	Type    string              `json:"type,omitempty"`
}

// DMBlocking represents the DM blocking data
type DMBlocking struct {
	DMBlocking bool `json:"dm_blocking"`
}

// PostDMBlockingResponse represents the response for blocking DMs from a user
type PostDMBlockingResponse struct {
	DMBlocking *DMBlocking         `json:"data,omitempty"`
	Errors     []*APIResponseError `json:"errors,omitempty"`
	Title      string              `json:"title,omitempty"`
	Detail     string              `json:"detail,omitempty"`
	Type       string              `json:"type,omitempty"`
}

// UndoDMBlockingResponse represents the response for unblocking DMs from a user
type UndoDMBlockingResponse struct {
	DMBlocking *DMBlocking         `json:"data,omitempty"`
	Errors     []*APIResponseError `json:"errors,omitempty"`
	Title      string              `json:"title,omitempty"`
	Detail     string              `json:"detail,omitempty"`
	Type       string              `json:"type,omitempty"`
}

// DMBlockingBody represents the request body for DM blocking operations
type DMBlockingBody struct {
	TargetUserID string `json:"target_user_id"`
}

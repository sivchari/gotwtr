package gotwtr

type Message struct {
	Attachments []MessageAttachment `json:"attachments,omitempty"`
	Text        string              `json:"text,omitempty"`
}

type MessageAttachment struct {
	MediaID string `json:"media_id"`
}

type CreateOneToOneDMBody struct {
	Text        string              `json:"text,omitempty"`
	Attachments []MessageAttachment `json:"attachments,omitempty"`
}

type CreateOneToOneDMResponse struct {
	DMConversationID string              `json:"dm_conversation_id,omitempty"`
	DMEventID        string              `json:"dm_event_id,omitempty"`
	Errors           []*APIResponseError `json:"errors,omitempty"`
	Title            string              `json:"title,omitempty"`
	Detail           string              `json:"detail,omitempty"`
	Type             string              `json:"type,omitempty"`
}

type CreateNewGroupDMBody struct {
	Text        string              `json:"text,omitempty"`
	Attachments []MessageAttachment `json:"attachments,omitempty"`
}

type CreateNewGroupDMResponse struct {
	DMConversationID string              `json:"dm_conversation_id,omitempty"`
	DMEventID        string              `json:"dm_event_id,omitempty"`
	Errors           []*APIResponseError `json:"errors,omitempty"`
	Title            string              `json:"title,omitempty"`
	Detail           string              `json:"detail,omitempty"`
	Type             string              `json:"type,omitempty"`
}

type PostDMBody struct {
	ConversationType string   `json:"conversation_type"`
	Message          *Message `json:"message"`
	ParticipantIDs   []string `json:"participant_id"`
}

type PostDMResponse struct {
	DMConversationID string              `json:"dm_conversation_id,omitempty"`
	DMEventID        string              `json:"dm_event_id,omitempty"`
	Errors           []*APIResponseError `json:"errors,omitempty"`
	Title            string              `json:"title,omitempty"`
	Detail           string              `json:"detail,omitempty"`
	Type             string              `json:"type,omitempty"`
}

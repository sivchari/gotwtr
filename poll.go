package twitter

type PollField string

/*
	Poll field will only return
	if you've also included the expansions=attachments.poll_ids query parameter in your request.
*/
const (
	PollFieldDurationMinutes PollField = "duration_minutes"
	PollFieldEndDateTime     PollField = "end_datetime"
	PollFieldID              PollField = "id"
	PollFieldOptions         PollField = "options"
	PollFieldVotingStatus    PollField = "voting_status"
)

type Poll struct {
	ID              string        `json:"id"`
	Options         []*PollOption `json:"options"`
	DurationMinutes int           `json:"duration_minutes,omitempty"`
	EndDatetime     string        `json:"end_datetime,omitempty"`
	VotingStatus    string        `json:"voting_status,omitempty"`
}

type PollOption struct {
	Position int    `json:"position"`
	Label    string `json:"label"`
	Votes    int    `json:"votes"`
}

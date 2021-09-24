package field

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

package gotwtr

type TopicField string

const (
	TopicFieldID          TopicField = "id"
	TopicFieldName        TopicField = "name"
	TopicFieldDescription TopicField = "description"
)

func topicFieldsToString(tfs []TopicField) []string {
	slice := make([]string, len(tfs))
	for i, tf := range tfs {
		slice[i] = string(tf)
	}
	return slice
}

type Topic struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

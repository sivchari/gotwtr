package gotwtr

// Exclude used in the timeline parameters
type Exclude string

const (
	ExcludeRetweets Exclude = "retweets"
	ExcludeReplies  Exclude = "replies"
)

func excludeToString(es []Exclude) []string {
	slice := make([]string, len(es))
	for i, e := range es {
		slice[i] = string(e)
	}
	return slice
}

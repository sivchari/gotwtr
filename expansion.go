package gotwtr

type Expansion string

const (
	ExpansionAttachmentsPollIDs         Expansion = "attachments.poll_ids"
	ExpansionAttachmentsMediaKeys       Expansion = "attachments.media_keys"
	ExpansionAuthorID                   Expansion = "author_id"
	ExpansionEntitiesMentionsUserName   Expansion = "entities.mentions.username"
	ExpansionGeoPlaceID                 Expansion = "geo.place_id"
	ExpansionInReplyToUserID            Expansion = "in_reply_to_user_id"
	ExpansionReferencedTweetsID         Expansion = "referenced_tweets.id"
	ExpansionReferencedTweetsIDAuthorID Expansion = "referenced_tweets.id.author_id"
	ExpansionPinnedTweetID              Expansion = "pinned_tweet_id"
	ExpansionInvitedUserIDs             Expansion = "invited_user_ids"
	ExpansionSpeakerIDs                 Expansion = "speaker_ids"
	ExpansionCreatorID                  Expansion = "creator_id"
	ExpansionHostIDs                    Expansion = "host_ids"
	ExpansionContextAnnotations         Expansion = "context_annotations"
	ExpansionTopicIDs                   Expansion = "topic_ids"
	ExpansionOwnerID                    Expansion = "list.owner_id"
	// ExpansionTopics                     Expansion = "topics" TODO: must check
)

func expansionsToString(es []Expansion) []string {
	slice := make([]string, len(es))
	for i, e := range es {
		slice[i] = string(e)
	}
	return slice
}

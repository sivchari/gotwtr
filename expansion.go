package gotwtr

type Expansion string

const (
	// Tweet payloads
	ExpansionAuthorID                   Expansion = "author_id"
	ExpansionReferencedTweetsID         Expansion = "referenced_tweets.id"
	ExpansionEditHistoryTweetIDs        Expansion = "edit_history_tweet_ids"
	ExpansionInReplyToUserID            Expansion = "in_reply_to_user_id"
	ExpansionAttachmentsMediaKeys       Expansion = "attachments.media_keys"
	ExpansionAttachmentsPollIDs         Expansion = "attachments.poll_ids"
	ExpansionGeoPlaceID                 Expansion = "geo.place_id"
	ExpansionEntitiesMentionsUserName   Expansion = "entities.mentions.username"
	ExpansionReferencedTweetsIDAuthorID Expansion = "referenced_tweets.id.author_id"
	ExpansionContextAnnotations         Expansion = "context_annotations"
	// USer payloads
	ExpansionPinnedTweetID Expansion = "pinned_tweet_id"
	// Direct Message event payloads with attachments.media_keys + referenced_tweets.id
	ExpansionSenderID       Expansion = "sender_id"
	ExpansionParticipantIDs Expansion = "participant_ids"
	// Space payloads
	ExpansionInvitedUserIDs Expansion = "invited_user_ids"
	ExpansionSpeakerIDs     Expansion = "speaker_ids"
	ExpansionCreatorID      Expansion = "creator_id"
	ExpansionHostIDs        Expansion = "host_ids"
	ExpansionTopicIDs       Expansion = "topic_ids"
	// List payloads
	ExpansionOwnerID Expansion = "owner_id"
)

func expansionsToString(es []Expansion) []string {
	slice := make([]string, len(es))
	for i, e := range es {
		slice[i] = string(e)
	}
	return slice
}

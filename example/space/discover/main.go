package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	dsr, err := client.DiscoverSpacesByUserIDs(context.Background(), []string{"id"}, &gotwtr.DiscoverSpacesOption{
		Expansions: []gotwtr.Expansion{
			gotwtr.ExpansionHostIDs,
			gotwtr.ExpansionCreatorID,
			gotwtr.ExpansionInvitedUserIDs,
			gotwtr.ExpansionSpeakerIDs,
			gotwtr.ExpansionTopicIDs,
		},
		TopicFields: []gotwtr.TopicField{
			gotwtr.TopicFieldName,
			gotwtr.TopicFieldID,
			gotwtr.TopicFieldDescription,
		},
		UserFields: []gotwtr.UserField{
			gotwtr.UserFieldCreatedAt,
			gotwtr.UserFieldDescription,
			gotwtr.UserFieldEntities,
			gotwtr.UserFieldID,
			gotwtr.UserFieldLocation,
			gotwtr.UserFieldName,
			gotwtr.UserFieldPinnedTweetID,
			gotwtr.UserFieldProfileImageURL,
			gotwtr.UserFieldProtected,
			gotwtr.UserFieldPublicMetrics,
			gotwtr.UserFieldURL,
			gotwtr.UserFieldUserName,
			gotwtr.UserFieldVerified,
			gotwtr.UserFieldWithHeld,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("---")
	for _, t := range dsr.Spaces {
		fmt.Println(t)
	}
	fmt.Println("---")
	for _, t := range dsr.Includes.Topics {
		fmt.Println(t)
	}
	fmt.Println("---")
	for _, t := range dsr.Includes.Users {
		fmt.Println(t)
	}
}

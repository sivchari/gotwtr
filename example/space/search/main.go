package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	ssr, err := client.SearchSpaces(context.Background(), "hello", &gotwtr.SearchSpacesOption{
		SpaceFields: []gotwtr.SpaceField{
			gotwtr.SpaceFieldHostIDs,
			gotwtr.SpaceFieldCreatedAt,
			gotwtr.SpaceFieldCreatorID,
			gotwtr.SpaceFieldID,
			gotwtr.SpaceFieldLanguage,
			gotwtr.SpaceFieldInvittedUserIDs,
			gotwtr.SpaceFieldParticipantCount,
			gotwtr.SpaceFieldSpeakerIDs,
			gotwtr.SpaceFieldStartedAt,
			gotwtr.SpaceFieldState,
			gotwtr.SpaceFieldTitle,
			gotwtr.SpaceFieldUpdatedAt,
			gotwtr.SpaceFieldScheduledStart,
			gotwtr.SpaceFieldIsTicketed,
		},
	})
	if err != nil {
		panic(err)
	}
	for _, s := range ssr.Spaces {
		fmt.Println(s)
	}
}

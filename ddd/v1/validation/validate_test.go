package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_checkUserStory(t *testing.T) {
	tests := []struct {
		story   string
		want    userStory
		wantErr bool
	}{
		{
			"As a moderator I want to create a new sprint by entering a name and an optional comment so that I can start planning the stories.",
			userStory{
				Role:   "moderator",
				Goal:   "create a new sprint by entering a name and an optional comment",
				Reason: "I can start planning the stories",
			},
			false,
		},
		{
			"I want to create a new sprint by entering a name and an optional comment so that I can start planning the stories.",
			userStory{},
			true,
		},

		{
			"",
			userStory{},
			true,
		},

		{
			"asdf",
			userStory{},
			true,
		},
		{
			"As a I want to so that.",
			userStory{},
			true,
		},

		{
			"so that I can start planning the stories I want to create a new sprint by entering a name and an optional comment so.",
			userStory{},
			true,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := CheckUserStory(tt.story)
			if err != nil {
				fmt.Println(err.Error())
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("checkUserStory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkUserStory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Copyright 2016 Netflix, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package tracker provides an entry point for instantiating Trackers
package tracker

import (
	"fmt"
	"os"

	"github.com/sbasgall/chaosmonkey"
	"github.com/sbasgall/chaosmonkey/config"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

var slacktracker chaosmonkey.Tracker
// SlackObject yes
type SlackObject struct {}

// Track is a great function
func (s *SlackObject) Track(chaosmonkey.Termination) error {
	fmt.Println("inside")
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage("custodian", "Some text", params)
	if err != nil {
			fmt.Printf("%s\n", err)
			return err
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	return nil
}

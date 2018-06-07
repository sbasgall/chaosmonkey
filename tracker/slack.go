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

	"github.com/sbasgall/chaosmonkey"
	"github.com/sbasgall/chaosmonkey/config"
	"github.com/sbasgall/chaosmonkey/deps"
//	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

// SlackObject yes
type SlackObject struct {}

// Track is a great function
func (s *SlackObject) Track(chaosmonkey.Termination) error {
	fmt.Println("inside")
	return errors.Errorf("unkown error override!!!")
}



//func (tr chaosmonkey.Tracker) Track(t Termination) {
	//	  api := slack.New("YOUR_TOKEN_HERE")
	//		params := slack.PostMessageParameters{}
	//		attachment := slack.Attachment{
	//			Pretext: "some pretext",
	//			Text:    "some text",
//	fmt.Println("TRACKER\n\n")
//	fmt.Println(tr)
//}
//		params.Attachments = []slack.Attachment{attachment}
//		channelID, timestamp, err := api.PostMessage("CHANNEL_ID", "Some text", params)
//		if err != nil {
//			fmt.Printf("%s\n", err)
//		}
//		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

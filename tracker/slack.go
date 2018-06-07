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
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/pkg/errors"
)

var slacktracker chaosmonkey.Tracker
// SlackObject yes
type SlackObject struct {}

// Track is a great function
func (s *SlackObject) Track(chaosmonkey.Termination) error {
	fmt.Println("inside")

	webhookURL := os.Getenv("SLACK_WEBHOOK")

	    attachment1 := slack.Attachment {}
	    attachment1.AddField(slack.Field { Title: "Chaos Termination", Value: "Yes" }).AddField(slack.Field { Title: "Status", Value: "Completed" })
	    payload := slack.Payload {
	      Text: "Hello from <https://github.com/ashwanthkumar/slack-go-webhook|slack-go-webhook>, a Go-Lang library to send slack webhook messages.\n<https://golangschool.com/wp-content/uploads/golang-teach.jpg|golang-img>",
	      Username: "custodian",
	      Channel: "#custodian",
	      IconEmoji: ":monkey_face:",
	      Attachments: []slack.Attachment{attachment1},
	    }
	    err := slack.Send(webhookURL, "", payload)
	    if len(err) > 0 {
	      return errors.Errorf("error posting")
	    }
	return nil
}

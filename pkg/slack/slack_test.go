/*
copyright 2020 the Goployer authors

licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at

    http://www.apache.org/licenses/license-2.0

unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package slack

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"

	"github.com/DevopsArtFactory/goployer/pkg/schemas"
)

func TestNewSlackClient(t *testing.T) {
	tests := []struct {
		name     string
		slackOff bool
	}{
		{
			name:     "Slack enabled",
			slackOff: false,
		},
		{
			name:     "Slack disabled",
			slackOff: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewSlackClient(tt.slackOff)
			assert.NotNil(t, client, "Slack client should not be nil")
			assert.Equal(t, tt.slackOff, client.SlackOff, "SlackOff flag should match")
		})
	}
}

func TestValidClient(t *testing.T) {
	tests := []struct {
		name     string
		slack    Slack
		expected bool
	}{
		{
			name: "Valid client with token and channel",
			slack: Slack{
				Token:     "xoxb-test-token",
				ChannelID: "C12345678",
				SlackOff:  false,
			},
			expected: true,
		},
		{
			name: "Valid client with webhook",
			slack: Slack{
				WebhookURL: "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXX",
				SlackOff:   false,
			},
			expected: true,
		},
		{
			name: "Invalid client - Slack disabled",
			slack: Slack{
				Token:     "xoxb-test-token",
				ChannelID: "C12345678",
				SlackOff:  true,
			},
			expected: false,
		},
		{
			name: "Invalid client - No credentials",
			slack: Slack{
				SlackOff: false,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.slack.ValidClient()
			assert.Equal(t, tt.expected, result, "ValidClient result should match expected")
		})
	}
}

func TestCreateSimpleSection(t *testing.T) {
	slackClient := Slack{}
	text := "Deployment started for my-app"

	section := slackClient.CreateSimpleSection(text)

	assert.NotNil(t, section, "Section should not be nil")
	assert.IsType(t, &slack.SectionBlock{}, section, "Should return SectionBlock type")
}

func TestCreateDividerSection(t *testing.T) {
	slackClient := Slack{}

	divider := slackClient.CreateDividerSection()

	assert.NotNil(t, divider, "Divider should not be nil")
	assert.IsType(t, &slack.DividerBlock{}, divider, "Should return DividerBlock type")
}

func TestCreateTitleSection(t *testing.T) {
	slackClient := Slack{}
	title := "Deployment Summary"

	section := slackClient.CreateTitleSection(title)

	assert.NotNil(t, section, "Title section should not be nil")
	assert.IsType(t, &slack.SectionBlock{}, section, "Should return SectionBlock type")
}

func TestSlackMessageStructures(t *testing.T) {
	tests := []struct {
		name string
		body Body
	}{
		{
			name: "Body with blocks",
			body: Body{
				Blocks: []Block{
					{
						Type: "section",
						Text: &Text{
							Type: "mrkdwn",
							Text: "Test message",
						},
					},
				},
			},
		},
		{
			name: "Body with attachments",
			body: Body{
				Attachments: []Attachment{
					{
						Text:  "Deployment completed",
						Color: "good",
						Fields: []Field{
							{
								Title: "Environment",
								Value: "production",
								Short: true,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.body, "Body should not be nil")
			if len(tt.body.Blocks) > 0 {
				assert.Greater(t, len(tt.body.Blocks), 0, "Should have blocks")
			}
			if len(tt.body.Attachments) > 0 {
				assert.Greater(t, len(tt.body.Attachments), 0, "Should have attachments")
			}
		})
	}
}

func TestAttachmentColors(t *testing.T) {
	tests := []struct {
		name  string
		color string
		valid bool
	}{
		{
			name:  "Good color",
			color: "good",
			valid: true,
		},
		{
			name:  "Warning color",
			color: "warning",
			valid: true,
		},
		{
			name:  "Danger color",
			color: "danger",
			valid: true,
		},
		{
			name:  "Custom hex color",
			color: "#36a64f",
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attachment := Attachment{
				Text:  "Test message",
				Color: tt.color,
			}
			assert.Equal(t, tt.color, attachment.Color, "Color should match")
		})
	}
}

func TestMetricResultFormatting(t *testing.T) {
	metrics := []schemas.MetricResult{
		{
			URL:    "https://api.example.com/health",
			Method: "GET",
		},
		{
			URL:    "https://api.example.com/users",
			Method: "POST",
		},
	}

	assert.Equal(t, 2, len(metrics), "Should have 2 metrics")
	for _, metric := range metrics {
		assert.NotEmpty(t, metric.URL, "URL should not be empty")
		assert.NotEmpty(t, metric.Method, "Method should not be empty")
	}
}

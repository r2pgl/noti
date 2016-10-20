// +build !darwin
// +build !windows

package config

import (
	"github.com/variadico/noti/ntypes/espeak"
	"github.com/variadico/noti/ntypes/freedesktop"
	"github.com/variadico/noti/ntypes/slack"
)

type Options struct {
	DefaultNotifications []string
	DefaultTriggers      []string
	Banner               *freedesktop.Notification
	Speech               *espeak.Notification
	Slack                *slack.Notification
}

func NewOptions() Options {
	return Options{
		DefaultNotifications: make([]string, 0),
		DefaultTriggers:      make([]string, 0),
		Banner:               new(freedesktop.Notification),
		Speech:               new(espeak.Notification),
		Slack:                new(slack.Notification),
	}
}
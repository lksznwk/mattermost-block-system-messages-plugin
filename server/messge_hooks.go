package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	configuration := p.getConfiguration()
	if err := configuration.isValid(); err != nil {
		p.API.LogWarn("Plugin is not configured. " + err.Error())
		return post, ""
	}

	if post.ChannelId == configuration.ChannelId && strings.Contains(post.Type, configuration.EventType) {
		return nil, fmt.Sprintf("Blocking %s notification for channel %s", configuration.EventType, configuration.ChannelId)
	}
	return post, ""

}

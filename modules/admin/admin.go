package admin

import (
	"strings"
	"log"
	"github.com/mattermost/platform/model"
	"fmt"
)

const DEFAULT_ERROR_MSG = "I couldn't service your request"

func HandleCommand(client *model.Client, command string) string {
	tokens := strings.Split(command, " ")

	switch tokens[0] {
	case "channels":
		return listJoinedChannels(client)
	case "join":
		if len(tokens) < 2 {
			return "You must specify a channel"
		}
		return joinChannel(client, tokens[1])
	case "leave":
		if len(tokens) < 2 {
			return "You must specify a channel"
		}
		return leaveChannel(client, tokens[1])
	case "":
		return "I am here to serve"
	}
	return DEFAULT_ERROR_MSG
}

func listJoinedChannels(client *model.Client) string {
	result, err := client.GetChannels("")
	if err != nil {
		log.Println(err)
		return DEFAULT_ERROR_MSG;
	}
	msg := "Channels I belong to: \n"
	channels := result.Data.(*model.ChannelList)
	for _, channel := range *channels {
		msg += channel.Name
		msg += "\n"
	}
	return msg
}

func joinChannel(client *model.Client, channelName string) string {
	_, err := client.JoinChannelByName(channelName)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("Could not join %s\n%s", channelName, err)
	}
	return fmt.Sprintf("Joined %s", channelName)
}

func leaveChannel(client *model.Client, channelName string) string {
	result, err := client.GetChannelByName(channelName)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("Channel %s doesn't exist", channelName)
	}
	channelToLeave := result.Data.(*model.Channel)
	log.Println("Leaving:", channelToLeave.Name)
	result, err = client.LeaveChannel(channelToLeave.Id)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("Could not leave %s\n%s", channelName, err)
	}
	return fmt.Sprintf("Left %s", channelName)
}
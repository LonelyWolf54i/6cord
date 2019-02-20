package main

import (
	"fmt"

	"github.com/rumblefrog/discordgo"
)

func isRegularMessage(m *discordgo.Message) bool {
	var messageText string

	// https://github.com/Bios-Marcel/cordless
	switch m.Type {
	case discordgo.MessageTypeGuildMemberJoin:
		messageText = "joined the server."
	case discordgo.MessageTypeCall:
		messageText = "is calling you."
	case discordgo.MessageTypeChannelIconChange:
		messageText = "changed the channel icon."
	case discordgo.MessageTypeChannelNameChange:
		messageText = "changed the channel name to " + m.Content + "."
	case discordgo.MessageTypeChannelPinnedMessage:
		messageText = "pinned a message."
	case discordgo.MessageTypeRecipientAdd:
		messageText = "added " + m.Mentions[0].Username + " to the group."
	case discordgo.MessageTypeRecipientRemove:
		messageText = "removed " + m.Mentions[0].Username + " from the group."
	}

	if messageText != "" {
		setLastAuthor(0)

		msg := fmt.Sprintf(
			"\n\n[::d]%s %s[::-]\n",
			m.Author.Username, messageText,
		)

		// Writing it directly for performance
		messagesView.Write([]byte(msg))
		messageStore = append(messageStore, msg)

		return false
	}

	return true
}
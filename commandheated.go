package main

func cmdHeated(text []string) {
	if Channel == nil {
		Message("You're not in a channel!")
		return
	}

	heatedChannelsAdd(Channel.ID)
	Message("Added this channel. We'll warn you when there's a message.")
}

package mixerstructs

type DiscordBotCollection []DiscordBot
type DiscordBotStringMap map[string]DiscordBot
type DiscordBotIntMap map[int]DiscordBot
type DiscordBotUIntMap map[uint]DiscordBot

// Lists the configuration for this Discord bot.
type DiscordBot struct {
	ChannelId           uint     `json:"channelId"`
	GuildId             string   `json:"guildId"`
	Id                  uint     `json:"id"`
	InviteChannel       string   `json:"inviteChannel"`
	InviteSetting       string   `json:"inviteSetting"`
	LiveAnnounceChannel string   `json:"liveAnnounceChannel"`
	LiveChatChannel     string   `json:"liveChatChannel"`
	LiveUpdateState     bool     `json:"liveUpdateState"`
	Roles               []string `json:"roles"`
	SyncEmoteRoles      uint     `json:"syncEmoteRoles"`
}

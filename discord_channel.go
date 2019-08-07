package structs

type DiscordChannelCollection []DiscordChannel
type DiscordChannelStringMap map[string]DiscordChannel
type DiscordChannelIntMap map[int]DiscordChannel
type DiscordChannelUIntMap map[uint]DiscordChannel

// Represents a single channel within a Discord server.
type DiscordChannel struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
	Type    string `json:"type"`
}

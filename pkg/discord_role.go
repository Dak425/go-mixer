package mixer

type DiscordRoleCollection []DiscordRole
type DiscordRoleStringMap map[string]DiscordRole
type DiscordRoleIntMap map[int]DiscordRole
type DiscordRoleUIntMap map[uint]DiscordRole

// Represents a role within a Discord server.
type DiscordRole struct {
	Assignable bool   `json:"assignable"`
	Color      string `json:"color"`
	Id         string `json:"id"`
	Name       string `json:"name"`
}

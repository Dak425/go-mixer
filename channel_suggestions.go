package mixer

type ChannelSuggestionsCollection []ChannelSuggestions
type ChannelSuggestionsStringMap map[string]ChannelSuggestions
type ChannelSuggestionsIntMap map[int]ChannelSuggestions
type ChannelSuggestionsUIntMap map[uint]ChannelSuggestions

// Suggested channels with their beacon URLs
type ChannelSuggestions struct {
	ActionsUrl    string  `json:"actionsUrl"`
	Channels      Channel `json:"channels"`
	ImpressionUrl string  `json:"impressionUrl"`
}

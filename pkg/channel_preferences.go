package mixer

type ChannelPreferencesCollection []ChannelPreferences
type ChannelPreferencesStringMap map[string]ChannelPreferences
type ChannelPreferencesIntMap map[int]ChannelPreferences
type ChannelPreferencesUIntMap map[uint]ChannelPreferences

// Channel preferences are a list of options and attributes which control behaviour for the channel. Please see each property for more details.
type ChannelPreferences struct {
	ChannelLinksAllowed                bool   `json:"channel:links:allowed"`
	ChannelLinksClickable              bool   `json:"channel:links:clickable"`
	ChannelNotifyDirectPurchaseMessage string `json:"channel:notify:directPurchaseMessage"`
	ChannelNotifyFollow                bool   `json:"channel:notify:follow"`
	ChannelNotifyFollowmessage         string `json:"channel:notify:followmessage"`
	ChannelNotifyHostedBy              string `json:"channel:notify:hostedBy"`
	ChannelNotifyHosting               string `json:"channel:notify:hosting"`
	ChannelNotifySubscribe             bool   `json:"channel:notify:subscribe"`
	ChannelNotifySubscribemessage      string `json:"channel:notify:subscribemessage"`
	ChannelOfflineAutoplayVod          bool   `json:"channel:offline:autoplayVod"`
	ChannelPartnerSubmail              string `json:"channel:partner:submail"`
	ChannelPlayerMuteOwn               bool   `json:"channel:player:muteOwn"`
	ChannelSlowchat                    int    `json:"channel:slowchat"`
	ChannelTweetBody                   string `json:"channel:tweet:body"`
	ChannelTweetEnabled                bool   `json:"channel:tweet:enabled"`
	CostreamAllow                      bool   `json:"costream:allow"`
	HostingAllow                       bool   `json:"hosting:allow"`
	HypezoneAllow                      bool   `json:"hypezone:allow"`
	ShareText                          string `json:"sharetext"`
}

package mixer

type PartyChatSessionCollection []PartyChatSession
type PartyChatSessionStringMap map[string]PartyChatSession
type PartyChatSessionIntMap map[int]PartyChatSession
type PartyChatSessionUIntMap map[uint]PartyChatSession

type PartyChatMemberCollection []PartyChatMember

type PartyChatMember struct {
	Gamertag              string             `json:"gamertag"`
	SimpleConnectionState int                `json:"simpleConnectionState"`
	MixerUser             User               `json:"mixerUser"`
	Ref                   PartyChatMemberRef `json:"ref"`
}

// A Party Chat voice session.
type PartyChatSession struct {
	Members                    PartyChatMemberCollection `json:"members"`
	PartyId                    string                    `json:"partyId"`
	ServerConnectionCandidates []string                  `json:"serverConnectionCandidates"`
}

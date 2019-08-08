package mixerstructs

type PartyChatMemberRefCollection []PartyChatMemberRef
type PartyChatMemberRefStringMap map[string]PartyChatMemberRef
type PartyChatMemberRefIntMap map[int]PartyChatMemberRef
type PartyChatMemberRefUIntMap map[uint]PartyChatMemberRef

// Reference one user session in a party chat.
type PartyChatMemberRef struct {
	Gamertag string `json:"gamertag"`
	Index    int    `json:"index"`
	PartyId  string `json:"partyId"`
}

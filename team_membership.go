package mixerstructs

type TeamMembershipCollection []TeamMembership
type TeamMembershipStringMap map[string]TeamMembership
type TeamMembershipIntMap map[int]TeamMembership
type TeamMembershipUIntMap map[uint]TeamMembership

// Provides details about the membership of a user to a team.
type TeamMembership struct {
	TimeStamped
	Accepted bool `json:"accepted"`
	Id       uint `json:"id"`
	TeamId   uint `json:"teamId"`
	UserId   uint `json:"userId"`
}

package mixerstructs

type TeamMembershipExpandedCollection []TeamMembershipExpanded
type TeamMembershipExpandedStringMap map[string]TeamMembershipExpanded
type TeamMembershipExpandedIntMap map[int]TeamMembershipExpanded
type TeamMembershipExpandedUIntMap map[uint]TeamMembershipExpanded

// A team membership expanded with additional properties including nested relations.
type TeamMembershipExpanded struct {
	TeamMembership
	IsPrimary bool `json:"isPrimary"`
	Owner     User `json:"owner"`
}

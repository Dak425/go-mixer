package structs

type UserWithGroupsCollection []UserWithGroups
type UserWithGroupsStringMap map[string]UserWithGroups
type UserWithGroupsIntMap map[int]UserWithGroups
type UserWithGroupsUIntMap map[uint]UserWithGroups

// A User object with an embedded array of groups they belong to.
type UserWithGroups struct {
	User
	Groups UserGroup `json:"groups"`
}

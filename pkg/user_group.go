package mixer

type UserGroupCollection []UserGroup
type UserGroupStringMap map[string]UserGroup
type UserGroupIntMap map[int]UserGroup
type UserGroupUIntMap map[uint]UserGroup

// A Group which a user can belong to can control features or access controls throughout Mixer.
type UserGroup struct {
	TimeStamped
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

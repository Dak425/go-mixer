package mixer

type UserLogCollection []UserLog
type UserLogStringMap map[string]UserLog
type UserLogIntMap map[int]UserLog
type UserLogUIntMap map[uint]UserLog

// A UserLog entry lists events and changes made to a user.
type UserLog struct {
	CreatedAt  string   `json:"createdAt"`
	Event      string   `json:"event"`
	EventData  struct{} `json:"eventData"`
	Id         uint     `json:"id"`
	Source     string   `json:"source"`
	SourceData struct{} `json:"sourceData"`
	UserId     uint     `json:"userId"`
}

package mixerstructs

type NotificationPreferenceCollection []NotificationPreference
type NotificationPreferenceStringMap map[string]NotificationPreference
type NotificationPreferenceIntMap map[int]NotificationPreference
type NotificationPreferenceUIntMap map[uint]NotificationPreference

// A user's notification preferences.
type NotificationPreference struct {
	AllowsEmailMarketing bool       `json:"allowsEmailMarketing"`
	Health               struct{}   `json:"health"`
	ID                   string     `json:"id"`
	LastRead             string     `json:"lastRead"`
	LiveNotifications    []struct{} `json:"liveNotifications"`
	LiveOnByDefault      []string   `json:"liveOnByDefault"`
	NotifyFollower       []string   `json:"notifyFollower"`
	NotifySubscriber     []string   `json:"notifySubscriber"`
	Transports           []struct{} `json:"transports"`
}

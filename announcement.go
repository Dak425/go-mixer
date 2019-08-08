package mixerstructs

type AnnouncementCollection []Announcement
type AnnouncementStringMap map[string]Announcement
type AnnouncementIntMap map[int]Announcement
type AnnouncementUIntMap map[uint]Announcement

// Announcements are triggered by Mixer staff and inform users of important information and news on Mixer.
type Announcement struct {
	Confetti uint   `json:"confetti"`
	Level    string `json:"level"`
	Message  string `json:"message"`
	Sound    string `json:"sound"`
	Timeout  int    `json:"timeout"`
}

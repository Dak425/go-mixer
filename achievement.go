package mixer

type AchievementCollection []Achievement
type AchievementStringMap map[string]Achievement
type AchievementIntMap map[int]Achievement
type AchievementUIntMap map[uint]Achievement

// Achievements are earnt by users and have various requirements.
type Achievement struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Sparks      uint   `json:"sparks"`
}

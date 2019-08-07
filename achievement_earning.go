package structs

type AchievementEarningCollection []AchievementEarning
type AchievementEarningStringMap map[string]AchievementEarning
type AchievementEarningIntMap map[int]AchievementEarning
type AchievementEarningUIntMap map[uint]AchievementEarning

// An AchievementEarning represents a user's progress towards an achievement.
type AchievementEarning struct {
	TimeStamped
	AchievementStruct Achievement `json:"Achievement"`
	Achievement       string      `json:"achievement"`
	Earned            bool        `json:"earned"`
	ID                uint        `json:"id"`
	Progress          int         `json:"progress"`
	User              uint        `json:"user"`
}

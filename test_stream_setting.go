package mixerstructs

type TestStreamSettingCollection []TestStreamSetting
type TestStreamSettingStringMap map[string]TestStreamSetting
type TestStreamSettingIntMap map[int]TestStreamSetting
type TestStreamSettingUIntMap map[uint]TestStreamSetting

// Lists the test stream settings for a Channel.
type TestStreamSetting struct {
	HoursLastReset           string `json:"hoursLastReset"`
	HoursQuota               int    `json:"hoursQuota"`
	HoursRemaining           int    `json:"hoursRemaining"`
	HoursResetIntervalInDays uint   `json:"hoursResetIntervalInDays"`
	Id                       uint   `json:"id"`
	IsActive                 bool   `json:"isActive"`
}

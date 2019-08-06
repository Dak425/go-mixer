package mixer

type ProgressionStatisticCollection []ProgressionStatistic
type ProgressionStatisticStringMap map[string]ProgressionStatistic
type ProgressionStatisticIntMap map[int]ProgressionStatistic
type ProgressionStatisticUIntMap map[uint]ProgressionStatistic

// A type that contains information about a single value in progression storage.
type ProgressionStatistic struct {
	Amount int    `json:"amount"`
	Stat   string `json:"stat"`
}

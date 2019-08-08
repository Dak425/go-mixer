package mixerstructs

type ProgressionViewerStatisticValueCollection []ProgressionViewerStatisticValue
type ProgressionViewerStatisticValueStringMap map[string]ProgressionViewerStatisticValue
type ProgressionViewerStatisticValueIntMap map[int]ProgressionViewerStatisticValue
type ProgressionViewerStatisticValueUIntMap map[uint]ProgressionViewerStatisticValue

// A type that contains information about a single value in progression storage.
type ProgressionViewerStatisticValue struct {
	SessionId string `json:"sessionId"`
	Value     int    `json:"value"`
}

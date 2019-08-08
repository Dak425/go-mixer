package mixerstructs

type ProgressionViewersStatisticValuesCollection []ProgressionViewersStatisticValues
type ProgressionViewersStatisticValuesStringMap map[string]ProgressionViewersStatisticValues
type ProgressionViewersStatisticValuesIntMap map[int]ProgressionViewersStatisticValues
type ProgressionViewersStatisticValuesUIntMap map[uint]ProgressionViewersStatisticValues

// A type that contains information about a single value in progression storage.
type ProgressionViewersStatisticValues struct {
	Results ProgressionViewerStatisticValue `json:"results"`
	Stat    string                          `json:"stat"`
}

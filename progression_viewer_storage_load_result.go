package mixerstructs

type ProgressionViewerStorageLoadResultCollection []ProgressionViewerStorageLoadResult
type ProgressionViewerStorageLoadResultStringMap map[string]ProgressionViewerStorageLoadResult
type ProgressionViewerStorageLoadResultIntMap map[int]ProgressionViewerStorageLoadResult
type ProgressionViewerStorageLoadResultUIntMap map[uint]ProgressionViewerStorageLoadResult

// A type that contains information about a single value in progression storage.
type ProgressionViewerStorageLoadResult struct {
	ChannelId  uint                            `json:"channelId"`
	Project    ProgressionActiveStorageProject `json:"project"`
	Result     string                          `json:"result"`
	StatusCode uint                            `json:"statusCode"`
}

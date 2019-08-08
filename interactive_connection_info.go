package mixerstructs

type InteractiveConnectionInfoCollection []InteractiveConnectionInfo
type InteractiveConnectionInfoStringMap map[string]InteractiveConnectionInfo
type InteractiveConnectionInfoIntMap map[int]InteractiveConnectionInfo
type InteractiveConnectionInfoUIntMap map[uint]InteractiveConnectionInfo

// Provides connection details that can be used to connect to an interactive session on a channel.
type InteractiveConnectionInfo struct {
	Address   string                     `json:"address"`
	Influence int                        `json:"influence"`
	Key       string                     `json:"key"`
	Version   InteractiveVersionExpanded `json:"version"`
}

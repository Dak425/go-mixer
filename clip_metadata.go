package mixerstructs

type ClipMetadataCollection []ClipMetadata
type ClipMetadataStringMap map[string]ClipMetadata
type ClipMetadataIntMap map[int]ClipMetadata
type ClipMetadataUIntMap map[uint]ClipMetadata

// Metadata about a clip which can be updated
type ClipMetadata struct {
	Title string `json:"title"`
}

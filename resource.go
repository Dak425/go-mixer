package structs

type ResourceCollection []Resource
type ResourceStringMap map[string]Resource
type ResourceIntMap map[int]Resource
type ResourceUIntMap map[uint]Resource

// Resources are general use items that are stored and linked to other entities within Mixer.They usually represent images or backgrounds for a channel or user.
type Resource struct {
	Id         uint     `json:"id"`
	Meta       struct{} `json:"meta"`
	Relid      uint     `json:"relid"`
	RemotePath string   `json:"remotePath"`
	Store      string   `json:"store"`
	Type       string   `json:"type"`
	Url        string   `json:"url"`
}

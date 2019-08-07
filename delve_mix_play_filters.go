package structs

type DelveMixPlayFiltersCollection []DelveMixPlayFilters
type DelveMixPlayFiltersStringMap map[string]DelveMixPlayFilters
type DelveMixPlayFiltersIntMap map[int]DelveMixPlayFilters
type DelveMixPlayFiltersUIntMap map[uint]DelveMixPlayFilters

// Dictates the MixPlay games which are filterable on the browse page.
type DelveMixPlayFilters struct {
	Description    struct{} `json:"description"`
	IntegrationIds []uint   `json:"integrationIds"`
	Name           struct{} `json:"name"`
}

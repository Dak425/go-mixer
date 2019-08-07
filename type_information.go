package structs

type TypeInformationCollection []TypeInformation
type TypeInformationStringMap map[string]TypeInformation
type TypeInformationIntMap map[int]TypeInformation
type TypeInformationUIntMap map[uint]TypeInformation

// Describes a type stored in the Mixer backend.
type TypeInformation struct {
	Aliases        []string  `json:"aliases"`
	AvailableAt    string    `json:"availableAt"`
	BackgroundUrl  string    `json:"backgroundUrl"`
	CoverUrl       string    `json:"coverUrl"`
	Description    string    `json:"description"`
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Online         int       `json:"online"`
	Parent         string    `json:"parent"`
	Source         string    `json:"source"`
	Titles         TypeTitle `json:"titles"`
	Verified       bool      `json:"verified"`
	ViewersCurrent int       `json:"viewersCurrent"`
}

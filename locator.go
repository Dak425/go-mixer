package mixerstructs

type LocatorCollection []Locator
type LocatorStringMap map[string]Locator
type LocatorIntMap map[int]Locator
type LocatorUIntMap map[uint]Locator

// Information required to find a clip.
type Locator struct {
	LocatorType string `json:"locatorType"`
	Uri         string `json:"uri"`
}

package structs

type DelveCarouselCollection []DelveCarousel
type DelveCarouselStringMap map[string]DelveCarousel
type DelveCarouselIntMap map[int]DelveCarousel
type DelveCarouselUIntMap map[uint]DelveCarousel

// A large carousel for display on the homepage.
type DelveCarousel struct {
	Channels ChannelCollection `json:"channels"`
	Type     string            `json:"type"`
}

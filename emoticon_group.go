package mixerstructs

type EmoticonGroupCollection []EmoticonGroup
type EmoticonGroupStringMap map[string]EmoticonGroup
type EmoticonGroupIntMap map[int]EmoticonGroup
type EmoticonGroupUIntMap map[uint]EmoticonGroup

// A list of emoticons and their coordinates within an EmoticonPack.
type EmoticonGroup struct {
	EmoticonStringMap
}

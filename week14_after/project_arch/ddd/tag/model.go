package tag

type TagSet struct {
	Items []*Tag
}

type Tag struct {
	Key string
	Value string
}
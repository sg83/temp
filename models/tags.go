package data

type Tag struct {
	// Tag name
	//
	// required: true
	Tag string `json:"tag"`
	// Number of tags for that day.
	//
	// required: true
	Count int `json:"count"`
	// List of ids for the last 10 articles entered for that day.
	//
	// required: true
	Articles []string `json:"articles"`
	// List of tags that are on the articles that the current tag is on for the same day.
	//
	// required: true
	RelatedTags []string `json:"related_tags"`
}

// TagDatabase represents an in-memory database for tags
type TagDb struct {
	Tags map[string]map[string][]string // map[tagName]map[date][]articleIDs
}



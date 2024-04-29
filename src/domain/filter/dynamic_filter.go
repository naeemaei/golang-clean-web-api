package filter

type Sort struct {
	ColId string `json:"colId"`
	Sort  string `json:"sort"`
}

type Filter struct {
	// contains notContains equals notEqual startsWith lessThan lessThanOrEqual greaterThan greaterThanOrEqual inRange endsWith
	Type string `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`
	// text number
	FilterType string `json:"filterType"`
}

type DynamicFilter struct {
	Sort   *[]Sort           `json:"sort"`
	Filter map[string]Filter `json:"filter"`
}

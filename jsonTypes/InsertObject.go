package jsontypes

// InsertObject describes a json schema for inserting an item from store
type InsertObject struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

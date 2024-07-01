package csvutil

// CSVCol is a struct that holds the key and name of a column in a CSV file.
type CSVCol struct {
	Key   string
	Name  *string
	index int
}

// Col creates a new CSVCol struct.
func Col(key string, altNames ...string) *CSVCol {
	var name *string
	if len(altNames) > 0 {
		name = &altNames[0]
	}

	return &CSVCol{Key: key, Name: name}
}

// ColFromKeys creates a slice of CSVCol structs.
func ColFromKeys(keys ...string) []*CSVCol {
	recs := make([]*CSVCol, 0, len(keys))
	for _, key := range keys {
		recs = append(recs, Col(key))
	}

	return recs
}

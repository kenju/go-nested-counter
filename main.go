package nestedcounter

type Count int

type RootKey interface {}

type NestedKey interface {}

// Counter count incremented value per RootKey and NestedKey.
// Counter has a mapping table internally to count the values.
type Counter struct {
	RootKeyMap map[RootKey]map[NestedKey]Count
}

// NewCounter returns the reference to a new Counter.
func NewCounter() *Counter {
	return &Counter{
		RootKeyMap: map[RootKey]map[NestedKey]Count{},
	}
}

// Increment increment the count value by one for given rootKey and nestedKey.
// Create a new mapping if the given rootKey/nestedKey pair does not exist yet.
func (c *Counter) Increment(rootKey RootKey, nestedKey NestedKey, impression Count) {
	if _, ok := c.RootKeyMap[rootKey]; ok {
		c.RootKeyMap[rootKey][nestedKey] += impression
	} else {
		c.RootKeyMap[rootKey] = map[NestedKey]Count{nestedKey: impression}
	}
}

// GetCount return the count value by given rootKey/nestedKey.
func (c *Counter) GetCount(rootKey RootKey, nestedKey NestedKey) int {
	return int(c.RootKeyMap[rootKey][nestedKey])
}

// Size return the number of RootKey in the mapping table.
func (c *Counter) Size() int {
	return len(c.RootKeyMap)
}

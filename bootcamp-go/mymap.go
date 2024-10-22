package bootcamp

// MyMap structure definition
type MyMap struct {
	keys   []string
	values []interface{}
}

// Set method to set the value v associated with key k
func (m *MyMap) Set(k string, v interface{}) {
	// Check if the key already exists
	for i, key := range m.keys {
		if key == k {
			m.values[i] = v // Update existing value
			return
		}
	}
	// Add new key-value pair
	m.keys = append(m.keys, k)
	m.values = append(m.values, v)
}

// Get method to retrieve the value associated with key k
func (m *MyMap) Get(k string) interface{} {
	for i, key := range m.keys {
		if key == k {
			return m.values[i] // Return associated value
		}
	}
	return nil // Return nil if key does not exist
}

// Has method to check if the map contains key k
func (m *MyMap) Has(k string) bool {
	for _, key := range m.keys {
		if key == k {
			return true // Key exists
		}
	}
	return false // Key does not exist
}

// Delete method to remove the key k and its associated value
func (m *MyMap) Delete(k string) {
	for i, key := range m.keys {
		if key == k {
			// Remove the key-value pair by slicing
			m.keys = append(m.keys[:i], m.keys[i+1:]...)       // Remove key
			m.values = append(m.values[:i], m.values[i+1:]...) // Remove value
			return
		}
	}
}

// Items method to return a slice of key-value pairs
func (m *MyMap) Items() []struct {
	Key   string
	Value interface{}
} {
	items := make([]struct {
		Key   string
		Value interface{}
	}, len(m.keys))
	for i := range m.keys {
		items[i] = struct {
			Key   string
			Value interface{}
		}{Key: m.keys[i], Value: m.values[i]}
	}
	return items
}

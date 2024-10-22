package bootcamp

func MapEqual(m1 map[string]int, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if w, ok := m2[k]; !ok || v != w {
			return false
		}
	}

	return true
}

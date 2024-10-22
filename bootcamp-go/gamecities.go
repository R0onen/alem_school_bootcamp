package bootcamp

func GameCities(cities []string) []string {
	used := make(map[string]bool)
	ans := []string{}
	temp := []string{}

	solve_GameCities(&ans, used, cities, temp)

	return ans
}

func solve_GameCities(ans *[]string, used map[string]bool, cities []string, temp []string) {
	if len(temp) > len(*ans) {
		*ans = make([]string, len(temp))
		copy(*ans, temp)
	}

	for _, city := range cities {
		last := getLast(temp)
		if !used[city] && (last == "" || toupper(last[len(last)-1]) == toupper(city[0])) {
			used[city] = true
			temp = append(temp, city)
			solve_GameCities(ans, used, cities, temp)
			temp = temp[:len(temp)-1]
			used[city] = false
		}
	}
}

func getLast(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

func toupper(s byte) byte {
	if s >= 'a' && s <= 'z' {
		return s - 32
	}
	return s
}

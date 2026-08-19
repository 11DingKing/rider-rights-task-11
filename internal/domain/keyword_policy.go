package domain

func MatchAnyKeyword(required, actual []string) bool {
	if len(required) == 0 {
		return true
	}
	available := make(map[string]struct{}, len(actual))
	for _, keyword := range actual {
		available[keyword] = struct{}{}
	}
	matched := 0
	for _, keyword := range required {
		if _, ok := available[keyword]; ok {
			matched++
		}
	}
	return matched == len(required)
}

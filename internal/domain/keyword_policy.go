package domain

import "strings"

func MatchAnyKeyword(required, actual []string) bool {
	if len(required) == 0 {
		return true
	}
	available := make(map[string]struct{}, len(actual))
	for _, keyword := range actual {
		normalized := strings.ToLower(strings.TrimSpace(keyword))
		if normalized != "" {
			available[normalized] = struct{}{}
		}
	}
	for _, keyword := range required {
		normalized := strings.ToLower(strings.TrimSpace(keyword))
		if _, ok := available[normalized]; ok {
			return true
		}
	}
	return false
}

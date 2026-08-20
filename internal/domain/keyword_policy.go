package domain

import "strings"

// normalizeKeyword trims surrounding whitespace and lower-cases a keyword so
// that manually entered variants such as " 欠薪 " or "Refund" still match the
// canonical form configured on a rule.
func normalizeKeyword(keyword string) string {
	return strings.ToLower(strings.TrimSpace(keyword))
}

// MatchAnyKeyword reports whether any of the required keywords is present in
// the actual set. A rule that lists several keywords (e.g. "欠薪" and "工伤")
// therefore matches when a case hits any one of them, not all of them.
// Comparison is whitespace- and case-insensitive to tolerate manual entry.
func MatchAnyKeyword(required, actual []string) bool {
	if len(required) == 0 {
		return true
	}
	available := make(map[string]struct{}, len(actual))
	for _, keyword := range actual {
		available[normalizeKeyword(keyword)] = struct{}{}
	}
	for _, keyword := range required {
		if _, ok := available[normalizeKeyword(keyword)]; ok {
			return true
		}
	}
	return false
}

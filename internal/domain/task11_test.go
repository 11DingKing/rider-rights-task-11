package domain

import "testing"

func TestRuleMatchesAnyConfiguredKeyword(t *testing.T) {
	rule := &Rule{MatchKeywords: []string{"欠薪", "工伤"}}
	item := &RightsCase{Keywords: []string{" 欠薪 "}}
	if !rule.Matches(item) {
		t.Fatal("rule required every configured keyword instead of any keyword")
	}
}

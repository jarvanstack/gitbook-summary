package matcher

import "regexp"

// 字符串匹配器
type Marcher interface {
	Match(str string) bool
}

// 正则表达式匹配器
type RegexMatcher struct {
	regexs []string
}

func NewRegexMatcher(regexs []string) *RegexMatcher {
	return &RegexMatcher{
		regexs: regexs,
	}
}

func (r *RegexMatcher) Match(str string) bool {
	for _, regex := range r.regexs {
		if ok, _ := regexp.MatchString(regex, str); ok {
			return true
		}
	}
	return false
}

package clip

import "regexp"

const (
	regexShortOptionPattern = "^[-]{1}[a-zA-Z]+$"
	regexLongOptionPattern  = "^[-]{2}[a-zA-Z]{1}[a-zA-Z0-9-_]{0,14}[a-zA-Z0-9]{1}$"
)

var (
	regexShortOption = regexp.MustCompile(regexShortOptionPattern)
	regexLongOption  = regexp.MustCompile(regexLongOptionPattern)
)

func matchShortOption(short string) bool {
	return regexShortOption.MatchString(short)
}

func matchLongOption(long string) bool {
	return regexLongOption.MatchString(long)
}

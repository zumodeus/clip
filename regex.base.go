package clip

import "regexp"

const (
	regexNameStringPattern  = "^[a-zA-Z]{1}[a-zA-Z0-9-.]{0,14}[a-zA-Z0-9]{1}$"
	regexShortStringPattern = "^[a-zA-Z]{1}$"
	regexLongStringPattern  = "^[a-zA-Z]{1}[a-zA-Z0-9-_]{0,14}[a-zA-Z0-9]{1}$"
)

var (
	regexNameString  = regexp.MustCompile(regexNameStringPattern)
	regexShortString = regexp.MustCompile(regexShortStringPattern)
	regexLongString  = regexp.MustCompile(regexLongStringPattern)
)

func matchNameString(name string) bool {
	return regexNameString.MatchString(name)
}

func matchShortString(short string) bool {
	return regexShortString.MatchString(short)
}

func matchLongString(long string) bool {
	return regexLongString.MatchString(long)
}

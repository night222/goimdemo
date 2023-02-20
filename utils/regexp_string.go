package utils

import "regexp"

//正则匹配字符串
func RegexpString(Regexp, Data string) []string {
	var result []string
	re, err := regexp.Compile(Regexp)
	if err != nil {
		return result
	}
	result = re.FindStringSubmatch(Data)
	return result
}

package cudnn_log_parser

import (
	"regexp"
	"strings"
)

func trimSuffix(s string, suffs ...string) string {
	if len(suffs) == 0 {
		return s
	}
	for _, suff := range suffs {
		s = strings.TrimSuffix(s, suff)
	}
	return s
}

func trimPrefix(s string, suffs ...string) string {
	if len(suffs) == 0 {
		return s
	}
	for _, suff := range suffs {
		s = strings.TrimPrefix(s, suff)
	}
	return s
}

func indentOf(line string) int {
	indentN := 0
	line = trimPrefix(line, "I! ", "i! ")
	for _, e := range line {
		if e != ' ' {
			break
		}
		indentN++
	}
	return indentN
}

func getRegexParams(regEx, s string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(s)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

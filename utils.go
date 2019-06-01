package cudnn_log_parser

import "strings"

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

func indentOf(line string) {
	indentN := 0
	line = trimPrefix(line, "I! ", "i! ")
	for _, e := range line {
		if e != " " {
			break
		}
		indentN++
	}
	return indentN
}

package cudnn_log_parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/Unknwon/com"
	"github.com/pkg/errors"
)

func indent(n int) string {
	if n == 0 {
		return ""
	}

	s := make([]byte, n)
	for ii := range s {
		s[ii] = ' '
	}
	return string(s)
}

func processLine(line string) {
	indentN := indentOf(line)
	if strings.HasPrefix(line, "I! ") {
		re := regexp.MustCompile(`I!.*function\s+(\w+)\(.*`)
		for i, match := range re.FindAllString(line, -1) {
			return indent(indentN) + match + ":"
		}
	}

	var re = regexp.MustCompile(`(?m)i!\s+(.*):\s((.*)=(.*);)*`)
	for i, match := range re.FindAllString(line, -1) {
		fmt.Println(match, "found at index", i)
	}
}

func toYaml(log string) ([]byte, error) {
	lines := strings.Split(log, "\n")
	for ii, line := range lines {
		line = processLine(line)
	}
	return strings.Join(lines, "\n")
}

func Parse(log string) (*Info, error) {
	bts, err := toYaml(log)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse into yaml")
	}
	info := new(Info)
	err = yaml.Unmashal(bts, info)
	return info, err
}

func ParseFile(file string) (*Info, error) {
	if !com.IsFile(file) {
		return nil, errors.Errorf("unable to find %v", file)
	}
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return Parse(string(bts))
}

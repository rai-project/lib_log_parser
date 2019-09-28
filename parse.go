package cudnn_log_parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/atotto/clipboard"
	"github.com/gocarina/gocsv"
	"github.com/k0kubun/pp"

	"github.com/Unknwon/com"
	"github.com/cydev/zero"
	"github.com/pkg/errors"
)

var verboseDebug = false

func indent(n int) string {
	if n == 0 {
		return ""
	}

	s := make([]byte, 2*n)
	for ii := range s {
		s[ii] = ' '
	}
	return string(s)
}

func parseFunctionName(line string) (string, error) {
	if !strings.HasPrefix(line, "I! CuDNN") {
		return "", errors.New("not a function line")
	}

	params := getRegexParams(`I!.*function\s+(?P<functionName>\w+)\(.*`, line)
	if f, ok := params["functionName"]; ok {
		return f, nil
	}

	return "", errors.New("cannot find function")
}

func parseTimeStamp(line string) (time.Time, error) {
	if !strings.HasPrefix(line, "i! Time") {
		return time.Time{}, errors.New("not a function line")
	}

	params := getRegexParams(`i!\s+Time:\s+(?P<time>[\d-:T.]+)+\s+.*`, line)
	if t, ok := params["time"]; ok {
		// 2019-06-04T18:01:43.967769
		return time.Parse("2006-01-02T15:04:05.999999999", t)
	}

	return time.Time{}, errors.New("cannot find time")
}

func processLine(info *Info, line string) {
	// indentN := indentOf(line)

	if t, err := parseTimeStamp(line); err == nil {
		info.TimeStamp = t
		return
	}

	if f, err := parseFunctionName(line); err == nil {
		info.FunctionKind = strings.TrimPrefix(f, "cudnn")
		info.FunctionName = f
		return
	}

	// var re = regexp.MustCompile(`(?m)i!\s+(.*):\s((.*)=(.*);)*`)
	// for i, match := range re.FindAllString(line, -1) {
	// 	fmt.Println(match, "found at index", i)
	// }
}

var indentattion = "     "
var nestIndentation = "i!" + indentattion

func processNest(info *Info, log string) {
	lines := strings.Split(log, "\n")
	for ii, line := range lines {
		lines[ii] = strings.TrimPrefix(line, nestIndentation)
	}
	// procesedLines := []string{}
	// if err := copier.Copy(procesedLines, lines); err != nil {
	//   panic("unable to copy lines")
	// }
	endOfLineColonRe := regexp.MustCompile(`(?m).*:.*:$`)
	typeAssignmentRe := regexp.MustCompile(`(?m)\s+type=`)
	processedLines := []string{}
	for _, line := range lines {
		indentN := indentOf(line)
		nextIndent := "\n" + indent(indentN+1)

		if endOfLineColonRe.MatchString(line) {
			line = strings.TrimSuffix(line, ":")
		}
		if strings.Contains(line, ": ") {
			line = strings.ReplaceAll(line, ": ", ": {") + "}"
		}
		if indentN == 0 {
			line = strings.ReplaceAll(line, ": {", ":"+nextIndent+indent(indentN+1)+"opinfo: {")
		}
		line = strings.ReplaceAll(line, "=", ": ")
		line = strings.ReplaceAll(line, ";", ", ")
		line = strings.ReplaceAll(line, ", }", " }")
		if false {
			if typeAssignmentRe.MatchString(line) {
				line = typeAssignmentRe.ReplaceAllString(line, nextIndent+"type: ")
			} else {
				line = strings.ReplaceAll(line, ":", ":"+nextIndent)
			}
			if endOfLineColonRe.MatchString(line) {
				line = strings.TrimSuffix(line, ":")
				// line = strings.ReplaceAll(line, " type=", "type: ")
				fmt.Println(line)
				fmt.Println("----------------------")
			}
			line = strings.ReplaceAll(line, "=", ": ")
			line = strings.ReplaceAll(line, ";", nextIndent)
		}
		for _, newLine := range strings.Split(line, "\n") {
			if strings.TrimSpace(newLine) == "" {
				continue
			}
			processedLines = append(processedLines, newLine)
		}
	}
	ymlBlock := strings.Join(processedLines, "\n")
	// clipboard.WriteAll(ymlBlock)
	attributes := new(Attributes)
	err := yaml.Unmarshal([]byte(ymlBlock), &attributes)
	if err != nil {
		clipboard.WriteAll(ymlBlock)
		panic(err)
	}
	if verboseDebug {
		pp.Println(attributes)
		yml, err := yaml.Marshal(&attributes)
		if err != nil {
			panic(err)
		}
		_ = yml
	}
	info.Attributes = attributes
}

func toInfos(log string) (Infos, error) {
	lines := strings.Split(log, "\n")
	nst := ""
	infos := make([]Info, 0, len(lines))
	var currentInfo Info
	for _, line := range lines {
		if strings.HasPrefix(line, "I! CuDNN") && !zero.IsZero(currentInfo) {
			infos = append(infos, currentInfo)
		}
		if !strings.HasPrefix(line, nestIndentation) {
			processLine(&currentInfo, line)
			continue
		}
		nst += line + "\n"
	}
	if nst != "" {
		processNest(&currentInfo, nst)
	}
	infos = append(infos, currentInfo)
	return infos, nil
}

func (infos0 *Infos) computeDurations() {
	infos := *infos0
	for ii := 0; ii < len(infos)-1; ii++ {
		nextTimeStamp := infos[ii+1].TimeStamp
		currTimeStamp := infos[ii].TimeStamp
		infos[ii].Duration = Duration(nextTimeStamp.Sub(currTimeStamp))
	}
	infos = infos[1:]
}

func (infos Infos) ToCSV(filename string) error {
	os.MkdirAll(path.Dir(filename), os.ModePerm)
	cvsFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer cvsFile.Close()
	return gocsv.MarshalFile(infos, cvsFile)
}

func ParseBlock(log string) (Infos, error) {
	infos, err := toInfos(log)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse into infos")
	}
	infos.computeDurations()
	return infos, err
}

func Parse(log string) (Infos, error) {
	lines := strings.Split(log, "\n")
	var infos Infos
	buf := ""
	for _, line := range lines {
		if strings.TrimSpace(line) == "" && buf != "" {
			blk, err := toInfos(buf)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse into infos")
			}
			infos = append(infos, blk...)
			buf = ""
			continue
		}
		buf += line + "\n"
	}

	infos.computeDurations()
	return infos, nil
}

func ParseFile(file string) ([]Info, error) {
	if !com.IsFile(file) {
		return nil, errors.Errorf("unable to find %v", file)
	}
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return Parse(string(bts))
}

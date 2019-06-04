package cudnn_log_parser

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gocarina/gocsv"

	"github.com/Unknwon/com"
	"github.com/cydev/zero"
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
	}

	if f, err := parseFunctionName(line); err == nil {
		info.FunctionName = f
	}

	// var re = regexp.MustCompile(`(?m)i!\s+(.*):\s((.*)=(.*);)*`)
	// for i, match := range re.FindAllString(line, -1) {
	// 	fmt.Println(match, "found at index", i)
	// }
}

func toInfos(log string) (Infos, error) {
	lines := strings.Split(log, "\n")
	infos := make([]Info, 0, len(lines))
	var currentInfo Info
	for _, line := range lines {
		if strings.HasPrefix(line, "I! CuDNN") && !zero.IsZero(currentInfo) {
			infos = append(infos, currentInfo)
		}
		processLine(&currentInfo, line)
	}
	infos = append(infos, currentInfo)
	return infos, nil
}

func (infos0 *Infos) computeDurations() {
	infos := *infos0
	for ii := 1; ii < len(infos); ii++ {
		infos[ii].Duration = Duration(infos[ii].TimeStamp.Sub(infos[ii-1].TimeStamp))
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

func Parse(log string) (Infos, error) {
	infos, err := toInfos(log)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse into infos")
	}
	infos.computeDurations()
	return infos, err
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

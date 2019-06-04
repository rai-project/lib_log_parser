package cudnn_log_parser

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/GeertJohan/go-sourcepath"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

var simpleTest string

func TestParse(t *testing.T) {
	readData()
	info, err := Parse(simpleTest)
	assert.NoError(t, err)

	pp.Println(info)
}

func readSimpleData() {
	bts, err := ioutil.ReadFile(filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "simple.log"))
	if err != nil {
		assert(err)
	}
	simpleTest = string(bts)
}

func readData() {
	readSimpleData()
}

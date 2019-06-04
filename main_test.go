package cudnn_log_parser

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/GeertJohan/go-sourcepath"
	"github.com/k0kubun/pp"
)

var (
	simpleTestData  string
	alexnetTestData string
)

func readSimpleData() {
	bts, err := ioutil.ReadFile(filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "simple.log"))
	if err != nil {
		panic(err)
	}
	simpleTestData = string(bts)
}

func readAlexNetCUDNNData() {
	bts, err := ioutil.ReadFile(filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "alexnet_cudnn.log"))
	if err != nil {
		panic(err)
	}
	alexnetTestData = string(bts)
}

func readData() {
	readSimpleData()
	readAlexNetCUDNNData()
}

func TestMain(m *testing.M) {
	readData()
	pp.WithLineInfo = true
	os.Exit(m.Run())
}

package cudnn_log_parser

import (
	"testing"

	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

func DISABLED_TestParse(t *testing.T) {
	info, err := Parse(simpleTestData)
	assert.NoError(t, err)

	pp.Println(info)
}

func TestParseToCSV(t *testing.T) {
	info, err := Parse(alexnetTestData)
	assert.NoError(t, err)

	info.ToCSV("_fixtures/alexnet_cudnn.csv")
}

func TestParseToJSON(t *testing.T) {
	info, err := Parse(alexnetTestData)
	assert.NoError(t, err)

	info.ToJSON("_fixtures/alexnet_cudnn.json")
}

package cudnn_log_parser

import (
	"testing"

	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

var simpleTest string

func TestParse(t *testing.T) {
	info, err := Parse(simpleTest)
	assert.NoError(t, err)

	pp.Println(info)
}

func TestParseToCSV(t *testing.T) {
	info, err := Parse(simpleTest)
	assert.NoError(t, err)

	info.ToCSV("_fixtures/alexnet_cudnn.csv")
}

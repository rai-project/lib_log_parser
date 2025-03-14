package cudnn_log_parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/k0kubun/pp"
	"github.com/spf13/cast"
	"github.com/iancoleman/strcase"
)

type Duration time.Duration
type Attributes map[string]interface{}

type Info struct {
	FunctionKind string      `csv:"function_kind" json:"function_kind,omitempty" yaml:"function_kind,omitempty"`
	FunctionName string      `csv:"function_name" json:"function_name,omitempty" yaml:"function_name,omitempty"`
	Duration     Duration    `csv:"duration_ms"  json:"duration,omitempty" yaml:"duration,omitempty"`
	TimeStamp    time.Time   `csv:"time_stamp" json:"time_stamp,omitempty" yaml:"time_stamp,omitempty"`
	Attributes   *Attributes `csv:"attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

type Infos []Info

func (dur *Duration) MarshalCSV() (string, error) {
	return fmt.Sprintf("%v", float64(time.Duration(*dur))/float64(time.Millisecond)), nil
}

func (attrs0 *Attributes) MarshalJSON() ([]byte, error) {
	attrs := *attrs0
	// pp.Println(attrs)
	buffer := bytes.NewBufferString("{")
	length := len(attrs)
	count := 0
	for key, value := range attrs {
		var err error
		jsonValue := []byte("<<ERROR>>")
		switch value := value.(type) {
		case map[interface{}]interface{}:
			mp := make(Attributes)
			for k, v := range value {
				mp[strcase.ToSnake(cast.ToString(k))] = v
			}
			// pp.Println(mp)
			jsonValue, err = json.Marshal(&mp)
			if err != nil {
				panic(err)
				return nil, err
			}
		case map[interface{}]string:
			mp := make(Attributes)
			for k, v := range value {
				mp[strcase.ToSnake(cast.ToString(k))] = v
			}
			// pp.Println(mp)
			jsonValue, err = json.Marshal(&mp)
			if err != nil {
				panic(err)
				return nil, err
			}
		default:
			jsonValue, err = json.Marshal(value)
			if err != nil {
				pp.Println(err)
				return nil, err
			}
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s", strcase.ToSnake(cast.ToString(key)), cast.ToString(jsonValue)))
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func (attrs *Attributes) MarshalCSV() (string, error) {
	bts, err := json.Marshal(attrs)
	if err != nil {
		panic(err)
		return "", err
	}
	return string(bts), nil
}

package cudnn_log_parser

import (
	"fmt"
	"time"
)

type Duration time.Duration

type Info struct {
	FunctionName string    `csv:"function_name" json:"function_name,omitempty"`
	Duration     Duration  `csv:"duration"  json:"duration,omitempty"`
	TimeStamp    time.Time `csv:"time_stamp" json:"time_stamp,omitempty"`
}

type Infos []Info

func (dur *Duration) MarshalCSV() (string, error) {
	return fmt.Sprintf("%v", float64(time.Duration(*dur))/float64(time.Microsecond)), nil
}

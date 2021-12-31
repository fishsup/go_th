package pipefilter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data shoud be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	f := new(SplitFilter)
	f.delimiter = delimiter
	return f
}

func (s *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, s.delimiter)
	return parts, nil
}

package pipefilter

import (
	"errors"
)

var SumFilterWrongFormatError = errors.New("Parameter type error, should be int slice")

type SumFilter struct{}

func NewSumFilter() *SumFilter {
	return new(SumFilter)
}

func (s *SumFilter) Process(data Request) (Response, error) {
	slice, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, item := range slice {
		ret += item
	}
	return ret, nil
}

package pipefilter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("Parameter type error, should be string slice")

type ToIntFilter struct{}

func NewToIntFilter() *ToIntFilter {
	return new(ToIntFilter)
}

func (s *ToIntFilter) Process(data Request) (Response, error) {
	slice, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, item := range slice {
		s, err := strconv.Atoi(item)
		if err != nil {
			return nil, ToIntFilterWrongFormatError
		}
		ret = append(ret, s)
	}
	return ret, nil
}

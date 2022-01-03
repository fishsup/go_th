package pipefilter

import "testing"

func TestFilter(t *testing.T) {
	sf := NewSplitFilter(",")
	cf := NewToIntFilter()
	sumf := NewSumFilter()
	sp := NewStraightPipeline("TESTPL", sf, cf, sumf)
	ret, err := sp.Process("1,2,3,4,5")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 15 {
		t.Fatal("result error")
	}

}

type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{Name: name, Filters: &filters}
}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}

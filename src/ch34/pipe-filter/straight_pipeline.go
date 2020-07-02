package pipefilter

type StraightPipeline struct {
	Name string
	Filters *[]Filter
}

// 参数1 名字， 参数二 可变参数 filter
func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline {
		Name: name,
		Filters: &filters,
	}
}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	// 轮询 filters
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
package pipefilter

import (
	"errors"
	"strings"
)

var SpliFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct{
	delimiter string // 分隔符
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string) // 检查数据格式/类型，是否可以处理
	if !ok {
		return nil, SpliFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
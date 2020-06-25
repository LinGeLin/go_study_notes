package client

import (
	"testing"
	"ch15/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	//t.Log(series.square(1))
}
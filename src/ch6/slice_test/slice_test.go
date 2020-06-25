package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	for i :=0; i<10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceCap(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	t.Log(len(arr), cap(arr))
	s := arr[:2]
	t.Log(len(s), cap(s))
} 

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Log(len(year), cap(year))
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	Summer := year[4:7]
	t.Log(Summer, len(Summer), cap(Summer))
}

func TestSliceShareMemory1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	s1 := nums[:2]
	t.Log(s1, len(s1), cap(s1))
	s2 := nums[1:3]
	t.Log(s2, len(s2), cap(s2))
	s1 = append(s1, 5)
	t.Log(s1, len(s1), cap(s1))
	t.Log(s2, len(s2), cap(s2))

	s2 = append(s2, 6)
	t.Log("s1", s1, len(s1), cap(s1))
	t.Log("s2", s2, len(s2), cap(s2))
	t.Log("nums", nums, len(nums), cap(nums))

	s2 = append(s2, 7)
	t.Log("s1", s1, len(s1), cap(s1))
	t.Log("s2", s2, len(s2), cap(s2))
	t.Log("nums", nums, len(nums), cap(nums))

	s2 = append(s2, 8)
	t.Log("s1", s1, len(s1), cap(s1))
	t.Log("s2", s2, len(s2), cap(s2))
	t.Log("nums", nums, len(nums), cap(nums))
}

func TestSliceComparing(t *testing.T) {
	a :=[]int{1, 2, 3, 4}
	b :=[]int{1, 2, 3, 4}
	if a == nil && b == nil {
		t.Log("equal")
	}
}
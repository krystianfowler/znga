package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int                // s is a slice of int
	fmt.Println("len", len(s)) // len is "nil safe"
	if s == nil {
		fmt.Println("nil slice")
		fmt.Print(s)
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] // slicing operation, half-open range
	fmt.Printf("s3 = %#v\n", s3)

	// fmt.Println(s2[:100]) // out of bound, panic
	s3 = append(s3, 100)
	fmt.Printf("s3 (append)= %#v\n", s3)
	fmt.Printf("s2 (append)= %#v\n", s2) // s2 is changed as well
	fmt.Printf("s2: len=%d, cap=%d'\n", len(s2), cap(s2))
	fmt.Printf("s2: len=%d, cap=%d'\n", len(s3), cap(s3))

	var s4 []int
	// s4 := make([]int, 0, 1_000) // single allocation
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Println("s4", len(s4), cap(s4))
	fmt.Println(concat([]string{"a", "b"}, []string{"c", "d"}))

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))
	vs = []float64{2, 1, 3, 4}
	fmt.Println(median(vs))
	fmt.Println(vs)

	fmt.Println(median(nil))
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slices")
	}
	// Copy in order not to change values
	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	const n = 2
	v := (values[i-1] + values[i]) / n // type of constant is generally defined at place of usage
	return v, nil
}

func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) { // enough space in underlying array
		s = s[:len(s)+1]
	} else { // need to re-allocate and copy
		fmt.Printf("reallocate: %d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]

	}

	s[i] = v
	return s
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum1(arr []int, target int) []int {
	if len(arr) < 2 {
		return nil
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func sum2(arr []int, target int) []int {
	if len(arr) < 2 {
		return nil
	}

	m := make(map[int]int)
	for i := range arr {
		need := target - arr[i]
		if x, ok := m[need]; ok {
			return []int{x, i}
		}
		m[arr[i]] = i
	}

	return nil
}

const max = 10000000

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	arr := make([]int, max)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(max)
	}

	for i := 10; i <= max; i = i * 10 {
		t1 := time.Duration(0)
		t2 := time.Duration(0)

		for k := 0; k < 100; k++ {
			target := rand.Intn(max)

			s1 := time.Now()
			sum1(arr[0:i], target)
			t1 += time.Since(s1)

			s2 := time.Now()
			sum2(arr[0:i], target)
			t2 += time.Since(s2)
		}

		fmt.Printf("n = %d \n\t t1 = %v \n\t t2 = %v \n\n", i, t1, t2)
	}
}

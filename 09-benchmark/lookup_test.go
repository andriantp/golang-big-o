package main

import (
	"fmt"
	"testing"
)

func benchmarkSliceLookup(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

func benchmarkMapLookup(items map[string]struct{}, target string) bool {
	_, ok := items[target]
	return ok
}

func BenchmarkSliceLookup(b *testing.B) {
	for _, size := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("n=%d", size), func(b *testing.B) {
			items := make([]string, size)
			for i := range items {
				items[i] = fmt.Sprintf("camera-%d", i)
			}

			target := items[size-1]
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				benchmarkSliceLookup(items, target)
			}
		})
	}
}

func BenchmarkMapLookup(b *testing.B) {
	for _, size := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("n=%d", size), func(b *testing.B) {
			items := make(map[string]struct{}, size)
			for i := 0; i < size; i++ {
				items[fmt.Sprintf("camera-%d", i)] = struct{}{}
			}

			target := fmt.Sprintf("camera-%d", size-1)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				benchmarkMapLookup(items, target)
			}
		})
	}
}

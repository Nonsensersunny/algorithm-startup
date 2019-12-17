package main

import (
	"fmt"
	"testing"
)

func TestMaxpoints(t *testing.T) {
	var N int64
	fmt.Scanln(&N)
	var points []*point
	for i := int64(0); i < N; i++ {
		p := new(point)
		fmt.Scanf("%d %d\n", p.X, p.Y)
		points = append(points, p)
	}
	SortpointsByX(points)
	for _, p := range points {
		p.Printpoint()
	}
}

package main

import (
	"fmt"
)

type point struct {
	X int64
	Y int64
}

func (p *point) lt(q *point) bool {
	if q.X > p.X && q.Y > p.Y {
		return true
	}
	return false
}

func (p *point) GetMax(q []*point) []*point {
	var res []*point
	for _, pt := range q {
		isMax := true
		for i := 0; i < len(q); i++ {
			if pt.lt(q[i]) {
				isMax = false
				break
			}
		}
		if isMax {
			res = append(res, pt)
		}
	}
	return res
}

func SortpointsByX(pts []*point) {
	for i := 0; i < len(pts)-1; i++ {
		tmp := i
		for j := i; j < len(pts); j++ {
			if pts[j].X < pts[tmp].X {
				tmp = j
			}
		}
		pts[tmp], pts[i] = pts[i], pts[tmp]
	}
}

func (p *point) Printpoint() {
	fmt.Printf("%d %d\n", p.X, p.Y)
}

func main() {
	fmt.Printf("Input the base:")
	base := new(point)
	fmt.Scanf("%d %d\n", &base.X, &base.Y)
	var N int64
	fmt.Scanln(&N)
	var points []*point
	for i := int64(0); i < N; i++ {
		p := new(point)
		fmt.Scanf("%d %d\n", &p.X, &p.Y)
		points = append(points, p)
	}
	fmt.Println("==START==")
	for _, p := range points {
		p.Printpoint()
	}
	fmt.Println("==END==")
	max := base.GetMax(points)
	SortpointsByX(max)
	fmt.Println("==START==")
	for _, p := range max {
		p.Printpoint()
	}
	fmt.Println("==END==")
}

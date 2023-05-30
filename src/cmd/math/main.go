package main

import (
	"github.com/surdeus/gox/src/gx"
	"fmt"
)

func main() {
	lines := []gx.Line{
		gx.LineSegment{
			gx.Point{0, 1},
			gx.Point{1, 2},
		}.Line(),
		gx.LineSegment{
			gx.Point{0, 5},
			gx.Point{1, 2},
		}.Line(),
		gx.LineSegment{
			gx.Point{-1, -1},
			gx.Point{1, 50},
		}.Line(),
	}
	
	for _, l := range lines { fmt.Println(l) }
	
	l1 := gx.LineSegment{
		gx.Point{0, 0},
		gx.Point{1, 1},
	}.Line()
	
	l2 := gx.LineSegment{
		gx.Point{0, 1},
		gx.Point{1, 0},
	}.Line()
	fmt.Println(l1.Crosses(l2))
	fmt.Println(l1.ContainsPoint(gx.Point{1, 4}))
	
	t := gx.Rectangle{
		Transform: gx.Transform{
			S: gx.Vector{1, 1},
			P: gx.Point{0, 200},
		},
		W: 100,
		H: 200,
	}
	
	points := []gx.Point{
		gx.Point{},
		gx.Point{100, 0},
		gx.Point{0, 99},
		gx.Point{.1, .1},
		gx.Point{-1, -1},
		gx.Point{1, 1},
		gx.Point{101, 1},
		gx.Point{100, 1},
		gx.Point{50, 1},
	}
	
	ts := t.Triangles()
	t1 := ts[0]
	t2 := ts[1]
	fmt.Printf("Rectangle triangles:\n\t%v\n\t%v\n", t1, t2)
	for _, p := range points {
		fmt.Println(p, t.ContainsPoint(p))
	}
	
}




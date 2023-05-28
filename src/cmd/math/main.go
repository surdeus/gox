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
	
	for _, l := range lines {
		fmt.Println(l.Equation())
	}
	/*t := gx.Triangle{
		gx.Point{0, 0},
		gx.Point{0, 100},
		gx.Point{100, 0},
	}
	
	points := []gx.Point{
		gx.Point{},
		gx.Point{.1, .1},
		gx.Point{-1, -1},
		gx.Point{1, 1},
		gx.Point{101, 1},
		gx.Point{100, 1},
		gx.Point{50, 1},
	}
	
	for _, p := range points {
		fmt.Println(p, t.PointIsIn(p))
	}*/
}




// https://leetcode.com/problems/max-points-on-a-line/
// FIXME: Failing on first test below
package main

import (
	"fmt"
	"math"
	"regexp"
)

func lineEquation(x1, y1, x2, y2 int) (float64, float64) {
	m := float64(y1-y2) / float64(x1-x2)
	b := float64(y1) - m*float64(x1)

	return m, b
}

func onLine(x, y int, m, b float64) bool {
	yl := m*float64(x) + b
	return math.Round(yl) == float64(y)
}

func numPoints(points [][]int, m, b float64) int {
	count := 0
	for _, pt := range points {
		if onLine(pt[0], pt[1], m, b) {
			count++
		}
	}
	return count
}

func sameX(x int, points [][]int) int {
	count := 0
	for _, pt := range points {
		if pt[0] == x {
			count++
		}
	}
	return count
}

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}
	max := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			m, b := lineEquation(x1, y1, x2, y2)
			var count int
			if math.IsInf(m, 0) || math.IsNaN(m) {
				count = sameX(x1, points)
			} else {
				count = numPoints(points, m, b)
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}

func parse(s string) [][]int {
	var points [][]int
	re, err := regexp.Compile(`\[-?\d+,-?\d+\]`)
	if err != nil {
		panic(err)
	}

	for _, pt := range re.FindAllString(s, -1) {
		var x, y int
		if _, err := fmt.Sscanf(pt, "[%d,%d]", &x, &y); err != nil {
			panic(err)
		}
		points = append(points, []int{x, y})
	}
	return points
}

func main() {
	cases := []struct {
		points [][]int
		count  int
	}{
		{parse("[[560,248],[0,16],[30,250],[950,187],[630,277],[950,187],[-212,-268],[-287,-222],[53,37],[-280,-100],[-1,-14],[-5,4],[-35,-387],[-95,11],[-70,-13],[-700,-274],[-95,11],[-2,-33],[3,62],[-4,-47],[106,98],[-7,-65],[-8,-71],[-8,-147],[5,5],[-5,-90],[-420,-158],[-420,-158],[-350,-129],[-475,-53],[-4,-47],[-380,-37],[0,-24],[35,299],[-8,-71],[-2,-6],[8,25],[6,13],[-106,-146],[53,37],[-7,-128],[-5,-1],[-318,-390],[-15,-191],[-665,-85],[318,342],[7,138],[-570,-69],[-9,-4],[0,-9],[1,-7],[-51,23],[4,1],[-7,5],[-280,-100],[700,306],[0,-23],[-7,-4],[-246,-184],[350,161],[-424,-512],[35,299],[0,-24],[-140,-42],[-760,-101],[-9,-9],[140,74],[-285,-21],[-350,-129],[-6,9],[-630,-245],[700,306],[1,-17],[0,16],[-70,-13],[1,24],[-328,-260],[-34,26],[7,-5],[-371,-451],[-570,-69],[0,27],[-7,-65],[-9,-166],[-475,-53],[-68,20],[210,103],[700,306],[7,-6],[-3,-52],[-106,-146],[560,248],[10,6],[6,119],[0,2],[-41,6],[7,19],[30,250]]"), 22},
		{parse("[[3,1],[12,3],[3,1],[-6,-1]]"), 4},
		{parse("[[1,1],[1,1],[1,1]]"), 3},
		{parse("[[1,1],[2,2],[3,3]]"), 3},
		{parse("[[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]"), 4},
		{parse("[[4,0],[4,-1],[4,5]]"), 3},
	}
	for _, tc := range cases {
		count := maxPoints(tc.points)
		if count == tc.count {
			continue
		}
		fmt.Printf("%v expected %d, got %d\n", tc.points, tc.count, count)
	}
}

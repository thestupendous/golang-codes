package main

import ( "fmt" ; "math")

type Point struct {
	x int
	y int
}

func main(){
	fmt.Println(Point{0,0})

	PointA := Point{1,2}
	PointB := Point{1,90}

	dist := math.Sqrt( float64( (PointA.x-PointB.x)*(PointA.x-PointB.x) + (PointA.y-PointB.y)*(PointA.y-PointB.y) ) )
	fmt.Printf(" %v , %T \n",dist,dist)

	p := &PointB				//pointer to structure
	p.x = 1e7
	fmt.Println(PointB.x)
}
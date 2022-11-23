package placekey

// A polygon is created on a 2D place by a set of contours
// It can contain holes, and self-intersecting
type Polygon struct {
	points []*Point
}

func NewPolygon(points []*Point) *Polygon {
	return &Polygon{points: points}
}

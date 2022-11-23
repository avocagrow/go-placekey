package placekey

const (
	// The earth's radius is about 6371km
	EARTH_RADIUS = 6371
)

// Point represents a physical latitude and longitude coordinate pair
type Point struct {
	lat float64
	lng float64
}

// NewPoint returns a new Point populated with the given latitude and
// longitude coordinates
func NewPoint(lat, lng float64) *Point {
	return &Point{lat: lat, lng: lng}
}

// Latitude returns a Point's latitude
func (p *Point) Latitude() float64 {
	return p.lat
}

// Longitude returns a Point's longitude
func (p *Point) Longitude() float64 {
	return p.lng
}

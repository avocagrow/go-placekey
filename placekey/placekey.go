// Functionality for converting between Placekeys, geos (latitude/longitude), or H3 indicies.
// This package also includes additional utilities related to Placekeys.
package placekey

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/uber/h3-go/v4"
)

const (
	RESOLUTION        = 10
	BASE_RESOLUTION   = 12
	ALPHABET          = "23456789bcdfghjkmnpqrstvwxyz"
	ALPHABETLENGTH    = len(ALPHABET)
	CODELENGTH        = 9
	TUPLELENGTH       = 3
	PADDING_CHAR      = 'a'
	REPLACEMENT_CHARS = "eu"
)

var (
	replacementMap = map[string]string{
		"prn":   "pre",
		"f4nny": "f4nne",
		"tw4t":  "tw4e",
		"ngr":   "ngu", // 'u' avoids introducing 'gey'
		"dck":   "dce",
		"vjn":   "vju", // 'u' avoids introducing 'jew'
		"fck":   "fce",
		"pns":   "pne",
		"sht":   "she",
		"kkk":   "kke",
		"fgt":   "fgu", // 'u' avoids introducing 'gey'
		"dyk":   "dye",
		"bch":   "bce",
	}
	headerBits             = h3.LatLngToCell(h3.NewLatLng(0.0, 0.0), RESOLUTION)
	baseCellShift          = math.Pow(2, (3 * 15)) // this will increment the base cell value by 1
	unusedResolutionFiller = math.Pow(2, (3 * (15 - BASE_RESOLUTION)))
	firstTupleRegex        = fmt.Sprintf("[%s%s%c]", ALPHABET, REPLACEMENT_CHARS, PADDING_CHAR)
	tupleRegex             = fmt.Sprintf("%s%s", ALPHABET, REPLACEMENT_CHARS)
	whereRegex             = createRegex(fmt.Sprintf("^%s$", strings.Join([]string{firstTupleRegex, tupleRegex, tupleRegex}, "-")))
	whatRegex              = createRegex(fmt.Sprintf("^[%s]{3,}(-[%s]{3,})?$", ALPHABET, ALPHABET))
	prefixDistanceMap      = map[int]float32{
		0: 2.004e7,
		1: 2.004e7,
		2: 2.777e6,
		3: 1.065e6,
		4: 1.524e5,
		5: 2.177e4,
		6: 8227.0,
		7: 1176.0,
		8: 444.3,
		9: 63.47,
	}
	headerInt = getHeaderInt()
)

type Placekey string

func createRegex(pattern string) *regexp.Regexp {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return r
}

// returns the an integer corresponding to the header of an H3 integer
func getHeaderInt() int64 {
	headerInt := 0
	return int64(headerInt)
}

// GeoToPlacekey converts latitude and longitude coordinates into a Placekey string
func GeoToPlacekey(lat, lng float64) Placekey {
	return ""
}

// PlacekeyToGeo converts a placekey into latitude and longitude coordinates.
func PlacekeyToGeo(placekey Placekey) (lat, lng float64) {
	return 0.0, 0.0
}

// PlacekeyToH3 converts a Placekey string into an H3 Hexidecimal string
func PlacekeyToH3(p Placekey) h3.Cell {
	return h3.LatLngToCell(h3.NewLatLng(0.0, 0.0), RESOLUTION)
}

// H3ToPlacekey converts an H3 hexidecimal into a Placekey string
func H3ToPlacekey(s string) Placekey {
	return ""
}

// H3CellToPlacekey converts an H3 cell integer into a Placekey string
func H3CellToPlacekey(cell h3.Cell) Placekey {
	return ""
}

// PlacekeyToH3Cell converts a Placekey string into an H3 cell integer
func PlacekeyToH3Cell(p Placekey) h3.Cell {
	return h3.Cell(0)
}

// GetNeighboringPlacekeys returns an unordered slice of Placekeys whose grid distance
// is `<= dist from the given Placekey. In this context, grid distance referes to the number of H3 cells
// between two H3 cells, so that neighboring cells have a distance of 1.
// Neighbors of Neighbors have a distance of 2, etc.
func GetNeighboringPlacekeys(p Placekey, dist int) []Placekey {
	return nil
}

// PlacekeyToHexBoundary convertrs a Placekey string into the coordinates of
func PlacekeyToHexBoundary(p Placekey) [][2]float64 {
	return nil
}

// PlacekeyToPolygon returns the Polygon boundary shape for a Placekeuy string
func PlacekeyToPolygon(p Placekey) *Polygon {
	return &Polygon{points: nil}
}

// Converts a Placekey string into a Well-Known Text string for the corresponding
// hexagon. Coordinates are (longitude, latitude)
func PlacekeyToWKT(p Placekey) string {
	return ""
}

// PlacekeyToGeoJSON converts a Placekey string into a GeoJSON map.
// GeoJSON uses (longitude, latitude) points, and the first and last points are identical.
func PlacekeyToGeoJSON(p Placekey) []byte {
	return nil
}

// PolygonToPlacekeys converts a Polygon shape into a set of Placekey strings
// if inclTouching is true, then Placekeys whose hexagon boundary only touches that of the
// input polygon are included in the set of boundary Placekeys
func PolygonToPlacekeys(p *Polygon, inclTouching bool) []Placekey {
	return nil
}

// WKTToPlacekeys converts a WKT description of a polygon into a set of Placekey strings
func WKTToPlacekeys(w string) []Placekey {
	return nil
}

// GeoJSONToPlacekeys converts a GeoJSON description of a polygon into a set of Placekeys
func GeoJSONToPlacekeys(g []byte, inclTouching bool) []Placekey {
	return nil
}

// ValidatePlacekey determines if the given Placekey string is valid,
// includes checking for valid encoding of location
func ValidatePlacekey(p Placekey) bool {
	return false
}

// PlacekeyDistance calculates the distance between two Placekeys in meters
func PlacekeyDistance(p1, p2 Placekey) float64 {
	return 0.0
}

// parsePlacekey splits a Placekey string into what and where parts
func parsePlacekey(p Placekey) (what, where string) {
	if strings.Contains(string(p), "@") {
		res := strings.Split(string(p), "@")
		what, where = res[0], res[1]
	} else {
		what, where = "", string(p)
	}
	return what, where
}

// validateWhereClause determines if the given where string is valid
func validateWhereClause(w string) bool {
	return false
}

func geoDistance(p1, p2 Point) float64 {
	return 0.0
}

func encodeH3Cell(c h3.Cell) Placekey {
	return ""
}

func encodeShortH3Cell(c h3.Cell) string {
	return ""
}

func decodeToH3Cell(where string) h3.Cell {
	return 0
}

func decodeString(s string) int64 {
	return 0
}

func stripEncoding(s string) string {
	return ""
}

func shortenH3Cell(c h3.Cell) int64 {
	return 0
}

func lengthenH3Cell(c int64) h3.Cell {
	unshiftedInt := c << (3 * (15 - BASE_RESOLUTION))
	rebuiltCell := int64(headerBits) + int64(unusedResolutionFiller) - BASE_CELL_SHIFT + unshiftedInt
	return rebuiltCell
}

func cleanString(s string) string {
	// Replacement should be in order
	for k, v := range replacementMap {
		if strings.Contains(s, k) {
			s = strings.ReplaceAll(s, k, v)
		}
	}
	return s
}

func dirtyString(s string) string {
	// Replacement should be reversed
	for k, v := range replacementMap {
		if strings.Contains(s, v) {
			s = strings.ReplaceAll(s, v, k)
		}
	}
	return s
}

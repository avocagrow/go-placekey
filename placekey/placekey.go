package placekey

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/uber/h3-go/v4"
)

const (
	RESOLUTION       = 10
	BASERESOLUTION   = 12
	ALPHABET         = "23456789bcdfghjkmnpqrstvwxyz"
	ALPHABETLENGTH   = len(ALPHABET)
	CODELENGTH       = 9
	TUPLELENGTH      = 3
	PADDINGCHAR      = 'a'
	REPLACEMENTCHARS = "eu"
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
	unusedResolutionFiller = math.Pow(2, (3 * (15 - BASERESOLUTION)))
	firstTupleRegex        = fmt.Sprintf("[%s%s%c]", ALPHABET, REPLACEMENTCHARS, PADDINGCHAR)
	tupleRegex             = fmt.Sprintf("%s%s", ALPHABET, REPLACEMENTCHARS)
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
)

type Placekey string

func createRegex(pattern string) *regexp.Regexp {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return r
}

func getHeaderInt() int {
	return 0
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

func PlacekeyToPolygon(placekey) *Polygon {
	return &Polygon{points: nil}
}

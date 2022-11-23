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
)

func createRegex(pattern string) *regexp.Regexp {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return r
}

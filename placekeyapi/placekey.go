package placekeyapi

import (
	"net/http"
	"net/url"
	"sync"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

const (
	defaultBaseURL = "https://api.placekey.io"
	defaultVersion = "v1"
	userAgent      = "go-placekey"
)

type Client struct {
	// client is the HTTP client used to communicate with the API.
	client *retryablehttp.Client

	// baseURL is the main URL for API requests.
	baseURL *url.URL

	// token is the apiToken used for authentication with the PlaceKey API.
	token string

	// tokenL protects the token from concurrent read/write access.
	tokenL sync.RWMutex

	// Services used for communicating with different parts of the Placekey API.
}

// Key is the Placekey API key used globally in the binding
var APIKey string

type APIResponse struct {
	// Header contains a map of all HTTP header keys
	Header http.Header

	RawJSON []byte

	// QueryID contains a string that identifies the Placekey request if one was provided
	QueryID string

	// Status is a status code and message. e.g. "200 OK"
	Status string

	// StatusCode is a status code as integer. e.g. 200
	StatusCode int
}

func newAPIResponse(res *http.Response, resBody []byte) *APIResponse {
	return &APIResponse{
		Header:  res.Header,
		RawJSON: resBody,
	}
}

type PlacekeyResponse struct {
	QueryID       string `json:"query_id omitempty"`
	Placekey      string `json:"placekey"`
	CorelogicClip string `json:"corelogic_clip omitempty"`
}

// Query is the structure used to define the placekey request
type Query struct {
	QueryID        string  `json:"query_id omitempty"`
	LocationName   string  `json:"location_name omitempty"`
	Lat            float64 `json:"latitude omitempty"`
	Long           float64 `json:"longitude omitempty"`
	StreetAddr     string  `json:"street_address omitempty"`
	Region         string  `json:"region omitempty"`
	PostalCode     string  `json:"postal_code omitempty"`
	IsoCountryCode string  `json:"iso_country_code omitempty"`
}

// QueryOptions are the various options that can be passed to the
// Placekey API.
//
// If StrictNameMatch is false, all LocationName matches will be fuzzy matched
// If StrictNameMatch is true, all LocationName matches are exact case-insensitve
//
// InclCorelogicClip is
// For more information on CLIP please refer to https://corelogic.com/data-solutions/property-data-solutions/clip/
type QueryOptions struct {
	StrictNameMatch   bool `json:"strict_name_match omitempty"`
	InclCorelogicClip bool `json:"include_corelogic_clip omitempty"`
}

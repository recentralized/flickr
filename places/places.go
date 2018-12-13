package places

import (
	"gopkg.in/masci/flickr.v2"
)

type PlacesGetInfoResponse struct {
	flickr.BasicResponse
	Place struct {
		PlaceID     string  `xml:"place_id,attr"`
		WoeID       string  `xml:"woeid,attr"`
		Lat         float64 `xml:"latitude,attr"`
		Long        float64 `xml:"longitude,attr"`
		PlaceType   string  `xml:"place_type,attr"`
		PlaceTypeID int     `xml:"place_type_id,attr"`
		Timezone    string  `xml:"timezone,attr"`
		Name        string  `xml:"name,attr"`
		WoeName     string  `xml:"woe_name,attr"`

		Neighborhood Place `xml:"neighbourhood"`
		Locality     Place `xml:"locality"`
		County       Place `xml:"county"`
		Region       Place `xml:"region"`
		Country      Place `xml:"country"`
		Continent    Place `xml:"continent"`

		// TODO: shapedata
	} `xml:"place"`
}

type Place struct {
	PlaceID string  `xml:"place_id,attr"`
	WoeID   string  `xml:"woeid,attr"`
	Lat     float64 `xml:"latitude,attr"`
	Long    float64 `xml:"longitude,attr"`
	URL     string  `xml:"place_url"`
	Name    string  `xml:",chardata"`
}

// Get info about a place. Pass either placeID or woeID.
func GetInfo(client *flickr.FlickrClient, placeID string, woeID string) (*PlacesGetInfoResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.places.getInfo")
	if placeID != "" {
		client.Args.Set("place_id", placeID)
	}
	if woeID != "" {
		client.Args.Set("woe_id", woeID)
	}
	client.OAuthSign()

	response := &PlacesGetInfoResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

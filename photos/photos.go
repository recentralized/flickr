package photos

import (
	"gopkg.in/masci/flickr.v2"
)

type PhotoInfo struct {
	Id           string `xml:"id,attr"`
	Secret       string `xml:"secret,attr"`
	Server       string `xml:"server,attr"`
	Farm         string `xml:"farm,attr"`
	DateUploaded string `xml:"dateuploaded,attr"`
	IsFavorite   bool   `xml:"isfavorite,attr"`
	License      string `xml:"license,attr"`
	// NOTE: one less than safety level set on upload (ie, here 0 = safe, 1 = moderate, 2 = restricted)
	//       while on upload, 1 = safe, 2 = moderate, 3 = restricted
	SafetyLevel    int    `xml:"safety_level,attr"`
	Rotation       int    `xml:"rotation,attr"`
	OriginalSecret string `xml:"originalsecret,attr"`
	OriginalFormat string `xml:"originalformat,attr"`
	Views          int    `xml:"views,attr"`
	Media          string `xml:"media,attr"`
	Title          string `xml:"title"`
	Description    string `xml:"description"`
	Owner          struct {
		NSID     string `xml:"nsid,attr"`
		Username string `xml:"username,attr"`
		Realname string `xml:"realname,attr"`
		Location string `xml:"location,attr"`
	} `xml:"owner"`
	Visibility struct {
		IsPublic bool `xml:"ispublic,attr"`
		IsFriend bool `xml:"isfriend,attr"`
		IsFamily bool `xml:"isfamily,attr"`
	} `xml:"visibility"`
	Dates struct {
		Posted           string `xml:"posted,attr"`
		Taken            string `xml:"taken,attr"`
		TakenGranularity string `xml:"takengranularity,attr"`
		TakenUnknown     string `xml:"takenunknown,attr"`
		LastUpdate       string `xml:"lastupdate,attr"`
	} `xml:"dates"`
	Permissions struct {
		PermComment string `xml:"permcomment,attr"`
		PermAdMeta  string `xml:"permadmeta,attr"`
	} `xml:"permissions"`
	Editability struct {
		CanComment string `xml:"cancomment,attr"`
		CanAddMeta string `xml:"canaddmeta,attr"`
	} `xml:"editability"`
	PublicEditability struct {
		CanComment string `xml:"cancomment,attr"`
		CanAddMeta string `xml:"canaddmeta,attr"`
	} `xml:"publiceditability"`
	Usage struct {
		CanDownload string `xml:"candownload,attr"`
		CanBlog     string `xml:"canblog,attr"`
		CanPrint    string `xml:"canprint,attr"`
		CanShare    string `xml:"canshare,attr"`
	} `xml:"usage"`
	Comments int `xml:"comments"`
	People   struct {
		HasPeople bool `xml:"haspeople,attr"`
	} `xml:"people"`
	NoteList struct {
		Notes []struct {
			ID       string `xml:"id,attr"`
			NSID     string `xml:"author,attr"`
			Username string `xml:"authorname,attr"`
			X        int    `xml:"x,attr"`
			Y        int    `xml:"y,attr"`
			W        int    `xml:"w,attr"`
			H        int    `xml:"h,attr"`
		} `xml:"note"`
	} `xml:"notes"`
	TagList struct {
		Tags []struct {
			ID         string `xml:"id,attr"`
			NSID       string `xml:"author,attr"`
			Username   string `xml:"authorname,attr"`
			MachineTag string `xml:"machine_tag,attr"`
			Raw        string `xml:"raw,attr"`
			Tag        string `xml:",chardata"`
		} `xml:"tag"`
	} `xml:"tags"`
	URLList struct {
		URLs []struct {
			Type  string `xml:"type,attr"`
			Value string `xml:",chardata"`
		} `xml:"url"`
	} `xml:"urls"`
}

type PhotoInfoResponse struct {
	flickr.BasicResponse
	Photo PhotoInfo `xml:"photo"`
}

// Delete a photo from Flickr
// This method requires authentication with 'delete' permission.
func Delete(client *flickr.FlickrClient, id string) (*flickr.BasicResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.delete")
	client.Args.Set("photo_id", id)
	client.OAuthSign()

	response := &flickr.BasicResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

// Get information about a Flickr photo
func GetInfo(client *flickr.FlickrClient, id string, secret string) (*PhotoInfoResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.getInfo")
	client.Args.Set("photo_id", id)
	if secret != "" {
		client.Args.Set("secret", secret)
	}
	client.OAuthSign()

	response := &PhotoInfoResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

// Set date posted and date taken on a Flickr photo
// datePosted and dateTaken are optional and may be set to ""
func SetDates(client *flickr.FlickrClient, id string, datePosted string, dateTaken string) (*flickr.BasicResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.setDates")
	client.Args.Set("photo_id", id)
	if datePosted != "" {
		client.Args.Set("date_posted", datePosted)
	}
	if dateTaken != "" {
		client.Args.Set("date_taken", dateTaken)
	}
	client.OAuthSign()

	response := &flickr.BasicResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

type FavoritesList struct {
	Page      int `xml:"page,attr"`
	Pages     int `xml:"pages,attr"`
	PerPage   int `xml:"perpage,attr"`
	Total     int `xml:"total,attr"`
	Favorites []struct {
		NSID     string `xml:"nsid,attr"`
		Username string `xml:"username,attr"`
		Date     string `xml:"favedate,attr"`
	} `xml:"person"`
}

type PhotoFavoritesResponse struct {
	flickr.BasicResponse
	Photo FavoritesList `xml:"photo"`
}

type GetFavoritesOptionalArgs struct {
	PerPage int // 0 to ignore
	Page    int // 0 to ignore
}

// Get favorites for a photo
func GetFavorites(client *flickr.FlickrClient, id string, opts GetFavoritesOptionalArgs) (*PhotoFavoritesResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.getFavorites")
	client.Args.Set("photo_id", id)
	client.OAuthSign()

	response := &PhotoFavoritesResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

type CommentsList struct {
	PhotoID  string `json:"photo_id,attr"`
	Comments []struct {
		ID        string `xml:"id,attr"`
		NSID      string `xml:"author,attr"`
		Username  string `xml:"authorname,attr"`
		Date      string `xml:"datecreate,attr"`
		Permalink string `xml:"permalink,attr"`
		Text      string `xml:",chardata"`
	} `xml:"comment"`
}

type PhotoCommentsResponse struct {
	flickr.BasicResponse
	Comments CommentsList `xml:"comments"`
}

// Get comments for a photo
func GetComments(client *flickr.FlickrClient, id string) (*PhotoCommentsResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.getComments")
	client.Args.Set("photo_id", id)
	client.OAuthSign()

	response := &PhotoCommentsResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

type PhotoContextsResponse struct {
	flickr.BasicResponse
	Sets []struct {
		ID    string `xml:"id,attr"`
		Title string `xml:"title,attr"`
	} `xml:"set"`
	Pools []struct {
		ID    string `xml:"id,attr"`
		Title string `xml:"title,attr"`
		URL   string `xml:"url,attr"`
	} `xml:"pool"`
}

// Get all of the contexts that a photo appears in
func GetAllContexts(client *flickr.FlickrClient, id string) (*PhotoContextsResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.getAllContexts")
	client.Args.Set("photo_id", id)
	client.OAuthSign()

	response := &PhotoContextsResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

type PhotoPeopleGetListResponse struct {
	flickr.BasicResponse
	People struct {
		Persons []struct {
			NSID        string `xml:"nsid,attr"`
			Username    string `xml:"username,attr"`
			AddedByNSID string `xml:"added_by,attr"`
		} `xml:"person"`
	} `xml:"people"`
}

// Get all of the contexts that a photo appears in
func PeopleGetList(client *flickr.FlickrClient, id string) (*PhotoPeopleGetListResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.people.getList")
	client.Args.Set("photo_id", id)
	client.OAuthSign()

	response := &PhotoPeopleGetListResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

type PhotoGetSizesResponse struct {
	flickr.BasicResponse
	SizeInfo struct {
		Sizes []struct {
			Label  string `xml:"label,attr"`
			Width  int    `xml:"width,attr"`
			Height int    `xml:"height,attr"`
			Source string `xml:"source,attr"`
			URL    string `xml:"url,attr"`
			Media  string `xml:"media,attr"`
		} `xml:"size"`
	} `xml:"sizes"`
}

// Get all of the sizes for a photo
func GetSizes(client *flickr.FlickrClient, id string) (*PhotoGetSizesResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.getSizes")
	client.Args.Set("photo_id", id)
	client.OAuthSign()

	response := &PhotoGetSizesResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

type PhotosGeoGetLocationResponse struct {
	flickr.BasicResponse
	Photo struct {
		Location struct {
			Lat          float64 `xml:"latitude,attr"`
			Long         float64 `xml:"longitude,attr"`
			Accuracy     int     `xml:"accuracy,attr"`
			Context      int     `xml:"context,attr"`
			PlaceID      string  `xml:"place_id,attr"`
			WoeID        string  `xml:"woeid,attr"`
			Neighborhood Place   `xml:"neighbourhood"`
			Locality     Place   `xml:"locality"`
			County       Place   `xml:"county"`
			Region       Place   `xml:"region"`
			Country      Place   `xml:"country"`
			Continent    Place   `xml:"continent"`
		} `xml:"location"`
	} `xml:"photo"`
}

type Place struct {
	PlaceID string `xml:"place_id,attr"`
	WoeID   string `xml:"woeid,attr"`
	Name    string `xml:",chardata"`
}

// Get geo location information about a photo.
func GeoGetLocation(client *flickr.FlickrClient, id string) (*PhotosGeoGetLocationResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.geo.getLocation")
	client.Args.Set("photo_id", id)
	client.OAuthSign()

	response := &PhotosGeoGetLocationResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

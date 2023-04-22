package unsplash

// struct for response from Unsplash API
type UnsplashResponse struct {
	URL         struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}
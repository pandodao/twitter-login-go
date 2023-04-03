package twitter

type (
	// twttier user
	User struct {
		ID    uint32 `json:"id"`
		IDStr string `json:"id_str"`
		Name  string `json:"name"`
		// screen name
		ScreenName           string `json:"screen_name"`
		ProfileImageURLHttps string `json:"profile_image_url_https"`
	}
)

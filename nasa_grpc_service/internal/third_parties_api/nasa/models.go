package nasa

type RandomSpaseImageResponse struct {
	Copyright      string `json:"copyright" omitempty:"true"`
	Date           string `json:"date" omitempty:"true"`
	Explanation    string `json:"explanation" omitempty:"true"`
	Hdurl          string `json:"hdurl" omitempty:"true"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version" omitempty:"true"`
	Title          string `json:"title" omitempty:"true"`
	Url            string `json:"url" omitempty:"true"`
	Data           []byte `json:"-"`
}

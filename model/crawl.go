package model

type CrawlRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type HeadingCount struct {
	H1 int `json:"h1"`
	H2 int `json:"h2"`
	H3 int `json:"h3"`
	H4 int `json:"h4"`
	H5 int `json:"h5"`
	H6 int `json:"h6"`
}

type CrawlResponse struct {
	Url               string       `json:"url"`
	HTMLVersion       string       `json:"html_version"`
	Title             string       `json:"title"`
	Headings          HeadingCount `json:"headings"`
	InternalLinks     int          `json:"internal_links"`
	ExternalLinks     int          `json:"external_links"`
	InaccessibleLinks int          `json:"inaccessible_links"`
	HasLoginForm      bool         `json:"has_login_form"`
	Checkbox          bool         `json:"checkbox"`
	Status            string       `json:"status"`
}

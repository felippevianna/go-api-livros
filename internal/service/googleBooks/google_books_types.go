package service

type GoogleBooksResponse struct {
	Items []struct {
		ID         string `json:"id"`
		VolumeInfo struct {
			Title               string   `json:"title"`
			Authors             []string `json:"authors"`
			Description         string   `json:"description"`
			PageCount           int      `json:"pageCount"`
			Categories          []string `json:"categories"`
			ImageLinks          struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
		} `json:"volumeInfo"`
	} `json:"items"`
}
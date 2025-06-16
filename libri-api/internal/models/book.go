package models

type Book struct {
  ID            string   `json:"id"`
  Title         string   `json:"title"`
  Authors       []string `json:"authors"`
  Publisher     string   `json:"publisher,omitempty"`
  PublishedDate string   `json:"published_date,omitempty"`
  Description   string   `json:"description,omitempty"`
  InfoLink      string   `json:"info_link"`
}


type googleVolumeInfo struct {
  Title         string   `json:"title"`
  Authors       []string `json:"authors"`
  Publisher     string   `json:"publisher"`
  PublishedDate string   `json:"publishedDate"`
  Description   string   `json:"description"`
  InfoLink      string   `json:"infoLink"`
}

type googleBookItem struct {
  ID         string           `json:"id"`
  VolumeInfo googleVolumeInfo `json:"volumeInfo"`
}

type googleBooksResponse struct {
  Items []googleBookItem `json:"items"`
}

package service

import (
  "context"
  "encoding/json"
  "fmt"
  "net/http"
  "time"

  "github.com/spf13/viper"
  "github.com/brunnossanttos/libri/internal/models"
)

type BookService struct {
  httpClient *http.Client
  apiKey     string
}

type googleBooksResponse struct {
  Items []googleBookItem `json:"items"`
}

type googleBookItem struct {
  ID         string           `json:"id"`
  VolumeInfo googleVolumeInfo `json:"volumeInfo"`
}

type googleVolumeInfo struct {
  Title         string   `json:"title"`
  Authors       []string `json:"authors"`
  Publisher     string   `json:"publisher"`
  PublishedDate string   `json:"publishedDate"`
  Description   string   `json:"description"`
  InfoLink      string   `json:"infoLink"`
}

func NewBookService() *BookService {
  return &BookService{
    httpClient: &http.Client{Timeout: 5 * time.Second},
    apiKey:     viper.GetString("GOOGLE_BOOKS_API_KEY"),
  }
}

func (s *BookService) SearchBooks(ctx context.Context, q string) ([]models.Book, error) {
  if s.apiKey == "" {
    return nil, fmt.Errorf("GOOGLE_BOOKS_API_KEY não configurada")
  }

  url := fmt.Sprintf(
    "https://www.googleapis.com/books/v1/volumes?q=%s&key=%s",
    q, s.apiKey,
  )

  req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
  resp, err := s.httpClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  var gr googleBooksResponse
  if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
    return nil, err
  }

  var books []models.Book
  for _, item := range gr.Items {
    vi := item.VolumeInfo
    books = append(books, models.Book{
      ID:            item.ID,
      Title:         vi.Title,
      Authors:       vi.Authors,
      Publisher:     vi.Publisher,
      PublishedDate: vi.PublishedDate,
      Description:   vi.Description,
      InfoLink:      vi.InfoLink,
    })
  }
  return books, nil
}

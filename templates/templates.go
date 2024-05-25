package templates

import (
	"net/url"
	"tadeobennett/celebrateease/pkg/models"
)

type TemplateData struct {
	CSRFToken       string
	Flash           string
	ErrorsFromForm  map[string]string //map[key]//value
	FormData        url.Values
	IsAuthenticated bool
}

type UserTemplate struct {
	Users           []*models.User
	User            *models.User
	Flash           string
	ErrorsFromForm  map[string]string
	FormData        url.Values
	IsAuthenticated bool
	CSRFToken       string
}
type EventTemplate struct {
	Events          []*models.Event
	Event           *models.Event
	Flash           string
	ErrorsFromForm  map[string]string
	FormData        url.Values
	IsAuthenticated bool
	CSRFToken       string
}
type PageTemplate struct {
	Pages           []*models.Page
	Page            *models.Page
	Flash           string
	ErrorsFromForm  map[string]string
	FormData        url.Values
	IsAuthenticated bool
	CSRFToken       string
}
type CelebrantTemplate struct {
	Celebrants      []*models.Celebrant
	Celebrant       *models.Celebrant
	Flash           string
	ErrorsFromForm  map[string]string
	FormData        url.Values
	IsAuthenticated bool
	CSRFToken       string
}

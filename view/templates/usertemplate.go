package templates

import (
	"net/url"
	"tadeobennett/celebrateease/model"
)

type UserTemplate struct {
	Users          []*model.User
	User           *model.User
	Flash          string
	ErrorsFromForm map[string]string //map[key]//value
	FormData       url.Values        //asking go to get the values from the form
}
type EventTemplate struct {
	Events         []*model.Event
	Event          *model.Event
	Flash          string
	ErrorsFromForm map[string]string //map[key]//value
	FormData       url.Values        //asking go to get the values from the form
}
type PageTemplate struct {
	Pages         []*model.Page
	Page           *model.Page
	Flash          string
	ErrorsFromForm map[string]string //map[key]//value
	FormData       url.Values        //asking go to get the values from the form
}
type CelebrantTemplate struct {
	Celebrants         []*model.Celebrant
	Celebrant      *model.Celebrant
	Flash          string
	ErrorsFromForm map[string]string //map[key]//value
	FormData       url.Values        //asking go to get the values from the form
}

package templates

import (
	// "net/url"

	"tadeobennett/celebrateease/model"
)

type UserTemplate struct {
	Users         []*model.User
	User          *model.User
	Flash          string
	//ErrorsFromForm map[string]string //map[key]//value
	//FormData       url.Values        //asking go to get the values from the form
}

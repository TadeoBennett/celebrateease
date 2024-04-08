package shared

import "net/http"


//this interface is used to allow these functions to run and display errorlogs within the controllers and views
type ErrorHandler interface {
	ServerError(w http.ResponseWriter, err error)
	ClientError(w http.ResponseWriter, status int)
	NotFound(w http.ResponseWriter, err error)
	TemplateError(w http.ResponseWriter, err string)
}

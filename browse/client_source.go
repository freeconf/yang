package browse

import "net/http"

type ClientSource interface {
	GetHttpClient() *http.Client
}

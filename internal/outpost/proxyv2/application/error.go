package application

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ErrorPageData struct {
	Title       string
	Message     string
	ProxyPrefix string
}

func (a *Application) ErrorPage(rw http.ResponseWriter, r *http.Request, err string) {
	claims, _ := a.getClaims(r)
	data := ErrorPageData{
		Title:       "Bad Gateway",
		Message:     "Error proxying to upstream server",
		ProxyPrefix: "/goauthentik.io",
	}
	if claims != nil && len(err) > 0 {
		data.Message = err
	}
	er := a.errorTemplates.Execute(rw, data)
	if er != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}
}

// NewProxyErrorHandler creates a ProxyErrorHandler using the template given.
func (a *Application) newProxyErrorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(rw http.ResponseWriter, req *http.Request, proxyErr error) {
		log.WithError(proxyErr).Warning("Error proxying to upstream server")
		rw.WriteHeader(http.StatusBadGateway)
		a.ErrorPage(rw, req, fmt.Sprintf("Error proxying to upstream server: %s", proxyErr.Error()))
	}
}

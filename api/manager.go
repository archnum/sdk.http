/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import "net/http"

type (
	Manager interface {
		http.Handler
		Router() Router
	}

	implManager struct {
		router Router
	}
)

func New() *implManager {
	return &implManager{
		router: newRouter(&segment{}),
	}
}

func (impl *implManager) Router() Router {
	return impl.router
}

func (impl *implManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = newContext(w, r)
	w.WriteHeader(204) // TODO
}

/*
####### END ############################################################################################################
*/

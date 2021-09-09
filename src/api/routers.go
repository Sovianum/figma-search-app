package api

import (
	"encoding/json"
	"net/http"

	"github.com/davyzhang/agw"
	"github.com/gorilla/mux"
)

func NewAPI(
	tagEndpoints *TagEndpoints,
) *API {

	return &API{
		TagEndpoints: tagEndpoints,
	}
}

type API struct {
	TagEndpoints *TagEndpoints
}

func (api *API) NewRouter() *mux.Router {
	r := mux.NewRouter()

	api.post(r, "/files/{file_id}/tags/get", api.TagEndpoints.GetTags)
	api.post(r, "/files/{file_id}/tags/create", api.TagEndpoints.CreateTags)
	api.post(r, "/files/{file_id}/tags/remove", api.TagEndpoints.RemoveTags)

	return r
}

func (api *API) post(r *mux.Router, path string, f func(r *http.Request) (interface{}, error)) *mux.Route {
	return r.HandleFunc(path, api.wrapReturning(f)).
		Methods("POST")
}

func (api *API) wrapReturning(f func(r *http.Request) (interface{}, error)) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		iface, err := f(r)
		if err != nil {
			api.writeError(rw, err)
		} else {
			api.writeContent(rw, iface)
		}
	}
}

type errorResponse struct {
	Error string `json:"error"`
}

func (api *API) writeError(rw http.ResponseWriter, err error) {
	b, err := json.Marshal(errorResponse{Error: err.Error()})
	if err != nil {
		panic(err) // todo handle it somehow
	}

	typedRW := rw.(*agw.LPResponse)

	typedRW.WriteHeader(http.StatusInternalServerError) // TODO add error wrappers
	typedRW.WriteBody(b, false)
}

func (api *API) writeContent(rw http.ResponseWriter, content interface{}) {
	b, err := json.Marshal(content)
	if err != nil {
		api.writeError(rw, err)
		return
	}

	typedRW := rw.(*agw.LPResponse)
	typedRW.WriteBody(b, false)
}

package general

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

const (
	CmeDateFormat = "20060102"
)

func (api *API) GetArticles(w rest.ResponseWriter, r *rest.Request) {
	data, err := api.DB.GetAllArticles()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteJson(data)
}

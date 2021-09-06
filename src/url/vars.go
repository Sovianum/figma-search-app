package url

import (
	"errors"
	"net/http"

	"github.com/Sovianum/figma-search-app/src/domain/files/fileid"
	"github.com/gorilla/mux"
)

func FileIDFromRequest(r *http.Request) (fileid.ID, error) {
	vars := mux.Vars(r)
	fileIDStr, ok := vars["id"]

	if !ok {
		return fileid.ID{}, errors.New("file id not found")
	}

	return fileid.FromString(fileIDStr)
}

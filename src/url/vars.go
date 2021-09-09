package url

import (
	"net/http"

	"github.com/Sovianum/figma-search-app/src/domain/files/fileid"
	"github.com/gorilla/mux"
	"github.com/joomcode/errorx"
)

func FileIDFromRequest(r *http.Request) (fileid.ID, error) {
	vars := mux.Vars(r)
	fileIDStr, ok := vars["file_id"]

	if !ok {
		return fileid.ID{}, errorx.IllegalArgument.New("file id not present in path")
	}

	return fileid.FromString(fileIDStr)
}

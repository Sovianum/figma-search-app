package url

import (
	"net/http"

	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/gorilla/mux"
	"github.com/joomcode/errorx"
)

func ProjectIDFromRequest(r *http.Request) (projectid.ID, error) {
	vars := mux.Vars(r)
	projectIDStr, ok := vars["projectId"]

	if !ok {
		return projectid.ID{}, errorx.IllegalArgument.New("project id not present in path")
	}

	return projectid.FromString(projectIDStr)
}

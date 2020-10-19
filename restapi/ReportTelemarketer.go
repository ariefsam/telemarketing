package restapi

import (
	"net/http"

	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) ReportTelemarketer(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)
	_, err := api.authAdminOrResponseError(post, w)
	if err != nil {
		return
	}

	if post.FilterReportTelemarketer == nil {
		post.FilterReportTelemarketer = &entity.FilterReportTelemarketer{}
	}

	telemarketers, err := api.Usecase.ReportTelemarketer(*post.FilterReportTelemarketer, post.Limit)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}
	response := map[string]interface{}{
		"Telemarketers": telemarketers,
	}
	JSONView(w, response, http.StatusOK)
	return
}

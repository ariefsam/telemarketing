package restapi

import (
	"log"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) ImportCustomer(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()

	var token []string
	var ok bool
	if token, ok = keys["token"]; !ok {
		return
	}
	var source []string
	if source, ok = keys["source"]; !ok {
		return
	}

	post := RequestModel{
		Token: &token[0],
	}
	telemarketer, err := api.authAdminOrResponseError(post, w)
	if err != nil {
		log.Println(err)
		return
	}

	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		return
	}
	defer file.Close()
	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Println(err)
		return
	}
	rows := f.GetRows("Sheet1")
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		customer := entity.Customer{
			ID:               api.Usecase.GenerateID(),
			PhoneNumber:      row[1],
			Name:             row[0],
			TimestampCreated: api.Usecase.CurrentTimestamp(),
			CreatedBy:        telemarketer.ID,
			DataSource:       source[0],
		}
		go func() {
			err := api.Usecase.CreateCustomer(customer)
			if err != nil {
				log.Println(err)
			}
		}()
	}
}

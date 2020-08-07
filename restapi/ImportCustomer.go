package restapi

import (
	"log"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) ImportCustomer(w http.ResponseWriter, r *http.Request) {
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
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		customer := entity.Customer{
			PhoneNumber: row[1],
			Name:        row[0],
		}
		api.Usecase.SaveCustomer(customer)
	}
}

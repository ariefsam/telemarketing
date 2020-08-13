package restapi

func (api *RestAPI) parseInteger(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

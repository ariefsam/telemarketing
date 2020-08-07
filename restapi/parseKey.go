package restapi

import "errors"

func (api *RestAPI) parseKey(post map[string]interface{}, key string) (val interface{}, err error) {
	if _, ok := post[key]; !ok {
		err = errors.New(key + " is needed")
		return
	}
	if post[key] == nil {
		err = errors.New(key + " is needed")
		return
	}
	val = post[key]
	return
}

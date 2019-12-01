package services

import (
	"encoding/json"
	"net/http"
)

func HandlePostData(r http.Request) (map[string]interface{}, error) {
	postData := map[string]interface{}{}
	decodeErr := json.NewDecoder(r.Body).Decode(&postData)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return postData, nil
}

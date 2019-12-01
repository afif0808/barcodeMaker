package mocks

import (
	"barcodeMaker/model"
	"net/http"
)

func HandlePostDataMock(data map[string]interface{}, err error) model.HandlePostData {
	return func(http.Request) (map[string]interface{}, error) {
		return data, err
	}
}

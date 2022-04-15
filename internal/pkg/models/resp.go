package models

import "encoding/json"

type Response struct {
	DocumentationUrl string      `json:"documentation_url"`
	Message          string      `json:"message"`
	Data             interface{} `json:"data"`
	Status           int         `json:"status"`
}

func (r *Response) ToMap() Map {
	return Map{"documentation_url": r.DocumentationUrl, "message": r.Message, "data": r.Data, "status": r.Status}
}

func (r *Response) ToByteArray() ([]byte, error) {
	var returnStatus string
	if (200 <= r.Status) && (r.Status < 300) {
		returnStatus = "ok"
	} else {
		returnStatus = "error"
	}

	output, err := json.Marshal(Map{"documentation_url": r.DocumentationUrl, "message": r.Message, "data": r.Data, "status": returnStatus})
	if err != nil {
		return nil, err
	}
	return output, nil
}

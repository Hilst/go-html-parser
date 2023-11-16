package models

import "encoding/json"

// DATA RESP
type DataResponse struct {
	Ok    *DataContent
	Error error
}

func NewDataRespFree(ok map[string]any, err error) DataResponse {
	return DataResponse{&DataContent{"free", "free", ok}, err}
}

func NewDataResp(middle MiddleDataResp, err error) DataResponse {
	content, mappingErr := middle.dataContent()
	if err == nil && mappingErr != nil {
		err = mappingErr
	}
	return DataResponse{&content, err}
}

func NewDataRespError(err error) DataResponse {
	return DataResponse{nil, err}
}

type DataContent struct {
	Id         string
	LayoutRoot string
	Data       map[string]any
}

type MiddleDataResp struct {
	Id         string          `json:"name"`
	LayoutRoot string          `json:"layout"`
	Data       json.RawMessage `json:"data"`
}

func (m MiddleDataResp) dataContent() (DataContent, error) {
	var mappedJson map[string]any
	mappingError := json.Unmarshal(m.Data, &mappedJson)
	content := DataContent{
		m.Id,
		m.LayoutRoot,
		mappedJson,
	}
	return content, mappingError
}

// LAYOUT RESP
type LayoutResponse struct {
	Ok    *layoutContent
	Error error
}

func NewLayoutResp(tmpl string, name string) LayoutResponse {
	return LayoutResponse{&layoutContent{tmpl, name}, nil}
}

func NewLayoutRespError(err error) LayoutResponse {
	return LayoutResponse{nil, err}
}

type layoutContent struct {
	Tmpl string
	Name string
}

// TEST REQ
type TestRequest struct {
	LayoutHTML string `form:"html"`
	Data       string `form:"data"`
}

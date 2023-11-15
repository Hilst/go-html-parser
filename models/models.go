package models

type DataResponse struct {
	Ok    map[string]any
	Error error
}

func NewDataResp(ok map[string]any, err error) DataResponse {
	return DataResponse{ok, err}
}

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

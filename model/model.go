package model

// ResponsePayload is the basic response payload
type ResponsePayload struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// GhGqlPayload is the basic response payload
type GhGqlPayload struct {
	Data interface{} `json:"data"`
}

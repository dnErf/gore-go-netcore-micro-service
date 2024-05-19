package models

type DataObject struct {
	Message string `json:"message"`
}

type RequestPayload struct {
	Action string     `json:"action"`
	Data   DataObject `json:"data"`
}

type ResponsePaystub struct {
	IsFine  bool   `json:"is_fine"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

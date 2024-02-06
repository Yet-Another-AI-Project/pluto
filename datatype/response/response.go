package response

import "pluto/datatype/pluto_error"

const (
	STATUSOK    = "ok"
	STATUSERROR = "error"
)

type Reponse struct {
	Status string                  `json:"status" swaggertype:"string" enum:"ok, error"`
	Error  *pluto_error.PlutoError `json:"error"`
	Body   interface{}             `json:"body"`
}

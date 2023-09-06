package entities

type (
	JsonResponse struct {
		Message string      `json:"message"`
		Body    interface{} `json:"body"`
	}
)

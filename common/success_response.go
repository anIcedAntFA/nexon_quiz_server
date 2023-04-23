package common

type successResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Metadata   interface{} `json:"metadata,omitempty"`
}

func NewSuccessResponse(
	statusCode int,
	message string,
	data,
	metadata interface{},
) *successResponse {
	return &successResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Metadata:   metadata,
	}
}

func SimpleSuccessResponse(
	statusCode int,
	message string,
	data interface{},
) *successResponse {
	return NewSuccessResponse(
		statusCode,
		message,
		data,
		nil,
	)
}

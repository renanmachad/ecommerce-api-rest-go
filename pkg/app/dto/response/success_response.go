package response

type SuccessResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details"`
	Code    string   `json:"code"`
}

func NewSuccessResponse(message string, details []string, code string) *SuccessResponse {
	return &SuccessResponse{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func (s *SuccessResponse) getMessage() string {
	return s.Message
}

func (s *SuccessResponse) getCode() string {
	return s.Code
}

func (s *SuccessResponse) getDetails() []string {
	return s.Details
}

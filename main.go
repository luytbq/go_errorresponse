package go_errorresponse

import (
	"fmt"
)

type ErrorResponse struct {
	Name    string `json:"name,omitempty"`
	Details string `json:"details,omitempty"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprint(r.ToMap())
}

func NewErrorWithName(name string) ErrorResponse {
	return ErrorResponse{Name: name}
}

func (r ErrorResponse) WithDetails(details string) ErrorResponse {
	return ErrorResponse{Name: r.Name, Details: details}
}

func (r ErrorResponse) ToMap() map[string]any {
	return map[string]any{
		"name":    r.Name,
		"details": r.Details,
	}
}

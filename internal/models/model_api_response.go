package models

type APIResponse struct {
	Code int32 `json:"code,omitempty"`

	Type string `json:"type,omitempty"`

	Message string `json:"message,omitempty"`
}

// Code generated by goa v3.0.6, DO NOT EDIT.
//
// files HTTP client types
//
// Command:
// $ goa gen github.com/wild-mouse/go-playground/design/goa_files

package client

import (
	files "github.com/wild-mouse/go-playground/gen/files"
)

// AddRequestBody is the type of the "files" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	File *string `form:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "files" service.
func NewAddRequestBody(p *files.AddPayload) *AddRequestBody {
	body := &AddRequestBody{
		File: p.File,
	}
	return body
}
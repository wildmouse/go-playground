// Code generated by goa v3.0.6, DO NOT EDIT.
//
// files HTTP server types
//
// Command:
// $ goa gen github.com/wild-mouse/go-playground/design/goa_files

package server

import (
	files "github.com/wild-mouse/go-playground/gen/files"
)

// AddRequestBody is the type of the "files" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	File *string `form:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
}

// NewAddPayload builds a files service add endpoint payload.
func NewAddPayload(body *AddRequestBody) *files.AddPayload {
	v := &files.AddPayload{
		File: body.File,
	}
	return v
}

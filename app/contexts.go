// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application Contexts
//
// Command:
// $ goagen
// --design=go-playground/design
// --out=$(GOPATH)/src/go-playground
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// CreateBottleContext provides the bottle create action context.
type CreateBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *BottlePayload
}

// NewCreateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller create action.
func NewCreateBottleContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateBottleContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// ShowBottleContext provides the bottle show action context.
type ShowBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *BottlePayload
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(r *Bottle) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gophercon.goa.bottle")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

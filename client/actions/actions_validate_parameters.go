//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewActionsValidateParams creates a new ActionsValidateParams object
// with the default values initialized.
func NewActionsValidateParams() *ActionsValidateParams {
	var ()
	return &ActionsValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActionsValidateParamsWithTimeout creates a new ActionsValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActionsValidateParamsWithTimeout(timeout time.Duration) *ActionsValidateParams {
	var ()
	return &ActionsValidateParams{

		timeout: timeout,
	}
}

// NewActionsValidateParamsWithContext creates a new ActionsValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewActionsValidateParamsWithContext(ctx context.Context) *ActionsValidateParams {
	var ()
	return &ActionsValidateParams{

		Context: ctx,
	}
}

// NewActionsValidateParamsWithHTTPClient creates a new ActionsValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActionsValidateParamsWithHTTPClient(client *http.Client) *ActionsValidateParams {
	var ()
	return &ActionsValidateParams{
		HTTPClient: client,
	}
}

/*ActionsValidateParams contains all the parameters to send to the API endpoint
for the actions validate operation typically these are written to a http.Request
*/
type ActionsValidateParams struct {

	/*Body*/
	Body *models.Action

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the actions validate params
func (o *ActionsValidateParams) WithTimeout(timeout time.Duration) *ActionsValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the actions validate params
func (o *ActionsValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the actions validate params
func (o *ActionsValidateParams) WithContext(ctx context.Context) *ActionsValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the actions validate params
func (o *ActionsValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the actions validate params
func (o *ActionsValidateParams) WithHTTPClient(client *http.Client) *ActionsValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the actions validate params
func (o *ActionsValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the actions validate params
func (o *ActionsValidateParams) WithBody(body *models.Action) *ActionsValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the actions validate params
func (o *ActionsValidateParams) SetBody(body *models.Action) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ActionsValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
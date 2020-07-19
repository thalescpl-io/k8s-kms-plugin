// Code generated by go-swagger; DO NOT EDIT.

package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// SimpleenrollOKCode is the HTTP code returned for type SimpleenrollOK
const SimpleenrollOKCode int = 200

/*SimpleenrollOK successful operation

swagger:response simpleenrollOK
*/
type SimpleenrollOK struct {
	/*

	  Default: "base64"
	*/
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	/*

	  Default: "application/pkcs7-mime"
	*/
	ContentType string `json:"Content-Type"`

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSimpleenrollOK creates SimpleenrollOK with default headers values
func NewSimpleenrollOK() *SimpleenrollOK {

	var (
		// initialize headers with default values

		contentTransferEncodingDefault = string("base64")
		contentTypeDefault             = string("application/pkcs7-mime")
	)

	return &SimpleenrollOK{

		ContentTransferEncoding: contentTransferEncodingDefault,

		ContentType: contentTypeDefault,
	}
}

// WithContentTransferEncoding adds the contentTransferEncoding to the simpleenroll o k response
func (o *SimpleenrollOK) WithContentTransferEncoding(contentTransferEncoding string) *SimpleenrollOK {
	o.ContentTransferEncoding = contentTransferEncoding
	return o
}

// SetContentTransferEncoding sets the contentTransferEncoding to the simpleenroll o k response
func (o *SimpleenrollOK) SetContentTransferEncoding(contentTransferEncoding string) {
	o.ContentTransferEncoding = contentTransferEncoding
}

// WithContentType adds the contentType to the simpleenroll o k response
func (o *SimpleenrollOK) WithContentType(contentType string) *SimpleenrollOK {
	o.ContentType = contentType
	return o
}

// SetContentType sets the contentType to the simpleenroll o k response
func (o *SimpleenrollOK) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithPayload adds the payload to the simpleenroll o k response
func (o *SimpleenrollOK) WithPayload(payload string) *SimpleenrollOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the simpleenroll o k response
func (o *SimpleenrollOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SimpleenrollOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Transfer-Encoding

	contentTransferEncoding := o.ContentTransferEncoding
	if contentTransferEncoding != "" {
		rw.Header().Set("Content-Transfer-Encoding", contentTransferEncoding)
	}

	// response header Content-Type

	contentType := o.ContentType
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SimpleenrollBadRequestCode is the HTTP code returned for type SimpleenrollBadRequest
const SimpleenrollBadRequestCode int = 400

/*SimpleenrollBadRequest invalid request

swagger:response simpleenrollBadRequest
*/
type SimpleenrollBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSimpleenrollBadRequest creates SimpleenrollBadRequest with default headers values
func NewSimpleenrollBadRequest() *SimpleenrollBadRequest {

	return &SimpleenrollBadRequest{}
}

// WithPayload adds the payload to the simpleenroll bad request response
func (o *SimpleenrollBadRequest) WithPayload(payload string) *SimpleenrollBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the simpleenroll bad request response
func (o *SimpleenrollBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SimpleenrollBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SimpleenrollUnauthorizedCode is the HTTP code returned for type SimpleenrollUnauthorized
const SimpleenrollUnauthorizedCode int = 401

/*SimpleenrollUnauthorized unauthorized

swagger:response simpleenrollUnauthorized
*/
type SimpleenrollUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSimpleenrollUnauthorized creates SimpleenrollUnauthorized with default headers values
func NewSimpleenrollUnauthorized() *SimpleenrollUnauthorized {

	return &SimpleenrollUnauthorized{}
}

// WithPayload adds the payload to the simpleenroll unauthorized response
func (o *SimpleenrollUnauthorized) WithPayload(payload string) *SimpleenrollUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the simpleenroll unauthorized response
func (o *SimpleenrollUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SimpleenrollUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SimpleenrollInternalServerErrorCode is the HTTP code returned for type SimpleenrollInternalServerError
const SimpleenrollInternalServerErrorCode int = 500

/*SimpleenrollInternalServerError something went wrong

swagger:response simpleenrollInternalServerError
*/
type SimpleenrollInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSimpleenrollInternalServerError creates SimpleenrollInternalServerError with default headers values
func NewSimpleenrollInternalServerError() *SimpleenrollInternalServerError {

	return &SimpleenrollInternalServerError{}
}

// WithPayload adds the payload to the simpleenroll internal server error response
func (o *SimpleenrollInternalServerError) WithPayload(payload string) *SimpleenrollInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the simpleenroll internal server error response
func (o *SimpleenrollInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SimpleenrollInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	models "github.com/guzenok/go-memkvdb-example/examples/service/models"
)

// GetValuesKeyOKCode is the HTTP code returned for type GetValuesKeyOK
const GetValuesKeyOKCode int = 200

/*GetValuesKeyOK Value by key

swagger:response getValuesKeyOK
*/
type GetValuesKeyOK struct {

	/*
	  In: Body
	*/
	Payload models.Value `json:"body,omitempty"`
}

// NewGetValuesKeyOK creates GetValuesKeyOK with default headers values
func NewGetValuesKeyOK() *GetValuesKeyOK {

	return &GetValuesKeyOK{}
}

// WithPayload adds the payload to the get values key o k response
func (o *GetValuesKeyOK) WithPayload(payload models.Value) *GetValuesKeyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get values key o k response
func (o *GetValuesKeyOK) SetPayload(payload models.Value) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetValuesKeyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetValuesKeyNotFoundCode is the HTTP code returned for type GetValuesKeyNotFound
const GetValuesKeyNotFoundCode int = 404

/*GetValuesKeyNotFound Key not found

swagger:response getValuesKeyNotFound
*/
type GetValuesKeyNotFound struct {
}

// NewGetValuesKeyNotFound creates GetValuesKeyNotFound with default headers values
func NewGetValuesKeyNotFound() *GetValuesKeyNotFound {

	return &GetValuesKeyNotFound{}
}

// WriteResponse to the client
func (o *GetValuesKeyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*GetValuesKeyDefault Unexpected error

swagger:response getValuesKeyDefault
*/
type GetValuesKeyDefault struct {
	_statusCode int
}

// NewGetValuesKeyDefault creates GetValuesKeyDefault with default headers values
func NewGetValuesKeyDefault(code int) *GetValuesKeyDefault {
	if code <= 0 {
		code = 500
	}

	return &GetValuesKeyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get values key default response
func (o *GetValuesKeyDefault) WithStatusCode(code int) *GetValuesKeyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get values key default response
func (o *GetValuesKeyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *GetValuesKeyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
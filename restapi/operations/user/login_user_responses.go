// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
)

// LoginUserOKCode is the HTTP code returned for type LoginUserOK
const LoginUserOKCode int = 200

/*
LoginUserOK successful operation

swagger:response loginUserOK
*/
type LoginUserOK struct {
	/*date in UTC when token expires

	 */
	XExpiresAfter strfmt.DateTime `json:"X-Expires-After"`
	/*calls per hour allowed by the user

	 */
	XRateLimit int32 `json:"X-Rate-Limit"`

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewLoginUserOK creates LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {

	return &LoginUserOK{}
}

// WithXExpiresAfter adds the xExpiresAfter to the login user o k response
func (o *LoginUserOK) WithXExpiresAfter(xExpiresAfter strfmt.DateTime) *LoginUserOK {
	o.XExpiresAfter = xExpiresAfter
	return o
}

// SetXExpiresAfter sets the xExpiresAfter to the login user o k response
func (o *LoginUserOK) SetXExpiresAfter(xExpiresAfter strfmt.DateTime) {
	o.XExpiresAfter = xExpiresAfter
}

// WithXRateLimit adds the xRateLimit to the login user o k response
func (o *LoginUserOK) WithXRateLimit(xRateLimit int32) *LoginUserOK {
	o.XRateLimit = xRateLimit
	return o
}

// SetXRateLimit sets the xRateLimit to the login user o k response
func (o *LoginUserOK) SetXRateLimit(xRateLimit int32) {
	o.XRateLimit = xRateLimit
}

// WithPayload adds the payload to the login user o k response
func (o *LoginUserOK) WithPayload(payload string) *LoginUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user o k response
func (o *LoginUserOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Expires-After

	xExpiresAfter := o.XExpiresAfter.String()
	if xExpiresAfter != "" {
		rw.Header().Set("X-Expires-After", xExpiresAfter)
	}

	// response header X-Rate-Limit

	xRateLimit := swag.FormatInt32(o.XRateLimit)
	if xRateLimit != "" {
		rw.Header().Set("X-Rate-Limit", xRateLimit)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// LoginUserBadRequestCode is the HTTP code returned for type LoginUserBadRequest
const LoginUserBadRequestCode int = 400

/*
LoginUserBadRequest Invalid username/password supplied

swagger:response loginUserBadRequest
*/
type LoginUserBadRequest struct {
}

// NewLoginUserBadRequest creates LoginUserBadRequest with default headers values
func NewLoginUserBadRequest() *LoginUserBadRequest {

	return &LoginUserBadRequest{}
}

// WriteResponse to the client
func (o *LoginUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

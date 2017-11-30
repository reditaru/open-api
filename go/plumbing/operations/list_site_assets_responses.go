// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// ListSiteAssetsReader is a Reader for the ListSiteAssets structure.
type ListSiteAssetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSiteAssetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSiteAssetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSiteAssetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSiteAssetsOK creates a ListSiteAssetsOK with default headers values
func NewListSiteAssetsOK() *ListSiteAssetsOK {
	return &ListSiteAssetsOK{}
}

/*ListSiteAssetsOK handles this case with default header values.

OK
*/
type ListSiteAssetsOK struct {
	Payload []*models.Asset
}

func (o *ListSiteAssetsOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/assets][%d] listSiteAssetsOK  %+v", 200, o.Payload)
}

func (o *ListSiteAssetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSiteAssetsDefault creates a ListSiteAssetsDefault with default headers values
func NewListSiteAssetsDefault(code int) *ListSiteAssetsDefault {
	return &ListSiteAssetsDefault{
		_statusCode: code,
	}
}

/*ListSiteAssetsDefault handles this case with default header values.

error
*/
type ListSiteAssetsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list site assets default response
func (o *ListSiteAssetsDefault) Code() int {
	return o._statusCode
}

func (o *ListSiteAssetsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/assets][%d] listSiteAssets default  %+v", o._statusCode, o.Payload)
}

func (o *ListSiteAssetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

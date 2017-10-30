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

// CreateSiteAssetReader is a Reader for the CreateSiteAsset structure.
type CreateSiteAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSiteAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSiteAssetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateSiteAssetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSiteAssetCreated creates a CreateSiteAssetCreated with default headers values
func NewCreateSiteAssetCreated() *CreateSiteAssetCreated {
	return &CreateSiteAssetCreated{}
}

/*CreateSiteAssetCreated handles this case with default header values.

Created
*/
type CreateSiteAssetCreated struct {
	Payload *models.AssetSignature
}

func (o *CreateSiteAssetCreated) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/assets][%d] createSiteAssetCreated  %+v", 201, o.Payload)
}

func (o *CreateSiteAssetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetSignature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteAssetDefault creates a CreateSiteAssetDefault with default headers values
func NewCreateSiteAssetDefault(code int) *CreateSiteAssetDefault {
	return &CreateSiteAssetDefault{
		_statusCode: code,
	}
}

/*CreateSiteAssetDefault handles this case with default header values.

error
*/
type CreateSiteAssetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create site asset default response
func (o *CreateSiteAssetDefault) Code() int {
	return o._statusCode
}

func (o *CreateSiteAssetDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/assets][%d] createSiteAsset default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSiteAssetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
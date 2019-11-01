// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VPNGateway v p n gateway
// swagger:model VPNGateway
type VPNGateway struct {

	// Collection of references to VPN connections
	Connections []*VPNGatewayConnectionReference `json:"connections"`

	// The date and time that this VPN gateway was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The VPN gateway's CRN
	Crn string `json:"crn,omitempty"`

	// The VPN gateway's canonical URL
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The VPN gateway's unique identifier
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The name given to this VPN gateway
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// public ip
	PublicIP *VPNGatewayPublicIP `json:"public_ip,omitempty"`

	// resource group
	ResourceGroup *VPNGatewayResourceGroup `json:"resource_group,omitempty"`

	// The status of the VPN gateway
	// Enum: [available failed pending]
	Status string `json:"status,omitempty"`

	// subnet
	Subnet *VPNGatewaySubnet `json:"subnet,omitempty"`
}

// Validate validates this v p n gateway
func (m *VPNGateway) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VPNGateway) validateConnections(formats strfmt.Registry) error {

	if swag.IsZero(m.Connections) { // not required
		return nil
	}

	for i := 0; i < len(m.Connections); i++ {
		if swag.IsZero(m.Connections[i]) { // not required
			continue
		}

		if m.Connections[i] != nil {
			if err := m.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VPNGateway) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VPNGateway) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *VPNGateway) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VPNGateway) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *VPNGateway) validatePublicIP(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicIP) { // not required
		return nil
	}

	if m.PublicIP != nil {
		if err := m.PublicIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_ip")
			}
			return err
		}
	}

	return nil
}

func (m *VPNGateway) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

var vPNGatewayTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["available","failed","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vPNGatewayTypeStatusPropEnum = append(vPNGatewayTypeStatusPropEnum, v)
	}
}

const (

	// VPNGatewayStatusAvailable captures enum value "available"
	VPNGatewayStatusAvailable string = "available"

	// VPNGatewayStatusFailed captures enum value "failed"
	VPNGatewayStatusFailed string = "failed"

	// VPNGatewayStatusPending captures enum value "pending"
	VPNGatewayStatusPending string = "pending"
)

// prop value enum
func (m *VPNGateway) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vPNGatewayTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VPNGateway) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *VPNGateway) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VPNGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNGateway) UnmarshalBinary(b []byte) error {
	var res VPNGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VPNGatewayPublicIP v p n gateway public IP
// swagger:model VPNGatewayPublicIP
type VPNGatewayPublicIP struct {

	// The IP address assigned to this VPN gateway
	Address string `json:"address,omitempty"`
}

// Validate validates this v p n gateway public IP
func (m *VPNGatewayPublicIP) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VPNGatewayPublicIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNGatewayPublicIP) UnmarshalBinary(b []byte) error {
	var res VPNGatewayPublicIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VPNGatewayResourceGroup ResourceGroupIdentity
// swagger:model VPNGatewayResourceGroup
type VPNGatewayResourceGroup struct {

	// The unique identifier for this resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this v p n gateway resource group
func (m *VPNGatewayResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VPNGatewayResourceGroup) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("resource_group"+"."+"id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VPNGatewayResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNGatewayResourceGroup) UnmarshalBinary(b []byte) error {
	var res VPNGatewayResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VPNGatewaySubnet SubnetReference
// swagger:model VPNGatewaySubnet
type VPNGatewaySubnet struct {

	// The CRN for this subnet
	Crn string `json:"crn,omitempty"`

	// The unique identifier for this subnet
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this subnet
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this v p n gateway subnet
func (m *VPNGatewaySubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VPNGatewaySubnet) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("subnet"+"."+"id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VPNGatewaySubnet) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("subnet"+"."+"name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VPNGatewaySubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNGatewaySubnet) UnmarshalBinary(b []byte) error {
	var res VPNGatewaySubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
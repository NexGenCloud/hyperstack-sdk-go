/*
Infrahub-API

Leverage the Infrahub API and Hyperstack platform to easily create, manage, and scale powerful GPU virtual machines and their associated resources.   Access this SDK to automate the deployment of your workloads and streamline your infrastructure management.  To contribute, please raise an issue with a bug report, feature request, feedback, or general inquiry.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
)

// checks if the GetPermissionsResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPermissionsResponseModel{}

// GetPermissionsResponseModel struct for GetPermissionsResponseModel
type GetPermissionsResponseModel struct {
	Message     *string            `json:"message,omitempty"`
	Permissions []PermissionFields `json:"permissions,omitempty"`
	Status      *bool              `json:"status,omitempty"`
}

// NewGetPermissionsResponseModel instantiates a new GetPermissionsResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPermissionsResponseModel() *GetPermissionsResponseModel {
	this := GetPermissionsResponseModel{}
	return &this
}

// NewGetPermissionsResponseModelWithDefaults instantiates a new GetPermissionsResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPermissionsResponseModelWithDefaults() *GetPermissionsResponseModel {
	this := GetPermissionsResponseModel{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetPermissionsResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionsResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetPermissionsResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetPermissionsResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GetPermissionsResponseModel) GetPermissions() []PermissionFields {
	if o == nil || IsNil(o.Permissions) {
		var ret []PermissionFields
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionsResponseModel) GetPermissionsOk() ([]PermissionFields, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GetPermissionsResponseModel) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []PermissionFields and assigns it to the Permissions field.
func (o *GetPermissionsResponseModel) SetPermissions(v []PermissionFields) {
	o.Permissions = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetPermissionsResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionsResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetPermissionsResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *GetPermissionsResponseModel) SetStatus(v bool) {
	o.Status = &v
}

func (o GetPermissionsResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPermissionsResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGetPermissionsResponseModel struct {
	value *GetPermissionsResponseModel
	isSet bool
}

func (v NullableGetPermissionsResponseModel) Get() *GetPermissionsResponseModel {
	return v.value
}

func (v *NullableGetPermissionsResponseModel) Set(val *GetPermissionsResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPermissionsResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPermissionsResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPermissionsResponseModel(val *GetPermissionsResponseModel) *NullableGetPermissionsResponseModel {
	return &NullableGetPermissionsResponseModel{value: val, isSet: true}
}

func (v NullableGetPermissionsResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPermissionsResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

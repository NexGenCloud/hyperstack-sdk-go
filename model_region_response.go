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

// checks if the RegionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionResponse{}

// RegionResponse struct for RegionResponse
type RegionResponse struct {
	Message *string       `json:"message,omitempty"`
	Region  *RegionFields `json:"region,omitempty"`
	Status  *bool         `json:"status,omitempty"`
}

// NewRegionResponse instantiates a new RegionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionResponse() *RegionResponse {
	this := RegionResponse{}
	return &this
}

// NewRegionResponseWithDefaults instantiates a new RegionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionResponseWithDefaults() *RegionResponse {
	this := RegionResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RegionResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RegionResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RegionResponse) SetMessage(v string) {
	o.Message = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RegionResponse) GetRegion() RegionFields {
	if o == nil || IsNil(o.Region) {
		var ret RegionFields
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionResponse) GetRegionOk() (*RegionFields, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RegionResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given RegionFields and assigns it to the Region field.
func (o *RegionResponse) SetRegion(v RegionFields) {
	o.Region = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RegionResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RegionResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *RegionResponse) SetStatus(v bool) {
	o.Status = &v
}

func (o RegionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableRegionResponse struct {
	value *RegionResponse
	isSet bool
}

func (v NullableRegionResponse) Get() *RegionResponse {
	return v.value
}

func (v *NullableRegionResponse) Set(val *RegionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionResponse(val *RegionResponse) *NullableRegionResponse {
	return &NullableRegionResponse{value: val, isSet: true}
}

func (v NullableRegionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

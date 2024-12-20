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

// checks if the NewStockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewStockResponse{}

// NewStockResponse struct for NewStockResponse
type NewStockResponse struct {
	Models    []NewModelResponse `json:"models,omitempty"`
	Region    *string            `json:"region,omitempty"`
	StockType *string            `json:"stock-type,omitempty"`
}

// NewNewStockResponse instantiates a new NewStockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewStockResponse() *NewStockResponse {
	this := NewStockResponse{}
	return &this
}

// NewNewStockResponseWithDefaults instantiates a new NewStockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewStockResponseWithDefaults() *NewStockResponse {
	this := NewStockResponse{}
	return &this
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *NewStockResponse) GetModels() []NewModelResponse {
	if o == nil || IsNil(o.Models) {
		var ret []NewModelResponse
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewStockResponse) GetModelsOk() ([]NewModelResponse, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *NewStockResponse) HasModels() bool {
	if o != nil && !IsNil(o.Models) {
		return true
	}

	return false
}

// SetModels gets a reference to the given []NewModelResponse and assigns it to the Models field.
func (o *NewStockResponse) SetModels(v []NewModelResponse) {
	o.Models = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *NewStockResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewStockResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *NewStockResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *NewStockResponse) SetRegion(v string) {
	o.Region = &v
}

// GetStockType returns the StockType field value if set, zero value otherwise.
func (o *NewStockResponse) GetStockType() string {
	if o == nil || IsNil(o.StockType) {
		var ret string
		return ret
	}
	return *o.StockType
}

// GetStockTypeOk returns a tuple with the StockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewStockResponse) GetStockTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StockType) {
		return nil, false
	}
	return o.StockType, true
}

// HasStockType returns a boolean if a field has been set.
func (o *NewStockResponse) HasStockType() bool {
	if o != nil && !IsNil(o.StockType) {
		return true
	}

	return false
}

// SetStockType gets a reference to the given string and assigns it to the StockType field.
func (o *NewStockResponse) SetStockType(v string) {
	o.StockType = &v
}

func (o NewStockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewStockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Models) {
		toSerialize["models"] = o.Models
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.StockType) {
		toSerialize["stock-type"] = o.StockType
	}
	return toSerialize, nil
}

type NullableNewStockResponse struct {
	value *NewStockResponse
	isSet bool
}

func (v NullableNewStockResponse) Get() *NewStockResponse {
	return v.value
}

func (v *NullableNewStockResponse) Set(val *NewStockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewStockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewStockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewStockResponse(val *NewStockResponse) *NullableNewStockResponse {
	return &NullableNewStockResponse{value: val, isSet: true}
}

func (v NullableNewStockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewStockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

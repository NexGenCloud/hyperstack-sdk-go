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

// checks if the NewStockRetriveResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewStockRetriveResponse{}

// NewStockRetriveResponse struct for NewStockRetriveResponse
type NewStockRetriveResponse struct {
	Stocks []NewStockResponse `json:"stocks,omitempty"`
}

// NewNewStockRetriveResponse instantiates a new NewStockRetriveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewStockRetriveResponse() *NewStockRetriveResponse {
	this := NewStockRetriveResponse{}
	return &this
}

// NewNewStockRetriveResponseWithDefaults instantiates a new NewStockRetriveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewStockRetriveResponseWithDefaults() *NewStockRetriveResponse {
	this := NewStockRetriveResponse{}
	return &this
}

// GetStocks returns the Stocks field value if set, zero value otherwise.
func (o *NewStockRetriveResponse) GetStocks() []NewStockResponse {
	if o == nil || IsNil(o.Stocks) {
		var ret []NewStockResponse
		return ret
	}
	return o.Stocks
}

// GetStocksOk returns a tuple with the Stocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewStockRetriveResponse) GetStocksOk() ([]NewStockResponse, bool) {
	if o == nil || IsNil(o.Stocks) {
		return nil, false
	}
	return o.Stocks, true
}

// HasStocks returns a boolean if a field has been set.
func (o *NewStockRetriveResponse) HasStocks() bool {
	if o != nil && !IsNil(o.Stocks) {
		return true
	}

	return false
}

// SetStocks gets a reference to the given []NewStockResponse and assigns it to the Stocks field.
func (o *NewStockRetriveResponse) SetStocks(v []NewStockResponse) {
	o.Stocks = v
}

func (o NewStockRetriveResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewStockRetriveResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stocks) {
		toSerialize["stocks"] = o.Stocks
	}
	return toSerialize, nil
}

type NullableNewStockRetriveResponse struct {
	value *NewStockRetriveResponse
	isSet bool
}

func (v NullableNewStockRetriveResponse) Get() *NewStockRetriveResponse {
	return v.value
}

func (v *NullableNewStockRetriveResponse) Set(val *NewStockRetriveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewStockRetriveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewStockRetriveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewStockRetriveResponse(val *NewStockRetriveResponse) *NullableNewStockRetriveResponse {
	return &NullableNewStockRetriveResponse{value: val, isSet: true}
}

func (v NullableNewStockRetriveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewStockRetriveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

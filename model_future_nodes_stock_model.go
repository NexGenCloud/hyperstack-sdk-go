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

// checks if the FutureNodesStockModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FutureNodesStockModel{}

// FutureNodesStockModel struct for FutureNodesStockModel
type FutureNodesStockModel struct {
	FutureStocks []FutureNodeResponseModel `json:"future_stocks,omitempty"`
}

// NewFutureNodesStockModel instantiates a new FutureNodesStockModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFutureNodesStockModel() *FutureNodesStockModel {
	this := FutureNodesStockModel{}
	return &this
}

// NewFutureNodesStockModelWithDefaults instantiates a new FutureNodesStockModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFutureNodesStockModelWithDefaults() *FutureNodesStockModel {
	this := FutureNodesStockModel{}
	return &this
}

// GetFutureStocks returns the FutureStocks field value if set, zero value otherwise.
func (o *FutureNodesStockModel) GetFutureStocks() []FutureNodeResponseModel {
	if o == nil || IsNil(o.FutureStocks) {
		var ret []FutureNodeResponseModel
		return ret
	}
	return o.FutureStocks
}

// GetFutureStocksOk returns a tuple with the FutureStocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureNodesStockModel) GetFutureStocksOk() ([]FutureNodeResponseModel, bool) {
	if o == nil || IsNil(o.FutureStocks) {
		return nil, false
	}
	return o.FutureStocks, true
}

// HasFutureStocks returns a boolean if a field has been set.
func (o *FutureNodesStockModel) HasFutureStocks() bool {
	if o != nil && !IsNil(o.FutureStocks) {
		return true
	}

	return false
}

// SetFutureStocks gets a reference to the given []FutureNodeResponseModel and assigns it to the FutureStocks field.
func (o *FutureNodesStockModel) SetFutureStocks(v []FutureNodeResponseModel) {
	o.FutureStocks = v
}

func (o FutureNodesStockModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FutureNodesStockModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FutureStocks) {
		toSerialize["future_stocks"] = o.FutureStocks
	}
	return toSerialize, nil
}

type NullableFutureNodesStockModel struct {
	value *FutureNodesStockModel
	isSet bool
}

func (v NullableFutureNodesStockModel) Get() *FutureNodesStockModel {
	return v.value
}

func (v *NullableFutureNodesStockModel) Set(val *FutureNodesStockModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFutureNodesStockModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFutureNodesStockModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFutureNodesStockModel(val *FutureNodesStockModel) *NullableFutureNodesStockModel {
	return &NullableFutureNodesStockModel{value: val, isSet: true}
}

func (v NullableFutureNodesStockModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFutureNodesStockModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
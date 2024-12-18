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

// checks if the NodeStockResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeStockResponseModel{}

// NodeStockResponseModel struct for NodeStockResponseModel
type NodeStockResponseModel struct {
	Stocks []NodeResponseModel `json:"stocks,omitempty"`
}

// NewNodeStockResponseModel instantiates a new NodeStockResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeStockResponseModel() *NodeStockResponseModel {
	this := NodeStockResponseModel{}
	return &this
}

// NewNodeStockResponseModelWithDefaults instantiates a new NodeStockResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeStockResponseModelWithDefaults() *NodeStockResponseModel {
	this := NodeStockResponseModel{}
	return &this
}

// GetStocks returns the Stocks field value if set, zero value otherwise.
func (o *NodeStockResponseModel) GetStocks() []NodeResponseModel {
	if o == nil || IsNil(o.Stocks) {
		var ret []NodeResponseModel
		return ret
	}
	return o.Stocks
}

// GetStocksOk returns a tuple with the Stocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeStockResponseModel) GetStocksOk() ([]NodeResponseModel, bool) {
	if o == nil || IsNil(o.Stocks) {
		return nil, false
	}
	return o.Stocks, true
}

// HasStocks returns a boolean if a field has been set.
func (o *NodeStockResponseModel) HasStocks() bool {
	if o != nil && !IsNil(o.Stocks) {
		return true
	}

	return false
}

// SetStocks gets a reference to the given []NodeResponseModel and assigns it to the Stocks field.
func (o *NodeStockResponseModel) SetStocks(v []NodeResponseModel) {
	o.Stocks = v
}

func (o NodeStockResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeStockResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stocks) {
		toSerialize["stocks"] = o.Stocks
	}
	return toSerialize, nil
}

type NullableNodeStockResponseModel struct {
	value *NodeStockResponseModel
	isSet bool
}

func (v NullableNodeStockResponseModel) Get() *NodeStockResponseModel {
	return v.value
}

func (v *NullableNodeStockResponseModel) Set(val *NodeStockResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeStockResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeStockResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeStockResponseModel(val *NodeStockResponseModel) *NullableNodeStockResponseModel {
	return &NullableNodeStockResponseModel{value: val, isSet: true}
}

func (v NullableNodeStockResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeStockResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

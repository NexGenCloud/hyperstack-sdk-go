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

// checks if the MetricItemFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricItemFields{}

// MetricItemFields struct for MetricItemFields
type MetricItemFields struct {
	Columns []string                   `json:"columns,omitempty"`
	Data    [][]map[string]interface{} `json:"data,omitempty"`
	Unit    *string                    `json:"unit,omitempty"`
}

// NewMetricItemFields instantiates a new MetricItemFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricItemFields() *MetricItemFields {
	this := MetricItemFields{}
	return &this
}

// NewMetricItemFieldsWithDefaults instantiates a new MetricItemFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricItemFieldsWithDefaults() *MetricItemFields {
	this := MetricItemFields{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *MetricItemFields) GetColumns() []string {
	if o == nil || IsNil(o.Columns) {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItemFields) GetColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *MetricItemFields) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *MetricItemFields) SetColumns(v []string) {
	o.Columns = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MetricItemFields) GetData() [][]map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret [][]map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItemFields) GetDataOk() ([][]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MetricItemFields) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given [][]map[string]interface{} and assigns it to the Data field.
func (o *MetricItemFields) SetData(v [][]map[string]interface{}) {
	o.Data = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricItemFields) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItemFields) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricItemFields) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MetricItemFields) SetUnit(v string) {
	o.Unit = &v
}

func (o MetricItemFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricItemFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableMetricItemFields struct {
	value *MetricItemFields
	isSet bool
}

func (v NullableMetricItemFields) Get() *MetricItemFields {
	return v.value
}

func (v *NullableMetricItemFields) Set(val *MetricItemFields) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricItemFields) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricItemFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricItemFields(val *MetricItemFields) *NullableMetricItemFields {
	return &NullableMetricItemFields{value: val, isSet: true}
}

func (v NullableMetricItemFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricItemFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the BillingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingResponse{}

// BillingResponse struct for BillingResponse
type BillingResponse struct {
	CalculatedBills []OrganizationObjectResponse `json:"calculated_bills,omitempty"`
	CalculationTime *CustomTime                  `json:"calculation_time,omitempty"`
	Type            *string                      `json:"type,omitempty"`
}

// NewBillingResponse instantiates a new BillingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingResponse() *BillingResponse {
	this := BillingResponse{}
	return &this
}

// NewBillingResponseWithDefaults instantiates a new BillingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingResponseWithDefaults() *BillingResponse {
	this := BillingResponse{}
	return &this
}

// GetCalculatedBills returns the CalculatedBills field value if set, zero value otherwise.
func (o *BillingResponse) GetCalculatedBills() []OrganizationObjectResponse {
	if o == nil || IsNil(o.CalculatedBills) {
		var ret []OrganizationObjectResponse
		return ret
	}
	return o.CalculatedBills
}

// GetCalculatedBillsOk returns a tuple with the CalculatedBills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponse) GetCalculatedBillsOk() ([]OrganizationObjectResponse, bool) {
	if o == nil || IsNil(o.CalculatedBills) {
		return nil, false
	}
	return o.CalculatedBills, true
}

// HasCalculatedBills returns a boolean if a field has been set.
func (o *BillingResponse) HasCalculatedBills() bool {
	if o != nil && !IsNil(o.CalculatedBills) {
		return true
	}

	return false
}

// SetCalculatedBills gets a reference to the given []OrganizationObjectResponse and assigns it to the CalculatedBills field.
func (o *BillingResponse) SetCalculatedBills(v []OrganizationObjectResponse) {
	o.CalculatedBills = v
}

// GetCalculationTime returns the CalculationTime field value if set, zero value otherwise.
func (o *BillingResponse) GetCalculationTime() CustomTime {
	if o == nil || IsNil(o.CalculationTime) {
		var ret CustomTime
		return ret
	}
	return *o.CalculationTime
}

// GetCalculationTimeOk returns a tuple with the CalculationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponse) GetCalculationTimeOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.CalculationTime) {
		return nil, false
	}
	return o.CalculationTime, true
}

// HasCalculationTime returns a boolean if a field has been set.
func (o *BillingResponse) HasCalculationTime() bool {
	if o != nil && !IsNil(o.CalculationTime) {
		return true
	}

	return false
}

// SetCalculationTime gets a reference to the given CustomTime and assigns it to the CalculationTime field.
func (o *BillingResponse) SetCalculationTime(v CustomTime) {
	o.CalculationTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BillingResponse) SetType(v string) {
	o.Type = &v
}

func (o BillingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CalculatedBills) {
		toSerialize["calculated_bills"] = o.CalculatedBills
	}
	if !IsNil(o.CalculationTime) {
		toSerialize["calculation_time"] = o.CalculationTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableBillingResponse struct {
	value *BillingResponse
	isSet bool
}

func (v NullableBillingResponse) Get() *BillingResponse {
	return v.value
}

func (v *NullableBillingResponse) Set(val *BillingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingResponse(val *BillingResponse) *NullableBillingResponse {
	return &NullableBillingResponse{value: val, isSet: true}
}

func (v NullableBillingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

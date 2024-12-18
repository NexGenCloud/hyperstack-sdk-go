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

// checks if the InfrahubResourceObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfrahubResourceObjectResponse{}

// InfrahubResourceObjectResponse struct for InfrahubResourceObjectResponse
type InfrahubResourceObjectResponse struct {
	ActualHostPrice   *float32                          `json:"actual_host_price,omitempty"`
	ActualPrice       *float32                          `json:"actual_price,omitempty"`
	ContractId        *int32                            `json:"contract_id,omitempty"`
	Host              *string                           `json:"host,omitempty"`
	HostPrice         *float32                          `json:"host_price,omitempty"`
	InfrahubId        *int32                            `json:"infrahub_id,omitempty"`
	Name              *string                           `json:"name,omitempty"`
	NexgenActualPrice *float32                          `json:"nexgen_actual_price,omitempty"`
	NexgenPrice       *float32                          `json:"nexgen_price,omitempty"`
	Price             *float32                          `json:"price,omitempty"`
	Resources         []PricebookResourceObjectResponse `json:"resources,omitempty"`
	Status            *string                           `json:"status,omitempty"`
	Type              *string                           `json:"type,omitempty"`
}

// NewInfrahubResourceObjectResponse instantiates a new InfrahubResourceObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfrahubResourceObjectResponse() *InfrahubResourceObjectResponse {
	this := InfrahubResourceObjectResponse{}
	return &this
}

// NewInfrahubResourceObjectResponseWithDefaults instantiates a new InfrahubResourceObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfrahubResourceObjectResponseWithDefaults() *InfrahubResourceObjectResponse {
	this := InfrahubResourceObjectResponse{}
	return &this
}

// GetActualHostPrice returns the ActualHostPrice field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetActualHostPrice() float32 {
	if o == nil || IsNil(o.ActualHostPrice) {
		var ret float32
		return ret
	}
	return *o.ActualHostPrice
}

// GetActualHostPriceOk returns a tuple with the ActualHostPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetActualHostPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.ActualHostPrice) {
		return nil, false
	}
	return o.ActualHostPrice, true
}

// HasActualHostPrice returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasActualHostPrice() bool {
	if o != nil && !IsNil(o.ActualHostPrice) {
		return true
	}

	return false
}

// SetActualHostPrice gets a reference to the given float32 and assigns it to the ActualHostPrice field.
func (o *InfrahubResourceObjectResponse) SetActualHostPrice(v float32) {
	o.ActualHostPrice = &v
}

// GetActualPrice returns the ActualPrice field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetActualPrice() float32 {
	if o == nil || IsNil(o.ActualPrice) {
		var ret float32
		return ret
	}
	return *o.ActualPrice
}

// GetActualPriceOk returns a tuple with the ActualPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetActualPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.ActualPrice) {
		return nil, false
	}
	return o.ActualPrice, true
}

// HasActualPrice returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasActualPrice() bool {
	if o != nil && !IsNil(o.ActualPrice) {
		return true
	}

	return false
}

// SetActualPrice gets a reference to the given float32 and assigns it to the ActualPrice field.
func (o *InfrahubResourceObjectResponse) SetActualPrice(v float32) {
	o.ActualPrice = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetContractId() int32 {
	if o == nil || IsNil(o.ContractId) {
		var ret int32
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetContractIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given int32 and assigns it to the ContractId field.
func (o *InfrahubResourceObjectResponse) SetContractId(v int32) {
	o.ContractId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *InfrahubResourceObjectResponse) SetHost(v string) {
	o.Host = &v
}

// GetHostPrice returns the HostPrice field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetHostPrice() float32 {
	if o == nil || IsNil(o.HostPrice) {
		var ret float32
		return ret
	}
	return *o.HostPrice
}

// GetHostPriceOk returns a tuple with the HostPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetHostPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.HostPrice) {
		return nil, false
	}
	return o.HostPrice, true
}

// HasHostPrice returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasHostPrice() bool {
	if o != nil && !IsNil(o.HostPrice) {
		return true
	}

	return false
}

// SetHostPrice gets a reference to the given float32 and assigns it to the HostPrice field.
func (o *InfrahubResourceObjectResponse) SetHostPrice(v float32) {
	o.HostPrice = &v
}

// GetInfrahubId returns the InfrahubId field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetInfrahubId() int32 {
	if o == nil || IsNil(o.InfrahubId) {
		var ret int32
		return ret
	}
	return *o.InfrahubId
}

// GetInfrahubIdOk returns a tuple with the InfrahubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetInfrahubIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InfrahubId) {
		return nil, false
	}
	return o.InfrahubId, true
}

// HasInfrahubId returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasInfrahubId() bool {
	if o != nil && !IsNil(o.InfrahubId) {
		return true
	}

	return false
}

// SetInfrahubId gets a reference to the given int32 and assigns it to the InfrahubId field.
func (o *InfrahubResourceObjectResponse) SetInfrahubId(v int32) {
	o.InfrahubId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InfrahubResourceObjectResponse) SetName(v string) {
	o.Name = &v
}

// GetNexgenActualPrice returns the NexgenActualPrice field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetNexgenActualPrice() float32 {
	if o == nil || IsNil(o.NexgenActualPrice) {
		var ret float32
		return ret
	}
	return *o.NexgenActualPrice
}

// GetNexgenActualPriceOk returns a tuple with the NexgenActualPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetNexgenActualPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.NexgenActualPrice) {
		return nil, false
	}
	return o.NexgenActualPrice, true
}

// HasNexgenActualPrice returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasNexgenActualPrice() bool {
	if o != nil && !IsNil(o.NexgenActualPrice) {
		return true
	}

	return false
}

// SetNexgenActualPrice gets a reference to the given float32 and assigns it to the NexgenActualPrice field.
func (o *InfrahubResourceObjectResponse) SetNexgenActualPrice(v float32) {
	o.NexgenActualPrice = &v
}

// GetNexgenPrice returns the NexgenPrice field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetNexgenPrice() float32 {
	if o == nil || IsNil(o.NexgenPrice) {
		var ret float32
		return ret
	}
	return *o.NexgenPrice
}

// GetNexgenPriceOk returns a tuple with the NexgenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetNexgenPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.NexgenPrice) {
		return nil, false
	}
	return o.NexgenPrice, true
}

// HasNexgenPrice returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasNexgenPrice() bool {
	if o != nil && !IsNil(o.NexgenPrice) {
		return true
	}

	return false
}

// SetNexgenPrice gets a reference to the given float32 and assigns it to the NexgenPrice field.
func (o *InfrahubResourceObjectResponse) SetNexgenPrice(v float32) {
	o.NexgenPrice = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *InfrahubResourceObjectResponse) SetPrice(v float32) {
	o.Price = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetResources() []PricebookResourceObjectResponse {
	if o == nil || IsNil(o.Resources) {
		var ret []PricebookResourceObjectResponse
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetResourcesOk() ([]PricebookResourceObjectResponse, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []PricebookResourceObjectResponse and assigns it to the Resources field.
func (o *InfrahubResourceObjectResponse) SetResources(v []PricebookResourceObjectResponse) {
	o.Resources = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InfrahubResourceObjectResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InfrahubResourceObjectResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrahubResourceObjectResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InfrahubResourceObjectResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InfrahubResourceObjectResponse) SetType(v string) {
	o.Type = &v
}

func (o InfrahubResourceObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfrahubResourceObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActualHostPrice) {
		toSerialize["actual_host_price"] = o.ActualHostPrice
	}
	if !IsNil(o.ActualPrice) {
		toSerialize["actual_price"] = o.ActualPrice
	}
	if !IsNil(o.ContractId) {
		toSerialize["contract_id"] = o.ContractId
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.HostPrice) {
		toSerialize["host_price"] = o.HostPrice
	}
	if !IsNil(o.InfrahubId) {
		toSerialize["infrahub_id"] = o.InfrahubId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NexgenActualPrice) {
		toSerialize["nexgen_actual_price"] = o.NexgenActualPrice
	}
	if !IsNil(o.NexgenPrice) {
		toSerialize["nexgen_price"] = o.NexgenPrice
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInfrahubResourceObjectResponse struct {
	value *InfrahubResourceObjectResponse
	isSet bool
}

func (v NullableInfrahubResourceObjectResponse) Get() *InfrahubResourceObjectResponse {
	return v.value
}

func (v *NullableInfrahubResourceObjectResponse) Set(val *InfrahubResourceObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInfrahubResourceObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInfrahubResourceObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfrahubResourceObjectResponse(val *InfrahubResourceObjectResponse) *NullableInfrahubResourceObjectResponse {
	return &NullableInfrahubResourceObjectResponse{value: val, isSet: true}
}

func (v NullableInfrahubResourceObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfrahubResourceObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

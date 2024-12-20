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

// checks if the UsersInfoFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersInfoFields{}

// UsersInfoFields struct for UsersInfoFields
type UsersInfoFields struct {
	BillingAddress1 *string     `json:"billing_address1,omitempty"`
	BillingAddress2 *string     `json:"billing_address2,omitempty"`
	Business        *bool       `json:"business,omitempty"`
	CompanyName     *string     `json:"company_name,omitempty"`
	Country         *string     `json:"country,omitempty"`
	CreatedAt       *CustomTime `json:"created_at,omitempty"`
	Email           *string     `json:"email,omitempty"`
	Id              *int32      `json:"id,omitempty"`
	Name            *string     `json:"name,omitempty"`
	OrganizationId  *int32      `json:"organization_id,omitempty"`
	Phone           *string     `json:"phone,omitempty"`
	State           *string     `json:"state,omitempty"`
	VatNumber       *string     `json:"vat_number,omitempty"`
	ZipCode         *string     `json:"zip_code,omitempty"`
}

// NewUsersInfoFields instantiates a new UsersInfoFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersInfoFields() *UsersInfoFields {
	this := UsersInfoFields{}
	return &this
}

// NewUsersInfoFieldsWithDefaults instantiates a new UsersInfoFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersInfoFieldsWithDefaults() *UsersInfoFields {
	this := UsersInfoFields{}
	return &this
}

// GetBillingAddress1 returns the BillingAddress1 field value if set, zero value otherwise.
func (o *UsersInfoFields) GetBillingAddress1() string {
	if o == nil || IsNil(o.BillingAddress1) {
		var ret string
		return ret
	}
	return *o.BillingAddress1
}

// GetBillingAddress1Ok returns a tuple with the BillingAddress1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetBillingAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.BillingAddress1) {
		return nil, false
	}
	return o.BillingAddress1, true
}

// HasBillingAddress1 returns a boolean if a field has been set.
func (o *UsersInfoFields) HasBillingAddress1() bool {
	if o != nil && !IsNil(o.BillingAddress1) {
		return true
	}

	return false
}

// SetBillingAddress1 gets a reference to the given string and assigns it to the BillingAddress1 field.
func (o *UsersInfoFields) SetBillingAddress1(v string) {
	o.BillingAddress1 = &v
}

// GetBillingAddress2 returns the BillingAddress2 field value if set, zero value otherwise.
func (o *UsersInfoFields) GetBillingAddress2() string {
	if o == nil || IsNil(o.BillingAddress2) {
		var ret string
		return ret
	}
	return *o.BillingAddress2
}

// GetBillingAddress2Ok returns a tuple with the BillingAddress2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetBillingAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.BillingAddress2) {
		return nil, false
	}
	return o.BillingAddress2, true
}

// HasBillingAddress2 returns a boolean if a field has been set.
func (o *UsersInfoFields) HasBillingAddress2() bool {
	if o != nil && !IsNil(o.BillingAddress2) {
		return true
	}

	return false
}

// SetBillingAddress2 gets a reference to the given string and assigns it to the BillingAddress2 field.
func (o *UsersInfoFields) SetBillingAddress2(v string) {
	o.BillingAddress2 = &v
}

// GetBusiness returns the Business field value if set, zero value otherwise.
func (o *UsersInfoFields) GetBusiness() bool {
	if o == nil || IsNil(o.Business) {
		var ret bool
		return ret
	}
	return *o.Business
}

// GetBusinessOk returns a tuple with the Business field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetBusinessOk() (*bool, bool) {
	if o == nil || IsNil(o.Business) {
		return nil, false
	}
	return o.Business, true
}

// HasBusiness returns a boolean if a field has been set.
func (o *UsersInfoFields) HasBusiness() bool {
	if o != nil && !IsNil(o.Business) {
		return true
	}

	return false
}

// SetBusiness gets a reference to the given bool and assigns it to the Business field.
func (o *UsersInfoFields) SetBusiness(v bool) {
	o.Business = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UsersInfoFields) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UsersInfoFields) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UsersInfoFields) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *UsersInfoFields) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *UsersInfoFields) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *UsersInfoFields) SetCountry(v string) {
	o.Country = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UsersInfoFields) GetCreatedAt() CustomTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret CustomTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetCreatedAtOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UsersInfoFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given CustomTime and assigns it to the CreatedAt field.
func (o *UsersInfoFields) SetCreatedAt(v CustomTime) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UsersInfoFields) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UsersInfoFields) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UsersInfoFields) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UsersInfoFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UsersInfoFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UsersInfoFields) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UsersInfoFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UsersInfoFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UsersInfoFields) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *UsersInfoFields) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *UsersInfoFields) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *UsersInfoFields) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UsersInfoFields) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UsersInfoFields) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UsersInfoFields) SetPhone(v string) {
	o.Phone = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UsersInfoFields) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UsersInfoFields) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UsersInfoFields) SetState(v string) {
	o.State = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *UsersInfoFields) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *UsersInfoFields) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *UsersInfoFields) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *UsersInfoFields) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersInfoFields) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *UsersInfoFields) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *UsersInfoFields) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o UsersInfoFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersInfoFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAddress1) {
		toSerialize["billing_address1"] = o.BillingAddress1
	}
	if !IsNil(o.BillingAddress2) {
		toSerialize["billing_address2"] = o.BillingAddress2
	}
	if !IsNil(o.Business) {
		toSerialize["business"] = o.Business
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vat_number"] = o.VatNumber
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zip_code"] = o.ZipCode
	}
	return toSerialize, nil
}

type NullableUsersInfoFields struct {
	value *UsersInfoFields
	isSet bool
}

func (v NullableUsersInfoFields) Get() *UsersInfoFields {
	return v.value
}

func (v *NullableUsersInfoFields) Set(val *UsersInfoFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersInfoFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersInfoFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersInfoFields(val *UsersInfoFields) *NullableUsersInfoFields {
	return &NullableUsersInfoFields{value: val, isSet: true}
}

func (v NullableUsersInfoFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersInfoFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

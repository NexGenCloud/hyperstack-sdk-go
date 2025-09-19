# UserInfoPostPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress1** | Pointer to **string** |  | [optional] 
**BillingAddress2** | Pointer to **string** |  | [optional] 
**Business** | **bool** |  | 
**CompanyName** | Pointer to **string** |  | [optional] 
**Country** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**ZipCode** | **string** |  | 

## Methods

### NewUserInfoPostPayload

`func NewUserInfoPostPayload(business bool, country string, zipCode string, ) *UserInfoPostPayload`

NewUserInfoPostPayload instantiates a new UserInfoPostPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoPostPayloadWithDefaults

`func NewUserInfoPostPayloadWithDefaults() *UserInfoPostPayload`

NewUserInfoPostPayloadWithDefaults instantiates a new UserInfoPostPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress1

`func (o *UserInfoPostPayload) GetBillingAddress1() string`

GetBillingAddress1 returns the BillingAddress1 field if non-nil, zero value otherwise.

### GetBillingAddress1Ok

`func (o *UserInfoPostPayload) GetBillingAddress1Ok() (*string, bool)`

GetBillingAddress1Ok returns a tuple with the BillingAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress1

`func (o *UserInfoPostPayload) SetBillingAddress1(v string)`

SetBillingAddress1 sets BillingAddress1 field to given value.

### HasBillingAddress1

`func (o *UserInfoPostPayload) HasBillingAddress1() bool`

HasBillingAddress1 returns a boolean if a field has been set.

### GetBillingAddress2

`func (o *UserInfoPostPayload) GetBillingAddress2() string`

GetBillingAddress2 returns the BillingAddress2 field if non-nil, zero value otherwise.

### GetBillingAddress2Ok

`func (o *UserInfoPostPayload) GetBillingAddress2Ok() (*string, bool)`

GetBillingAddress2Ok returns a tuple with the BillingAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress2

`func (o *UserInfoPostPayload) SetBillingAddress2(v string)`

SetBillingAddress2 sets BillingAddress2 field to given value.

### HasBillingAddress2

`func (o *UserInfoPostPayload) HasBillingAddress2() bool`

HasBillingAddress2 returns a boolean if a field has been set.

### GetBusiness

`func (o *UserInfoPostPayload) GetBusiness() bool`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *UserInfoPostPayload) GetBusinessOk() (*bool, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *UserInfoPostPayload) SetBusiness(v bool)`

SetBusiness sets Business field to given value.


### GetCompanyName

`func (o *UserInfoPostPayload) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UserInfoPostPayload) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UserInfoPostPayload) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UserInfoPostPayload) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *UserInfoPostPayload) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserInfoPostPayload) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserInfoPostPayload) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *UserInfoPostPayload) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInfoPostPayload) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInfoPostPayload) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInfoPostPayload) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserInfoPostPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserInfoPostPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserInfoPostPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserInfoPostPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *UserInfoPostPayload) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserInfoPostPayload) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserInfoPostPayload) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserInfoPostPayload) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *UserInfoPostPayload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserInfoPostPayload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserInfoPostPayload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserInfoPostPayload) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVatNumber

`func (o *UserInfoPostPayload) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UserInfoPostPayload) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UserInfoPostPayload) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UserInfoPostPayload) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetZipCode

`func (o *UserInfoPostPayload) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UserInfoPostPayload) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UserInfoPostPayload) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



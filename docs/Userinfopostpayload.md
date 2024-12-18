# Userinfopostpayload

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

### NewUserinfopostpayload

`func NewUserinfopostpayload(business bool, country string, zipCode string, ) *Userinfopostpayload`

NewUserinfopostpayload instantiates a new Userinfopostpayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserinfopostpayloadWithDefaults

`func NewUserinfopostpayloadWithDefaults() *Userinfopostpayload`

NewUserinfopostpayloadWithDefaults instantiates a new Userinfopostpayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress1

`func (o *Userinfopostpayload) GetBillingAddress1() string`

GetBillingAddress1 returns the BillingAddress1 field if non-nil, zero value otherwise.

### GetBillingAddress1Ok

`func (o *Userinfopostpayload) GetBillingAddress1Ok() (*string, bool)`

GetBillingAddress1Ok returns a tuple with the BillingAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress1

`func (o *Userinfopostpayload) SetBillingAddress1(v string)`

SetBillingAddress1 sets BillingAddress1 field to given value.

### HasBillingAddress1

`func (o *Userinfopostpayload) HasBillingAddress1() bool`

HasBillingAddress1 returns a boolean if a field has been set.

### GetBillingAddress2

`func (o *Userinfopostpayload) GetBillingAddress2() string`

GetBillingAddress2 returns the BillingAddress2 field if non-nil, zero value otherwise.

### GetBillingAddress2Ok

`func (o *Userinfopostpayload) GetBillingAddress2Ok() (*string, bool)`

GetBillingAddress2Ok returns a tuple with the BillingAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress2

`func (o *Userinfopostpayload) SetBillingAddress2(v string)`

SetBillingAddress2 sets BillingAddress2 field to given value.

### HasBillingAddress2

`func (o *Userinfopostpayload) HasBillingAddress2() bool`

HasBillingAddress2 returns a boolean if a field has been set.

### GetBusiness

`func (o *Userinfopostpayload) GetBusiness() bool`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *Userinfopostpayload) GetBusinessOk() (*bool, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *Userinfopostpayload) SetBusiness(v bool)`

SetBusiness sets Business field to given value.


### GetCompanyName

`func (o *Userinfopostpayload) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Userinfopostpayload) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Userinfopostpayload) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Userinfopostpayload) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *Userinfopostpayload) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Userinfopostpayload) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Userinfopostpayload) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *Userinfopostpayload) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Userinfopostpayload) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Userinfopostpayload) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Userinfopostpayload) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Userinfopostpayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Userinfopostpayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Userinfopostpayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Userinfopostpayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *Userinfopostpayload) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Userinfopostpayload) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Userinfopostpayload) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Userinfopostpayload) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *Userinfopostpayload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Userinfopostpayload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Userinfopostpayload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Userinfopostpayload) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVatNumber

`func (o *Userinfopostpayload) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Userinfopostpayload) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Userinfopostpayload) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Userinfopostpayload) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetZipCode

`func (o *Userinfopostpayload) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Userinfopostpayload) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Userinfopostpayload) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



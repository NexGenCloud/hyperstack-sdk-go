# UsersInfoFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress1** | Pointer to **string** |  | [optional] 
**BillingAddress2** | Pointer to **string** |  | [optional] 
**Business** | Pointer to **bool** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StripeUserId** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**ZipCode** | Pointer to **string** |  | [optional] 

## Methods

### NewUsersInfoFields

`func NewUsersInfoFields() *UsersInfoFields`

NewUsersInfoFields instantiates a new UsersInfoFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersInfoFieldsWithDefaults

`func NewUsersInfoFieldsWithDefaults() *UsersInfoFields`

NewUsersInfoFieldsWithDefaults instantiates a new UsersInfoFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress1

`func (o *UsersInfoFields) GetBillingAddress1() string`

GetBillingAddress1 returns the BillingAddress1 field if non-nil, zero value otherwise.

### GetBillingAddress1Ok

`func (o *UsersInfoFields) GetBillingAddress1Ok() (*string, bool)`

GetBillingAddress1Ok returns a tuple with the BillingAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress1

`func (o *UsersInfoFields) SetBillingAddress1(v string)`

SetBillingAddress1 sets BillingAddress1 field to given value.

### HasBillingAddress1

`func (o *UsersInfoFields) HasBillingAddress1() bool`

HasBillingAddress1 returns a boolean if a field has been set.

### GetBillingAddress2

`func (o *UsersInfoFields) GetBillingAddress2() string`

GetBillingAddress2 returns the BillingAddress2 field if non-nil, zero value otherwise.

### GetBillingAddress2Ok

`func (o *UsersInfoFields) GetBillingAddress2Ok() (*string, bool)`

GetBillingAddress2Ok returns a tuple with the BillingAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress2

`func (o *UsersInfoFields) SetBillingAddress2(v string)`

SetBillingAddress2 sets BillingAddress2 field to given value.

### HasBillingAddress2

`func (o *UsersInfoFields) HasBillingAddress2() bool`

HasBillingAddress2 returns a boolean if a field has been set.

### GetBusiness

`func (o *UsersInfoFields) GetBusiness() bool`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *UsersInfoFields) GetBusinessOk() (*bool, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *UsersInfoFields) SetBusiness(v bool)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *UsersInfoFields) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetCompanyName

`func (o *UsersInfoFields) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UsersInfoFields) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UsersInfoFields) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UsersInfoFields) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *UsersInfoFields) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UsersInfoFields) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UsersInfoFields) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UsersInfoFields) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UsersInfoFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsersInfoFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsersInfoFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UsersInfoFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *UsersInfoFields) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersInfoFields) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersInfoFields) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UsersInfoFields) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *UsersInfoFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersInfoFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersInfoFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UsersInfoFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UsersInfoFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersInfoFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersInfoFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsersInfoFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *UsersInfoFields) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UsersInfoFields) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UsersInfoFields) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UsersInfoFields) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPhone

`func (o *UsersInfoFields) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UsersInfoFields) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UsersInfoFields) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UsersInfoFields) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *UsersInfoFields) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UsersInfoFields) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UsersInfoFields) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UsersInfoFields) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStripeUserId

`func (o *UsersInfoFields) GetStripeUserId() string`

GetStripeUserId returns the StripeUserId field if non-nil, zero value otherwise.

### GetStripeUserIdOk

`func (o *UsersInfoFields) GetStripeUserIdOk() (*string, bool)`

GetStripeUserIdOk returns a tuple with the StripeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUserId

`func (o *UsersInfoFields) SetStripeUserId(v string)`

SetStripeUserId sets StripeUserId field to given value.

### HasStripeUserId

`func (o *UsersInfoFields) HasStripeUserId() bool`

HasStripeUserId returns a boolean if a field has been set.

### GetVatNumber

`func (o *UsersInfoFields) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UsersInfoFields) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UsersInfoFields) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UsersInfoFields) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetZipCode

`func (o *UsersInfoFields) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UsersInfoFields) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UsersInfoFields) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UsersInfoFields) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



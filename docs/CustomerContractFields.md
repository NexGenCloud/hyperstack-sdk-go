# CustomerContractFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Discounts** | Pointer to [**[]ContractDiscountPlanFields**](ContractDiscountPlanFields.md) |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationPolicy** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerContractFields

`func NewCustomerContractFields() *CustomerContractFields`

NewCustomerContractFields instantiates a new CustomerContractFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerContractFieldsWithDefaults

`func NewCustomerContractFieldsWithDefaults() *CustomerContractFields`

NewCustomerContractFieldsWithDefaults instantiates a new CustomerContractFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CustomerContractFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerContractFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerContractFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerContractFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CustomerContractFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerContractFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerContractFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerContractFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscounts

`func (o *CustomerContractFields) GetDiscounts() []ContractDiscountPlanFields`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *CustomerContractFields) GetDiscountsOk() (*[]ContractDiscountPlanFields, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *CustomerContractFields) SetDiscounts(v []ContractDiscountPlanFields)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *CustomerContractFields) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetEndDate

`func (o *CustomerContractFields) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CustomerContractFields) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CustomerContractFields) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CustomerContractFields) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *CustomerContractFields) GetExpirationPolicy() int32`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *CustomerContractFields) GetExpirationPolicyOk() (*int32, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *CustomerContractFields) SetExpirationPolicy(v int32)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *CustomerContractFields) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetId

`func (o *CustomerContractFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerContractFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerContractFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerContractFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *CustomerContractFields) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CustomerContractFields) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CustomerContractFields) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *CustomerContractFields) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStartDate

`func (o *CustomerContractFields) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CustomerContractFields) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CustomerContractFields) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CustomerContractFields) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerContractFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerContractFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerContractFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerContractFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



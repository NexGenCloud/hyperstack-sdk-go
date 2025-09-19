# CreateContractFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiscountPlans** | Pointer to [**[]ContractDiscountPlanFields**](ContractDiscountPlanFields.md) |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationPolicy** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateContractFields

`func NewCreateContractFields() *CreateContractFields`

NewCreateContractFields instantiates a new CreateContractFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContractFieldsWithDefaults

`func NewCreateContractFieldsWithDefaults() *CreateContractFields`

NewCreateContractFieldsWithDefaults instantiates a new CreateContractFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CreateContractFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateContractFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateContractFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateContractFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CreateContractFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateContractFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateContractFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateContractFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountPlans

`func (o *CreateContractFields) GetDiscountPlans() []ContractDiscountPlanFields`

GetDiscountPlans returns the DiscountPlans field if non-nil, zero value otherwise.

### GetDiscountPlansOk

`func (o *CreateContractFields) GetDiscountPlansOk() (*[]ContractDiscountPlanFields, bool)`

GetDiscountPlansOk returns a tuple with the DiscountPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPlans

`func (o *CreateContractFields) SetDiscountPlans(v []ContractDiscountPlanFields)`

SetDiscountPlans sets DiscountPlans field to given value.

### HasDiscountPlans

`func (o *CreateContractFields) HasDiscountPlans() bool`

HasDiscountPlans returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateContractFields) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateContractFields) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateContractFields) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateContractFields) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *CreateContractFields) GetExpirationPolicy() int32`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *CreateContractFields) GetExpirationPolicyOk() (*int32, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *CreateContractFields) SetExpirationPolicy(v int32)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *CreateContractFields) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetId

`func (o *CreateContractFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateContractFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateContractFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateContractFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *CreateContractFields) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CreateContractFields) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CreateContractFields) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *CreateContractFields) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateContractFields) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateContractFields) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateContractFields) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateContractFields) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



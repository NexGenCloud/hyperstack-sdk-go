# CreateContractPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DiscountResources** | [**[]ContractResourcePayload**](ContractResourcePayload.md) |  | 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationPolicy** | **int32** |  | 
**OrgId** | **int32** |  | 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateContractPayload

`func NewCreateContractPayload(discountResources []ContractResourcePayload, expirationPolicy int32, orgId int32, ) *CreateContractPayload`

NewCreateContractPayload instantiates a new CreateContractPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContractPayloadWithDefaults

`func NewCreateContractPayloadWithDefaults() *CreateContractPayload`

NewCreateContractPayloadWithDefaults instantiates a new CreateContractPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateContractPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateContractPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateContractPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateContractPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountResources

`func (o *CreateContractPayload) GetDiscountResources() []ContractResourcePayload`

GetDiscountResources returns the DiscountResources field if non-nil, zero value otherwise.

### GetDiscountResourcesOk

`func (o *CreateContractPayload) GetDiscountResourcesOk() (*[]ContractResourcePayload, bool)`

GetDiscountResourcesOk returns a tuple with the DiscountResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountResources

`func (o *CreateContractPayload) SetDiscountResources(v []ContractResourcePayload)`

SetDiscountResources sets DiscountResources field to given value.


### GetEndDate

`func (o *CreateContractPayload) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateContractPayload) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateContractPayload) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateContractPayload) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *CreateContractPayload) GetExpirationPolicy() int32`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *CreateContractPayload) GetExpirationPolicyOk() (*int32, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *CreateContractPayload) SetExpirationPolicy(v int32)`

SetExpirationPolicy sets ExpirationPolicy field to given value.


### GetOrgId

`func (o *CreateContractPayload) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CreateContractPayload) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CreateContractPayload) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetStartDate

`func (o *CreateContractPayload) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateContractPayload) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateContractPayload) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateContractPayload) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ContractResourcePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountAmount** | Pointer to **float32** |  | [optional] 
**DiscountPercent** | Pointer to **float32** |  | [optional] 
**DiscountType** | **string** |  | 
**ResourceCount** | Pointer to **int32** |  | [optional] 
**ResourceId** | **int32** |  | 

## Methods

### NewContractResourcePayload

`func NewContractResourcePayload(discountType string, resourceId int32, ) *ContractResourcePayload`

NewContractResourcePayload instantiates a new ContractResourcePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractResourcePayloadWithDefaults

`func NewContractResourcePayloadWithDefaults() *ContractResourcePayload`

NewContractResourcePayloadWithDefaults instantiates a new ContractResourcePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountAmount

`func (o *ContractResourcePayload) GetDiscountAmount() float32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *ContractResourcePayload) GetDiscountAmountOk() (*float32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *ContractResourcePayload) SetDiscountAmount(v float32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *ContractResourcePayload) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercent

`func (o *ContractResourcePayload) GetDiscountPercent() float32`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *ContractResourcePayload) GetDiscountPercentOk() (*float32, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *ContractResourcePayload) SetDiscountPercent(v float32)`

SetDiscountPercent sets DiscountPercent field to given value.

### HasDiscountPercent

`func (o *ContractResourcePayload) HasDiscountPercent() bool`

HasDiscountPercent returns a boolean if a field has been set.

### GetDiscountType

`func (o *ContractResourcePayload) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *ContractResourcePayload) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *ContractResourcePayload) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.


### GetResourceCount

`func (o *ContractResourcePayload) GetResourceCount() int32`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *ContractResourcePayload) GetResourceCountOk() (*int32, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *ContractResourcePayload) SetResourceCount(v int32)`

SetResourceCount sets ResourceCount field to given value.

### HasResourceCount

`func (o *ContractResourcePayload) HasResourceCount() bool`

HasResourceCount returns a boolean if a field has been set.

### GetResourceId

`func (o *ContractResourcePayload) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ContractResourcePayload) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ContractResourcePayload) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



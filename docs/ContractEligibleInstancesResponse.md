# ContractEligibleInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCount** | Pointer to **int32** |  | [optional] 
**Instances** | Pointer to [**[]ContractEligibleInstanceFields**](ContractEligibleInstanceFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewContractEligibleInstancesResponse

`func NewContractEligibleInstancesResponse() *ContractEligibleInstancesResponse`

NewContractEligibleInstancesResponse instantiates a new ContractEligibleInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractEligibleInstancesResponseWithDefaults

`func NewContractEligibleInstancesResponseWithDefaults() *ContractEligibleInstancesResponse`

NewContractEligibleInstancesResponseWithDefaults instantiates a new ContractEligibleInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCount

`func (o *ContractEligibleInstancesResponse) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ContractEligibleInstancesResponse) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ContractEligibleInstancesResponse) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *ContractEligibleInstancesResponse) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstances

`func (o *ContractEligibleInstancesResponse) GetInstances() []ContractEligibleInstanceFields`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ContractEligibleInstancesResponse) GetInstancesOk() (*[]ContractEligibleInstanceFields, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ContractEligibleInstancesResponse) SetInstances(v []ContractEligibleInstanceFields)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ContractEligibleInstancesResponse) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMessage

`func (o *ContractEligibleInstancesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContractEligibleInstancesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContractEligibleInstancesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ContractEligibleInstancesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ContractEligibleInstancesResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractEligibleInstancesResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractEligibleInstancesResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContractEligibleInstancesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



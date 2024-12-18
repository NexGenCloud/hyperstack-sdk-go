# ContractInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]ContractInstanceFields**](ContractInstanceFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewContractInstancesResponse

`func NewContractInstancesResponse() *ContractInstancesResponse`

NewContractInstancesResponse instantiates a new ContractInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractInstancesResponseWithDefaults

`func NewContractInstancesResponseWithDefaults() *ContractInstancesResponse`

NewContractInstancesResponseWithDefaults instantiates a new ContractInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ContractInstancesResponse) GetInstances() []ContractInstanceFields`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ContractInstancesResponse) GetInstancesOk() (*[]ContractInstanceFields, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ContractInstancesResponse) SetInstances(v []ContractInstanceFields)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ContractInstancesResponse) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMessage

`func (o *ContractInstancesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContractInstancesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContractInstancesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ContractInstancesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ContractInstancesResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractInstancesResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractInstancesResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContractInstancesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



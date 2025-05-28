# SnapshotRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | Pointer to **int32** | Contract ID to assign to the newly restored VM | [optional] 
**NewVmName** | **string** | The name of the newly restored VM | 

## Methods

### NewSnapshotRestoreRequest

`func NewSnapshotRestoreRequest(newVmName string, ) *SnapshotRestoreRequest`

NewSnapshotRestoreRequest instantiates a new SnapshotRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRestoreRequestWithDefaults

`func NewSnapshotRestoreRequestWithDefaults() *SnapshotRestoreRequest`

NewSnapshotRestoreRequestWithDefaults instantiates a new SnapshotRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *SnapshotRestoreRequest) GetContractId() int32`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *SnapshotRestoreRequest) GetContractIdOk() (*int32, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *SnapshotRestoreRequest) SetContractId(v int32)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *SnapshotRestoreRequest) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetNewVmName

`func (o *SnapshotRestoreRequest) GetNewVmName() string`

GetNewVmName returns the NewVmName field if non-nil, zero value otherwise.

### GetNewVmNameOk

`func (o *SnapshotRestoreRequest) GetNewVmNameOk() (*string, bool)`

GetNewVmNameOk returns a tuple with the NewVmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVmName

`func (o *SnapshotRestoreRequest) SetNewVmName(v string)`

SetNewVmName sets NewVmName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



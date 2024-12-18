# VMUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**VirtualMachines** | Pointer to [**[]VirtualMachineUsage**](VirtualMachineUsage.md) |  | [optional] 

## Methods

### NewVMUsageResponse

`func NewVMUsageResponse() *VMUsageResponse`

NewVMUsageResponse instantiates a new VMUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMUsageResponseWithDefaults

`func NewVMUsageResponseWithDefaults() *VMUsageResponse`

NewVMUsageResponseWithDefaults instantiates a new VMUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VMUsageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VMUsageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VMUsageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VMUsageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrgId

`func (o *VMUsageResponse) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *VMUsageResponse) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *VMUsageResponse) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *VMUsageResponse) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStatus

`func (o *VMUsageResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VMUsageResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VMUsageResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VMUsageResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVirtualMachines

`func (o *VMUsageResponse) GetVirtualMachines() []VirtualMachineUsage`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *VMUsageResponse) GetVirtualMachinesOk() (*[]VirtualMachineUsage, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *VMUsageResponse) SetVirtualMachines(v []VirtualMachineUsage)`

SetVirtualMachines sets VirtualMachines field to given value.

### HasVirtualMachines

`func (o *VMUsageResponse) HasVirtualMachines() bool`

HasVirtualMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



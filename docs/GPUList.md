# GPUList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuList** | Pointer to [**[]GPUFields**](GPUFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGPUList

`func NewGPUList() *GPUList`

NewGPUList instantiates a new GPUList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPUListWithDefaults

`func NewGPUListWithDefaults() *GPUList`

NewGPUListWithDefaults instantiates a new GPUList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuList

`func (o *GPUList) GetGpuList() []GPUFields`

GetGpuList returns the GpuList field if non-nil, zero value otherwise.

### GetGpuListOk

`func (o *GPUList) GetGpuListOk() (*[]GPUFields, bool)`

GetGpuListOk returns a tuple with the GpuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuList

`func (o *GPUList) SetGpuList(v []GPUFields)`

SetGpuList sets GpuList field to given value.

### HasGpuList

`func (o *GPUList) HasGpuList() bool`

HasGpuList returns a boolean if a field has been set.

### GetMessage

`func (o *GPUList) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GPUList) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GPUList) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GPUList) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GPUList) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GPUList) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GPUList) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GPUList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



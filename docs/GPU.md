# GPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpu** | Pointer to [**GPUFields**](GPUFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGPU

`func NewGPU() *GPU`

NewGPU instantiates a new GPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPUWithDefaults

`func NewGPUWithDefaults() *GPU`

NewGPUWithDefaults instantiates a new GPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpu

`func (o *GPU) GetGpu() GPUFields`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *GPU) GetGpuOk() (*GPUFields, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *GPU) SetGpu(v GPUFields)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *GPU) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetMessage

`func (o *GPU) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GPU) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GPU) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GPU) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GPU) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GPU) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GPU) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GPU) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



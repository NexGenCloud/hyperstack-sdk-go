# GpuClusterRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuCount** | **int32** | Number of GPUs to allocate. | 
**Gpus** | **[]string** | List of GPU types to include in the cluster. | 
**Purpose** | **string** | Purpose of the GPU cluster. | 

## Methods

### NewGpuClusterRequestPayload

`func NewGpuClusterRequestPayload(gpuCount int32, gpus []string, purpose string, ) *GpuClusterRequestPayload`

NewGpuClusterRequestPayload instantiates a new GpuClusterRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuClusterRequestPayloadWithDefaults

`func NewGpuClusterRequestPayloadWithDefaults() *GpuClusterRequestPayload`

NewGpuClusterRequestPayloadWithDefaults instantiates a new GpuClusterRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuCount

`func (o *GpuClusterRequestPayload) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *GpuClusterRequestPayload) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *GpuClusterRequestPayload) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.


### GetGpus

`func (o *GpuClusterRequestPayload) GetGpus() []string`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *GpuClusterRequestPayload) GetGpusOk() (*[]string, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *GpuClusterRequestPayload) SetGpus(v []string)`

SetGpus sets Gpus field to given value.


### GetPurpose

`func (o *GpuClusterRequestPayload) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *GpuClusterRequestPayload) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *GpuClusterRequestPayload) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



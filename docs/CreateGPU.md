# CreateGPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExampleMetadata** | Pointer to **string** | A valid JSON string. | [optional] 
**Name** | **string** |  | 
**Regions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateGPU

`func NewCreateGPU(name string, ) *CreateGPU`

NewCreateGPU instantiates a new CreateGPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGPUWithDefaults

`func NewCreateGPUWithDefaults() *CreateGPU`

NewCreateGPUWithDefaults instantiates a new CreateGPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExampleMetadata

`func (o *CreateGPU) GetExampleMetadata() string`

GetExampleMetadata returns the ExampleMetadata field if non-nil, zero value otherwise.

### GetExampleMetadataOk

`func (o *CreateGPU) GetExampleMetadataOk() (*string, bool)`

GetExampleMetadataOk returns a tuple with the ExampleMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleMetadata

`func (o *CreateGPU) SetExampleMetadata(v string)`

SetExampleMetadata sets ExampleMetadata field to given value.

### HasExampleMetadata

`func (o *CreateGPU) HasExampleMetadata() bool`

HasExampleMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateGPU) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGPU) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGPU) SetName(v string)`

SetName sets Name field to given value.


### GetRegions

`func (o *CreateGPU) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CreateGPU) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CreateGPU) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CreateGPU) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



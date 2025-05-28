# GPUFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExampleMetadata** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Regions** | Pointer to [**[]GPURegionFields**](GPURegionFields.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGPUFields

`func NewGPUFields() *GPUFields`

NewGPUFields instantiates a new GPUFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPUFieldsWithDefaults

`func NewGPUFieldsWithDefaults() *GPUFields`

NewGPUFieldsWithDefaults instantiates a new GPUFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GPUFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GPUFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GPUFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GPUFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExampleMetadata

`func (o *GPUFields) GetExampleMetadata() string`

GetExampleMetadata returns the ExampleMetadata field if non-nil, zero value otherwise.

### GetExampleMetadataOk

`func (o *GPUFields) GetExampleMetadataOk() (*string, bool)`

GetExampleMetadataOk returns a tuple with the ExampleMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleMetadata

`func (o *GPUFields) SetExampleMetadata(v string)`

SetExampleMetadata sets ExampleMetadata field to given value.

### HasExampleMetadata

`func (o *GPUFields) HasExampleMetadata() bool`

HasExampleMetadata returns a boolean if a field has been set.

### GetId

`func (o *GPUFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GPUFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GPUFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GPUFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GPUFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GPUFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GPUFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GPUFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegions

`func (o *GPUFields) GetRegions() []GPURegionFields`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GPUFields) GetRegionsOk() (*[]GPURegionFields, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GPUFields) SetRegions(v []GPURegionFields)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GPUFields) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GPUFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GPUFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GPUFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GPUFields) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



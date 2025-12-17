# ObjectStorageAccessKeyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **map[string]interface{}** |  | [optional] 
**Region** | [**ObjectStorageRegionsEnum**](ObjectStorageRegionsEnum.md) |  | 

## Methods

### NewObjectStorageAccessKeyCreateRequest

`func NewObjectStorageAccessKeyCreateRequest(region ObjectStorageRegionsEnum, ) *ObjectStorageAccessKeyCreateRequest`

NewObjectStorageAccessKeyCreateRequest instantiates a new ObjectStorageAccessKeyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageAccessKeyCreateRequestWithDefaults

`func NewObjectStorageAccessKeyCreateRequestWithDefaults() *ObjectStorageAccessKeyCreateRequest`

NewObjectStorageAccessKeyCreateRequestWithDefaults instantiates a new ObjectStorageAccessKeyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ObjectStorageAccessKeyCreateRequest) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageAccessKeyCreateRequest) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageAccessKeyCreateRequest) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageAccessKeyCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRegion

`func (o *ObjectStorageAccessKeyCreateRequest) GetRegion() ObjectStorageRegionsEnum`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ObjectStorageAccessKeyCreateRequest) GetRegionOk() (*ObjectStorageRegionsEnum, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ObjectStorageAccessKeyCreateRequest) SetRegion(v ObjectStorageRegionsEnum)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ObjectStorageBucketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Endpoint** | **string** |  | 
**Name** | **string** |  | 
**NumObjects** | **int32** | Number of objects | 
**Region** | [**ObjectStorageRegionsEnum**](ObjectStorageRegionsEnum.md) |  | 
**SizeBytes** | **int32** | Accumulated size in bytes | 
**SizeBytesActual** | **int32** | Size utilized in bytes | 

## Methods

### NewObjectStorageBucketResponse

`func NewObjectStorageBucketResponse(createdAt time.Time, endpoint string, name string, numObjects int32, region ObjectStorageRegionsEnum, sizeBytes int32, sizeBytesActual int32, ) *ObjectStorageBucketResponse`

NewObjectStorageBucketResponse instantiates a new ObjectStorageBucketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketResponseWithDefaults

`func NewObjectStorageBucketResponseWithDefaults() *ObjectStorageBucketResponse`

NewObjectStorageBucketResponseWithDefaults instantiates a new ObjectStorageBucketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ObjectStorageBucketResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ObjectStorageBucketResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ObjectStorageBucketResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEndpoint

`func (o *ObjectStorageBucketResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ObjectStorageBucketResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ObjectStorageBucketResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *ObjectStorageBucketResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageBucketResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageBucketResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNumObjects

`func (o *ObjectStorageBucketResponse) GetNumObjects() int32`

GetNumObjects returns the NumObjects field if non-nil, zero value otherwise.

### GetNumObjectsOk

`func (o *ObjectStorageBucketResponse) GetNumObjectsOk() (*int32, bool)`

GetNumObjectsOk returns a tuple with the NumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjects

`func (o *ObjectStorageBucketResponse) SetNumObjects(v int32)`

SetNumObjects sets NumObjects field to given value.


### GetRegion

`func (o *ObjectStorageBucketResponse) GetRegion() ObjectStorageRegionsEnum`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ObjectStorageBucketResponse) GetRegionOk() (*ObjectStorageRegionsEnum, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ObjectStorageBucketResponse) SetRegion(v ObjectStorageRegionsEnum)`

SetRegion sets Region field to given value.


### GetSizeBytes

`func (o *ObjectStorageBucketResponse) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ObjectStorageBucketResponse) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ObjectStorageBucketResponse) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetSizeBytesActual

`func (o *ObjectStorageBucketResponse) GetSizeBytesActual() int32`

GetSizeBytesActual returns the SizeBytesActual field if non-nil, zero value otherwise.

### GetSizeBytesActualOk

`func (o *ObjectStorageBucketResponse) GetSizeBytesActualOk() (*int32, bool)`

GetSizeBytesActualOk returns a tuple with the SizeBytesActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytesActual

`func (o *ObjectStorageBucketResponse) SetSizeBytesActual(v int32)`

SetSizeBytesActual sets SizeBytesActual field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



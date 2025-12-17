# ObjectStorageAccessKeyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeys** | [**[]ObjectStorageAccessKeyResponse**](ObjectStorageAccessKeyResponse.md) |  | 
**Meta** | [**ObjectStoragePaginationMeta**](ObjectStoragePaginationMeta.md) |  | 

## Methods

### NewObjectStorageAccessKeyListResponse

`func NewObjectStorageAccessKeyListResponse(accessKeys []ObjectStorageAccessKeyResponse, meta ObjectStoragePaginationMeta, ) *ObjectStorageAccessKeyListResponse`

NewObjectStorageAccessKeyListResponse instantiates a new ObjectStorageAccessKeyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageAccessKeyListResponseWithDefaults

`func NewObjectStorageAccessKeyListResponseWithDefaults() *ObjectStorageAccessKeyListResponse`

NewObjectStorageAccessKeyListResponseWithDefaults instantiates a new ObjectStorageAccessKeyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeys

`func (o *ObjectStorageAccessKeyListResponse) GetAccessKeys() []ObjectStorageAccessKeyResponse`

GetAccessKeys returns the AccessKeys field if non-nil, zero value otherwise.

### GetAccessKeysOk

`func (o *ObjectStorageAccessKeyListResponse) GetAccessKeysOk() (*[]ObjectStorageAccessKeyResponse, bool)`

GetAccessKeysOk returns a tuple with the AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeys

`func (o *ObjectStorageAccessKeyListResponse) SetAccessKeys(v []ObjectStorageAccessKeyResponse)`

SetAccessKeys sets AccessKeys field to given value.


### GetMeta

`func (o *ObjectStorageAccessKeyListResponse) GetMeta() ObjectStoragePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ObjectStorageAccessKeyListResponse) GetMetaOk() (*ObjectStoragePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ObjectStorageAccessKeyListResponse) SetMeta(v ObjectStoragePaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



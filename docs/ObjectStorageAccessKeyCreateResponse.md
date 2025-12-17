# ObjectStorageAccessKeyCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | **int32** |  | 
**Region** | [**ObjectStorageRegionsEnum**](ObjectStorageRegionsEnum.md) |  | 
**SecretKey** | **string** |  | 
**UserId** | **int32** |  | 

## Methods

### NewObjectStorageAccessKeyCreateResponse

`func NewObjectStorageAccessKeyCreateResponse(accessKey string, createdAt time.Time, id int32, region ObjectStorageRegionsEnum, secretKey string, userId int32, ) *ObjectStorageAccessKeyCreateResponse`

NewObjectStorageAccessKeyCreateResponse instantiates a new ObjectStorageAccessKeyCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageAccessKeyCreateResponseWithDefaults

`func NewObjectStorageAccessKeyCreateResponseWithDefaults() *ObjectStorageAccessKeyCreateResponse`

NewObjectStorageAccessKeyCreateResponseWithDefaults instantiates a new ObjectStorageAccessKeyCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ObjectStorageAccessKeyCreateResponse) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ObjectStorageAccessKeyCreateResponse) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ObjectStorageAccessKeyCreateResponse) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetCreatedAt

`func (o *ObjectStorageAccessKeyCreateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ObjectStorageAccessKeyCreateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ObjectStorageAccessKeyCreateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *ObjectStorageAccessKeyCreateResponse) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageAccessKeyCreateResponse) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageAccessKeyCreateResponse) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageAccessKeyCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageAccessKeyCreateResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageAccessKeyCreateResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageAccessKeyCreateResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetRegion

`func (o *ObjectStorageAccessKeyCreateResponse) GetRegion() ObjectStorageRegionsEnum`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ObjectStorageAccessKeyCreateResponse) GetRegionOk() (*ObjectStorageRegionsEnum, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ObjectStorageAccessKeyCreateResponse) SetRegion(v ObjectStorageRegionsEnum)`

SetRegion sets Region field to given value.


### GetSecretKey

`func (o *ObjectStorageAccessKeyCreateResponse) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ObjectStorageAccessKeyCreateResponse) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ObjectStorageAccessKeyCreateResponse) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetUserId

`func (o *ObjectStorageAccessKeyCreateResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ObjectStorageAccessKeyCreateResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ObjectStorageAccessKeyCreateResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



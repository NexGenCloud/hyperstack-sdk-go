# ObjectStorageBucketListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]ObjectStorageBucketResponse**](ObjectStorageBucketResponse.md) |  | 

## Methods

### NewObjectStorageBucketListResponse

`func NewObjectStorageBucketListResponse(buckets []ObjectStorageBucketResponse, ) *ObjectStorageBucketListResponse`

NewObjectStorageBucketListResponse instantiates a new ObjectStorageBucketListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketListResponseWithDefaults

`func NewObjectStorageBucketListResponseWithDefaults() *ObjectStorageBucketListResponse`

NewObjectStorageBucketListResponseWithDefaults instantiates a new ObjectStorageBucketListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *ObjectStorageBucketListResponse) GetBuckets() []ObjectStorageBucketResponse`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ObjectStorageBucketListResponse) GetBucketsOk() (*[]ObjectStorageBucketResponse, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ObjectStorageBucketListResponse) SetBuckets(v []ObjectStorageBucketResponse)`

SetBuckets sets Buckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResourceLevelBucketBillingHistoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryBucket** | Pointer to [**ResourceLevelBillingHistory**](ResourceLevelBillingHistory.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceLevelBucketBillingHistoryResponseModel

`func NewResourceLevelBucketBillingHistoryResponseModel() *ResourceLevelBucketBillingHistoryResponseModel`

NewResourceLevelBucketBillingHistoryResponseModel instantiates a new ResourceLevelBucketBillingHistoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelBucketBillingHistoryResponseModelWithDefaults

`func NewResourceLevelBucketBillingHistoryResponseModelWithDefaults() *ResourceLevelBucketBillingHistoryResponseModel`

NewResourceLevelBucketBillingHistoryResponseModelWithDefaults instantiates a new ResourceLevelBucketBillingHistoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryBucket

`func (o *ResourceLevelBucketBillingHistoryResponseModel) GetBillingHistoryBucket() ResourceLevelBillingHistory`

GetBillingHistoryBucket returns the BillingHistoryBucket field if non-nil, zero value otherwise.

### GetBillingHistoryBucketOk

`func (o *ResourceLevelBucketBillingHistoryResponseModel) GetBillingHistoryBucketOk() (*ResourceLevelBillingHistory, bool)`

GetBillingHistoryBucketOk returns a tuple with the BillingHistoryBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryBucket

`func (o *ResourceLevelBucketBillingHistoryResponseModel) SetBillingHistoryBucket(v ResourceLevelBillingHistory)`

SetBillingHistoryBucket sets BillingHistoryBucket field to given value.

### HasBillingHistoryBucket

`func (o *ResourceLevelBucketBillingHistoryResponseModel) HasBillingHistoryBucket() bool`

HasBillingHistoryBucket returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceLevelBucketBillingHistoryResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceLevelBucketBillingHistoryResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceLevelBucketBillingHistoryResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceLevelBucketBillingHistoryResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceLevelBucketBillingHistoryResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceLevelBucketBillingHistoryResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceLevelBucketBillingHistoryResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceLevelBucketBillingHistoryResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



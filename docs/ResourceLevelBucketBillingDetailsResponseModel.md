# ResourceLevelBucketBillingDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryBucketDetails** | Pointer to [**ResourceLevelBillingBucketDetailsResources**](ResourceLevelBillingBucketDetailsResources.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceLevelBucketBillingDetailsResponseModel

`func NewResourceLevelBucketBillingDetailsResponseModel() *ResourceLevelBucketBillingDetailsResponseModel`

NewResourceLevelBucketBillingDetailsResponseModel instantiates a new ResourceLevelBucketBillingDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelBucketBillingDetailsResponseModelWithDefaults

`func NewResourceLevelBucketBillingDetailsResponseModelWithDefaults() *ResourceLevelBucketBillingDetailsResponseModel`

NewResourceLevelBucketBillingDetailsResponseModelWithDefaults instantiates a new ResourceLevelBucketBillingDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryBucketDetails

`func (o *ResourceLevelBucketBillingDetailsResponseModel) GetBillingHistoryBucketDetails() ResourceLevelBillingBucketDetailsResources`

GetBillingHistoryBucketDetails returns the BillingHistoryBucketDetails field if non-nil, zero value otherwise.

### GetBillingHistoryBucketDetailsOk

`func (o *ResourceLevelBucketBillingDetailsResponseModel) GetBillingHistoryBucketDetailsOk() (*ResourceLevelBillingBucketDetailsResources, bool)`

GetBillingHistoryBucketDetailsOk returns a tuple with the BillingHistoryBucketDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryBucketDetails

`func (o *ResourceLevelBucketBillingDetailsResponseModel) SetBillingHistoryBucketDetails(v ResourceLevelBillingBucketDetailsResources)`

SetBillingHistoryBucketDetails sets BillingHistoryBucketDetails field to given value.

### HasBillingHistoryBucketDetails

`func (o *ResourceLevelBucketBillingDetailsResponseModel) HasBillingHistoryBucketDetails() bool`

HasBillingHistoryBucketDetails returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceLevelBucketBillingDetailsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceLevelBucketBillingDetailsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceLevelBucketBillingDetailsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceLevelBucketBillingDetailsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceLevelBucketBillingDetailsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceLevelBucketBillingDetailsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceLevelBucketBillingDetailsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceLevelBucketBillingDetailsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



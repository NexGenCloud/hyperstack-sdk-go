# ResourceLevelClusterBillingDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryClusterDetails** | Pointer to [**ResourceLevelBillingHistoryResourcesCluster**](ResourceLevelBillingHistoryResourcesCluster.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceLevelClusterBillingDetailsResponseModel

`func NewResourceLevelClusterBillingDetailsResponseModel() *ResourceLevelClusterBillingDetailsResponseModel`

NewResourceLevelClusterBillingDetailsResponseModel instantiates a new ResourceLevelClusterBillingDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelClusterBillingDetailsResponseModelWithDefaults

`func NewResourceLevelClusterBillingDetailsResponseModelWithDefaults() *ResourceLevelClusterBillingDetailsResponseModel`

NewResourceLevelClusterBillingDetailsResponseModelWithDefaults instantiates a new ResourceLevelClusterBillingDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryClusterDetails

`func (o *ResourceLevelClusterBillingDetailsResponseModel) GetBillingHistoryClusterDetails() ResourceLevelBillingHistoryResourcesCluster`

GetBillingHistoryClusterDetails returns the BillingHistoryClusterDetails field if non-nil, zero value otherwise.

### GetBillingHistoryClusterDetailsOk

`func (o *ResourceLevelClusterBillingDetailsResponseModel) GetBillingHistoryClusterDetailsOk() (*ResourceLevelBillingHistoryResourcesCluster, bool)`

GetBillingHistoryClusterDetailsOk returns a tuple with the BillingHistoryClusterDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryClusterDetails

`func (o *ResourceLevelClusterBillingDetailsResponseModel) SetBillingHistoryClusterDetails(v ResourceLevelBillingHistoryResourcesCluster)`

SetBillingHistoryClusterDetails sets BillingHistoryClusterDetails field to given value.

### HasBillingHistoryClusterDetails

`func (o *ResourceLevelClusterBillingDetailsResponseModel) HasBillingHistoryClusterDetails() bool`

HasBillingHistoryClusterDetails returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceLevelClusterBillingDetailsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceLevelClusterBillingDetailsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceLevelClusterBillingDetailsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceLevelClusterBillingDetailsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceLevelClusterBillingDetailsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceLevelClusterBillingDetailsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceLevelClusterBillingDetailsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceLevelClusterBillingDetailsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResourceLevelVolumeBillingDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryVolumeDetails** | Pointer to [**ResourceLevelBillingDetailsVolume**](ResourceLevelBillingDetailsVolume.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceLevelVolumeBillingDetailsResponseModel

`func NewResourceLevelVolumeBillingDetailsResponseModel() *ResourceLevelVolumeBillingDetailsResponseModel`

NewResourceLevelVolumeBillingDetailsResponseModel instantiates a new ResourceLevelVolumeBillingDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelVolumeBillingDetailsResponseModelWithDefaults

`func NewResourceLevelVolumeBillingDetailsResponseModelWithDefaults() *ResourceLevelVolumeBillingDetailsResponseModel`

NewResourceLevelVolumeBillingDetailsResponseModelWithDefaults instantiates a new ResourceLevelVolumeBillingDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryVolumeDetails

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) GetBillingHistoryVolumeDetails() ResourceLevelBillingDetailsVolume`

GetBillingHistoryVolumeDetails returns the BillingHistoryVolumeDetails field if non-nil, zero value otherwise.

### GetBillingHistoryVolumeDetailsOk

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) GetBillingHistoryVolumeDetailsOk() (*ResourceLevelBillingDetailsVolume, bool)`

GetBillingHistoryVolumeDetailsOk returns a tuple with the BillingHistoryVolumeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryVolumeDetails

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) SetBillingHistoryVolumeDetails(v ResourceLevelBillingDetailsVolume)`

SetBillingHistoryVolumeDetails sets BillingHistoryVolumeDetails field to given value.

### HasBillingHistoryVolumeDetails

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) HasBillingHistoryVolumeDetails() bool`

HasBillingHistoryVolumeDetails returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceLevelVolumeBillingDetailsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



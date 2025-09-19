# ResourceLevelVMBillingDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryVmDetails** | Pointer to [**ResourceLevelBillingDetailsVM**](ResourceLevelBillingDetailsVM.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceLevelVMBillingDetailsResponseModel

`func NewResourceLevelVMBillingDetailsResponseModel() *ResourceLevelVMBillingDetailsResponseModel`

NewResourceLevelVMBillingDetailsResponseModel instantiates a new ResourceLevelVMBillingDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelVMBillingDetailsResponseModelWithDefaults

`func NewResourceLevelVMBillingDetailsResponseModelWithDefaults() *ResourceLevelVMBillingDetailsResponseModel`

NewResourceLevelVMBillingDetailsResponseModelWithDefaults instantiates a new ResourceLevelVMBillingDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryVmDetails

`func (o *ResourceLevelVMBillingDetailsResponseModel) GetBillingHistoryVmDetails() ResourceLevelBillingDetailsVM`

GetBillingHistoryVmDetails returns the BillingHistoryVmDetails field if non-nil, zero value otherwise.

### GetBillingHistoryVmDetailsOk

`func (o *ResourceLevelVMBillingDetailsResponseModel) GetBillingHistoryVmDetailsOk() (*ResourceLevelBillingDetailsVM, bool)`

GetBillingHistoryVmDetailsOk returns a tuple with the BillingHistoryVmDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryVmDetails

`func (o *ResourceLevelVMBillingDetailsResponseModel) SetBillingHistoryVmDetails(v ResourceLevelBillingDetailsVM)`

SetBillingHistoryVmDetails sets BillingHistoryVmDetails field to given value.

### HasBillingHistoryVmDetails

`func (o *ResourceLevelVMBillingDetailsResponseModel) HasBillingHistoryVmDetails() bool`

HasBillingHistoryVmDetails returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceLevelVMBillingDetailsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceLevelVMBillingDetailsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceLevelVMBillingDetailsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceLevelVMBillingDetailsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceLevelVMBillingDetailsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceLevelVMBillingDetailsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceLevelVMBillingDetailsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceLevelVMBillingDetailsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



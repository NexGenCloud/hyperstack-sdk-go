# BetaAccessStatusResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetaAccessRequests** | Pointer to [**[]BetaAccessStatusItem**](BetaAccessStatusItem.md) | List of beta access requests | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewBetaAccessStatusResponseModel

`func NewBetaAccessStatusResponseModel() *BetaAccessStatusResponseModel`

NewBetaAccessStatusResponseModel instantiates a new BetaAccessStatusResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAccessStatusResponseModelWithDefaults

`func NewBetaAccessStatusResponseModelWithDefaults() *BetaAccessStatusResponseModel`

NewBetaAccessStatusResponseModelWithDefaults instantiates a new BetaAccessStatusResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetaAccessRequests

`func (o *BetaAccessStatusResponseModel) GetBetaAccessRequests() []BetaAccessStatusItem`

GetBetaAccessRequests returns the BetaAccessRequests field if non-nil, zero value otherwise.

### GetBetaAccessRequestsOk

`func (o *BetaAccessStatusResponseModel) GetBetaAccessRequestsOk() (*[]BetaAccessStatusItem, bool)`

GetBetaAccessRequestsOk returns a tuple with the BetaAccessRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaAccessRequests

`func (o *BetaAccessStatusResponseModel) SetBetaAccessRequests(v []BetaAccessStatusItem)`

SetBetaAccessRequests sets BetaAccessRequests field to given value.

### HasBetaAccessRequests

`func (o *BetaAccessStatusResponseModel) HasBetaAccessRequests() bool`

HasBetaAccessRequests returns a boolean if a field has been set.

### GetMessage

`func (o *BetaAccessStatusResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BetaAccessStatusResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BetaAccessStatusResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BetaAccessStatusResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *BetaAccessStatusResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BetaAccessStatusResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BetaAccessStatusResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BetaAccessStatusResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



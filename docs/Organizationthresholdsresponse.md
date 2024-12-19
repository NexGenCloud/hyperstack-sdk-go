# Organizationthresholdsresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Thresholds** | Pointer to [**[]OrganizationThresholdfields**](OrganizationThresholdfields.md) |  | [optional] 

## Methods

### NewOrganizationthresholdsresponse

`func NewOrganizationthresholdsresponse() *Organizationthresholdsresponse`

NewOrganizationthresholdsresponse instantiates a new Organizationthresholdsresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationthresholdsresponseWithDefaults

`func NewOrganizationthresholdsresponseWithDefaults() *Organizationthresholdsresponse`

NewOrganizationthresholdsresponseWithDefaults instantiates a new Organizationthresholdsresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Organizationthresholdsresponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Organizationthresholdsresponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Organizationthresholdsresponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Organizationthresholdsresponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *Organizationthresholdsresponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Organizationthresholdsresponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Organizationthresholdsresponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Organizationthresholdsresponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThresholds

`func (o *Organizationthresholdsresponse) GetThresholds() []OrganizationThresholdfields`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Organizationthresholdsresponse) GetThresholdsOk() (*[]OrganizationThresholdfields, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Organizationthresholdsresponse) SetThresholds(v []OrganizationThresholdfields)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *Organizationthresholdsresponse) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


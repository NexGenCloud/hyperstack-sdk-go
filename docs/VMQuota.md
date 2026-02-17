# VMQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableVms** | Pointer to **int32** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**CurrentVms** | Pointer to **int32** |  | [optional] 
**MaxVms** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PercentageUsed** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewVMQuota

`func NewVMQuota() *VMQuota`

NewVMQuota instantiates a new VMQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMQuotaWithDefaults

`func NewVMQuotaWithDefaults() *VMQuota`

NewVMQuotaWithDefaults instantiates a new VMQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableVms

`func (o *VMQuota) GetAvailableVms() int32`

GetAvailableVms returns the AvailableVms field if non-nil, zero value otherwise.

### GetAvailableVmsOk

`func (o *VMQuota) GetAvailableVmsOk() (*int32, bool)`

GetAvailableVmsOk returns a tuple with the AvailableVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVms

`func (o *VMQuota) SetAvailableVms(v int32)`

SetAvailableVms sets AvailableVms field to given value.

### HasAvailableVms

`func (o *VMQuota) HasAvailableVms() bool`

HasAvailableVms returns a boolean if a field has been set.

### GetCidr

`func (o *VMQuota) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *VMQuota) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *VMQuota) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *VMQuota) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCurrentVms

`func (o *VMQuota) GetCurrentVms() int32`

GetCurrentVms returns the CurrentVms field if non-nil, zero value otherwise.

### GetCurrentVmsOk

`func (o *VMQuota) GetCurrentVmsOk() (*int32, bool)`

GetCurrentVmsOk returns a tuple with the CurrentVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVms

`func (o *VMQuota) SetCurrentVms(v int32)`

SetCurrentVms sets CurrentVms field to given value.

### HasCurrentVms

`func (o *VMQuota) HasCurrentVms() bool`

HasCurrentVms returns a boolean if a field has been set.

### GetMaxVms

`func (o *VMQuota) GetMaxVms() int32`

GetMaxVms returns the MaxVms field if non-nil, zero value otherwise.

### GetMaxVmsOk

`func (o *VMQuota) GetMaxVmsOk() (*int32, bool)`

GetMaxVmsOk returns a tuple with the MaxVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVms

`func (o *VMQuota) SetMaxVms(v int32)`

SetMaxVms sets MaxVms field to given value.

### HasMaxVms

`func (o *VMQuota) HasMaxVms() bool`

HasMaxVms returns a boolean if a field has been set.

### GetMessage

`func (o *VMQuota) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VMQuota) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VMQuota) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VMQuota) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPercentageUsed

`func (o *VMQuota) GetPercentageUsed() float32`

GetPercentageUsed returns the PercentageUsed field if non-nil, zero value otherwise.

### GetPercentageUsedOk

`func (o *VMQuota) GetPercentageUsedOk() (*float32, bool)`

GetPercentageUsedOk returns a tuple with the PercentageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageUsed

`func (o *VMQuota) SetPercentageUsed(v float32)`

SetPercentageUsed sets PercentageUsed field to given value.

### HasPercentageUsed

`func (o *VMQuota) HasPercentageUsed() bool`

HasPercentageUsed returns a boolean if a field has been set.

### GetStatus

`func (o *VMQuota) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VMQuota) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VMQuota) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VMQuota) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



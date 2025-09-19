# OrganizationThresholdUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Threshold** | Pointer to [**OrganizationThresholdFields**](OrganizationThresholdFields.md) |  | [optional] 

## Methods

### NewOrganizationThresholdUpdateResponse

`func NewOrganizationThresholdUpdateResponse() *OrganizationThresholdUpdateResponse`

NewOrganizationThresholdUpdateResponse instantiates a new OrganizationThresholdUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationThresholdUpdateResponseWithDefaults

`func NewOrganizationThresholdUpdateResponseWithDefaults() *OrganizationThresholdUpdateResponse`

NewOrganizationThresholdUpdateResponseWithDefaults instantiates a new OrganizationThresholdUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *OrganizationThresholdUpdateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrganizationThresholdUpdateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrganizationThresholdUpdateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrganizationThresholdUpdateResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationThresholdUpdateResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationThresholdUpdateResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationThresholdUpdateResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationThresholdUpdateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThreshold

`func (o *OrganizationThresholdUpdateResponse) GetThreshold() OrganizationThresholdFields`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *OrganizationThresholdUpdateResponse) GetThresholdOk() (*OrganizationThresholdFields, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *OrganizationThresholdUpdateResponse) SetThreshold(v OrganizationThresholdFields)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *OrganizationThresholdUpdateResponse) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



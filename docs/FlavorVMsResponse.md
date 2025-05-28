# FlavorVMsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlavorVms** | Pointer to [**[]FlavorVMFields**](FlavorVMFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlavorVMsResponse

`func NewFlavorVMsResponse() *FlavorVMsResponse`

NewFlavorVMsResponse instantiates a new FlavorVMsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorVMsResponseWithDefaults

`func NewFlavorVMsResponseWithDefaults() *FlavorVMsResponse`

NewFlavorVMsResponseWithDefaults instantiates a new FlavorVMsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavorVms

`func (o *FlavorVMsResponse) GetFlavorVms() []FlavorVMFields`

GetFlavorVms returns the FlavorVms field if non-nil, zero value otherwise.

### GetFlavorVmsOk

`func (o *FlavorVMsResponse) GetFlavorVmsOk() (*[]FlavorVMFields, bool)`

GetFlavorVmsOk returns a tuple with the FlavorVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorVms

`func (o *FlavorVMsResponse) SetFlavorVms(v []FlavorVMFields)`

SetFlavorVms sets FlavorVms field to given value.

### HasFlavorVms

`func (o *FlavorVMsResponse) HasFlavorVms() bool`

HasFlavorVms returns a boolean if a field has been set.

### GetMessage

`func (o *FlavorVMsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FlavorVMsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FlavorVMsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FlavorVMsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *FlavorVMsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlavorVMsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlavorVMsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlavorVMsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



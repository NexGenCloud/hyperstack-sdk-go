# FlavorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavor** | Pointer to [**FlavorFields**](FlavorFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlavorResponse

`func NewFlavorResponse() *FlavorResponse`

NewFlavorResponse instantiates a new FlavorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorResponseWithDefaults

`func NewFlavorResponseWithDefaults() *FlavorResponse`

NewFlavorResponseWithDefaults instantiates a new FlavorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *FlavorResponse) GetFlavor() FlavorFields`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *FlavorResponse) GetFlavorOk() (*FlavorFields, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *FlavorResponse) SetFlavor(v FlavorFields)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *FlavorResponse) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetMessage

`func (o *FlavorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FlavorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FlavorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FlavorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *FlavorResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlavorResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlavorResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlavorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



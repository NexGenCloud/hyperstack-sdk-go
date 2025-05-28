# FlavorListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]FlavorItemGetResponse**](FlavorItemGetResponse.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlavorListResponse

`func NewFlavorListResponse() *FlavorListResponse`

NewFlavorListResponse instantiates a new FlavorListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorListResponseWithDefaults

`func NewFlavorListResponseWithDefaults() *FlavorListResponse`

NewFlavorListResponseWithDefaults instantiates a new FlavorListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FlavorListResponse) GetData() []FlavorItemGetResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FlavorListResponse) GetDataOk() (*[]FlavorItemGetResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FlavorListResponse) SetData(v []FlavorItemGetResponse)`

SetData sets Data field to given value.

### HasData

`func (o *FlavorListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *FlavorListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FlavorListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FlavorListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FlavorListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *FlavorListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlavorListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlavorListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlavorListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



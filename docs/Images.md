# Images

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]ImageGetResponse**](ImageGetResponse.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewImages

`func NewImages() *Images`

NewImages instantiates a new Images object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagesWithDefaults

`func NewImagesWithDefaults() *Images`

NewImagesWithDefaults instantiates a new Images object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *Images) GetImages() []ImageGetResponse`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Images) GetImagesOk() (*[]ImageGetResponse, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Images) SetImages(v []ImageGetResponse)`

SetImages sets Images field to given value.

### HasImages

`func (o *Images) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetMessage

`func (o *Images) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Images) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Images) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Images) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *Images) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Images) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Images) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Images) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



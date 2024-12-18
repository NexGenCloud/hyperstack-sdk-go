# ImageLogos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logos** | Pointer to [**[]LogoGetResponse**](LogoGetResponse.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewImageLogos

`func NewImageLogos() *ImageLogos`

NewImageLogos instantiates a new ImageLogos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageLogosWithDefaults

`func NewImageLogosWithDefaults() *ImageLogos`

NewImageLogosWithDefaults instantiates a new ImageLogos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogos

`func (o *ImageLogos) GetLogos() []LogoGetResponse`

GetLogos returns the Logos field if non-nil, zero value otherwise.

### GetLogosOk

`func (o *ImageLogos) GetLogosOk() (*[]LogoGetResponse, bool)`

GetLogosOk returns a tuple with the Logos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogos

`func (o *ImageLogos) SetLogos(v []LogoGetResponse)`

SetLogos sets Logos field to given value.

### HasLogos

`func (o *ImageLogos) HasLogos() bool`

HasLogos returns a boolean if a field has been set.

### GetMessage

`func (o *ImageLogos) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImageLogos) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImageLogos) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ImageLogos) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ImageLogos) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageLogos) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageLogos) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImageLogos) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



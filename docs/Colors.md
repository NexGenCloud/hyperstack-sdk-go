# Colors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to [**PrimaryColor**](PrimaryColor.md) |  | [optional] 
**Secondary** | Pointer to [**SecondaryColor**](SecondaryColor.md) |  | [optional] 

## Methods

### NewColors

`func NewColors() *Colors`

NewColors instantiates a new Colors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColorsWithDefaults

`func NewColorsWithDefaults() *Colors`

NewColorsWithDefaults instantiates a new Colors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *Colors) GetPrimary() PrimaryColor`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Colors) GetPrimaryOk() (*PrimaryColor, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Colors) SetPrimary(v PrimaryColor)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *Colors) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSecondary

`func (o *Colors) GetSecondary() SecondaryColor`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *Colors) GetSecondaryOk() (*SecondaryColor, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *Colors) SetSecondary(v SecondaryColor)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *Colors) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



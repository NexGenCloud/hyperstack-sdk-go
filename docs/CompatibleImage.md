# CompatibleImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to **map[string]interface{}** | JSON constraints object | [optional] 
**ImageId** | Pointer to **int32** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**LinkType** | Pointer to **string** | Either &#39;hard&#39; or &#39;soft&#39; | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewCompatibleImage

`func NewCompatibleImage() *CompatibleImage`

NewCompatibleImage instantiates a new CompatibleImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompatibleImageWithDefaults

`func NewCompatibleImageWithDefaults() *CompatibleImage`

NewCompatibleImageWithDefaults instantiates a new CompatibleImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *CompatibleImage) GetConstraints() map[string]interface{}`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *CompatibleImage) GetConstraintsOk() (*map[string]interface{}, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *CompatibleImage) SetConstraints(v map[string]interface{})`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *CompatibleImage) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetImageId

`func (o *CompatibleImage) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CompatibleImage) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CompatibleImage) SetImageId(v int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CompatibleImage) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetImageName

`func (o *CompatibleImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CompatibleImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CompatibleImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *CompatibleImage) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetLinkType

`func (o *CompatibleImage) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *CompatibleImage) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *CompatibleImage) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *CompatibleImage) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetReason

`func (o *CompatibleImage) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CompatibleImage) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CompatibleImage) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CompatibleImage) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



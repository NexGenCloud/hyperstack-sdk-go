# ImageRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibleImages** | Pointer to [**[]CompatibleImage**](CompatibleImage.md) | List of images this flavor is allowed to launch, with link metadata | [optional] 
**HasImageRestrictions** | Pointer to **bool** | Whether the flavor is restricted to a set of images | [optional] 
**RestrictionType** | Pointer to **string** | Either &#39;hard&#39;, &#39;soft&#39;, or null if no restrictions | [optional] 

## Methods

### NewImageRestrictions

`func NewImageRestrictions() *ImageRestrictions`

NewImageRestrictions instantiates a new ImageRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageRestrictionsWithDefaults

`func NewImageRestrictionsWithDefaults() *ImageRestrictions`

NewImageRestrictionsWithDefaults instantiates a new ImageRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibleImages

`func (o *ImageRestrictions) GetCompatibleImages() []CompatibleImage`

GetCompatibleImages returns the CompatibleImages field if non-nil, zero value otherwise.

### GetCompatibleImagesOk

`func (o *ImageRestrictions) GetCompatibleImagesOk() (*[]CompatibleImage, bool)`

GetCompatibleImagesOk returns a tuple with the CompatibleImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleImages

`func (o *ImageRestrictions) SetCompatibleImages(v []CompatibleImage)`

SetCompatibleImages sets CompatibleImages field to given value.

### HasCompatibleImages

`func (o *ImageRestrictions) HasCompatibleImages() bool`

HasCompatibleImages returns a boolean if a field has been set.

### GetHasImageRestrictions

`func (o *ImageRestrictions) GetHasImageRestrictions() bool`

GetHasImageRestrictions returns the HasImageRestrictions field if non-nil, zero value otherwise.

### GetHasImageRestrictionsOk

`func (o *ImageRestrictions) GetHasImageRestrictionsOk() (*bool, bool)`

GetHasImageRestrictionsOk returns a tuple with the HasImageRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImageRestrictions

`func (o *ImageRestrictions) SetHasImageRestrictions(v bool)`

SetHasImageRestrictions sets HasImageRestrictions field to given value.

### HasHasImageRestrictions

`func (o *ImageRestrictions) HasHasImageRestrictions() bool`

HasHasImageRestrictions returns a boolean if a field has been set.

### GetRestrictionType

`func (o *ImageRestrictions) GetRestrictionType() string`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *ImageRestrictions) GetRestrictionTypeOk() (*string, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *ImageRestrictions) SetRestrictionType(v string)`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *ImageRestrictions) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



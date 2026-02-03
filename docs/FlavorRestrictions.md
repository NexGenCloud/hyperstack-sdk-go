# FlavorRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibleFlavors** | Pointer to [**[]CompatibleFlavor**](CompatibleFlavor.md) | List of compatible flavors with their link metadata | [optional] 
**HasFlavorRestrictions** | Pointer to **bool** | Whether the image has any flavor restrictions | [optional] 
**RestrictionType** | Pointer to **string** | Either &#39;hard&#39;, &#39;soft&#39;, or null if no restrictions | [optional] 

## Methods

### NewFlavorRestrictions

`func NewFlavorRestrictions() *FlavorRestrictions`

NewFlavorRestrictions instantiates a new FlavorRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorRestrictionsWithDefaults

`func NewFlavorRestrictionsWithDefaults() *FlavorRestrictions`

NewFlavorRestrictionsWithDefaults instantiates a new FlavorRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibleFlavors

`func (o *FlavorRestrictions) GetCompatibleFlavors() []CompatibleFlavor`

GetCompatibleFlavors returns the CompatibleFlavors field if non-nil, zero value otherwise.

### GetCompatibleFlavorsOk

`func (o *FlavorRestrictions) GetCompatibleFlavorsOk() (*[]CompatibleFlavor, bool)`

GetCompatibleFlavorsOk returns a tuple with the CompatibleFlavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleFlavors

`func (o *FlavorRestrictions) SetCompatibleFlavors(v []CompatibleFlavor)`

SetCompatibleFlavors sets CompatibleFlavors field to given value.

### HasCompatibleFlavors

`func (o *FlavorRestrictions) HasCompatibleFlavors() bool`

HasCompatibleFlavors returns a boolean if a field has been set.

### GetHasFlavorRestrictions

`func (o *FlavorRestrictions) GetHasFlavorRestrictions() bool`

GetHasFlavorRestrictions returns the HasFlavorRestrictions field if non-nil, zero value otherwise.

### GetHasFlavorRestrictionsOk

`func (o *FlavorRestrictions) GetHasFlavorRestrictionsOk() (*bool, bool)`

GetHasFlavorRestrictionsOk returns a tuple with the HasFlavorRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFlavorRestrictions

`func (o *FlavorRestrictions) SetHasFlavorRestrictions(v bool)`

SetHasFlavorRestrictions sets HasFlavorRestrictions field to given value.

### HasHasFlavorRestrictions

`func (o *FlavorRestrictions) HasHasFlavorRestrictions() bool`

HasHasFlavorRestrictions returns a boolean if a field has been set.

### GetRestrictionType

`func (o *FlavorRestrictions) GetRestrictionType() string`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *FlavorRestrictions) GetRestrictionTypeOk() (*string, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *FlavorRestrictions) SetRestrictionType(v string)`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *FlavorRestrictions) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



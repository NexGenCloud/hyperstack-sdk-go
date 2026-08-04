# EmailCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Childs** | Pointer to [**[]EmailCategoryChild**](EmailCategoryChild.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**OptedIn** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailCategory

`func NewEmailCategory() *EmailCategory`

NewEmailCategory instantiates a new EmailCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailCategoryWithDefaults

`func NewEmailCategoryWithDefaults() *EmailCategory`

NewEmailCategoryWithDefaults instantiates a new EmailCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChilds

`func (o *EmailCategory) GetChilds() []EmailCategoryChild`

GetChilds returns the Childs field if non-nil, zero value otherwise.

### GetChildsOk

`func (o *EmailCategory) GetChildsOk() (*[]EmailCategoryChild, bool)`

GetChildsOk returns a tuple with the Childs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChilds

`func (o *EmailCategory) SetChilds(v []EmailCategoryChild)`

SetChilds sets Childs field to given value.

### HasChilds

`func (o *EmailCategory) HasChilds() bool`

HasChilds returns a boolean if a field has been set.

### GetDescription

`func (o *EmailCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailCategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailCategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *EmailCategory) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EmailCategory) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EmailCategory) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EmailCategory) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIcon

`func (o *EmailCategory) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EmailCategory) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EmailCategory) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EmailCategory) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetOptedIn

`func (o *EmailCategory) GetOptedIn() bool`

GetOptedIn returns the OptedIn field if non-nil, zero value otherwise.

### GetOptedInOk

`func (o *EmailCategory) GetOptedInOk() (*bool, bool)`

GetOptedInOk returns a tuple with the OptedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedIn

`func (o *EmailCategory) SetOptedIn(v bool)`

SetOptedIn sets OptedIn field to given value.

### HasOptedIn

`func (o *EmailCategory) HasOptedIn() bool`

HasOptedIn returns a boolean if a field has been set.

### GetPosition

`func (o *EmailCategory) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EmailCategory) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EmailCategory) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *EmailCategory) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetRequired

`func (o *EmailCategory) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EmailCategory) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EmailCategory) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *EmailCategory) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSlug

`func (o *EmailCategory) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EmailCategory) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EmailCategory) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *EmailCategory) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



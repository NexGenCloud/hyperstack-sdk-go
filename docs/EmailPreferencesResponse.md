# EmailPreferencesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailCategories** | Pointer to [**[]EmailCategory**](EmailCategory.md) |  | [optional] 

## Methods

### NewEmailPreferencesResponse

`func NewEmailPreferencesResponse() *EmailPreferencesResponse`

NewEmailPreferencesResponse instantiates a new EmailPreferencesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailPreferencesResponseWithDefaults

`func NewEmailPreferencesResponseWithDefaults() *EmailPreferencesResponse`

NewEmailPreferencesResponseWithDefaults instantiates a new EmailPreferencesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailCategories

`func (o *EmailPreferencesResponse) GetEmailCategories() []EmailCategory`

GetEmailCategories returns the EmailCategories field if non-nil, zero value otherwise.

### GetEmailCategoriesOk

`func (o *EmailPreferencesResponse) GetEmailCategoriesOk() (*[]EmailCategory, bool)`

GetEmailCategoriesOk returns a tuple with the EmailCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCategories

`func (o *EmailPreferencesResponse) SetEmailCategories(v []EmailCategory)`

SetEmailCategories sets EmailCategories field to given value.

### HasEmailCategories

`func (o *EmailPreferencesResponse) HasEmailCategories() bool`

HasEmailCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



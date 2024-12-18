# GetOrganizationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**OrganizationFields**](OrganizationFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetOrganizationResponseModel

`func NewGetOrganizationResponseModel() *GetOrganizationResponseModel`

NewGetOrganizationResponseModel instantiates a new GetOrganizationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationResponseModelWithDefaults

`func NewGetOrganizationResponseModelWithDefaults() *GetOrganizationResponseModel`

NewGetOrganizationResponseModelWithDefaults instantiates a new GetOrganizationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetOrganizationResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetOrganizationResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetOrganizationResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetOrganizationResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganization

`func (o *GetOrganizationResponseModel) GetOrganization() OrganizationFields`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetOrganizationResponseModel) GetOrganizationOk() (*OrganizationFields, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetOrganizationResponseModel) SetOrganization(v OrganizationFields)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetOrganizationResponseModel) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



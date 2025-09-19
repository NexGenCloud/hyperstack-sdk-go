# UserOrganizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | Pointer to [**[]OrganizationFields**](OrganizationFields.md) |  | [optional] 

## Methods

### NewUserOrganizationsResponse

`func NewUserOrganizationsResponse() *UserOrganizationsResponse`

NewUserOrganizationsResponse instantiates a new UserOrganizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOrganizationsResponseWithDefaults

`func NewUserOrganizationsResponseWithDefaults() *UserOrganizationsResponse`

NewUserOrganizationsResponseWithDefaults instantiates a new UserOrganizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *UserOrganizationsResponse) GetOrganizations() []OrganizationFields`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *UserOrganizationsResponse) GetOrganizationsOk() (*[]OrganizationFields, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *UserOrganizationsResponse) SetOrganizations(v []OrganizationFields)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *UserOrganizationsResponse) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



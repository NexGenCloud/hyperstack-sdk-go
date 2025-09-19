# OrganizationUserResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] 
**LastLogin** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RbacRoles** | Pointer to [**[]RbacRoleField**](RbacRoleField.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationUserResponseModel

`func NewOrganizationUserResponseModel() *OrganizationUserResponseModel`

NewOrganizationUserResponseModel instantiates a new OrganizationUserResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserResponseModelWithDefaults

`func NewOrganizationUserResponseModelWithDefaults() *OrganizationUserResponseModel`

NewOrganizationUserResponseModelWithDefaults instantiates a new OrganizationUserResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationUserResponseModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationUserResponseModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationUserResponseModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationUserResponseModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *OrganizationUserResponseModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationUserResponseModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationUserResponseModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationUserResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJoinedAt

`func (o *OrganizationUserResponseModel) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrganizationUserResponseModel) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrganizationUserResponseModel) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *OrganizationUserResponseModel) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLastLogin

`func (o *OrganizationUserResponseModel) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *OrganizationUserResponseModel) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *OrganizationUserResponseModel) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *OrganizationUserResponseModel) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetName

`func (o *OrganizationUserResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationUserResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationUserResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationUserResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRbacRoles

`func (o *OrganizationUserResponseModel) GetRbacRoles() []RbacRoleField`

GetRbacRoles returns the RbacRoles field if non-nil, zero value otherwise.

### GetRbacRolesOk

`func (o *OrganizationUserResponseModel) GetRbacRolesOk() (*[]RbacRoleField, bool)`

GetRbacRolesOk returns a tuple with the RbacRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbacRoles

`func (o *OrganizationUserResponseModel) SetRbacRoles(v []RbacRoleField)`

SetRbacRoles sets RbacRoles field to given value.

### HasRbacRoles

`func (o *OrganizationUserResponseModel) HasRbacRoles() bool`

HasRbacRoles returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationUserResponseModel) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationUserResponseModel) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationUserResponseModel) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationUserResponseModel) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSub

`func (o *OrganizationUserResponseModel) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *OrganizationUserResponseModel) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *OrganizationUserResponseModel) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *OrganizationUserResponseModel) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetUsername

`func (o *OrganizationUserResponseModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OrganizationUserResponseModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OrganizationUserResponseModel) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OrganizationUserResponseModel) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



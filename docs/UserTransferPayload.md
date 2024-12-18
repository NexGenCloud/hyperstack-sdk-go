# UserTransferPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **int32** |  | 
**Role** | **string** |  | 

## Methods

### NewUserTransferPayload

`func NewUserTransferPayload(orgId int32, role string, ) *UserTransferPayload`

NewUserTransferPayload instantiates a new UserTransferPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTransferPayloadWithDefaults

`func NewUserTransferPayloadWithDefaults() *UserTransferPayload`

NewUserTransferPayloadWithDefaults instantiates a new UserTransferPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *UserTransferPayload) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UserTransferPayload) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UserTransferPayload) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetRole

`func (o *UserTransferPayload) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserTransferPayload) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserTransferPayload) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



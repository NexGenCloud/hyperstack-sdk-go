# OrganizationObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** |  | [optional] 
**Resources** | Pointer to [**[]InfrahubResourceObjectResponse**](InfrahubResourceObjectResponse.md) |  | [optional] 

## Methods

### NewOrganizationObjectResponse

`func NewOrganizationObjectResponse() *OrganizationObjectResponse`

NewOrganizationObjectResponse instantiates a new OrganizationObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationObjectResponseWithDefaults

`func NewOrganizationObjectResponseWithDefaults() *OrganizationObjectResponse`

NewOrganizationObjectResponseWithDefaults instantiates a new OrganizationObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrganizationObjectResponse) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrganizationObjectResponse) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrganizationObjectResponse) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrganizationObjectResponse) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetResources

`func (o *OrganizationObjectResponse) GetResources() []InfrahubResourceObjectResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *OrganizationObjectResponse) GetResourcesOk() (*[]InfrahubResourceObjectResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *OrganizationObjectResponse) SetResources(v []InfrahubResourceObjectResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *OrganizationObjectResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



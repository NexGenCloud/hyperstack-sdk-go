# ContractChangePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to [**[]FieldChange**](FieldChange.md) | List of field changes for &#39;updated&#39; type | [optional] 
**Id** | **int32** | The ID of the contract | 
**OrgId** | **int32** | The ORG ID of the contract | 
**Type** | **string** | Purpose of the change: created, deleted, expired, or updated | 

## Methods

### NewContractChangePayload

`func NewContractChangePayload(id int32, orgId int32, type_ string, ) *ContractChangePayload`

NewContractChangePayload instantiates a new ContractChangePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractChangePayloadWithDefaults

`func NewContractChangePayloadWithDefaults() *ContractChangePayload`

NewContractChangePayloadWithDefaults instantiates a new ContractChangePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *ContractChangePayload) GetChanges() []FieldChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ContractChangePayload) GetChangesOk() (*[]FieldChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ContractChangePayload) SetChanges(v []FieldChange)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *ContractChangePayload) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetId

`func (o *ContractChangePayload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractChangePayload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractChangePayload) SetId(v int32)`

SetId sets Id field to given value.


### GetOrgId

`func (o *ContractChangePayload) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ContractChangePayload) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ContractChangePayload) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetType

`func (o *ContractChangePayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractChangePayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractChangePayload) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



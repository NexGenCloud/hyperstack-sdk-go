# FirewallFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**FirewallEnvironmentFields**](FirewallEnvironmentFields.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewFirewallFields

`func NewFirewallFields() *FirewallFields`

NewFirewallFields instantiates a new FirewallFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallFieldsWithDefaults

`func NewFirewallFieldsWithDefaults() *FirewallFields`

NewFirewallFieldsWithDefaults instantiates a new FirewallFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FirewallFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *FirewallFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *FirewallFields) GetEnvironment() FirewallEnvironmentFields`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FirewallFields) GetEnvironmentOk() (*FirewallEnvironmentFields, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FirewallFields) SetEnvironment(v FirewallEnvironmentFields)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FirewallFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *FirewallFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FirewallFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *FirewallFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



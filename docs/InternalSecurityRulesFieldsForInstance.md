# InternalSecurityRulesFieldsForInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Ethertype** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**PortRangeMax** | Pointer to **int32** |  | [optional] 
**PortRangeMin** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**RemoteIpPrefix** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalSecurityRulesFieldsForInstance

`func NewInternalSecurityRulesFieldsForInstance() *InternalSecurityRulesFieldsForInstance`

NewInternalSecurityRulesFieldsForInstance instantiates a new InternalSecurityRulesFieldsForInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalSecurityRulesFieldsForInstanceWithDefaults

`func NewInternalSecurityRulesFieldsForInstanceWithDefaults() *InternalSecurityRulesFieldsForInstance`

NewInternalSecurityRulesFieldsForInstanceWithDefaults instantiates a new InternalSecurityRulesFieldsForInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *InternalSecurityRulesFieldsForInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InternalSecurityRulesFieldsForInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InternalSecurityRulesFieldsForInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InternalSecurityRulesFieldsForInstance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *InternalSecurityRulesFieldsForInstance) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InternalSecurityRulesFieldsForInstance) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InternalSecurityRulesFieldsForInstance) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InternalSecurityRulesFieldsForInstance) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEthertype

`func (o *InternalSecurityRulesFieldsForInstance) GetEthertype() string`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *InternalSecurityRulesFieldsForInstance) GetEthertypeOk() (*string, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *InternalSecurityRulesFieldsForInstance) SetEthertype(v string)`

SetEthertype sets Ethertype field to given value.

### HasEthertype

`func (o *InternalSecurityRulesFieldsForInstance) HasEthertype() bool`

HasEthertype returns a boolean if a field has been set.

### GetId

`func (o *InternalSecurityRulesFieldsForInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalSecurityRulesFieldsForInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalSecurityRulesFieldsForInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InternalSecurityRulesFieldsForInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortRangeMax

`func (o *InternalSecurityRulesFieldsForInstance) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *InternalSecurityRulesFieldsForInstance) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *InternalSecurityRulesFieldsForInstance) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *InternalSecurityRulesFieldsForInstance) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### GetPortRangeMin

`func (o *InternalSecurityRulesFieldsForInstance) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *InternalSecurityRulesFieldsForInstance) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *InternalSecurityRulesFieldsForInstance) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *InternalSecurityRulesFieldsForInstance) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### GetProtocol

`func (o *InternalSecurityRulesFieldsForInstance) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InternalSecurityRulesFieldsForInstance) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InternalSecurityRulesFieldsForInstance) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InternalSecurityRulesFieldsForInstance) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteIpPrefix

`func (o *InternalSecurityRulesFieldsForInstance) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *InternalSecurityRulesFieldsForInstance) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *InternalSecurityRulesFieldsForInstance) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *InternalSecurityRulesFieldsForInstance) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### GetStatus

`func (o *InternalSecurityRulesFieldsForInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalSecurityRulesFieldsForInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalSecurityRulesFieldsForInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalSecurityRulesFieldsForInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



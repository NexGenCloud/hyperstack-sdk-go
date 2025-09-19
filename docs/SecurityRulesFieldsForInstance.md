# SecurityRulesFieldsForInstance

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

### NewSecurityRulesFieldsForInstance

`func NewSecurityRulesFieldsForInstance() *SecurityRulesFieldsForInstance`

NewSecurityRulesFieldsForInstance instantiates a new SecurityRulesFieldsForInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRulesFieldsForInstanceWithDefaults

`func NewSecurityRulesFieldsForInstanceWithDefaults() *SecurityRulesFieldsForInstance`

NewSecurityRulesFieldsForInstanceWithDefaults instantiates a new SecurityRulesFieldsForInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SecurityRulesFieldsForInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityRulesFieldsForInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityRulesFieldsForInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityRulesFieldsForInstance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *SecurityRulesFieldsForInstance) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityRulesFieldsForInstance) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityRulesFieldsForInstance) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SecurityRulesFieldsForInstance) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEthertype

`func (o *SecurityRulesFieldsForInstance) GetEthertype() string`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *SecurityRulesFieldsForInstance) GetEthertypeOk() (*string, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *SecurityRulesFieldsForInstance) SetEthertype(v string)`

SetEthertype sets Ethertype field to given value.

### HasEthertype

`func (o *SecurityRulesFieldsForInstance) HasEthertype() bool`

HasEthertype returns a boolean if a field has been set.

### GetId

`func (o *SecurityRulesFieldsForInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityRulesFieldsForInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityRulesFieldsForInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityRulesFieldsForInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortRangeMax

`func (o *SecurityRulesFieldsForInstance) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *SecurityRulesFieldsForInstance) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *SecurityRulesFieldsForInstance) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *SecurityRulesFieldsForInstance) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### GetPortRangeMin

`func (o *SecurityRulesFieldsForInstance) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *SecurityRulesFieldsForInstance) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *SecurityRulesFieldsForInstance) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *SecurityRulesFieldsForInstance) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### GetProtocol

`func (o *SecurityRulesFieldsForInstance) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityRulesFieldsForInstance) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityRulesFieldsForInstance) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SecurityRulesFieldsForInstance) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteIpPrefix

`func (o *SecurityRulesFieldsForInstance) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *SecurityRulesFieldsForInstance) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *SecurityRulesFieldsForInstance) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *SecurityRulesFieldsForInstance) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityRulesFieldsForInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityRulesFieldsForInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityRulesFieldsForInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityRulesFieldsForInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



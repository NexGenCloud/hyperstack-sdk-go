# SecurityRulesFieldsforInstance

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

### NewSecurityRulesFieldsforInstance

`func NewSecurityRulesFieldsforInstance() *SecurityRulesFieldsforInstance`

NewSecurityRulesFieldsforInstance instantiates a new SecurityRulesFieldsforInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRulesFieldsforInstanceWithDefaults

`func NewSecurityRulesFieldsforInstanceWithDefaults() *SecurityRulesFieldsforInstance`

NewSecurityRulesFieldsforInstanceWithDefaults instantiates a new SecurityRulesFieldsforInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SecurityRulesFieldsforInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityRulesFieldsforInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityRulesFieldsforInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityRulesFieldsforInstance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *SecurityRulesFieldsforInstance) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityRulesFieldsforInstance) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityRulesFieldsforInstance) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SecurityRulesFieldsforInstance) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEthertype

`func (o *SecurityRulesFieldsforInstance) GetEthertype() string`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *SecurityRulesFieldsforInstance) GetEthertypeOk() (*string, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *SecurityRulesFieldsforInstance) SetEthertype(v string)`

SetEthertype sets Ethertype field to given value.

### HasEthertype

`func (o *SecurityRulesFieldsforInstance) HasEthertype() bool`

HasEthertype returns a boolean if a field has been set.

### GetId

`func (o *SecurityRulesFieldsforInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityRulesFieldsforInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityRulesFieldsforInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityRulesFieldsforInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortRangeMax

`func (o *SecurityRulesFieldsforInstance) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *SecurityRulesFieldsforInstance) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *SecurityRulesFieldsforInstance) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *SecurityRulesFieldsforInstance) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### GetPortRangeMin

`func (o *SecurityRulesFieldsforInstance) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *SecurityRulesFieldsforInstance) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *SecurityRulesFieldsforInstance) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *SecurityRulesFieldsforInstance) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### GetProtocol

`func (o *SecurityRulesFieldsforInstance) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityRulesFieldsforInstance) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityRulesFieldsforInstance) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SecurityRulesFieldsforInstance) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteIpPrefix

`func (o *SecurityRulesFieldsforInstance) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *SecurityRulesFieldsforInstance) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *SecurityRulesFieldsforInstance) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *SecurityRulesFieldsforInstance) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityRulesFieldsforInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityRulesFieldsforInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityRulesFieldsforInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityRulesFieldsforInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



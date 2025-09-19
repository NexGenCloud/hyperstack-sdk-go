# FirewallDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firewall** | Pointer to [**FirewallDetailFields**](FirewallDetailFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewFirewallDetailResponse

`func NewFirewallDetailResponse() *FirewallDetailResponse`

NewFirewallDetailResponse instantiates a new FirewallDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallDetailResponseWithDefaults

`func NewFirewallDetailResponseWithDefaults() *FirewallDetailResponse`

NewFirewallDetailResponseWithDefaults instantiates a new FirewallDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewall

`func (o *FirewallDetailResponse) GetFirewall() FirewallDetailFields`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *FirewallDetailResponse) GetFirewallOk() (*FirewallDetailFields, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *FirewallDetailResponse) SetFirewall(v FirewallDetailFields)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *FirewallDetailResponse) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetMessage

`func (o *FirewallDetailResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallDetailResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallDetailResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FirewallDetailResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *FirewallDetailResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallDetailResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallDetailResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FirewallsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Firewalls** | Pointer to [**[]FirewallDetailFields**](FirewallDetailFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewFirewallsListResponse

`func NewFirewallsListResponse() *FirewallsListResponse`

NewFirewallsListResponse instantiates a new FirewallsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallsListResponseWithDefaults

`func NewFirewallsListResponseWithDefaults() *FirewallsListResponse`

NewFirewallsListResponseWithDefaults instantiates a new FirewallsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FirewallsListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FirewallsListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FirewallsListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FirewallsListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFirewalls

`func (o *FirewallsListResponse) GetFirewalls() []FirewallDetailFields`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *FirewallsListResponse) GetFirewallsOk() (*[]FirewallDetailFields, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *FirewallsListResponse) SetFirewalls(v []FirewallDetailFields)`

SetFirewalls sets Firewalls field to given value.

### HasFirewalls

`func (o *FirewallsListResponse) HasFirewalls() bool`

HasFirewalls returns a boolean if a field has been set.

### GetMessage

`func (o *FirewallsListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallsListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallsListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FirewallsListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPage

`func (o *FirewallsListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FirewallsListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FirewallsListResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *FirewallsListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *FirewallsListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *FirewallsListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *FirewallsListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *FirewallsListResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStatus

`func (o *FirewallsListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallsListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallsListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallsListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



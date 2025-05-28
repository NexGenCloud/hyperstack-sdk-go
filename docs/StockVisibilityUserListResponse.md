# StockVisibilityUserListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Users** | Pointer to [**[]SingleVisibilityUserResponse**](SingleVisibilityUserResponse.md) |  | [optional] 

## Methods

### NewStockVisibilityUserListResponse

`func NewStockVisibilityUserListResponse() *StockVisibilityUserListResponse`

NewStockVisibilityUserListResponse instantiates a new StockVisibilityUserListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockVisibilityUserListResponseWithDefaults

`func NewStockVisibilityUserListResponseWithDefaults() *StockVisibilityUserListResponse`

NewStockVisibilityUserListResponseWithDefaults instantiates a new StockVisibilityUserListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *StockVisibilityUserListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StockVisibilityUserListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StockVisibilityUserListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StockVisibilityUserListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *StockVisibilityUserListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StockVisibilityUserListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StockVisibilityUserListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StockVisibilityUserListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsers

`func (o *StockVisibilityUserListResponse) GetUsers() []SingleVisibilityUserResponse`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *StockVisibilityUserListResponse) GetUsersOk() (*[]SingleVisibilityUserResponse, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *StockVisibilityUserListResponse) SetUsers(v []SingleVisibilityUserResponse)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *StockVisibilityUserListResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



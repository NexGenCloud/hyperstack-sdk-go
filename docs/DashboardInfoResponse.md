# DashboardInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to [**OverviewInfo**](OverviewInfo.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewDashboardInfoResponse

`func NewDashboardInfoResponse() *DashboardInfoResponse`

NewDashboardInfoResponse instantiates a new DashboardInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardInfoResponseWithDefaults

`func NewDashboardInfoResponseWithDefaults() *DashboardInfoResponse`

NewDashboardInfoResponseWithDefaults instantiates a new DashboardInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DashboardInfoResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DashboardInfoResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DashboardInfoResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DashboardInfoResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOverview

`func (o *DashboardInfoResponse) GetOverview() OverviewInfo`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *DashboardInfoResponse) GetOverviewOk() (*OverviewInfo, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *DashboardInfoResponse) SetOverview(v OverviewInfo)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *DashboardInfoResponse) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetStatus

`func (o *DashboardInfoResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DashboardInfoResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DashboardInfoResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DashboardInfoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



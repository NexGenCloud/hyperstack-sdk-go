# UserEnhancedMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallCommand** | Pointer to **string** | One-liner the user can run inside the VM to install the agent when enabling. Omitted when disabling. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**UserEnhancedMetricsResponseFields**](UserEnhancedMetricsResponseFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserEnhancedMetricsResponse

`func NewUserEnhancedMetricsResponse() *UserEnhancedMetricsResponse`

NewUserEnhancedMetricsResponse instantiates a new UserEnhancedMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEnhancedMetricsResponseWithDefaults

`func NewUserEnhancedMetricsResponseWithDefaults() *UserEnhancedMetricsResponse`

NewUserEnhancedMetricsResponseWithDefaults instantiates a new UserEnhancedMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallCommand

`func (o *UserEnhancedMetricsResponse) GetInstallCommand() string`

GetInstallCommand returns the InstallCommand field if non-nil, zero value otherwise.

### GetInstallCommandOk

`func (o *UserEnhancedMetricsResponse) GetInstallCommandOk() (*string, bool)`

GetInstallCommandOk returns a tuple with the InstallCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallCommand

`func (o *UserEnhancedMetricsResponse) SetInstallCommand(v string)`

SetInstallCommand sets InstallCommand field to given value.

### HasInstallCommand

`func (o *UserEnhancedMetricsResponse) HasInstallCommand() bool`

HasInstallCommand returns a boolean if a field has been set.

### GetMessage

`func (o *UserEnhancedMetricsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserEnhancedMetricsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserEnhancedMetricsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserEnhancedMetricsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetrics

`func (o *UserEnhancedMetricsResponse) GetMetrics() UserEnhancedMetricsResponseFields`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *UserEnhancedMetricsResponse) GetMetricsOk() (*UserEnhancedMetricsResponseFields, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *UserEnhancedMetricsResponse) SetMetrics(v UserEnhancedMetricsResponseFields)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *UserEnhancedMetricsResponse) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetStatus

`func (o *UserEnhancedMetricsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserEnhancedMetricsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserEnhancedMetricsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserEnhancedMetricsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



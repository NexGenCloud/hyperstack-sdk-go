# BetaAccessStatusItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramName** | Pointer to **string** | Name of the beta program | [optional] 
**RequestDate** | Pointer to **time.Time** | When the request was made | [optional] 
**Status** | Pointer to **string** | Status of the request (requested, approved, denied, revoked) | [optional] 

## Methods

### NewBetaAccessStatusItem

`func NewBetaAccessStatusItem() *BetaAccessStatusItem`

NewBetaAccessStatusItem instantiates a new BetaAccessStatusItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAccessStatusItemWithDefaults

`func NewBetaAccessStatusItemWithDefaults() *BetaAccessStatusItem`

NewBetaAccessStatusItemWithDefaults instantiates a new BetaAccessStatusItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramName

`func (o *BetaAccessStatusItem) GetProgramName() string`

GetProgramName returns the ProgramName field if non-nil, zero value otherwise.

### GetProgramNameOk

`func (o *BetaAccessStatusItem) GetProgramNameOk() (*string, bool)`

GetProgramNameOk returns a tuple with the ProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramName

`func (o *BetaAccessStatusItem) SetProgramName(v string)`

SetProgramName sets ProgramName field to given value.

### HasProgramName

`func (o *BetaAccessStatusItem) HasProgramName() bool`

HasProgramName returns a boolean if a field has been set.

### GetRequestDate

`func (o *BetaAccessStatusItem) GetRequestDate() time.Time`

GetRequestDate returns the RequestDate field if non-nil, zero value otherwise.

### GetRequestDateOk

`func (o *BetaAccessStatusItem) GetRequestDateOk() (*time.Time, bool)`

GetRequestDateOk returns a tuple with the RequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDate

`func (o *BetaAccessStatusItem) SetRequestDate(v time.Time)`

SetRequestDate sets RequestDate field to given value.

### HasRequestDate

`func (o *BetaAccessStatusItem) HasRequestDate() bool`

HasRequestDate returns a boolean if a field has been set.

### GetStatus

`func (o *BetaAccessStatusItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BetaAccessStatusItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BetaAccessStatusItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BetaAccessStatusItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BetaAccessRequestFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | When the request was made | [optional] 
**Id** | Pointer to **int32** | Unique identifier for the request | [optional] 
**ProgramId** | Pointer to **int32** | ID of the beta program | [optional] 
**ProgramName** | Pointer to **string** | Name of the beta program | [optional] 
**Status** | Pointer to **string** | Status of the request | [optional] 
**UserId** | Pointer to **int32** | ID of the user who made the request | [optional] 

## Methods

### NewBetaAccessRequestFields

`func NewBetaAccessRequestFields() *BetaAccessRequestFields`

NewBetaAccessRequestFields instantiates a new BetaAccessRequestFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAccessRequestFieldsWithDefaults

`func NewBetaAccessRequestFieldsWithDefaults() *BetaAccessRequestFields`

NewBetaAccessRequestFieldsWithDefaults instantiates a new BetaAccessRequestFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BetaAccessRequestFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BetaAccessRequestFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BetaAccessRequestFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BetaAccessRequestFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *BetaAccessRequestFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BetaAccessRequestFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BetaAccessRequestFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BetaAccessRequestFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProgramId

`func (o *BetaAccessRequestFields) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *BetaAccessRequestFields) GetProgramIdOk() (*int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *BetaAccessRequestFields) SetProgramId(v int32)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *BetaAccessRequestFields) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetProgramName

`func (o *BetaAccessRequestFields) GetProgramName() string`

GetProgramName returns the ProgramName field if non-nil, zero value otherwise.

### GetProgramNameOk

`func (o *BetaAccessRequestFields) GetProgramNameOk() (*string, bool)`

GetProgramNameOk returns a tuple with the ProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramName

`func (o *BetaAccessRequestFields) SetProgramName(v string)`

SetProgramName sets ProgramName field to given value.

### HasProgramName

`func (o *BetaAccessRequestFields) HasProgramName() bool`

HasProgramName returns a boolean if a field has been set.

### GetStatus

`func (o *BetaAccessRequestFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BetaAccessRequestFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BetaAccessRequestFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BetaAccessRequestFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *BetaAccessRequestFields) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BetaAccessRequestFields) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BetaAccessRequestFields) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BetaAccessRequestFields) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



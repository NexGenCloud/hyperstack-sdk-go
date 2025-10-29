# VoucherRedeemResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response message | 
**Status** | **bool** | Success status of the operation | 
**Voucher** | Pointer to [**Voucher**](Voucher.md) | Redeemed voucher details | [optional] 

## Methods

### NewVoucherRedeemResponseSchema

`func NewVoucherRedeemResponseSchema(message string, status bool, ) *VoucherRedeemResponseSchema`

NewVoucherRedeemResponseSchema instantiates a new VoucherRedeemResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherRedeemResponseSchemaWithDefaults

`func NewVoucherRedeemResponseSchemaWithDefaults() *VoucherRedeemResponseSchema`

NewVoucherRedeemResponseSchemaWithDefaults instantiates a new VoucherRedeemResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VoucherRedeemResponseSchema) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VoucherRedeemResponseSchema) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VoucherRedeemResponseSchema) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *VoucherRedeemResponseSchema) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoucherRedeemResponseSchema) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoucherRedeemResponseSchema) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetVoucher

`func (o *VoucherRedeemResponseSchema) GetVoucher() Voucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *VoucherRedeemResponseSchema) GetVoucherOk() (*Voucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *VoucherRedeemResponseSchema) SetVoucher(v Voucher)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *VoucherRedeemResponseSchema) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



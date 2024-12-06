/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TransferSubscriptionsRequest is the corresponding interface of TransferSubscriptionsRequest
type TransferSubscriptionsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetSubscriptionIds returns SubscriptionIds (property field)
	GetSubscriptionIds() []uint32
	// GetSendInitialValues returns SendInitialValues (property field)
	GetSendInitialValues() bool
	// IsTransferSubscriptionsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTransferSubscriptionsRequest()
	// CreateBuilder creates a TransferSubscriptionsRequestBuilder
	CreateTransferSubscriptionsRequestBuilder() TransferSubscriptionsRequestBuilder
}

// _TransferSubscriptionsRequest is the data-structure of this message
type _TransferSubscriptionsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader     RequestHeader
	SubscriptionIds   []uint32
	SendInitialValues bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ TransferSubscriptionsRequest = (*_TransferSubscriptionsRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TransferSubscriptionsRequest)(nil)

// NewTransferSubscriptionsRequest factory function for _TransferSubscriptionsRequest
func NewTransferSubscriptionsRequest(requestHeader RequestHeader, subscriptionIds []uint32, sendInitialValues bool) *_TransferSubscriptionsRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for TransferSubscriptionsRequest must not be nil")
	}
	_result := &_TransferSubscriptionsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionIds:                   subscriptionIds,
		SendInitialValues:                 sendInitialValues,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TransferSubscriptionsRequestBuilder is a builder for TransferSubscriptionsRequest
type TransferSubscriptionsRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, subscriptionIds []uint32, sendInitialValues bool) TransferSubscriptionsRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) TransferSubscriptionsRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) TransferSubscriptionsRequestBuilder
	// WithSubscriptionIds adds SubscriptionIds (property field)
	WithSubscriptionIds(...uint32) TransferSubscriptionsRequestBuilder
	// WithSendInitialValues adds SendInitialValues (property field)
	WithSendInitialValues(bool) TransferSubscriptionsRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the TransferSubscriptionsRequest or returns an error if something is wrong
	Build() (TransferSubscriptionsRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TransferSubscriptionsRequest
}

// NewTransferSubscriptionsRequestBuilder() creates a TransferSubscriptionsRequestBuilder
func NewTransferSubscriptionsRequestBuilder() TransferSubscriptionsRequestBuilder {
	return &_TransferSubscriptionsRequestBuilder{_TransferSubscriptionsRequest: new(_TransferSubscriptionsRequest)}
}

type _TransferSubscriptionsRequestBuilder struct {
	*_TransferSubscriptionsRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (TransferSubscriptionsRequestBuilder) = (*_TransferSubscriptionsRequestBuilder)(nil)

func (b *_TransferSubscriptionsRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._TransferSubscriptionsRequest
}

func (b *_TransferSubscriptionsRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, subscriptionIds []uint32, sendInitialValues bool) TransferSubscriptionsRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithSubscriptionIds(subscriptionIds...).WithSendInitialValues(sendInitialValues)
}

func (b *_TransferSubscriptionsRequestBuilder) WithRequestHeader(requestHeader RequestHeader) TransferSubscriptionsRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_TransferSubscriptionsRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) TransferSubscriptionsRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateRequestHeaderBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestHeaderBuilder failed"))
	}
	return b
}

func (b *_TransferSubscriptionsRequestBuilder) WithSubscriptionIds(subscriptionIds ...uint32) TransferSubscriptionsRequestBuilder {
	b.SubscriptionIds = subscriptionIds
	return b
}

func (b *_TransferSubscriptionsRequestBuilder) WithSendInitialValues(sendInitialValues bool) TransferSubscriptionsRequestBuilder {
	b.SendInitialValues = sendInitialValues
	return b
}

func (b *_TransferSubscriptionsRequestBuilder) Build() (TransferSubscriptionsRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TransferSubscriptionsRequest.deepCopy(), nil
}

func (b *_TransferSubscriptionsRequestBuilder) MustBuild() TransferSubscriptionsRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TransferSubscriptionsRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_TransferSubscriptionsRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_TransferSubscriptionsRequestBuilder) DeepCopy() any {
	_copy := b.CreateTransferSubscriptionsRequestBuilder().(*_TransferSubscriptionsRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTransferSubscriptionsRequestBuilder creates a TransferSubscriptionsRequestBuilder
func (b *_TransferSubscriptionsRequest) CreateTransferSubscriptionsRequestBuilder() TransferSubscriptionsRequestBuilder {
	if b == nil {
		return NewTransferSubscriptionsRequestBuilder()
	}
	return &_TransferSubscriptionsRequestBuilder{_TransferSubscriptionsRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransferSubscriptionsRequest) GetExtensionId() int32 {
	return int32(841)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransferSubscriptionsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransferSubscriptionsRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_TransferSubscriptionsRequest) GetSubscriptionIds() []uint32 {
	return m.SubscriptionIds
}

func (m *_TransferSubscriptionsRequest) GetSendInitialValues() bool {
	return m.SendInitialValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTransferSubscriptionsRequest(structType any) TransferSubscriptionsRequest {
	if casted, ok := structType.(TransferSubscriptionsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TransferSubscriptionsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TransferSubscriptionsRequest) GetTypeName() string {
	return "TransferSubscriptionsRequest"
}

func (m *_TransferSubscriptionsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfSubscriptionIds)
	lengthInBits += 32

	// Array field
	if len(m.SubscriptionIds) > 0 {
		lengthInBits += 32 * uint16(len(m.SubscriptionIds))
	}

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (sendInitialValues)
	lengthInBits += 1

	return lengthInBits
}

func (m *_TransferSubscriptionsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TransferSubscriptionsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__transferSubscriptionsRequest TransferSubscriptionsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TransferSubscriptionsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransferSubscriptionsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfSubscriptionIds, err := ReadImplicitField[int32](ctx, "noOfSubscriptionIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSubscriptionIds' field"))
	}
	_ = noOfSubscriptionIds

	subscriptionIds, err := ReadCountArrayField[uint32](ctx, "subscriptionIds", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfSubscriptionIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionIds' field"))
	}
	m.SubscriptionIds = subscriptionIds

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	sendInitialValues, err := ReadSimpleField(ctx, "sendInitialValues", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sendInitialValues' field"))
	}
	m.SendInitialValues = sendInitialValues

	if closeErr := readBuffer.CloseContext("TransferSubscriptionsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransferSubscriptionsRequest")
	}

	return m, nil
}

func (m *_TransferSubscriptionsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransferSubscriptionsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransferSubscriptionsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransferSubscriptionsRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}
		noOfSubscriptionIds := int32(utils.InlineIf(bool((m.GetSubscriptionIds()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSubscriptionIds()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSubscriptionIds", noOfSubscriptionIds, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSubscriptionIds' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "subscriptionIds", m.GetSubscriptionIds(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionIds' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "sendInitialValues", m.GetSendInitialValues(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sendInitialValues' field")
		}

		if popErr := writeBuffer.PopContext("TransferSubscriptionsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransferSubscriptionsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TransferSubscriptionsRequest) IsTransferSubscriptionsRequest() {}

func (m *_TransferSubscriptionsRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TransferSubscriptionsRequest) deepCopy() *_TransferSubscriptionsRequest {
	if m == nil {
		return nil
	}
	_TransferSubscriptionsRequestCopy := &_TransferSubscriptionsRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		utils.DeepCopySlice[uint32, uint32](m.SubscriptionIds),
		m.SendInitialValues,
		m.reservedField0,
	}
	_TransferSubscriptionsRequestCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _TransferSubscriptionsRequestCopy
}

func (m *_TransferSubscriptionsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}

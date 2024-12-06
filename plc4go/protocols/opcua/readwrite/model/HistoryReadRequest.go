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

// HistoryReadRequest is the corresponding interface of HistoryReadRequest
type HistoryReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetHistoryReadDetails returns HistoryReadDetails (property field)
	GetHistoryReadDetails() ExtensionObject
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetReleaseContinuationPoints returns ReleaseContinuationPoints (property field)
	GetReleaseContinuationPoints() bool
	// GetNodesToRead returns NodesToRead (property field)
	GetNodesToRead() []HistoryReadValueId
	// IsHistoryReadRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryReadRequest()
	// CreateBuilder creates a HistoryReadRequestBuilder
	CreateHistoryReadRequestBuilder() HistoryReadRequestBuilder
}

// _HistoryReadRequest is the data-structure of this message
type _HistoryReadRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader             RequestHeader
	HistoryReadDetails        ExtensionObject
	TimestampsToReturn        TimestampsToReturn
	ReleaseContinuationPoints bool
	NodesToRead               []HistoryReadValueId
	// Reserved Fields
	reservedField0 *uint8
}

var _ HistoryReadRequest = (*_HistoryReadRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryReadRequest)(nil)

// NewHistoryReadRequest factory function for _HistoryReadRequest
func NewHistoryReadRequest(requestHeader RequestHeader, historyReadDetails ExtensionObject, timestampsToReturn TimestampsToReturn, releaseContinuationPoints bool, nodesToRead []HistoryReadValueId) *_HistoryReadRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for HistoryReadRequest must not be nil")
	}
	if historyReadDetails == nil {
		panic("historyReadDetails of type ExtensionObject for HistoryReadRequest must not be nil")
	}
	_result := &_HistoryReadRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		HistoryReadDetails:                historyReadDetails,
		TimestampsToReturn:                timestampsToReturn,
		ReleaseContinuationPoints:         releaseContinuationPoints,
		NodesToRead:                       nodesToRead,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryReadRequestBuilder is a builder for HistoryReadRequest
type HistoryReadRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, historyReadDetails ExtensionObject, timestampsToReturn TimestampsToReturn, releaseContinuationPoints bool, nodesToRead []HistoryReadValueId) HistoryReadRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) HistoryReadRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) HistoryReadRequestBuilder
	// WithHistoryReadDetails adds HistoryReadDetails (property field)
	WithHistoryReadDetails(ExtensionObject) HistoryReadRequestBuilder
	// WithHistoryReadDetailsBuilder adds HistoryReadDetails (property field) which is build by the builder
	WithHistoryReadDetailsBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) HistoryReadRequestBuilder
	// WithTimestampsToReturn adds TimestampsToReturn (property field)
	WithTimestampsToReturn(TimestampsToReturn) HistoryReadRequestBuilder
	// WithReleaseContinuationPoints adds ReleaseContinuationPoints (property field)
	WithReleaseContinuationPoints(bool) HistoryReadRequestBuilder
	// WithNodesToRead adds NodesToRead (property field)
	WithNodesToRead(...HistoryReadValueId) HistoryReadRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the HistoryReadRequest or returns an error if something is wrong
	Build() (HistoryReadRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryReadRequest
}

// NewHistoryReadRequestBuilder() creates a HistoryReadRequestBuilder
func NewHistoryReadRequestBuilder() HistoryReadRequestBuilder {
	return &_HistoryReadRequestBuilder{_HistoryReadRequest: new(_HistoryReadRequest)}
}

type _HistoryReadRequestBuilder struct {
	*_HistoryReadRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryReadRequestBuilder) = (*_HistoryReadRequestBuilder)(nil)

func (b *_HistoryReadRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._HistoryReadRequest
}

func (b *_HistoryReadRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, historyReadDetails ExtensionObject, timestampsToReturn TimestampsToReturn, releaseContinuationPoints bool, nodesToRead []HistoryReadValueId) HistoryReadRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithHistoryReadDetails(historyReadDetails).WithTimestampsToReturn(timestampsToReturn).WithReleaseContinuationPoints(releaseContinuationPoints).WithNodesToRead(nodesToRead...)
}

func (b *_HistoryReadRequestBuilder) WithRequestHeader(requestHeader RequestHeader) HistoryReadRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_HistoryReadRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) HistoryReadRequestBuilder {
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

func (b *_HistoryReadRequestBuilder) WithHistoryReadDetails(historyReadDetails ExtensionObject) HistoryReadRequestBuilder {
	b.HistoryReadDetails = historyReadDetails
	return b
}

func (b *_HistoryReadRequestBuilder) WithHistoryReadDetailsBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) HistoryReadRequestBuilder {
	builder := builderSupplier(b.HistoryReadDetails.CreateExtensionObjectBuilder())
	var err error
	b.HistoryReadDetails, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_HistoryReadRequestBuilder) WithTimestampsToReturn(timestampsToReturn TimestampsToReturn) HistoryReadRequestBuilder {
	b.TimestampsToReturn = timestampsToReturn
	return b
}

func (b *_HistoryReadRequestBuilder) WithReleaseContinuationPoints(releaseContinuationPoints bool) HistoryReadRequestBuilder {
	b.ReleaseContinuationPoints = releaseContinuationPoints
	return b
}

func (b *_HistoryReadRequestBuilder) WithNodesToRead(nodesToRead ...HistoryReadValueId) HistoryReadRequestBuilder {
	b.NodesToRead = nodesToRead
	return b
}

func (b *_HistoryReadRequestBuilder) Build() (HistoryReadRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.HistoryReadDetails == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'historyReadDetails' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryReadRequest.deepCopy(), nil
}

func (b *_HistoryReadRequestBuilder) MustBuild() HistoryReadRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HistoryReadRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_HistoryReadRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryReadRequestBuilder) DeepCopy() any {
	_copy := b.CreateHistoryReadRequestBuilder().(*_HistoryReadRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryReadRequestBuilder creates a HistoryReadRequestBuilder
func (b *_HistoryReadRequest) CreateHistoryReadRequestBuilder() HistoryReadRequestBuilder {
	if b == nil {
		return NewHistoryReadRequestBuilder()
	}
	return &_HistoryReadRequestBuilder{_HistoryReadRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryReadRequest) GetExtensionId() int32 {
	return int32(664)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryReadRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryReadRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_HistoryReadRequest) GetHistoryReadDetails() ExtensionObject {
	return m.HistoryReadDetails
}

func (m *_HistoryReadRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_HistoryReadRequest) GetReleaseContinuationPoints() bool {
	return m.ReleaseContinuationPoints
}

func (m *_HistoryReadRequest) GetNodesToRead() []HistoryReadValueId {
	return m.NodesToRead
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryReadRequest(structType any) HistoryReadRequest {
	if casted, ok := structType.(HistoryReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryReadRequest) GetTypeName() string {
	return "HistoryReadRequest"
}

func (m *_HistoryReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (historyReadDetails)
	lengthInBits += m.HistoryReadDetails.GetLengthInBits(ctx)

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (releaseContinuationPoints)
	lengthInBits += 1

	// Implicit Field (noOfNodesToRead)
	lengthInBits += 32

	// Array field
	if len(m.NodesToRead) > 0 {
		for _curItem, element := range m.NodesToRead {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToRead), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryReadRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__historyReadRequest HistoryReadRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	historyReadDetails, err := ReadSimpleField[ExtensionObject](ctx, "historyReadDetails", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'historyReadDetails' field"))
	}
	m.HistoryReadDetails = historyReadDetails

	timestampsToReturn, err := ReadEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", ReadEnum(TimestampsToReturnByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestampsToReturn' field"))
	}
	m.TimestampsToReturn = timestampsToReturn

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	releaseContinuationPoints, err := ReadSimpleField(ctx, "releaseContinuationPoints", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'releaseContinuationPoints' field"))
	}
	m.ReleaseContinuationPoints = releaseContinuationPoints

	noOfNodesToRead, err := ReadImplicitField[int32](ctx, "noOfNodesToRead", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToRead' field"))
	}
	_ = noOfNodesToRead

	nodesToRead, err := ReadCountArrayField[HistoryReadValueId](ctx, "nodesToRead", ReadComplex[HistoryReadValueId](ExtensionObjectDefinitionParseWithBufferProducer[HistoryReadValueId]((int32)(int32(637))), readBuffer), uint64(noOfNodesToRead))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToRead' field"))
	}
	m.NodesToRead = nodesToRead

	if closeErr := readBuffer.CloseContext("HistoryReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryReadRequest")
	}

	return m, nil
}

func (m *_HistoryReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryReadRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "historyReadDetails", m.GetHistoryReadDetails(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'historyReadDetails' field")
		}

		if err := WriteSimpleEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", m.GetTimestampsToReturn(), WriteEnum[TimestampsToReturn, uint32](TimestampsToReturn.GetValue, TimestampsToReturn.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'timestampsToReturn' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "releaseContinuationPoints", m.GetReleaseContinuationPoints(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'releaseContinuationPoints' field")
		}
		noOfNodesToRead := int32(utils.InlineIf(bool((m.GetNodesToRead()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetNodesToRead()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfNodesToRead", noOfNodesToRead, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToRead' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToRead", m.GetNodesToRead(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToRead' field")
		}

		if popErr := writeBuffer.PopContext("HistoryReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryReadRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryReadRequest) IsHistoryReadRequest() {}

func (m *_HistoryReadRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryReadRequest) deepCopy() *_HistoryReadRequest {
	if m == nil {
		return nil
	}
	_HistoryReadRequestCopy := &_HistoryReadRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		utils.DeepCopy[ExtensionObject](m.HistoryReadDetails),
		m.TimestampsToReturn,
		m.ReleaseContinuationPoints,
		utils.DeepCopySlice[HistoryReadValueId, HistoryReadValueId](m.NodesToRead),
		m.reservedField0,
	}
	_HistoryReadRequestCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryReadRequestCopy
}

func (m *_HistoryReadRequest) String() string {
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
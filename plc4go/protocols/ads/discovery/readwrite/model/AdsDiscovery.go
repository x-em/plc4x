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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AdsDiscovery_HEADER uint32 = 0x71146603

// AdsDiscovery is the corresponding interface of AdsDiscovery
type AdsDiscovery interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetRequestId returns RequestId (property field)
	GetRequestId() uint32
	// GetOperation returns Operation (property field)
	GetOperation() Operation
	// GetAmsNetId returns AmsNetId (property field)
	GetAmsNetId() AmsNetId
	// GetPortNumber returns PortNumber (property field)
	GetPortNumber() AdsPortNumbers
	// GetBlocks returns Blocks (property field)
	GetBlocks() []AdsDiscoveryBlock
	// IsAdsDiscovery is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscovery()
	// CreateBuilder creates a AdsDiscoveryBuilder
	CreateAdsDiscoveryBuilder() AdsDiscoveryBuilder
}

// _AdsDiscovery is the data-structure of this message
type _AdsDiscovery struct {
	RequestId  uint32
	Operation  Operation
	AmsNetId   AmsNetId
	PortNumber AdsPortNumbers
	Blocks     []AdsDiscoveryBlock
}

var _ AdsDiscovery = (*_AdsDiscovery)(nil)

// NewAdsDiscovery factory function for _AdsDiscovery
func NewAdsDiscovery(requestId uint32, operation Operation, amsNetId AmsNetId, portNumber AdsPortNumbers, blocks []AdsDiscoveryBlock) *_AdsDiscovery {
	if amsNetId == nil {
		panic("amsNetId of type AmsNetId for AdsDiscovery must not be nil")
	}
	return &_AdsDiscovery{RequestId: requestId, Operation: operation, AmsNetId: amsNetId, PortNumber: portNumber, Blocks: blocks}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsDiscoveryBuilder is a builder for AdsDiscovery
type AdsDiscoveryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestId uint32, operation Operation, amsNetId AmsNetId, portNumber AdsPortNumbers, blocks []AdsDiscoveryBlock) AdsDiscoveryBuilder
	// WithRequestId adds RequestId (property field)
	WithRequestId(uint32) AdsDiscoveryBuilder
	// WithOperation adds Operation (property field)
	WithOperation(Operation) AdsDiscoveryBuilder
	// WithAmsNetId adds AmsNetId (property field)
	WithAmsNetId(AmsNetId) AdsDiscoveryBuilder
	// WithAmsNetIdBuilder adds AmsNetId (property field) which is build by the builder
	WithAmsNetIdBuilder(func(AmsNetIdBuilder) AmsNetIdBuilder) AdsDiscoveryBuilder
	// WithPortNumber adds PortNumber (property field)
	WithPortNumber(AdsPortNumbers) AdsDiscoveryBuilder
	// WithBlocks adds Blocks (property field)
	WithBlocks(...AdsDiscoveryBlock) AdsDiscoveryBuilder
	// Build builds the AdsDiscovery or returns an error if something is wrong
	Build() (AdsDiscovery, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsDiscovery
}

// NewAdsDiscoveryBuilder() creates a AdsDiscoveryBuilder
func NewAdsDiscoveryBuilder() AdsDiscoveryBuilder {
	return &_AdsDiscoveryBuilder{_AdsDiscovery: new(_AdsDiscovery)}
}

type _AdsDiscoveryBuilder struct {
	*_AdsDiscovery

	err *utils.MultiError
}

var _ (AdsDiscoveryBuilder) = (*_AdsDiscoveryBuilder)(nil)

func (b *_AdsDiscoveryBuilder) WithMandatoryFields(requestId uint32, operation Operation, amsNetId AmsNetId, portNumber AdsPortNumbers, blocks []AdsDiscoveryBlock) AdsDiscoveryBuilder {
	return b.WithRequestId(requestId).WithOperation(operation).WithAmsNetId(amsNetId).WithPortNumber(portNumber).WithBlocks(blocks...)
}

func (b *_AdsDiscoveryBuilder) WithRequestId(requestId uint32) AdsDiscoveryBuilder {
	b.RequestId = requestId
	return b
}

func (b *_AdsDiscoveryBuilder) WithOperation(operation Operation) AdsDiscoveryBuilder {
	b.Operation = operation
	return b
}

func (b *_AdsDiscoveryBuilder) WithAmsNetId(amsNetId AmsNetId) AdsDiscoveryBuilder {
	b.AmsNetId = amsNetId
	return b
}

func (b *_AdsDiscoveryBuilder) WithAmsNetIdBuilder(builderSupplier func(AmsNetIdBuilder) AmsNetIdBuilder) AdsDiscoveryBuilder {
	builder := builderSupplier(b.AmsNetId.CreateAmsNetIdBuilder())
	var err error
	b.AmsNetId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AmsNetIdBuilder failed"))
	}
	return b
}

func (b *_AdsDiscoveryBuilder) WithPortNumber(portNumber AdsPortNumbers) AdsDiscoveryBuilder {
	b.PortNumber = portNumber
	return b
}

func (b *_AdsDiscoveryBuilder) WithBlocks(blocks ...AdsDiscoveryBlock) AdsDiscoveryBuilder {
	b.Blocks = blocks
	return b
}

func (b *_AdsDiscoveryBuilder) Build() (AdsDiscovery, error) {
	if b.AmsNetId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'amsNetId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsDiscovery.deepCopy(), nil
}

func (b *_AdsDiscoveryBuilder) MustBuild() AdsDiscovery {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsDiscoveryBuilder) DeepCopy() any {
	_copy := b.CreateAdsDiscoveryBuilder().(*_AdsDiscoveryBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsDiscoveryBuilder creates a AdsDiscoveryBuilder
func (b *_AdsDiscovery) CreateAdsDiscoveryBuilder() AdsDiscoveryBuilder {
	if b == nil {
		return NewAdsDiscoveryBuilder()
	}
	return &_AdsDiscoveryBuilder{_AdsDiscovery: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscovery) GetRequestId() uint32 {
	return m.RequestId
}

func (m *_AdsDiscovery) GetOperation() Operation {
	return m.Operation
}

func (m *_AdsDiscovery) GetAmsNetId() AmsNetId {
	return m.AmsNetId
}

func (m *_AdsDiscovery) GetPortNumber() AdsPortNumbers {
	return m.PortNumber
}

func (m *_AdsDiscovery) GetBlocks() []AdsDiscoveryBlock {
	return m.Blocks
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsDiscovery) GetHeader() uint32 {
	return AdsDiscovery_HEADER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscovery(structType any) AdsDiscovery {
	if casted, ok := structType.(AdsDiscovery); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscovery); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscovery) GetTypeName() string {
	return "AdsDiscovery"
}

func (m *_AdsDiscovery) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (header)
	lengthInBits += 32

	// Simple field (requestId)
	lengthInBits += 32

	// Simple field (operation)
	lengthInBits += 32

	// Simple field (amsNetId)
	lengthInBits += m.AmsNetId.GetLengthInBits(ctx)

	// Simple field (portNumber)
	lengthInBits += 16

	// Implicit Field (numBlocks)
	lengthInBits += 32

	// Array field
	if len(m.Blocks) > 0 {
		for _curItem, element := range m.Blocks {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Blocks), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AdsDiscovery) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryParse(ctx context.Context, theBytes []byte) (AdsDiscovery, error) {
	return AdsDiscoveryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AdsDiscoveryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscovery, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscovery, error) {
		return AdsDiscoveryParseWithBuffer(ctx, readBuffer)
	}
}

func AdsDiscoveryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscovery, error) {
	v, err := (&_AdsDiscovery{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AdsDiscovery) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__adsDiscovery AdsDiscovery, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscovery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscovery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadConstField[uint32](ctx, "header", ReadUnsignedInt(readBuffer, uint8(32)), AdsDiscovery_HEADER, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	_ = header

	requestId, err := ReadSimpleField(ctx, "requestId", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestId' field"))
	}
	m.RequestId = requestId

	operation, err := ReadEnumField[Operation](ctx, "operation", "Operation", ReadEnum(OperationByValue, ReadUnsignedInt(readBuffer, uint8(32))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operation' field"))
	}
	m.Operation = operation

	amsNetId, err := ReadSimpleField[AmsNetId](ctx, "amsNetId", ReadComplex[AmsNetId](AmsNetIdParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'amsNetId' field"))
	}
	m.AmsNetId = amsNetId

	portNumber, err := ReadEnumField[AdsPortNumbers](ctx, "portNumber", "AdsPortNumbers", ReadEnum(AdsPortNumbersByValue, ReadUnsignedShort(readBuffer, uint8(16))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'portNumber' field"))
	}
	m.PortNumber = portNumber

	numBlocks, err := ReadImplicitField[uint32](ctx, "numBlocks", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numBlocks' field"))
	}
	_ = numBlocks

	blocks, err := ReadCountArrayField[AdsDiscoveryBlock](ctx, "blocks", ReadComplex[AdsDiscoveryBlock](AdsDiscoveryBlockParseWithBuffer, readBuffer), uint64(numBlocks), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'blocks' field"))
	}
	m.Blocks = blocks

	if closeErr := readBuffer.CloseContext("AdsDiscovery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscovery")
	}

	return m, nil
}

func (m *_AdsDiscovery) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscovery) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsDiscovery"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsDiscovery")
	}

	if err := WriteConstField(ctx, "header", AdsDiscovery_HEADER, WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[uint32](ctx, "requestId", m.GetRequestId(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'requestId' field")
	}

	if err := WriteSimpleEnumField[Operation](ctx, "operation", "Operation", m.GetOperation(), WriteEnum[Operation, uint32](Operation.GetValue, Operation.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32)), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'operation' field")
	}

	if err := WriteSimpleField[AmsNetId](ctx, "amsNetId", m.GetAmsNetId(), WriteComplex[AmsNetId](writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'amsNetId' field")
	}

	if err := WriteSimpleEnumField[AdsPortNumbers](ctx, "portNumber", "AdsPortNumbers", m.GetPortNumber(), WriteEnum[AdsPortNumbers, uint16](AdsPortNumbers.GetValue, AdsPortNumbers.PLC4XEnumName, WriteUnsignedShort(writeBuffer, 16)), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'portNumber' field")
	}
	numBlocks := uint32(uint32(len(m.GetBlocks())))
	if err := WriteImplicitField(ctx, "numBlocks", numBlocks, WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'numBlocks' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "blocks", m.GetBlocks(), writeBuffer, codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'blocks' field")
	}

	if popErr := writeBuffer.PopContext("AdsDiscovery"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsDiscovery")
	}
	return nil
}

func (m *_AdsDiscovery) IsAdsDiscovery() {}

func (m *_AdsDiscovery) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscovery) deepCopy() *_AdsDiscovery {
	if m == nil {
		return nil
	}
	_AdsDiscoveryCopy := &_AdsDiscovery{
		m.RequestId,
		m.Operation,
		utils.DeepCopy[AmsNetId](m.AmsNetId),
		m.PortNumber,
		utils.DeepCopySlice[AdsDiscoveryBlock, AdsDiscoveryBlock](m.Blocks),
	}
	return _AdsDiscoveryCopy
}

func (m *_AdsDiscovery) String() string {
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

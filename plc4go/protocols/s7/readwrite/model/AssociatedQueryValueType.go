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

// AssociatedQueryValueType is the corresponding interface of AssociatedQueryValueType
type AssociatedQueryValueType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetValueLength returns ValueLength (property field)
	GetValueLength() uint16
	// GetData returns Data (property field)
	GetData() []uint8
	// IsAssociatedQueryValueType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAssociatedQueryValueType()
	// CreateBuilder creates a AssociatedQueryValueTypeBuilder
	CreateAssociatedQueryValueTypeBuilder() AssociatedQueryValueTypeBuilder
}

// _AssociatedQueryValueType is the data-structure of this message
type _AssociatedQueryValueType struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	ValueLength   uint16
	Data          []uint8
}

var _ AssociatedQueryValueType = (*_AssociatedQueryValueType)(nil)

// NewAssociatedQueryValueType factory function for _AssociatedQueryValueType
func NewAssociatedQueryValueType(returnCode DataTransportErrorCode, transportSize DataTransportSize, valueLength uint16, data []uint8) *_AssociatedQueryValueType {
	return &_AssociatedQueryValueType{ReturnCode: returnCode, TransportSize: transportSize, ValueLength: valueLength, Data: data}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AssociatedQueryValueTypeBuilder is a builder for AssociatedQueryValueType
type AssociatedQueryValueTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(returnCode DataTransportErrorCode, transportSize DataTransportSize, valueLength uint16, data []uint8) AssociatedQueryValueTypeBuilder
	// WithReturnCode adds ReturnCode (property field)
	WithReturnCode(DataTransportErrorCode) AssociatedQueryValueTypeBuilder
	// WithTransportSize adds TransportSize (property field)
	WithTransportSize(DataTransportSize) AssociatedQueryValueTypeBuilder
	// WithValueLength adds ValueLength (property field)
	WithValueLength(uint16) AssociatedQueryValueTypeBuilder
	// WithData adds Data (property field)
	WithData(...uint8) AssociatedQueryValueTypeBuilder
	// Build builds the AssociatedQueryValueType or returns an error if something is wrong
	Build() (AssociatedQueryValueType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AssociatedQueryValueType
}

// NewAssociatedQueryValueTypeBuilder() creates a AssociatedQueryValueTypeBuilder
func NewAssociatedQueryValueTypeBuilder() AssociatedQueryValueTypeBuilder {
	return &_AssociatedQueryValueTypeBuilder{_AssociatedQueryValueType: new(_AssociatedQueryValueType)}
}

type _AssociatedQueryValueTypeBuilder struct {
	*_AssociatedQueryValueType

	err *utils.MultiError
}

var _ (AssociatedQueryValueTypeBuilder) = (*_AssociatedQueryValueTypeBuilder)(nil)

func (b *_AssociatedQueryValueTypeBuilder) WithMandatoryFields(returnCode DataTransportErrorCode, transportSize DataTransportSize, valueLength uint16, data []uint8) AssociatedQueryValueTypeBuilder {
	return b.WithReturnCode(returnCode).WithTransportSize(transportSize).WithValueLength(valueLength).WithData(data...)
}

func (b *_AssociatedQueryValueTypeBuilder) WithReturnCode(returnCode DataTransportErrorCode) AssociatedQueryValueTypeBuilder {
	b.ReturnCode = returnCode
	return b
}

func (b *_AssociatedQueryValueTypeBuilder) WithTransportSize(transportSize DataTransportSize) AssociatedQueryValueTypeBuilder {
	b.TransportSize = transportSize
	return b
}

func (b *_AssociatedQueryValueTypeBuilder) WithValueLength(valueLength uint16) AssociatedQueryValueTypeBuilder {
	b.ValueLength = valueLength
	return b
}

func (b *_AssociatedQueryValueTypeBuilder) WithData(data ...uint8) AssociatedQueryValueTypeBuilder {
	b.Data = data
	return b
}

func (b *_AssociatedQueryValueTypeBuilder) Build() (AssociatedQueryValueType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AssociatedQueryValueType.deepCopy(), nil
}

func (b *_AssociatedQueryValueTypeBuilder) MustBuild() AssociatedQueryValueType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AssociatedQueryValueTypeBuilder) DeepCopy() any {
	_copy := b.CreateAssociatedQueryValueTypeBuilder().(*_AssociatedQueryValueTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAssociatedQueryValueTypeBuilder creates a AssociatedQueryValueTypeBuilder
func (b *_AssociatedQueryValueType) CreateAssociatedQueryValueTypeBuilder() AssociatedQueryValueTypeBuilder {
	if b == nil {
		return NewAssociatedQueryValueTypeBuilder()
	}
	return &_AssociatedQueryValueTypeBuilder{_AssociatedQueryValueType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AssociatedQueryValueType) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_AssociatedQueryValueType) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_AssociatedQueryValueType) GetValueLength() uint16 {
	return m.ValueLength
}

func (m *_AssociatedQueryValueType) GetData() []uint8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAssociatedQueryValueType(structType any) AssociatedQueryValueType {
	if casted, ok := structType.(AssociatedQueryValueType); ok {
		return casted
	}
	if casted, ok := structType.(*AssociatedQueryValueType); ok {
		return *casted
	}
	return nil
}

func (m *_AssociatedQueryValueType) GetTypeName() string {
	return "AssociatedQueryValueType"
}

func (m *_AssociatedQueryValueType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Simple field (valueLength)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AssociatedQueryValueType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AssociatedQueryValueTypeParse(ctx context.Context, theBytes []byte) (AssociatedQueryValueType, error) {
	return AssociatedQueryValueTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AssociatedQueryValueTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AssociatedQueryValueType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AssociatedQueryValueType, error) {
		return AssociatedQueryValueTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AssociatedQueryValueTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AssociatedQueryValueType, error) {
	v, err := (&_AssociatedQueryValueType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AssociatedQueryValueType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__associatedQueryValueType AssociatedQueryValueType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AssociatedQueryValueType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AssociatedQueryValueType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	returnCode, err := ReadEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", ReadEnum(DataTransportErrorCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'returnCode' field"))
	}
	m.ReturnCode = returnCode

	transportSize, err := ReadEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", ReadEnum(DataTransportSizeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSize' field"))
	}
	m.TransportSize = transportSize

	valueLength, err := ReadSimpleField(ctx, "valueLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueLength' field"))
	}
	m.ValueLength = valueLength

	data, err := ReadCountArrayField[uint8](ctx, "data", ReadUnsignedByte(readBuffer, uint8(8)), uint64(valueLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AssociatedQueryValueType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AssociatedQueryValueType")
	}

	return m, nil
}

func (m *_AssociatedQueryValueType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AssociatedQueryValueType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AssociatedQueryValueType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AssociatedQueryValueType")
	}

	if err := WriteSimpleEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", m.GetReturnCode(), WriteEnum[DataTransportErrorCode, uint8](DataTransportErrorCode.GetValue, DataTransportErrorCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'returnCode' field")
	}

	if err := WriteSimpleEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", m.GetTransportSize(), WriteEnum[DataTransportSize, uint8](DataTransportSize.GetValue, DataTransportSize.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'transportSize' field")
	}

	if err := WriteSimpleField[uint16](ctx, "valueLength", m.GetValueLength(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'valueLength' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "data", m.GetData(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("AssociatedQueryValueType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AssociatedQueryValueType")
	}
	return nil
}

func (m *_AssociatedQueryValueType) IsAssociatedQueryValueType() {}

func (m *_AssociatedQueryValueType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AssociatedQueryValueType) deepCopy() *_AssociatedQueryValueType {
	if m == nil {
		return nil
	}
	_AssociatedQueryValueTypeCopy := &_AssociatedQueryValueType{
		m.ReturnCode,
		m.TransportSize,
		m.ValueLength,
		utils.DeepCopySlice[uint8, uint8](m.Data),
	}
	return _AssociatedQueryValueTypeCopy
}

func (m *_AssociatedQueryValueType) String() string {
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

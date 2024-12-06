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

// ApduDataGroupValueWrite is the corresponding interface of ApduDataGroupValueWrite
type ApduDataGroupValueWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// GetDataFirstByte returns DataFirstByte (property field)
	GetDataFirstByte() int8
	// GetData returns Data (property field)
	GetData() []byte
	// IsApduDataGroupValueWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataGroupValueWrite()
	// CreateBuilder creates a ApduDataGroupValueWriteBuilder
	CreateApduDataGroupValueWriteBuilder() ApduDataGroupValueWriteBuilder
}

// _ApduDataGroupValueWrite is the data-structure of this message
type _ApduDataGroupValueWrite struct {
	ApduDataContract
	DataFirstByte int8
	Data          []byte
}

var _ ApduDataGroupValueWrite = (*_ApduDataGroupValueWrite)(nil)
var _ ApduDataRequirements = (*_ApduDataGroupValueWrite)(nil)

// NewApduDataGroupValueWrite factory function for _ApduDataGroupValueWrite
func NewApduDataGroupValueWrite(dataFirstByte int8, data []byte, dataLength uint8) *_ApduDataGroupValueWrite {
	_result := &_ApduDataGroupValueWrite{
		ApduDataContract: NewApduData(dataLength),
		DataFirstByte:    dataFirstByte,
		Data:             data,
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataGroupValueWriteBuilder is a builder for ApduDataGroupValueWrite
type ApduDataGroupValueWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dataFirstByte int8, data []byte) ApduDataGroupValueWriteBuilder
	// WithDataFirstByte adds DataFirstByte (property field)
	WithDataFirstByte(int8) ApduDataGroupValueWriteBuilder
	// WithData adds Data (property field)
	WithData(...byte) ApduDataGroupValueWriteBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataBuilder
	// Build builds the ApduDataGroupValueWrite or returns an error if something is wrong
	Build() (ApduDataGroupValueWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataGroupValueWrite
}

// NewApduDataGroupValueWriteBuilder() creates a ApduDataGroupValueWriteBuilder
func NewApduDataGroupValueWriteBuilder() ApduDataGroupValueWriteBuilder {
	return &_ApduDataGroupValueWriteBuilder{_ApduDataGroupValueWrite: new(_ApduDataGroupValueWrite)}
}

type _ApduDataGroupValueWriteBuilder struct {
	*_ApduDataGroupValueWrite

	parentBuilder *_ApduDataBuilder

	err *utils.MultiError
}

var _ (ApduDataGroupValueWriteBuilder) = (*_ApduDataGroupValueWriteBuilder)(nil)

func (b *_ApduDataGroupValueWriteBuilder) setParent(contract ApduDataContract) {
	b.ApduDataContract = contract
	contract.(*_ApduData)._SubType = b._ApduDataGroupValueWrite
}

func (b *_ApduDataGroupValueWriteBuilder) WithMandatoryFields(dataFirstByte int8, data []byte) ApduDataGroupValueWriteBuilder {
	return b.WithDataFirstByte(dataFirstByte).WithData(data...)
}

func (b *_ApduDataGroupValueWriteBuilder) WithDataFirstByte(dataFirstByte int8) ApduDataGroupValueWriteBuilder {
	b.DataFirstByte = dataFirstByte
	return b
}

func (b *_ApduDataGroupValueWriteBuilder) WithData(data ...byte) ApduDataGroupValueWriteBuilder {
	b.Data = data
	return b
}

func (b *_ApduDataGroupValueWriteBuilder) Build() (ApduDataGroupValueWrite, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataGroupValueWrite.deepCopy(), nil
}

func (b *_ApduDataGroupValueWriteBuilder) MustBuild() ApduDataGroupValueWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataGroupValueWriteBuilder) Done() ApduDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataBuilder().(*_ApduDataBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataGroupValueWriteBuilder) buildForApduData() (ApduData, error) {
	return b.Build()
}

func (b *_ApduDataGroupValueWriteBuilder) DeepCopy() any {
	_copy := b.CreateApduDataGroupValueWriteBuilder().(*_ApduDataGroupValueWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataGroupValueWriteBuilder creates a ApduDataGroupValueWriteBuilder
func (b *_ApduDataGroupValueWrite) CreateApduDataGroupValueWriteBuilder() ApduDataGroupValueWriteBuilder {
	if b == nil {
		return NewApduDataGroupValueWriteBuilder()
	}
	return &_ApduDataGroupValueWriteBuilder{_ApduDataGroupValueWrite: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataGroupValueWrite) GetApciType() uint8 {
	return 0x2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataGroupValueWrite) GetParent() ApduDataContract {
	return m.ApduDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataGroupValueWrite) GetDataFirstByte() int8 {
	return m.DataFirstByte
}

func (m *_ApduDataGroupValueWrite) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataGroupValueWrite(structType any) ApduDataGroupValueWrite {
	if casted, ok := structType.(ApduDataGroupValueWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataGroupValueWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataGroupValueWrite) GetTypeName() string {
	return "ApduDataGroupValueWrite"
}

func (m *_ApduDataGroupValueWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	// Simple field (dataFirstByte)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataGroupValueWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataGroupValueWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataGroupValueWrite ApduDataGroupValueWrite, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataGroupValueWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataGroupValueWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataFirstByte, err := ReadSimpleField(ctx, "dataFirstByte", ReadSignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataFirstByte' field"))
	}
	m.DataFirstByte = dataFirstByte

	data, err := readBuffer.ReadByteArray("data", int(utils.InlineIf((bool((dataLength) < (1))), func() any { return int32(int32(0)) }, func() any { return int32(int32(dataLength) - int32(int32(1))) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ApduDataGroupValueWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataGroupValueWrite")
	}

	return m, nil
}

func (m *_ApduDataGroupValueWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataGroupValueWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataGroupValueWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataGroupValueWrite")
		}

		if err := WriteSimpleField[int8](ctx, "dataFirstByte", m.GetDataFirstByte(), WriteSignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataFirstByte' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataGroupValueWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataGroupValueWrite")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataGroupValueWrite) IsApduDataGroupValueWrite() {}

func (m *_ApduDataGroupValueWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataGroupValueWrite) deepCopy() *_ApduDataGroupValueWrite {
	if m == nil {
		return nil
	}
	_ApduDataGroupValueWriteCopy := &_ApduDataGroupValueWrite{
		m.ApduDataContract.(*_ApduData).deepCopy(),
		m.DataFirstByte,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	_ApduDataGroupValueWriteCopy.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataGroupValueWriteCopy
}

func (m *_ApduDataGroupValueWrite) String() string {
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
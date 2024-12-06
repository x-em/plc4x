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

// ModbusPDUReadFileRecordResponseItem is the corresponding interface of ModbusPDUReadFileRecordResponseItem
type ModbusPDUReadFileRecordResponseItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() uint8
	// GetData returns Data (property field)
	GetData() []byte
	// IsModbusPDUReadFileRecordResponseItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadFileRecordResponseItem()
	// CreateBuilder creates a ModbusPDUReadFileRecordResponseItemBuilder
	CreateModbusPDUReadFileRecordResponseItemBuilder() ModbusPDUReadFileRecordResponseItemBuilder
}

// _ModbusPDUReadFileRecordResponseItem is the data-structure of this message
type _ModbusPDUReadFileRecordResponseItem struct {
	ReferenceType uint8
	Data          []byte
}

var _ ModbusPDUReadFileRecordResponseItem = (*_ModbusPDUReadFileRecordResponseItem)(nil)

// NewModbusPDUReadFileRecordResponseItem factory function for _ModbusPDUReadFileRecordResponseItem
func NewModbusPDUReadFileRecordResponseItem(referenceType uint8, data []byte) *_ModbusPDUReadFileRecordResponseItem {
	return &_ModbusPDUReadFileRecordResponseItem{ReferenceType: referenceType, Data: data}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadFileRecordResponseItemBuilder is a builder for ModbusPDUReadFileRecordResponseItem
type ModbusPDUReadFileRecordResponseItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(referenceType uint8, data []byte) ModbusPDUReadFileRecordResponseItemBuilder
	// WithReferenceType adds ReferenceType (property field)
	WithReferenceType(uint8) ModbusPDUReadFileRecordResponseItemBuilder
	// WithData adds Data (property field)
	WithData(...byte) ModbusPDUReadFileRecordResponseItemBuilder
	// Build builds the ModbusPDUReadFileRecordResponseItem or returns an error if something is wrong
	Build() (ModbusPDUReadFileRecordResponseItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadFileRecordResponseItem
}

// NewModbusPDUReadFileRecordResponseItemBuilder() creates a ModbusPDUReadFileRecordResponseItemBuilder
func NewModbusPDUReadFileRecordResponseItemBuilder() ModbusPDUReadFileRecordResponseItemBuilder {
	return &_ModbusPDUReadFileRecordResponseItemBuilder{_ModbusPDUReadFileRecordResponseItem: new(_ModbusPDUReadFileRecordResponseItem)}
}

type _ModbusPDUReadFileRecordResponseItemBuilder struct {
	*_ModbusPDUReadFileRecordResponseItem

	err *utils.MultiError
}

var _ (ModbusPDUReadFileRecordResponseItemBuilder) = (*_ModbusPDUReadFileRecordResponseItemBuilder)(nil)

func (b *_ModbusPDUReadFileRecordResponseItemBuilder) WithMandatoryFields(referenceType uint8, data []byte) ModbusPDUReadFileRecordResponseItemBuilder {
	return b.WithReferenceType(referenceType).WithData(data...)
}

func (b *_ModbusPDUReadFileRecordResponseItemBuilder) WithReferenceType(referenceType uint8) ModbusPDUReadFileRecordResponseItemBuilder {
	b.ReferenceType = referenceType
	return b
}

func (b *_ModbusPDUReadFileRecordResponseItemBuilder) WithData(data ...byte) ModbusPDUReadFileRecordResponseItemBuilder {
	b.Data = data
	return b
}

func (b *_ModbusPDUReadFileRecordResponseItemBuilder) Build() (ModbusPDUReadFileRecordResponseItem, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadFileRecordResponseItem.deepCopy(), nil
}

func (b *_ModbusPDUReadFileRecordResponseItemBuilder) MustBuild() ModbusPDUReadFileRecordResponseItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUReadFileRecordResponseItemBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadFileRecordResponseItemBuilder().(*_ModbusPDUReadFileRecordResponseItemBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadFileRecordResponseItemBuilder creates a ModbusPDUReadFileRecordResponseItemBuilder
func (b *_ModbusPDUReadFileRecordResponseItem) CreateModbusPDUReadFileRecordResponseItemBuilder() ModbusPDUReadFileRecordResponseItemBuilder {
	if b == nil {
		return NewModbusPDUReadFileRecordResponseItemBuilder()
	}
	return &_ModbusPDUReadFileRecordResponseItemBuilder{_ModbusPDUReadFileRecordResponseItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFileRecordResponseItem) GetReferenceType() uint8 {
	return m.ReferenceType
}

func (m *_ModbusPDUReadFileRecordResponseItem) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFileRecordResponseItem(structType any) ModbusPDUReadFileRecordResponseItem {
	if casted, ok := structType.(ModbusPDUReadFileRecordResponseItem); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFileRecordResponseItem); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) GetTypeName() string {
	return "ModbusPDUReadFileRecordResponseItem"
}

func (m *_ModbusPDUReadFileRecordResponseItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (dataLength)
	lengthInBits += 8

	// Simple field (referenceType)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadFileRecordResponseItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadFileRecordResponseItemParse(ctx context.Context, theBytes []byte) (ModbusPDUReadFileRecordResponseItem, error) {
	return ModbusPDUReadFileRecordResponseItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusPDUReadFileRecordResponseItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadFileRecordResponseItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadFileRecordResponseItem, error) {
		return ModbusPDUReadFileRecordResponseItemParseWithBuffer(ctx, readBuffer)
	}
}

func ModbusPDUReadFileRecordResponseItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadFileRecordResponseItem, error) {
	v, err := (&_ModbusPDUReadFileRecordResponseItem{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__modbusPDUReadFileRecordResponseItem ModbusPDUReadFileRecordResponseItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordResponseItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFileRecordResponseItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataLength, err := ReadImplicitField[uint8](ctx, "dataLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataLength' field"))
	}
	_ = dataLength

	referenceType, err := ReadSimpleField(ctx, "referenceType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceType' field"))
	}
	m.ReferenceType = referenceType

	data, err := readBuffer.ReadByteArray("data", int(int32(dataLength)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordResponseItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFileRecordResponseItem")
	}

	return m, nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordResponseItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFileRecordResponseItem")
	}
	dataLength := uint8(uint8(uint8(len(m.GetData()))) + uint8(uint8(1)))
	if err := WriteImplicitField(ctx, "dataLength", dataLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataLength' field")
	}

	if err := WriteSimpleField[uint8](ctx, "referenceType", m.GetReferenceType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'referenceType' field")
	}

	if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordResponseItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDUReadFileRecordResponseItem")
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) IsModbusPDUReadFileRecordResponseItem() {}

func (m *_ModbusPDUReadFileRecordResponseItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadFileRecordResponseItem) deepCopy() *_ModbusPDUReadFileRecordResponseItem {
	if m == nil {
		return nil
	}
	_ModbusPDUReadFileRecordResponseItemCopy := &_ModbusPDUReadFileRecordResponseItem{
		m.ReferenceType,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	return _ModbusPDUReadFileRecordResponseItemCopy
}

func (m *_ModbusPDUReadFileRecordResponseItem) String() string {
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
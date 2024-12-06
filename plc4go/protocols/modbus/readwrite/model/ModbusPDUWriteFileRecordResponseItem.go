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

// ModbusPDUWriteFileRecordResponseItem is the corresponding interface of ModbusPDUWriteFileRecordResponseItem
type ModbusPDUWriteFileRecordResponseItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint16
	// GetRecordNumber returns RecordNumber (property field)
	GetRecordNumber() uint16
	// GetRecordData returns RecordData (property field)
	GetRecordData() []byte
	// IsModbusPDUWriteFileRecordResponseItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUWriteFileRecordResponseItem()
	// CreateBuilder creates a ModbusPDUWriteFileRecordResponseItemBuilder
	CreateModbusPDUWriteFileRecordResponseItemBuilder() ModbusPDUWriteFileRecordResponseItemBuilder
}

// _ModbusPDUWriteFileRecordResponseItem is the data-structure of this message
type _ModbusPDUWriteFileRecordResponseItem struct {
	ReferenceType uint8
	FileNumber    uint16
	RecordNumber  uint16
	RecordData    []byte
}

var _ ModbusPDUWriteFileRecordResponseItem = (*_ModbusPDUWriteFileRecordResponseItem)(nil)

// NewModbusPDUWriteFileRecordResponseItem factory function for _ModbusPDUWriteFileRecordResponseItem
func NewModbusPDUWriteFileRecordResponseItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordData []byte) *_ModbusPDUWriteFileRecordResponseItem {
	return &_ModbusPDUWriteFileRecordResponseItem{ReferenceType: referenceType, FileNumber: fileNumber, RecordNumber: recordNumber, RecordData: recordData}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUWriteFileRecordResponseItemBuilder is a builder for ModbusPDUWriteFileRecordResponseItem
type ModbusPDUWriteFileRecordResponseItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(referenceType uint8, fileNumber uint16, recordNumber uint16, recordData []byte) ModbusPDUWriteFileRecordResponseItemBuilder
	// WithReferenceType adds ReferenceType (property field)
	WithReferenceType(uint8) ModbusPDUWriteFileRecordResponseItemBuilder
	// WithFileNumber adds FileNumber (property field)
	WithFileNumber(uint16) ModbusPDUWriteFileRecordResponseItemBuilder
	// WithRecordNumber adds RecordNumber (property field)
	WithRecordNumber(uint16) ModbusPDUWriteFileRecordResponseItemBuilder
	// WithRecordData adds RecordData (property field)
	WithRecordData(...byte) ModbusPDUWriteFileRecordResponseItemBuilder
	// Build builds the ModbusPDUWriteFileRecordResponseItem or returns an error if something is wrong
	Build() (ModbusPDUWriteFileRecordResponseItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUWriteFileRecordResponseItem
}

// NewModbusPDUWriteFileRecordResponseItemBuilder() creates a ModbusPDUWriteFileRecordResponseItemBuilder
func NewModbusPDUWriteFileRecordResponseItemBuilder() ModbusPDUWriteFileRecordResponseItemBuilder {
	return &_ModbusPDUWriteFileRecordResponseItemBuilder{_ModbusPDUWriteFileRecordResponseItem: new(_ModbusPDUWriteFileRecordResponseItem)}
}

type _ModbusPDUWriteFileRecordResponseItemBuilder struct {
	*_ModbusPDUWriteFileRecordResponseItem

	err *utils.MultiError
}

var _ (ModbusPDUWriteFileRecordResponseItemBuilder) = (*_ModbusPDUWriteFileRecordResponseItemBuilder)(nil)

func (b *_ModbusPDUWriteFileRecordResponseItemBuilder) WithMandatoryFields(referenceType uint8, fileNumber uint16, recordNumber uint16, recordData []byte) ModbusPDUWriteFileRecordResponseItemBuilder {
	return b.WithReferenceType(referenceType).WithFileNumber(fileNumber).WithRecordNumber(recordNumber).WithRecordData(recordData...)
}

func (b *_ModbusPDUWriteFileRecordResponseItemBuilder) WithReferenceType(referenceType uint8) ModbusPDUWriteFileRecordResponseItemBuilder {
	b.ReferenceType = referenceType
	return b
}

func (b *_ModbusPDUWriteFileRecordResponseItemBuilder) WithFileNumber(fileNumber uint16) ModbusPDUWriteFileRecordResponseItemBuilder {
	b.FileNumber = fileNumber
	return b
}

func (b *_ModbusPDUWriteFileRecordResponseItemBuilder) WithRecordNumber(recordNumber uint16) ModbusPDUWriteFileRecordResponseItemBuilder {
	b.RecordNumber = recordNumber
	return b
}

func (b *_ModbusPDUWriteFileRecordResponseItemBuilder) WithRecordData(recordData ...byte) ModbusPDUWriteFileRecordResponseItemBuilder {
	b.RecordData = recordData
	return b
}

func (b *_ModbusPDUWriteFileRecordResponseItemBuilder) Build() (ModbusPDUWriteFileRecordResponseItem, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUWriteFileRecordResponseItem.deepCopy(), nil
}

func (b *_ModbusPDUWriteFileRecordResponseItemBuilder) MustBuild() ModbusPDUWriteFileRecordResponseItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUWriteFileRecordResponseItemBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUWriteFileRecordResponseItemBuilder().(*_ModbusPDUWriteFileRecordResponseItemBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUWriteFileRecordResponseItemBuilder creates a ModbusPDUWriteFileRecordResponseItemBuilder
func (b *_ModbusPDUWriteFileRecordResponseItem) CreateModbusPDUWriteFileRecordResponseItemBuilder() ModbusPDUWriteFileRecordResponseItemBuilder {
	if b == nil {
		return NewModbusPDUWriteFileRecordResponseItemBuilder()
	}
	return &_ModbusPDUWriteFileRecordResponseItemBuilder{_ModbusPDUWriteFileRecordResponseItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteFileRecordResponseItem) GetReferenceType() uint8 {
	return m.ReferenceType
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetFileNumber() uint16 {
	return m.FileNumber
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetRecordNumber() uint16 {
	return m.RecordNumber
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetRecordData() []byte {
	return m.RecordData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteFileRecordResponseItem(structType any) ModbusPDUWriteFileRecordResponseItem {
	if casted, ok := structType.(ModbusPDUWriteFileRecordResponseItem); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteFileRecordResponseItem); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetTypeName() string {
	return "ModbusPDUWriteFileRecordResponseItem"
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Implicit Field (recordLength)
	lengthInBits += 16

	// Array field
	if len(m.RecordData) > 0 {
		lengthInBits += 8 * uint16(len(m.RecordData))
	}

	return lengthInBits
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUWriteFileRecordResponseItemParse(ctx context.Context, theBytes []byte) (ModbusPDUWriteFileRecordResponseItem, error) {
	return ModbusPDUWriteFileRecordResponseItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusPDUWriteFileRecordResponseItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUWriteFileRecordResponseItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUWriteFileRecordResponseItem, error) {
		return ModbusPDUWriteFileRecordResponseItemParseWithBuffer(ctx, readBuffer)
	}
}

func ModbusPDUWriteFileRecordResponseItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUWriteFileRecordResponseItem, error) {
	v, err := (&_ModbusPDUWriteFileRecordResponseItem{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__modbusPDUWriteFileRecordResponseItem ModbusPDUWriteFileRecordResponseItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordResponseItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteFileRecordResponseItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceType, err := ReadSimpleField(ctx, "referenceType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceType' field"))
	}
	m.ReferenceType = referenceType

	fileNumber, err := ReadSimpleField(ctx, "fileNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileNumber' field"))
	}
	m.FileNumber = fileNumber

	recordNumber, err := ReadSimpleField(ctx, "recordNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordNumber' field"))
	}
	m.RecordNumber = recordNumber

	recordLength, err := ReadImplicitField[uint16](ctx, "recordLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordLength' field"))
	}
	_ = recordLength

	recordData, err := readBuffer.ReadByteArray("recordData", int(recordLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordData' field"))
	}
	m.RecordData = recordData

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordResponseItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteFileRecordResponseItem")
	}

	return m, nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordResponseItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteFileRecordResponseItem")
	}

	if err := WriteSimpleField[uint8](ctx, "referenceType", m.GetReferenceType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'referenceType' field")
	}

	if err := WriteSimpleField[uint16](ctx, "fileNumber", m.GetFileNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'fileNumber' field")
	}

	if err := WriteSimpleField[uint16](ctx, "recordNumber", m.GetRecordNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'recordNumber' field")
	}
	recordLength := uint16(uint16(uint16(len(m.GetRecordData()))) / uint16(uint16(2)))
	if err := WriteImplicitField(ctx, "recordLength", recordLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'recordLength' field")
	}

	if err := WriteByteArrayField(ctx, "recordData", m.GetRecordData(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'recordData' field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordResponseItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDUWriteFileRecordResponseItem")
	}
	return nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) IsModbusPDUWriteFileRecordResponseItem() {}

func (m *_ModbusPDUWriteFileRecordResponseItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUWriteFileRecordResponseItem) deepCopy() *_ModbusPDUWriteFileRecordResponseItem {
	if m == nil {
		return nil
	}
	_ModbusPDUWriteFileRecordResponseItemCopy := &_ModbusPDUWriteFileRecordResponseItem{
		m.ReferenceType,
		m.FileNumber,
		m.RecordNumber,
		utils.DeepCopySlice[byte, byte](m.RecordData),
	}
	return _ModbusPDUWriteFileRecordResponseItemCopy
}

func (m *_ModbusPDUWriteFileRecordResponseItem) String() string {
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
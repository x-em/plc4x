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

// AdsMultiRequestItemReadWrite is the corresponding interface of AdsMultiRequestItemReadWrite
type AdsMultiRequestItemReadWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsMultiRequestItem
	// GetItemIndexGroup returns ItemIndexGroup (property field)
	GetItemIndexGroup() uint32
	// GetItemIndexOffset returns ItemIndexOffset (property field)
	GetItemIndexOffset() uint32
	// GetItemReadLength returns ItemReadLength (property field)
	GetItemReadLength() uint32
	// GetItemWriteLength returns ItemWriteLength (property field)
	GetItemWriteLength() uint32
	// IsAdsMultiRequestItemReadWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsMultiRequestItemReadWrite()
	// CreateBuilder creates a AdsMultiRequestItemReadWriteBuilder
	CreateAdsMultiRequestItemReadWriteBuilder() AdsMultiRequestItemReadWriteBuilder
}

// _AdsMultiRequestItemReadWrite is the data-structure of this message
type _AdsMultiRequestItemReadWrite struct {
	AdsMultiRequestItemContract
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemReadLength  uint32
	ItemWriteLength uint32
}

var _ AdsMultiRequestItemReadWrite = (*_AdsMultiRequestItemReadWrite)(nil)
var _ AdsMultiRequestItemRequirements = (*_AdsMultiRequestItemReadWrite)(nil)

// NewAdsMultiRequestItemReadWrite factory function for _AdsMultiRequestItemReadWrite
func NewAdsMultiRequestItemReadWrite(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32, itemWriteLength uint32) *_AdsMultiRequestItemReadWrite {
	_result := &_AdsMultiRequestItemReadWrite{
		AdsMultiRequestItemContract: NewAdsMultiRequestItem(),
		ItemIndexGroup:              itemIndexGroup,
		ItemIndexOffset:             itemIndexOffset,
		ItemReadLength:              itemReadLength,
		ItemWriteLength:             itemWriteLength,
	}
	_result.AdsMultiRequestItemContract.(*_AdsMultiRequestItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsMultiRequestItemReadWriteBuilder is a builder for AdsMultiRequestItemReadWrite
type AdsMultiRequestItemReadWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32, itemWriteLength uint32) AdsMultiRequestItemReadWriteBuilder
	// WithItemIndexGroup adds ItemIndexGroup (property field)
	WithItemIndexGroup(uint32) AdsMultiRequestItemReadWriteBuilder
	// WithItemIndexOffset adds ItemIndexOffset (property field)
	WithItemIndexOffset(uint32) AdsMultiRequestItemReadWriteBuilder
	// WithItemReadLength adds ItemReadLength (property field)
	WithItemReadLength(uint32) AdsMultiRequestItemReadWriteBuilder
	// WithItemWriteLength adds ItemWriteLength (property field)
	WithItemWriteLength(uint32) AdsMultiRequestItemReadWriteBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AdsMultiRequestItemBuilder
	// Build builds the AdsMultiRequestItemReadWrite or returns an error if something is wrong
	Build() (AdsMultiRequestItemReadWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsMultiRequestItemReadWrite
}

// NewAdsMultiRequestItemReadWriteBuilder() creates a AdsMultiRequestItemReadWriteBuilder
func NewAdsMultiRequestItemReadWriteBuilder() AdsMultiRequestItemReadWriteBuilder {
	return &_AdsMultiRequestItemReadWriteBuilder{_AdsMultiRequestItemReadWrite: new(_AdsMultiRequestItemReadWrite)}
}

type _AdsMultiRequestItemReadWriteBuilder struct {
	*_AdsMultiRequestItemReadWrite

	parentBuilder *_AdsMultiRequestItemBuilder

	err *utils.MultiError
}

var _ (AdsMultiRequestItemReadWriteBuilder) = (*_AdsMultiRequestItemReadWriteBuilder)(nil)

func (b *_AdsMultiRequestItemReadWriteBuilder) setParent(contract AdsMultiRequestItemContract) {
	b.AdsMultiRequestItemContract = contract
	contract.(*_AdsMultiRequestItem)._SubType = b._AdsMultiRequestItemReadWrite
}

func (b *_AdsMultiRequestItemReadWriteBuilder) WithMandatoryFields(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32, itemWriteLength uint32) AdsMultiRequestItemReadWriteBuilder {
	return b.WithItemIndexGroup(itemIndexGroup).WithItemIndexOffset(itemIndexOffset).WithItemReadLength(itemReadLength).WithItemWriteLength(itemWriteLength)
}

func (b *_AdsMultiRequestItemReadWriteBuilder) WithItemIndexGroup(itemIndexGroup uint32) AdsMultiRequestItemReadWriteBuilder {
	b.ItemIndexGroup = itemIndexGroup
	return b
}

func (b *_AdsMultiRequestItemReadWriteBuilder) WithItemIndexOffset(itemIndexOffset uint32) AdsMultiRequestItemReadWriteBuilder {
	b.ItemIndexOffset = itemIndexOffset
	return b
}

func (b *_AdsMultiRequestItemReadWriteBuilder) WithItemReadLength(itemReadLength uint32) AdsMultiRequestItemReadWriteBuilder {
	b.ItemReadLength = itemReadLength
	return b
}

func (b *_AdsMultiRequestItemReadWriteBuilder) WithItemWriteLength(itemWriteLength uint32) AdsMultiRequestItemReadWriteBuilder {
	b.ItemWriteLength = itemWriteLength
	return b
}

func (b *_AdsMultiRequestItemReadWriteBuilder) Build() (AdsMultiRequestItemReadWrite, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsMultiRequestItemReadWrite.deepCopy(), nil
}

func (b *_AdsMultiRequestItemReadWriteBuilder) MustBuild() AdsMultiRequestItemReadWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsMultiRequestItemReadWriteBuilder) Done() AdsMultiRequestItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAdsMultiRequestItemBuilder().(*_AdsMultiRequestItemBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsMultiRequestItemReadWriteBuilder) buildForAdsMultiRequestItem() (AdsMultiRequestItem, error) {
	return b.Build()
}

func (b *_AdsMultiRequestItemReadWriteBuilder) DeepCopy() any {
	_copy := b.CreateAdsMultiRequestItemReadWriteBuilder().(*_AdsMultiRequestItemReadWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsMultiRequestItemReadWriteBuilder creates a AdsMultiRequestItemReadWriteBuilder
func (b *_AdsMultiRequestItemReadWrite) CreateAdsMultiRequestItemReadWriteBuilder() AdsMultiRequestItemReadWriteBuilder {
	if b == nil {
		return NewAdsMultiRequestItemReadWriteBuilder()
	}
	return &_AdsMultiRequestItemReadWriteBuilder{_AdsMultiRequestItemReadWrite: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsMultiRequestItemReadWrite) GetIndexGroup() uint32 {
	return uint32(61570)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsMultiRequestItemReadWrite) GetParent() AdsMultiRequestItemContract {
	return m.AdsMultiRequestItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsMultiRequestItemReadWrite) GetItemIndexGroup() uint32 {
	return m.ItemIndexGroup
}

func (m *_AdsMultiRequestItemReadWrite) GetItemIndexOffset() uint32 {
	return m.ItemIndexOffset
}

func (m *_AdsMultiRequestItemReadWrite) GetItemReadLength() uint32 {
	return m.ItemReadLength
}

func (m *_AdsMultiRequestItemReadWrite) GetItemWriteLength() uint32 {
	return m.ItemWriteLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItemReadWrite(structType any) AdsMultiRequestItemReadWrite {
	if casted, ok := structType.(AdsMultiRequestItemReadWrite); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItemReadWrite); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItemReadWrite) GetTypeName() string {
	return "AdsMultiRequestItemReadWrite"
}

func (m *_AdsMultiRequestItemReadWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).getLengthInBits(ctx))

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemReadLength)
	lengthInBits += 32

	// Simple field (itemWriteLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsMultiRequestItemReadWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsMultiRequestItemReadWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsMultiRequestItem, indexGroup uint32) (__adsMultiRequestItemReadWrite AdsMultiRequestItemReadWrite, err error) {
	m.AdsMultiRequestItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemReadWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItemReadWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemIndexGroup, err := ReadSimpleField(ctx, "itemIndexGroup", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndexGroup' field"))
	}
	m.ItemIndexGroup = itemIndexGroup

	itemIndexOffset, err := ReadSimpleField(ctx, "itemIndexOffset", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndexOffset' field"))
	}
	m.ItemIndexOffset = itemIndexOffset

	itemReadLength, err := ReadSimpleField(ctx, "itemReadLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemReadLength' field"))
	}
	m.ItemReadLength = itemReadLength

	itemWriteLength, err := ReadSimpleField(ctx, "itemWriteLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemWriteLength' field"))
	}
	m.ItemWriteLength = itemWriteLength

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemReadWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItemReadWrite")
	}

	return m, nil
}

func (m *_AdsMultiRequestItemReadWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsMultiRequestItemReadWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemReadWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItemReadWrite")
		}

		if err := WriteSimpleField[uint32](ctx, "itemIndexGroup", m.GetItemIndexGroup(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemIndexGroup' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemIndexOffset", m.GetItemIndexOffset(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemIndexOffset' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemReadLength", m.GetItemReadLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemReadLength' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemWriteLength", m.GetItemWriteLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemWriteLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemReadWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsMultiRequestItemReadWrite")
		}
		return nil
	}
	return m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsMultiRequestItemReadWrite) IsAdsMultiRequestItemReadWrite() {}

func (m *_AdsMultiRequestItemReadWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsMultiRequestItemReadWrite) deepCopy() *_AdsMultiRequestItemReadWrite {
	if m == nil {
		return nil
	}
	_AdsMultiRequestItemReadWriteCopy := &_AdsMultiRequestItemReadWrite{
		m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).deepCopy(),
		m.ItemIndexGroup,
		m.ItemIndexOffset,
		m.ItemReadLength,
		m.ItemWriteLength,
	}
	_AdsMultiRequestItemReadWriteCopy.AdsMultiRequestItemContract.(*_AdsMultiRequestItem)._SubType = m
	return _AdsMultiRequestItemReadWriteCopy
}

func (m *_AdsMultiRequestItemReadWrite) String() string {
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
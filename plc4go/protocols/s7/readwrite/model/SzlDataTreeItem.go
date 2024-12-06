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

// SzlDataTreeItem is the corresponding interface of SzlDataTreeItem
type SzlDataTreeItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetItemIndex returns ItemIndex (property field)
	GetItemIndex() uint16
	// GetMlfb returns Mlfb (property field)
	GetMlfb() []byte
	// GetModuleTypeId returns ModuleTypeId (property field)
	GetModuleTypeId() uint16
	// GetAusbg returns Ausbg (property field)
	GetAusbg() uint16
	// GetAusbe returns Ausbe (property field)
	GetAusbe() uint16
	// IsSzlDataTreeItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSzlDataTreeItem()
	// CreateBuilder creates a SzlDataTreeItemBuilder
	CreateSzlDataTreeItemBuilder() SzlDataTreeItemBuilder
}

// _SzlDataTreeItem is the data-structure of this message
type _SzlDataTreeItem struct {
	ItemIndex    uint16
	Mlfb         []byte
	ModuleTypeId uint16
	Ausbg        uint16
	Ausbe        uint16
}

var _ SzlDataTreeItem = (*_SzlDataTreeItem)(nil)

// NewSzlDataTreeItem factory function for _SzlDataTreeItem
func NewSzlDataTreeItem(itemIndex uint16, mlfb []byte, moduleTypeId uint16, ausbg uint16, ausbe uint16) *_SzlDataTreeItem {
	return &_SzlDataTreeItem{ItemIndex: itemIndex, Mlfb: mlfb, ModuleTypeId: moduleTypeId, Ausbg: ausbg, Ausbe: ausbe}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SzlDataTreeItemBuilder is a builder for SzlDataTreeItem
type SzlDataTreeItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(itemIndex uint16, mlfb []byte, moduleTypeId uint16, ausbg uint16, ausbe uint16) SzlDataTreeItemBuilder
	// WithItemIndex adds ItemIndex (property field)
	WithItemIndex(uint16) SzlDataTreeItemBuilder
	// WithMlfb adds Mlfb (property field)
	WithMlfb(...byte) SzlDataTreeItemBuilder
	// WithModuleTypeId adds ModuleTypeId (property field)
	WithModuleTypeId(uint16) SzlDataTreeItemBuilder
	// WithAusbg adds Ausbg (property field)
	WithAusbg(uint16) SzlDataTreeItemBuilder
	// WithAusbe adds Ausbe (property field)
	WithAusbe(uint16) SzlDataTreeItemBuilder
	// Build builds the SzlDataTreeItem or returns an error if something is wrong
	Build() (SzlDataTreeItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SzlDataTreeItem
}

// NewSzlDataTreeItemBuilder() creates a SzlDataTreeItemBuilder
func NewSzlDataTreeItemBuilder() SzlDataTreeItemBuilder {
	return &_SzlDataTreeItemBuilder{_SzlDataTreeItem: new(_SzlDataTreeItem)}
}

type _SzlDataTreeItemBuilder struct {
	*_SzlDataTreeItem

	err *utils.MultiError
}

var _ (SzlDataTreeItemBuilder) = (*_SzlDataTreeItemBuilder)(nil)

func (b *_SzlDataTreeItemBuilder) WithMandatoryFields(itemIndex uint16, mlfb []byte, moduleTypeId uint16, ausbg uint16, ausbe uint16) SzlDataTreeItemBuilder {
	return b.WithItemIndex(itemIndex).WithMlfb(mlfb...).WithModuleTypeId(moduleTypeId).WithAusbg(ausbg).WithAusbe(ausbe)
}

func (b *_SzlDataTreeItemBuilder) WithItemIndex(itemIndex uint16) SzlDataTreeItemBuilder {
	b.ItemIndex = itemIndex
	return b
}

func (b *_SzlDataTreeItemBuilder) WithMlfb(mlfb ...byte) SzlDataTreeItemBuilder {
	b.Mlfb = mlfb
	return b
}

func (b *_SzlDataTreeItemBuilder) WithModuleTypeId(moduleTypeId uint16) SzlDataTreeItemBuilder {
	b.ModuleTypeId = moduleTypeId
	return b
}

func (b *_SzlDataTreeItemBuilder) WithAusbg(ausbg uint16) SzlDataTreeItemBuilder {
	b.Ausbg = ausbg
	return b
}

func (b *_SzlDataTreeItemBuilder) WithAusbe(ausbe uint16) SzlDataTreeItemBuilder {
	b.Ausbe = ausbe
	return b
}

func (b *_SzlDataTreeItemBuilder) Build() (SzlDataTreeItem, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SzlDataTreeItem.deepCopy(), nil
}

func (b *_SzlDataTreeItemBuilder) MustBuild() SzlDataTreeItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SzlDataTreeItemBuilder) DeepCopy() any {
	_copy := b.CreateSzlDataTreeItemBuilder().(*_SzlDataTreeItemBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSzlDataTreeItemBuilder creates a SzlDataTreeItemBuilder
func (b *_SzlDataTreeItem) CreateSzlDataTreeItemBuilder() SzlDataTreeItemBuilder {
	if b == nil {
		return NewSzlDataTreeItemBuilder()
	}
	return &_SzlDataTreeItemBuilder{_SzlDataTreeItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SzlDataTreeItem) GetItemIndex() uint16 {
	return m.ItemIndex
}

func (m *_SzlDataTreeItem) GetMlfb() []byte {
	return m.Mlfb
}

func (m *_SzlDataTreeItem) GetModuleTypeId() uint16 {
	return m.ModuleTypeId
}

func (m *_SzlDataTreeItem) GetAusbg() uint16 {
	return m.Ausbg
}

func (m *_SzlDataTreeItem) GetAusbe() uint16 {
	return m.Ausbe
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSzlDataTreeItem(structType any) SzlDataTreeItem {
	if casted, ok := structType.(SzlDataTreeItem); ok {
		return casted
	}
	if casted, ok := structType.(*SzlDataTreeItem); ok {
		return *casted
	}
	return nil
}

func (m *_SzlDataTreeItem) GetTypeName() string {
	return "SzlDataTreeItem"
}

func (m *_SzlDataTreeItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (itemIndex)
	lengthInBits += 16

	// Array field
	if len(m.Mlfb) > 0 {
		lengthInBits += 8 * uint16(len(m.Mlfb))
	}

	// Simple field (moduleTypeId)
	lengthInBits += 16

	// Simple field (ausbg)
	lengthInBits += 16

	// Simple field (ausbe)
	lengthInBits += 16

	return lengthInBits
}

func (m *_SzlDataTreeItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SzlDataTreeItemParse(ctx context.Context, theBytes []byte) (SzlDataTreeItem, error) {
	return SzlDataTreeItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SzlDataTreeItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SzlDataTreeItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SzlDataTreeItem, error) {
		return SzlDataTreeItemParseWithBuffer(ctx, readBuffer)
	}
}

func SzlDataTreeItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SzlDataTreeItem, error) {
	v, err := (&_SzlDataTreeItem{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_SzlDataTreeItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__szlDataTreeItem SzlDataTreeItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SzlDataTreeItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SzlDataTreeItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemIndex, err := ReadSimpleField(ctx, "itemIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndex' field"))
	}
	m.ItemIndex = itemIndex

	mlfb, err := readBuffer.ReadByteArray("mlfb", int(int32(20)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mlfb' field"))
	}
	m.Mlfb = mlfb

	moduleTypeId, err := ReadSimpleField(ctx, "moduleTypeId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moduleTypeId' field"))
	}
	m.ModuleTypeId = moduleTypeId

	ausbg, err := ReadSimpleField(ctx, "ausbg", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ausbg' field"))
	}
	m.Ausbg = ausbg

	ausbe, err := ReadSimpleField(ctx, "ausbe", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ausbe' field"))
	}
	m.Ausbe = ausbe

	if closeErr := readBuffer.CloseContext("SzlDataTreeItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SzlDataTreeItem")
	}

	return m, nil
}

func (m *_SzlDataTreeItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SzlDataTreeItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SzlDataTreeItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SzlDataTreeItem")
	}

	if err := WriteSimpleField[uint16](ctx, "itemIndex", m.GetItemIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'itemIndex' field")
	}

	if err := WriteByteArrayField(ctx, "mlfb", m.GetMlfb(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'mlfb' field")
	}

	if err := WriteSimpleField[uint16](ctx, "moduleTypeId", m.GetModuleTypeId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'moduleTypeId' field")
	}

	if err := WriteSimpleField[uint16](ctx, "ausbg", m.GetAusbg(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'ausbg' field")
	}

	if err := WriteSimpleField[uint16](ctx, "ausbe", m.GetAusbe(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'ausbe' field")
	}

	if popErr := writeBuffer.PopContext("SzlDataTreeItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SzlDataTreeItem")
	}
	return nil
}

func (m *_SzlDataTreeItem) IsSzlDataTreeItem() {}

func (m *_SzlDataTreeItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SzlDataTreeItem) deepCopy() *_SzlDataTreeItem {
	if m == nil {
		return nil
	}
	_SzlDataTreeItemCopy := &_SzlDataTreeItem{
		m.ItemIndex,
		utils.DeepCopySlice[byte, byte](m.Mlfb),
		m.ModuleTypeId,
		m.Ausbg,
		m.Ausbe,
	}
	return _SzlDataTreeItemCopy
}

func (m *_SzlDataTreeItem) String() string {
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
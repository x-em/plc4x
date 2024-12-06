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

// S7PayloadUserDataItemCyclicServicesSubscribeResponse is the corresponding interface of S7PayloadUserDataItemCyclicServicesSubscribeResponse
type S7PayloadUserDataItemCyclicServicesSubscribeResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetItemsCount returns ItemsCount (property field)
	GetItemsCount() uint16
	// GetItems returns Items (property field)
	GetItems() []AssociatedValueType
	// IsS7PayloadUserDataItemCyclicServicesSubscribeResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCyclicServicesSubscribeResponse()
	// CreateBuilder creates a S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder
	CreateS7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder
}

// _S7PayloadUserDataItemCyclicServicesSubscribeResponse is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesSubscribeResponse struct {
	S7PayloadUserDataItemContract
	ItemsCount uint16
	Items      []AssociatedValueType
}

var _ S7PayloadUserDataItemCyclicServicesSubscribeResponse = (*_S7PayloadUserDataItemCyclicServicesSubscribeResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCyclicServicesSubscribeResponse)(nil)

// NewS7PayloadUserDataItemCyclicServicesSubscribeResponse factory function for _S7PayloadUserDataItemCyclicServicesSubscribeResponse
func NewS7PayloadUserDataItemCyclicServicesSubscribeResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, itemsCount uint16, items []AssociatedValueType) *_S7PayloadUserDataItemCyclicServicesSubscribeResponse {
	_result := &_S7PayloadUserDataItemCyclicServicesSubscribeResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		ItemsCount:                    itemsCount,
		Items:                         items,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder is a builder for S7PayloadUserDataItemCyclicServicesSubscribeResponse
type S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(itemsCount uint16, items []AssociatedValueType) S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder
	// WithItemsCount adds ItemsCount (property field)
	WithItemsCount(uint16) S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder
	// WithItems adds Items (property field)
	WithItems(...AssociatedValueType) S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadUserDataItemBuilder
	// Build builds the S7PayloadUserDataItemCyclicServicesSubscribeResponse or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCyclicServicesSubscribeResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCyclicServicesSubscribeResponse
}

// NewS7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder() creates a S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder
func NewS7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder {
	return &_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder{_S7PayloadUserDataItemCyclicServicesSubscribeResponse: new(_S7PayloadUserDataItemCyclicServicesSubscribeResponse)}
}

type _S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder struct {
	*_S7PayloadUserDataItemCyclicServicesSubscribeResponse

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) = (*_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder)(nil)

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
	contract.(*_S7PayloadUserDataItem)._SubType = b._S7PayloadUserDataItemCyclicServicesSubscribeResponse
}

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) WithMandatoryFields(itemsCount uint16, items []AssociatedValueType) S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder {
	return b.WithItemsCount(itemsCount).WithItems(items...)
}

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) WithItemsCount(itemsCount uint16) S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder {
	b.ItemsCount = itemsCount
	return b
}

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) WithItems(items ...AssociatedValueType) S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder {
	b.Items = items
	return b
}

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) Build() (S7PayloadUserDataItemCyclicServicesSubscribeResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemCyclicServicesSubscribeResponse.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) MustBuild() S7PayloadUserDataItemCyclicServicesSubscribeResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) Done() S7PayloadUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadUserDataItemBuilder().(*_S7PayloadUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder().(*_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder creates a S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder
func (b *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) CreateS7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder()
	}
	return &_S7PayloadUserDataItemCyclicServicesSubscribeResponseBuilder{_S7PayloadUserDataItemCyclicServicesSubscribeResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetItemsCount() uint16 {
	return m.ItemsCount
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetItems() []AssociatedValueType {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesSubscribeResponse(structType any) S7PayloadUserDataItemCyclicServicesSubscribeResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesSubscribeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesSubscribeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesSubscribeResponse"
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (itemsCount)
	lengthInBits += 16

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCyclicServicesSubscribeResponse S7PayloadUserDataItemCyclicServicesSubscribeResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesSubscribeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesSubscribeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemsCount, err := ReadSimpleField(ctx, "itemsCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemsCount' field"))
	}
	m.ItemsCount = itemsCount

	items, err := ReadCountArrayField[AssociatedValueType](ctx, "items", ReadComplex[AssociatedValueType](AssociatedValueTypeParseWithBuffer, readBuffer), uint64(itemsCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesSubscribeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesSubscribeResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesSubscribeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesSubscribeResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "itemsCount", m.GetItemsCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemsCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesSubscribeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesSubscribeResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) IsS7PayloadUserDataItemCyclicServicesSubscribeResponse() {
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) deepCopy() *_S7PayloadUserDataItemCyclicServicesSubscribeResponse {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCyclicServicesSubscribeResponseCopy := &_S7PayloadUserDataItemCyclicServicesSubscribeResponse{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.ItemsCount,
		utils.DeepCopySlice[AssociatedValueType, AssociatedValueType](m.Items),
	}
	_S7PayloadUserDataItemCyclicServicesSubscribeResponseCopy.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCyclicServicesSubscribeResponseCopy
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeResponse) String() string {
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
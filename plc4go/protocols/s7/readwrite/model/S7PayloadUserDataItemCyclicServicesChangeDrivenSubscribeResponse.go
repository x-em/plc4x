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

// S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse is the corresponding interface of S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse
type S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetItemsCount returns ItemsCount (property field)
	GetItemsCount() uint16
	// GetItems returns Items (property field)
	GetItems() []AssociatedQueryValueType
	// IsS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse()
	// CreateBuilder creates a S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder
	CreateS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder
}

// _S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse struct {
	S7PayloadUserDataItemContract
	ItemsCount uint16
	Items      []AssociatedQueryValueType
}

var _ S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse = (*_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse)(nil)

// NewS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse factory function for _S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse
func NewS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, itemsCount uint16, items []AssociatedQueryValueType) *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse {
	_result := &_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse{
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

// S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder is a builder for S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse
type S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(itemsCount uint16, items []AssociatedQueryValueType) S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder
	// WithItemsCount adds ItemsCount (property field)
	WithItemsCount(uint16) S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder
	// WithItems adds Items (property field)
	WithItems(...AssociatedQueryValueType) S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadUserDataItemBuilder
	// Build builds the S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse
}

// NewS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder() creates a S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder
func NewS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder {
	return &_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder{_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse: new(_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse)}
}

type _S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder struct {
	*_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) = (*_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder)(nil)

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
	contract.(*_S7PayloadUserDataItem)._SubType = b._S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse
}

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) WithMandatoryFields(itemsCount uint16, items []AssociatedQueryValueType) S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder {
	return b.WithItemsCount(itemsCount).WithItems(items...)
}

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) WithItemsCount(itemsCount uint16) S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder {
	b.ItemsCount = itemsCount
	return b
}

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) WithItems(items ...AssociatedQueryValueType) S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder {
	b.Items = items
	return b
}

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) Build() (S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) MustBuild() S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) Done() S7PayloadUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadUserDataItemBuilder().(*_S7PayloadUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder().(*_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder creates a S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder
func (b *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) CreateS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder()
	}
	return &_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseBuilder{_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetCpuSubfunction() uint8 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetItemsCount() uint16 {
	return m.ItemsCount
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetItems() []AssociatedQueryValueType {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse(structType any) S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse"
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemsCount, err := ReadSimpleField(ctx, "itemsCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemsCount' field"))
	}
	m.ItemsCount = itemsCount

	items, err := ReadCountArrayField[AssociatedQueryValueType](ctx, "items", ReadComplex[AssociatedQueryValueType](AssociatedQueryValueTypeParseWithBuffer, readBuffer), uint64(itemsCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "itemsCount", m.GetItemsCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemsCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) IsS7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse() {
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) deepCopy() *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseCopy := &_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.ItemsCount,
		utils.DeepCopySlice[AssociatedQueryValueType, AssociatedQueryValueType](m.Items),
	}
	_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseCopy.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseCopy
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) String() string {
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
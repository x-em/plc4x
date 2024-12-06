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

// ModbusPDUWriteFileRecordResponse is the corresponding interface of ModbusPDUWriteFileRecordResponse
type ModbusPDUWriteFileRecordResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetItems returns Items (property field)
	GetItems() []ModbusPDUWriteFileRecordResponseItem
	// IsModbusPDUWriteFileRecordResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUWriteFileRecordResponse()
	// CreateBuilder creates a ModbusPDUWriteFileRecordResponseBuilder
	CreateModbusPDUWriteFileRecordResponseBuilder() ModbusPDUWriteFileRecordResponseBuilder
}

// _ModbusPDUWriteFileRecordResponse is the data-structure of this message
type _ModbusPDUWriteFileRecordResponse struct {
	ModbusPDUContract
	Items []ModbusPDUWriteFileRecordResponseItem
}

var _ ModbusPDUWriteFileRecordResponse = (*_ModbusPDUWriteFileRecordResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUWriteFileRecordResponse)(nil)

// NewModbusPDUWriteFileRecordResponse factory function for _ModbusPDUWriteFileRecordResponse
func NewModbusPDUWriteFileRecordResponse(items []ModbusPDUWriteFileRecordResponseItem) *_ModbusPDUWriteFileRecordResponse {
	_result := &_ModbusPDUWriteFileRecordResponse{
		ModbusPDUContract: NewModbusPDU(),
		Items:             items,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUWriteFileRecordResponseBuilder is a builder for ModbusPDUWriteFileRecordResponse
type ModbusPDUWriteFileRecordResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(items []ModbusPDUWriteFileRecordResponseItem) ModbusPDUWriteFileRecordResponseBuilder
	// WithItems adds Items (property field)
	WithItems(...ModbusPDUWriteFileRecordResponseItem) ModbusPDUWriteFileRecordResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUWriteFileRecordResponse or returns an error if something is wrong
	Build() (ModbusPDUWriteFileRecordResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUWriteFileRecordResponse
}

// NewModbusPDUWriteFileRecordResponseBuilder() creates a ModbusPDUWriteFileRecordResponseBuilder
func NewModbusPDUWriteFileRecordResponseBuilder() ModbusPDUWriteFileRecordResponseBuilder {
	return &_ModbusPDUWriteFileRecordResponseBuilder{_ModbusPDUWriteFileRecordResponse: new(_ModbusPDUWriteFileRecordResponse)}
}

type _ModbusPDUWriteFileRecordResponseBuilder struct {
	*_ModbusPDUWriteFileRecordResponse

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUWriteFileRecordResponseBuilder) = (*_ModbusPDUWriteFileRecordResponseBuilder)(nil)

func (b *_ModbusPDUWriteFileRecordResponseBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUWriteFileRecordResponse
}

func (b *_ModbusPDUWriteFileRecordResponseBuilder) WithMandatoryFields(items []ModbusPDUWriteFileRecordResponseItem) ModbusPDUWriteFileRecordResponseBuilder {
	return b.WithItems(items...)
}

func (b *_ModbusPDUWriteFileRecordResponseBuilder) WithItems(items ...ModbusPDUWriteFileRecordResponseItem) ModbusPDUWriteFileRecordResponseBuilder {
	b.Items = items
	return b
}

func (b *_ModbusPDUWriteFileRecordResponseBuilder) Build() (ModbusPDUWriteFileRecordResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUWriteFileRecordResponse.deepCopy(), nil
}

func (b *_ModbusPDUWriteFileRecordResponseBuilder) MustBuild() ModbusPDUWriteFileRecordResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUWriteFileRecordResponseBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUWriteFileRecordResponseBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUWriteFileRecordResponseBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUWriteFileRecordResponseBuilder().(*_ModbusPDUWriteFileRecordResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUWriteFileRecordResponseBuilder creates a ModbusPDUWriteFileRecordResponseBuilder
func (b *_ModbusPDUWriteFileRecordResponse) CreateModbusPDUWriteFileRecordResponseBuilder() ModbusPDUWriteFileRecordResponseBuilder {
	if b == nil {
		return NewModbusPDUWriteFileRecordResponseBuilder()
	}
	return &_ModbusPDUWriteFileRecordResponseBuilder{_ModbusPDUWriteFileRecordResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteFileRecordResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteFileRecordResponse) GetFunctionFlag() uint8 {
	return 0x15
}

func (m *_ModbusPDUWriteFileRecordResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteFileRecordResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteFileRecordResponse) GetItems() []ModbusPDUWriteFileRecordResponseItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteFileRecordResponse(structType any) ModbusPDUWriteFileRecordResponse {
	if casted, ok := structType.(ModbusPDUWriteFileRecordResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteFileRecordResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteFileRecordResponse) GetTypeName() string {
	return "ModbusPDUWriteFileRecordResponse"
}

func (m *_ModbusPDUWriteFileRecordResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_ModbusPDUWriteFileRecordResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUWriteFileRecordResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUWriteFileRecordResponse ModbusPDUWriteFileRecordResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteFileRecordResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	items, err := ReadLengthArrayField[ModbusPDUWriteFileRecordResponseItem](ctx, "items", ReadComplex[ModbusPDUWriteFileRecordResponseItem](ModbusPDUWriteFileRecordResponseItemParseWithBuffer, readBuffer), int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteFileRecordResponse")
	}

	return m, nil
}

func (m *_ModbusPDUWriteFileRecordResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteFileRecordResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	itemsArraySizeInBytes := func(items []ModbusPDUWriteFileRecordResponseItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.GetLengthInBytes(ctx))
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteFileRecordResponse")
		}
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.GetItems())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteFileRecordResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteFileRecordResponse) IsModbusPDUWriteFileRecordResponse() {}

func (m *_ModbusPDUWriteFileRecordResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUWriteFileRecordResponse) deepCopy() *_ModbusPDUWriteFileRecordResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUWriteFileRecordResponseCopy := &_ModbusPDUWriteFileRecordResponse{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		utils.DeepCopySlice[ModbusPDUWriteFileRecordResponseItem, ModbusPDUWriteFileRecordResponseItem](m.Items),
	}
	_ModbusPDUWriteFileRecordResponseCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUWriteFileRecordResponseCopy
}

func (m *_ModbusPDUWriteFileRecordResponse) String() string {
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
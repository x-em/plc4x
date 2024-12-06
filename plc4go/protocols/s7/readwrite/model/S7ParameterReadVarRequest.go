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

// S7ParameterReadVarRequest is the corresponding interface of S7ParameterReadVarRequest
type S7ParameterReadVarRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7Parameter
	// GetItems returns Items (property field)
	GetItems() []S7VarRequestParameterItem
	// IsS7ParameterReadVarRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterReadVarRequest()
	// CreateBuilder creates a S7ParameterReadVarRequestBuilder
	CreateS7ParameterReadVarRequestBuilder() S7ParameterReadVarRequestBuilder
}

// _S7ParameterReadVarRequest is the data-structure of this message
type _S7ParameterReadVarRequest struct {
	S7ParameterContract
	Items []S7VarRequestParameterItem
}

var _ S7ParameterReadVarRequest = (*_S7ParameterReadVarRequest)(nil)
var _ S7ParameterRequirements = (*_S7ParameterReadVarRequest)(nil)

// NewS7ParameterReadVarRequest factory function for _S7ParameterReadVarRequest
func NewS7ParameterReadVarRequest(items []S7VarRequestParameterItem) *_S7ParameterReadVarRequest {
	_result := &_S7ParameterReadVarRequest{
		S7ParameterContract: NewS7Parameter(),
		Items:               items,
	}
	_result.S7ParameterContract.(*_S7Parameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7ParameterReadVarRequestBuilder is a builder for S7ParameterReadVarRequest
type S7ParameterReadVarRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(items []S7VarRequestParameterItem) S7ParameterReadVarRequestBuilder
	// WithItems adds Items (property field)
	WithItems(...S7VarRequestParameterItem) S7ParameterReadVarRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7ParameterBuilder
	// Build builds the S7ParameterReadVarRequest or returns an error if something is wrong
	Build() (S7ParameterReadVarRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7ParameterReadVarRequest
}

// NewS7ParameterReadVarRequestBuilder() creates a S7ParameterReadVarRequestBuilder
func NewS7ParameterReadVarRequestBuilder() S7ParameterReadVarRequestBuilder {
	return &_S7ParameterReadVarRequestBuilder{_S7ParameterReadVarRequest: new(_S7ParameterReadVarRequest)}
}

type _S7ParameterReadVarRequestBuilder struct {
	*_S7ParameterReadVarRequest

	parentBuilder *_S7ParameterBuilder

	err *utils.MultiError
}

var _ (S7ParameterReadVarRequestBuilder) = (*_S7ParameterReadVarRequestBuilder)(nil)

func (b *_S7ParameterReadVarRequestBuilder) setParent(contract S7ParameterContract) {
	b.S7ParameterContract = contract
	contract.(*_S7Parameter)._SubType = b._S7ParameterReadVarRequest
}

func (b *_S7ParameterReadVarRequestBuilder) WithMandatoryFields(items []S7VarRequestParameterItem) S7ParameterReadVarRequestBuilder {
	return b.WithItems(items...)
}

func (b *_S7ParameterReadVarRequestBuilder) WithItems(items ...S7VarRequestParameterItem) S7ParameterReadVarRequestBuilder {
	b.Items = items
	return b
}

func (b *_S7ParameterReadVarRequestBuilder) Build() (S7ParameterReadVarRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7ParameterReadVarRequest.deepCopy(), nil
}

func (b *_S7ParameterReadVarRequestBuilder) MustBuild() S7ParameterReadVarRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7ParameterReadVarRequestBuilder) Done() S7ParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7ParameterBuilder().(*_S7ParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_S7ParameterReadVarRequestBuilder) buildForS7Parameter() (S7Parameter, error) {
	return b.Build()
}

func (b *_S7ParameterReadVarRequestBuilder) DeepCopy() any {
	_copy := b.CreateS7ParameterReadVarRequestBuilder().(*_S7ParameterReadVarRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7ParameterReadVarRequestBuilder creates a S7ParameterReadVarRequestBuilder
func (b *_S7ParameterReadVarRequest) CreateS7ParameterReadVarRequestBuilder() S7ParameterReadVarRequestBuilder {
	if b == nil {
		return NewS7ParameterReadVarRequestBuilder()
	}
	return &_S7ParameterReadVarRequestBuilder{_S7ParameterReadVarRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterReadVarRequest) GetParameterType() uint8 {
	return 0x04
}

func (m *_S7ParameterReadVarRequest) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterReadVarRequest) GetParent() S7ParameterContract {
	return m.S7ParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterReadVarRequest) GetItems() []S7VarRequestParameterItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7ParameterReadVarRequest(structType any) S7ParameterReadVarRequest {
	if casted, ok := structType.(S7ParameterReadVarRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterReadVarRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterReadVarRequest) GetTypeName() string {
	return "S7ParameterReadVarRequest"
}

func (m *_S7ParameterReadVarRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterContract.(*_S7Parameter).getLengthInBits(ctx))

	// Implicit Field (numItems)
	lengthInBits += 8

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

func (m *_S7ParameterReadVarRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterReadVarRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Parameter, messageType uint8) (__s7ParameterReadVarRequest S7ParameterReadVarRequest, err error) {
	m.S7ParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterReadVarRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterReadVarRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numItems, err := ReadImplicitField[uint8](ctx, "numItems", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numItems' field"))
	}
	_ = numItems

	items, err := ReadCountArrayField[S7VarRequestParameterItem](ctx, "items", ReadComplex[S7VarRequestParameterItem](S7VarRequestParameterItemParseWithBuffer, readBuffer), uint64(numItems))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7ParameterReadVarRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterReadVarRequest")
	}

	return m, nil
}

func (m *_S7ParameterReadVarRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterReadVarRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterReadVarRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterReadVarRequest")
		}
		numItems := uint8(uint8(len(m.GetItems())))
		if err := WriteImplicitField(ctx, "numItems", numItems, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numItems' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterReadVarRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterReadVarRequest")
		}
		return nil
	}
	return m.S7ParameterContract.(*_S7Parameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterReadVarRequest) IsS7ParameterReadVarRequest() {}

func (m *_S7ParameterReadVarRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7ParameterReadVarRequest) deepCopy() *_S7ParameterReadVarRequest {
	if m == nil {
		return nil
	}
	_S7ParameterReadVarRequestCopy := &_S7ParameterReadVarRequest{
		m.S7ParameterContract.(*_S7Parameter).deepCopy(),
		utils.DeepCopySlice[S7VarRequestParameterItem, S7VarRequestParameterItem](m.Items),
	}
	_S7ParameterReadVarRequestCopy.S7ParameterContract.(*_S7Parameter)._SubType = m
	return _S7ParameterReadVarRequestCopy
}

func (m *_S7ParameterReadVarRequest) String() string {
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
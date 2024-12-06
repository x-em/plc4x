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

// DF1UnprotectedReadRequest is the corresponding interface of DF1UnprotectedReadRequest
type DF1UnprotectedReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	DF1Command
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetSize returns Size (property field)
	GetSize() uint8
	// IsDF1UnprotectedReadRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1UnprotectedReadRequest()
	// CreateBuilder creates a DF1UnprotectedReadRequestBuilder
	CreateDF1UnprotectedReadRequestBuilder() DF1UnprotectedReadRequestBuilder
}

// _DF1UnprotectedReadRequest is the data-structure of this message
type _DF1UnprotectedReadRequest struct {
	DF1CommandContract
	Address uint16
	Size    uint8
}

var _ DF1UnprotectedReadRequest = (*_DF1UnprotectedReadRequest)(nil)
var _ DF1CommandRequirements = (*_DF1UnprotectedReadRequest)(nil)

// NewDF1UnprotectedReadRequest factory function for _DF1UnprotectedReadRequest
func NewDF1UnprotectedReadRequest(status uint8, transactionCounter uint16, address uint16, size uint8) *_DF1UnprotectedReadRequest {
	_result := &_DF1UnprotectedReadRequest{
		DF1CommandContract: NewDF1Command(status, transactionCounter),
		Address:            address,
		Size:               size,
	}
	_result.DF1CommandContract.(*_DF1Command)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DF1UnprotectedReadRequestBuilder is a builder for DF1UnprotectedReadRequest
type DF1UnprotectedReadRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(address uint16, size uint8) DF1UnprotectedReadRequestBuilder
	// WithAddress adds Address (property field)
	WithAddress(uint16) DF1UnprotectedReadRequestBuilder
	// WithSize adds Size (property field)
	WithSize(uint8) DF1UnprotectedReadRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() DF1CommandBuilder
	// Build builds the DF1UnprotectedReadRequest or returns an error if something is wrong
	Build() (DF1UnprotectedReadRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DF1UnprotectedReadRequest
}

// NewDF1UnprotectedReadRequestBuilder() creates a DF1UnprotectedReadRequestBuilder
func NewDF1UnprotectedReadRequestBuilder() DF1UnprotectedReadRequestBuilder {
	return &_DF1UnprotectedReadRequestBuilder{_DF1UnprotectedReadRequest: new(_DF1UnprotectedReadRequest)}
}

type _DF1UnprotectedReadRequestBuilder struct {
	*_DF1UnprotectedReadRequest

	parentBuilder *_DF1CommandBuilder

	err *utils.MultiError
}

var _ (DF1UnprotectedReadRequestBuilder) = (*_DF1UnprotectedReadRequestBuilder)(nil)

func (b *_DF1UnprotectedReadRequestBuilder) setParent(contract DF1CommandContract) {
	b.DF1CommandContract = contract
	contract.(*_DF1Command)._SubType = b._DF1UnprotectedReadRequest
}

func (b *_DF1UnprotectedReadRequestBuilder) WithMandatoryFields(address uint16, size uint8) DF1UnprotectedReadRequestBuilder {
	return b.WithAddress(address).WithSize(size)
}

func (b *_DF1UnprotectedReadRequestBuilder) WithAddress(address uint16) DF1UnprotectedReadRequestBuilder {
	b.Address = address
	return b
}

func (b *_DF1UnprotectedReadRequestBuilder) WithSize(size uint8) DF1UnprotectedReadRequestBuilder {
	b.Size = size
	return b
}

func (b *_DF1UnprotectedReadRequestBuilder) Build() (DF1UnprotectedReadRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DF1UnprotectedReadRequest.deepCopy(), nil
}

func (b *_DF1UnprotectedReadRequestBuilder) MustBuild() DF1UnprotectedReadRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DF1UnprotectedReadRequestBuilder) Done() DF1CommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewDF1CommandBuilder().(*_DF1CommandBuilder)
	}
	return b.parentBuilder
}

func (b *_DF1UnprotectedReadRequestBuilder) buildForDF1Command() (DF1Command, error) {
	return b.Build()
}

func (b *_DF1UnprotectedReadRequestBuilder) DeepCopy() any {
	_copy := b.CreateDF1UnprotectedReadRequestBuilder().(*_DF1UnprotectedReadRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDF1UnprotectedReadRequestBuilder creates a DF1UnprotectedReadRequestBuilder
func (b *_DF1UnprotectedReadRequest) CreateDF1UnprotectedReadRequestBuilder() DF1UnprotectedReadRequestBuilder {
	if b == nil {
		return NewDF1UnprotectedReadRequestBuilder()
	}
	return &_DF1UnprotectedReadRequestBuilder{_DF1UnprotectedReadRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1UnprotectedReadRequest) GetCommandCode() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1UnprotectedReadRequest) GetParent() DF1CommandContract {
	return m.DF1CommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1UnprotectedReadRequest) GetAddress() uint16 {
	return m.Address
}

func (m *_DF1UnprotectedReadRequest) GetSize() uint8 {
	return m.Size
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDF1UnprotectedReadRequest(structType any) DF1UnprotectedReadRequest {
	if casted, ok := structType.(DF1UnprotectedReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DF1UnprotectedReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DF1UnprotectedReadRequest) GetTypeName() string {
	return "DF1UnprotectedReadRequest"
}

func (m *_DF1UnprotectedReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1CommandContract.(*_DF1Command).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 16

	// Simple field (size)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DF1UnprotectedReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1UnprotectedReadRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1Command) (__dF1UnprotectedReadRequest DF1UnprotectedReadRequest, err error) {
	m.DF1CommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1UnprotectedReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1UnprotectedReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	size, err := ReadSimpleField(ctx, "size", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'size' field"))
	}
	m.Size = size

	if closeErr := readBuffer.CloseContext("DF1UnprotectedReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1UnprotectedReadRequest")
	}

	return m, nil
}

func (m *_DF1UnprotectedReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1UnprotectedReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1UnprotectedReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1UnprotectedReadRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "address", m.GetAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[uint8](ctx, "size", m.GetSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'size' field")
		}

		if popErr := writeBuffer.PopContext("DF1UnprotectedReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1UnprotectedReadRequest")
		}
		return nil
	}
	return m.DF1CommandContract.(*_DF1Command).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1UnprotectedReadRequest) IsDF1UnprotectedReadRequest() {}

func (m *_DF1UnprotectedReadRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DF1UnprotectedReadRequest) deepCopy() *_DF1UnprotectedReadRequest {
	if m == nil {
		return nil
	}
	_DF1UnprotectedReadRequestCopy := &_DF1UnprotectedReadRequest{
		m.DF1CommandContract.(*_DF1Command).deepCopy(),
		m.Address,
		m.Size,
	}
	_DF1UnprotectedReadRequestCopy.DF1CommandContract.(*_DF1Command)._SubType = m
	return _DF1UnprotectedReadRequestCopy
}

func (m *_DF1UnprotectedReadRequest) String() string {
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
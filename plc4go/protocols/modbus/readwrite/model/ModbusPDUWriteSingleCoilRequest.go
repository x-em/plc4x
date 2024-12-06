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

// ModbusPDUWriteSingleCoilRequest is the corresponding interface of ModbusPDUWriteSingleCoilRequest
type ModbusPDUWriteSingleCoilRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetValue returns Value (property field)
	GetValue() uint16
	// IsModbusPDUWriteSingleCoilRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUWriteSingleCoilRequest()
	// CreateBuilder creates a ModbusPDUWriteSingleCoilRequestBuilder
	CreateModbusPDUWriteSingleCoilRequestBuilder() ModbusPDUWriteSingleCoilRequestBuilder
}

// _ModbusPDUWriteSingleCoilRequest is the data-structure of this message
type _ModbusPDUWriteSingleCoilRequest struct {
	ModbusPDUContract
	Address uint16
	Value   uint16
}

var _ ModbusPDUWriteSingleCoilRequest = (*_ModbusPDUWriteSingleCoilRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUWriteSingleCoilRequest)(nil)

// NewModbusPDUWriteSingleCoilRequest factory function for _ModbusPDUWriteSingleCoilRequest
func NewModbusPDUWriteSingleCoilRequest(address uint16, value uint16) *_ModbusPDUWriteSingleCoilRequest {
	_result := &_ModbusPDUWriteSingleCoilRequest{
		ModbusPDUContract: NewModbusPDU(),
		Address:           address,
		Value:             value,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUWriteSingleCoilRequestBuilder is a builder for ModbusPDUWriteSingleCoilRequest
type ModbusPDUWriteSingleCoilRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(address uint16, value uint16) ModbusPDUWriteSingleCoilRequestBuilder
	// WithAddress adds Address (property field)
	WithAddress(uint16) ModbusPDUWriteSingleCoilRequestBuilder
	// WithValue adds Value (property field)
	WithValue(uint16) ModbusPDUWriteSingleCoilRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUWriteSingleCoilRequest or returns an error if something is wrong
	Build() (ModbusPDUWriteSingleCoilRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUWriteSingleCoilRequest
}

// NewModbusPDUWriteSingleCoilRequestBuilder() creates a ModbusPDUWriteSingleCoilRequestBuilder
func NewModbusPDUWriteSingleCoilRequestBuilder() ModbusPDUWriteSingleCoilRequestBuilder {
	return &_ModbusPDUWriteSingleCoilRequestBuilder{_ModbusPDUWriteSingleCoilRequest: new(_ModbusPDUWriteSingleCoilRequest)}
}

type _ModbusPDUWriteSingleCoilRequestBuilder struct {
	*_ModbusPDUWriteSingleCoilRequest

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUWriteSingleCoilRequestBuilder) = (*_ModbusPDUWriteSingleCoilRequestBuilder)(nil)

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUWriteSingleCoilRequest
}

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) WithMandatoryFields(address uint16, value uint16) ModbusPDUWriteSingleCoilRequestBuilder {
	return b.WithAddress(address).WithValue(value)
}

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) WithAddress(address uint16) ModbusPDUWriteSingleCoilRequestBuilder {
	b.Address = address
	return b
}

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) WithValue(value uint16) ModbusPDUWriteSingleCoilRequestBuilder {
	b.Value = value
	return b
}

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) Build() (ModbusPDUWriteSingleCoilRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUWriteSingleCoilRequest.deepCopy(), nil
}

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) MustBuild() ModbusPDUWriteSingleCoilRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUWriteSingleCoilRequestBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUWriteSingleCoilRequestBuilder().(*_ModbusPDUWriteSingleCoilRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUWriteSingleCoilRequestBuilder creates a ModbusPDUWriteSingleCoilRequestBuilder
func (b *_ModbusPDUWriteSingleCoilRequest) CreateModbusPDUWriteSingleCoilRequestBuilder() ModbusPDUWriteSingleCoilRequestBuilder {
	if b == nil {
		return NewModbusPDUWriteSingleCoilRequestBuilder()
	}
	return &_ModbusPDUWriteSingleCoilRequestBuilder{_ModbusPDUWriteSingleCoilRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteSingleCoilRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteSingleCoilRequest) GetFunctionFlag() uint8 {
	return 0x05
}

func (m *_ModbusPDUWriteSingleCoilRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteSingleCoilRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteSingleCoilRequest) GetAddress() uint16 {
	return m.Address
}

func (m *_ModbusPDUWriteSingleCoilRequest) GetValue() uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteSingleCoilRequest(structType any) ModbusPDUWriteSingleCoilRequest {
	if casted, ok := structType.(ModbusPDUWriteSingleCoilRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteSingleCoilRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteSingleCoilRequest) GetTypeName() string {
	return "ModbusPDUWriteSingleCoilRequest"
}

func (m *_ModbusPDUWriteSingleCoilRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUWriteSingleCoilRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUWriteSingleCoilRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUWriteSingleCoilRequest ModbusPDUWriteSingleCoilRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleCoilRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteSingleCoilRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	value, err := ReadSimpleField(ctx, "value", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleCoilRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteSingleCoilRequest")
	}

	return m, nil
}

func (m *_ModbusPDUWriteSingleCoilRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteSingleCoilRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleCoilRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteSingleCoilRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "address", m.GetAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[uint16](ctx, "value", m.GetValue(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleCoilRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteSingleCoilRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteSingleCoilRequest) IsModbusPDUWriteSingleCoilRequest() {}

func (m *_ModbusPDUWriteSingleCoilRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUWriteSingleCoilRequest) deepCopy() *_ModbusPDUWriteSingleCoilRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUWriteSingleCoilRequestCopy := &_ModbusPDUWriteSingleCoilRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.Address,
		m.Value,
	}
	_ModbusPDUWriteSingleCoilRequestCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUWriteSingleCoilRequestCopy
}

func (m *_ModbusPDUWriteSingleCoilRequest) String() string {
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
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

// ModbusPDUGetComEventCounterResponse is the corresponding interface of ModbusPDUGetComEventCounterResponse
type ModbusPDUGetComEventCounterResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetEventCount returns EventCount (property field)
	GetEventCount() uint16
	// IsModbusPDUGetComEventCounterResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUGetComEventCounterResponse()
	// CreateBuilder creates a ModbusPDUGetComEventCounterResponseBuilder
	CreateModbusPDUGetComEventCounterResponseBuilder() ModbusPDUGetComEventCounterResponseBuilder
}

// _ModbusPDUGetComEventCounterResponse is the data-structure of this message
type _ModbusPDUGetComEventCounterResponse struct {
	ModbusPDUContract
	Status     uint16
	EventCount uint16
}

var _ ModbusPDUGetComEventCounterResponse = (*_ModbusPDUGetComEventCounterResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUGetComEventCounterResponse)(nil)

// NewModbusPDUGetComEventCounterResponse factory function for _ModbusPDUGetComEventCounterResponse
func NewModbusPDUGetComEventCounterResponse(status uint16, eventCount uint16) *_ModbusPDUGetComEventCounterResponse {
	_result := &_ModbusPDUGetComEventCounterResponse{
		ModbusPDUContract: NewModbusPDU(),
		Status:            status,
		EventCount:        eventCount,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUGetComEventCounterResponseBuilder is a builder for ModbusPDUGetComEventCounterResponse
type ModbusPDUGetComEventCounterResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(status uint16, eventCount uint16) ModbusPDUGetComEventCounterResponseBuilder
	// WithStatus adds Status (property field)
	WithStatus(uint16) ModbusPDUGetComEventCounterResponseBuilder
	// WithEventCount adds EventCount (property field)
	WithEventCount(uint16) ModbusPDUGetComEventCounterResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUGetComEventCounterResponse or returns an error if something is wrong
	Build() (ModbusPDUGetComEventCounterResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUGetComEventCounterResponse
}

// NewModbusPDUGetComEventCounterResponseBuilder() creates a ModbusPDUGetComEventCounterResponseBuilder
func NewModbusPDUGetComEventCounterResponseBuilder() ModbusPDUGetComEventCounterResponseBuilder {
	return &_ModbusPDUGetComEventCounterResponseBuilder{_ModbusPDUGetComEventCounterResponse: new(_ModbusPDUGetComEventCounterResponse)}
}

type _ModbusPDUGetComEventCounterResponseBuilder struct {
	*_ModbusPDUGetComEventCounterResponse

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUGetComEventCounterResponseBuilder) = (*_ModbusPDUGetComEventCounterResponseBuilder)(nil)

func (b *_ModbusPDUGetComEventCounterResponseBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUGetComEventCounterResponse
}

func (b *_ModbusPDUGetComEventCounterResponseBuilder) WithMandatoryFields(status uint16, eventCount uint16) ModbusPDUGetComEventCounterResponseBuilder {
	return b.WithStatus(status).WithEventCount(eventCount)
}

func (b *_ModbusPDUGetComEventCounterResponseBuilder) WithStatus(status uint16) ModbusPDUGetComEventCounterResponseBuilder {
	b.Status = status
	return b
}

func (b *_ModbusPDUGetComEventCounterResponseBuilder) WithEventCount(eventCount uint16) ModbusPDUGetComEventCounterResponseBuilder {
	b.EventCount = eventCount
	return b
}

func (b *_ModbusPDUGetComEventCounterResponseBuilder) Build() (ModbusPDUGetComEventCounterResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUGetComEventCounterResponse.deepCopy(), nil
}

func (b *_ModbusPDUGetComEventCounterResponseBuilder) MustBuild() ModbusPDUGetComEventCounterResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUGetComEventCounterResponseBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUGetComEventCounterResponseBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUGetComEventCounterResponseBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUGetComEventCounterResponseBuilder().(*_ModbusPDUGetComEventCounterResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUGetComEventCounterResponseBuilder creates a ModbusPDUGetComEventCounterResponseBuilder
func (b *_ModbusPDUGetComEventCounterResponse) CreateModbusPDUGetComEventCounterResponseBuilder() ModbusPDUGetComEventCounterResponseBuilder {
	if b == nil {
		return NewModbusPDUGetComEventCounterResponseBuilder()
	}
	return &_ModbusPDUGetComEventCounterResponseBuilder{_ModbusPDUGetComEventCounterResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUGetComEventCounterResponse) GetFunctionFlag() uint8 {
	return 0x0B
}

func (m *_ModbusPDUGetComEventCounterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) GetStatus() uint16 {
	return m.Status
}

func (m *_ModbusPDUGetComEventCounterResponse) GetEventCount() uint16 {
	return m.EventCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUGetComEventCounterResponse(structType any) ModbusPDUGetComEventCounterResponse {
	if casted, ok := structType.(ModbusPDUGetComEventCounterResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventCounterResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUGetComEventCounterResponse) GetTypeName() string {
	return "ModbusPDUGetComEventCounterResponse"
}

func (m *_ModbusPDUGetComEventCounterResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUGetComEventCounterResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUGetComEventCounterResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUGetComEventCounterResponse ModbusPDUGetComEventCounterResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventCounterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUGetComEventCounterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	eventCount, err := ReadSimpleField(ctx, "eventCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventCount' field"))
	}
	m.EventCount = eventCount

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventCounterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUGetComEventCounterResponse")
	}

	return m, nil
}

func (m *_ModbusPDUGetComEventCounterResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUGetComEventCounterResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventCounterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUGetComEventCounterResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "status", m.GetStatus(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint16](ctx, "eventCount", m.GetEventCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventCount' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventCounterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUGetComEventCounterResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUGetComEventCounterResponse) IsModbusPDUGetComEventCounterResponse() {}

func (m *_ModbusPDUGetComEventCounterResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUGetComEventCounterResponse) deepCopy() *_ModbusPDUGetComEventCounterResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUGetComEventCounterResponseCopy := &_ModbusPDUGetComEventCounterResponse{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.Status,
		m.EventCount,
	}
	_ModbusPDUGetComEventCounterResponseCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUGetComEventCounterResponseCopy
}

func (m *_ModbusPDUGetComEventCounterResponse) String() string {
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
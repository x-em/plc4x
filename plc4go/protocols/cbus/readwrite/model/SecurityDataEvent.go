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

// SecurityDataEvent is the corresponding interface of SecurityDataEvent
type SecurityDataEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetData returns Data (property field)
	GetData() []byte
	// IsSecurityDataEvent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataEvent()
	// CreateBuilder creates a SecurityDataEventBuilder
	CreateSecurityDataEventBuilder() SecurityDataEventBuilder
}

// _SecurityDataEvent is the data-structure of this message
type _SecurityDataEvent struct {
	SecurityDataContract
	Data []byte
}

var _ SecurityDataEvent = (*_SecurityDataEvent)(nil)
var _ SecurityDataRequirements = (*_SecurityDataEvent)(nil)

// NewSecurityDataEvent factory function for _SecurityDataEvent
func NewSecurityDataEvent(commandTypeContainer SecurityCommandTypeContainer, argument byte, data []byte) *_SecurityDataEvent {
	_result := &_SecurityDataEvent{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		Data:                 data,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataEventBuilder is a builder for SecurityDataEvent
type SecurityDataEventBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(data []byte) SecurityDataEventBuilder
	// WithData adds Data (property field)
	WithData(...byte) SecurityDataEventBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataEvent or returns an error if something is wrong
	Build() (SecurityDataEvent, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataEvent
}

// NewSecurityDataEventBuilder() creates a SecurityDataEventBuilder
func NewSecurityDataEventBuilder() SecurityDataEventBuilder {
	return &_SecurityDataEventBuilder{_SecurityDataEvent: new(_SecurityDataEvent)}
}

type _SecurityDataEventBuilder struct {
	*_SecurityDataEvent

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataEventBuilder) = (*_SecurityDataEventBuilder)(nil)

func (b *_SecurityDataEventBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataEvent
}

func (b *_SecurityDataEventBuilder) WithMandatoryFields(data []byte) SecurityDataEventBuilder {
	return b.WithData(data...)
}

func (b *_SecurityDataEventBuilder) WithData(data ...byte) SecurityDataEventBuilder {
	b.Data = data
	return b
}

func (b *_SecurityDataEventBuilder) Build() (SecurityDataEvent, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataEvent.deepCopy(), nil
}

func (b *_SecurityDataEventBuilder) MustBuild() SecurityDataEvent {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataEventBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataEventBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataEventBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataEventBuilder().(*_SecurityDataEventBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataEventBuilder creates a SecurityDataEventBuilder
func (b *_SecurityDataEvent) CreateSecurityDataEventBuilder() SecurityDataEventBuilder {
	if b == nil {
		return NewSecurityDataEventBuilder()
	}
	return &_SecurityDataEventBuilder{_SecurityDataEvent: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataEvent) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataEvent) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataEvent(structType any) SecurityDataEvent {
	if casted, ok := structType.(SecurityDataEvent); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataEvent); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataEvent) GetTypeName() string {
	return "SecurityDataEvent"
}

func (m *_SecurityDataEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_SecurityDataEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataEvent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData, commandTypeContainer SecurityCommandTypeContainer) (__securityDataEvent SecurityDataEvent, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := readBuffer.ReadByteArray("data", int(int32(commandTypeContainer.NumBytes())-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("SecurityDataEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataEvent")
	}

	return m, nil
}

func (m *_SecurityDataEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataEvent")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataEvent")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataEvent) IsSecurityDataEvent() {}

func (m *_SecurityDataEvent) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataEvent) deepCopy() *_SecurityDataEvent {
	if m == nil {
		return nil
	}
	_SecurityDataEventCopy := &_SecurityDataEvent{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	_SecurityDataEventCopy.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataEventCopy
}

func (m *_SecurityDataEvent) String() string {
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
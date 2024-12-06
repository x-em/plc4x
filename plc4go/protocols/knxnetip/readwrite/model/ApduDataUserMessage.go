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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataUserMessage is the corresponding interface of ApduDataUserMessage
type ApduDataUserMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// IsApduDataUserMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataUserMessage()
	// CreateBuilder creates a ApduDataUserMessageBuilder
	CreateApduDataUserMessageBuilder() ApduDataUserMessageBuilder
}

// _ApduDataUserMessage is the data-structure of this message
type _ApduDataUserMessage struct {
	ApduDataContract
}

var _ ApduDataUserMessage = (*_ApduDataUserMessage)(nil)
var _ ApduDataRequirements = (*_ApduDataUserMessage)(nil)

// NewApduDataUserMessage factory function for _ApduDataUserMessage
func NewApduDataUserMessage(dataLength uint8) *_ApduDataUserMessage {
	_result := &_ApduDataUserMessage{
		ApduDataContract: NewApduData(dataLength),
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataUserMessageBuilder is a builder for ApduDataUserMessage
type ApduDataUserMessageBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataUserMessageBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataBuilder
	// Build builds the ApduDataUserMessage or returns an error if something is wrong
	Build() (ApduDataUserMessage, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataUserMessage
}

// NewApduDataUserMessageBuilder() creates a ApduDataUserMessageBuilder
func NewApduDataUserMessageBuilder() ApduDataUserMessageBuilder {
	return &_ApduDataUserMessageBuilder{_ApduDataUserMessage: new(_ApduDataUserMessage)}
}

type _ApduDataUserMessageBuilder struct {
	*_ApduDataUserMessage

	parentBuilder *_ApduDataBuilder

	err *utils.MultiError
}

var _ (ApduDataUserMessageBuilder) = (*_ApduDataUserMessageBuilder)(nil)

func (b *_ApduDataUserMessageBuilder) setParent(contract ApduDataContract) {
	b.ApduDataContract = contract
	contract.(*_ApduData)._SubType = b._ApduDataUserMessage
}

func (b *_ApduDataUserMessageBuilder) WithMandatoryFields() ApduDataUserMessageBuilder {
	return b
}

func (b *_ApduDataUserMessageBuilder) Build() (ApduDataUserMessage, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataUserMessage.deepCopy(), nil
}

func (b *_ApduDataUserMessageBuilder) MustBuild() ApduDataUserMessage {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataUserMessageBuilder) Done() ApduDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataBuilder().(*_ApduDataBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataUserMessageBuilder) buildForApduData() (ApduData, error) {
	return b.Build()
}

func (b *_ApduDataUserMessageBuilder) DeepCopy() any {
	_copy := b.CreateApduDataUserMessageBuilder().(*_ApduDataUserMessageBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataUserMessageBuilder creates a ApduDataUserMessageBuilder
func (b *_ApduDataUserMessage) CreateApduDataUserMessageBuilder() ApduDataUserMessageBuilder {
	if b == nil {
		return NewApduDataUserMessageBuilder()
	}
	return &_ApduDataUserMessageBuilder{_ApduDataUserMessage: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataUserMessage) GetApciType() uint8 {
	return 0xB
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataUserMessage) GetParent() ApduDataContract {
	return m.ApduDataContract
}

// Deprecated: use the interface for direct cast
func CastApduDataUserMessage(structType any) ApduDataUserMessage {
	if casted, ok := structType.(ApduDataUserMessage); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataUserMessage); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataUserMessage) GetTypeName() string {
	return "ApduDataUserMessage"
}

func (m *_ApduDataUserMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataUserMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataUserMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataUserMessage ApduDataUserMessage, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataUserMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataUserMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataUserMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataUserMessage")
	}

	return m, nil
}

func (m *_ApduDataUserMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataUserMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataUserMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataUserMessage")
		}

		if popErr := writeBuffer.PopContext("ApduDataUserMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataUserMessage")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataUserMessage) IsApduDataUserMessage() {}

func (m *_ApduDataUserMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataUserMessage) deepCopy() *_ApduDataUserMessage {
	if m == nil {
		return nil
	}
	_ApduDataUserMessageCopy := &_ApduDataUserMessage{
		m.ApduDataContract.(*_ApduData).deepCopy(),
	}
	_ApduDataUserMessageCopy.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataUserMessageCopy
}

func (m *_ApduDataUserMessage) String() string {
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
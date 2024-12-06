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

// FirmataCommandSetPinMode is the corresponding interface of FirmataCommandSetPinMode
type FirmataCommandSetPinMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetMode returns Mode (property field)
	GetMode() PinMode
	// IsFirmataCommandSetPinMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataCommandSetPinMode()
	// CreateBuilder creates a FirmataCommandSetPinModeBuilder
	CreateFirmataCommandSetPinModeBuilder() FirmataCommandSetPinModeBuilder
}

// _FirmataCommandSetPinMode is the data-structure of this message
type _FirmataCommandSetPinMode struct {
	FirmataCommandContract
	Pin  uint8
	Mode PinMode
}

var _ FirmataCommandSetPinMode = (*_FirmataCommandSetPinMode)(nil)
var _ FirmataCommandRequirements = (*_FirmataCommandSetPinMode)(nil)

// NewFirmataCommandSetPinMode factory function for _FirmataCommandSetPinMode
func NewFirmataCommandSetPinMode(pin uint8, mode PinMode, response bool) *_FirmataCommandSetPinMode {
	_result := &_FirmataCommandSetPinMode{
		FirmataCommandContract: NewFirmataCommand(response),
		Pin:                    pin,
		Mode:                   mode,
	}
	_result.FirmataCommandContract.(*_FirmataCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FirmataCommandSetPinModeBuilder is a builder for FirmataCommandSetPinMode
type FirmataCommandSetPinModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(pin uint8, mode PinMode) FirmataCommandSetPinModeBuilder
	// WithPin adds Pin (property field)
	WithPin(uint8) FirmataCommandSetPinModeBuilder
	// WithMode adds Mode (property field)
	WithMode(PinMode) FirmataCommandSetPinModeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() FirmataCommandBuilder
	// Build builds the FirmataCommandSetPinMode or returns an error if something is wrong
	Build() (FirmataCommandSetPinMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FirmataCommandSetPinMode
}

// NewFirmataCommandSetPinModeBuilder() creates a FirmataCommandSetPinModeBuilder
func NewFirmataCommandSetPinModeBuilder() FirmataCommandSetPinModeBuilder {
	return &_FirmataCommandSetPinModeBuilder{_FirmataCommandSetPinMode: new(_FirmataCommandSetPinMode)}
}

type _FirmataCommandSetPinModeBuilder struct {
	*_FirmataCommandSetPinMode

	parentBuilder *_FirmataCommandBuilder

	err *utils.MultiError
}

var _ (FirmataCommandSetPinModeBuilder) = (*_FirmataCommandSetPinModeBuilder)(nil)

func (b *_FirmataCommandSetPinModeBuilder) setParent(contract FirmataCommandContract) {
	b.FirmataCommandContract = contract
	contract.(*_FirmataCommand)._SubType = b._FirmataCommandSetPinMode
}

func (b *_FirmataCommandSetPinModeBuilder) WithMandatoryFields(pin uint8, mode PinMode) FirmataCommandSetPinModeBuilder {
	return b.WithPin(pin).WithMode(mode)
}

func (b *_FirmataCommandSetPinModeBuilder) WithPin(pin uint8) FirmataCommandSetPinModeBuilder {
	b.Pin = pin
	return b
}

func (b *_FirmataCommandSetPinModeBuilder) WithMode(mode PinMode) FirmataCommandSetPinModeBuilder {
	b.Mode = mode
	return b
}

func (b *_FirmataCommandSetPinModeBuilder) Build() (FirmataCommandSetPinMode, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FirmataCommandSetPinMode.deepCopy(), nil
}

func (b *_FirmataCommandSetPinModeBuilder) MustBuild() FirmataCommandSetPinMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FirmataCommandSetPinModeBuilder) Done() FirmataCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewFirmataCommandBuilder().(*_FirmataCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_FirmataCommandSetPinModeBuilder) buildForFirmataCommand() (FirmataCommand, error) {
	return b.Build()
}

func (b *_FirmataCommandSetPinModeBuilder) DeepCopy() any {
	_copy := b.CreateFirmataCommandSetPinModeBuilder().(*_FirmataCommandSetPinModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFirmataCommandSetPinModeBuilder creates a FirmataCommandSetPinModeBuilder
func (b *_FirmataCommandSetPinMode) CreateFirmataCommandSetPinModeBuilder() FirmataCommandSetPinModeBuilder {
	if b == nil {
		return NewFirmataCommandSetPinModeBuilder()
	}
	return &_FirmataCommandSetPinModeBuilder{_FirmataCommandSetPinMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataCommandSetPinMode) GetCommandCode() uint8 {
	return 0x4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataCommandSetPinMode) GetParent() FirmataCommandContract {
	return m.FirmataCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataCommandSetPinMode) GetPin() uint8 {
	return m.Pin
}

func (m *_FirmataCommandSetPinMode) GetMode() PinMode {
	return m.Mode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFirmataCommandSetPinMode(structType any) FirmataCommandSetPinMode {
	if casted, ok := structType.(FirmataCommandSetPinMode); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommandSetPinMode); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommandSetPinMode) GetTypeName() string {
	return "FirmataCommandSetPinMode"
}

func (m *_FirmataCommandSetPinMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataCommandContract.(*_FirmataCommand).getLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 8

	// Simple field (mode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_FirmataCommandSetPinMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataCommandSetPinMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataCommand, response bool) (__firmataCommandSetPinMode FirmataCommandSetPinMode, err error) {
	m.FirmataCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataCommandSetPinMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommandSetPinMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pin, err := ReadSimpleField(ctx, "pin", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pin' field"))
	}
	m.Pin = pin

	mode, err := ReadEnumField[PinMode](ctx, "mode", "PinMode", ReadEnum(PinModeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mode' field"))
	}
	m.Mode = mode

	if closeErr := readBuffer.CloseContext("FirmataCommandSetPinMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommandSetPinMode")
	}

	return m, nil
}

func (m *_FirmataCommandSetPinMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandSetPinMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSetPinMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandSetPinMode")
		}

		if err := WriteSimpleField[uint8](ctx, "pin", m.GetPin(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'pin' field")
		}

		if err := WriteSimpleEnumField[PinMode](ctx, "mode", "PinMode", m.GetMode(), WriteEnum[PinMode, uint8](PinMode.GetValue, PinMode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'mode' field")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSetPinMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandSetPinMode")
		}
		return nil
	}
	return m.FirmataCommandContract.(*_FirmataCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataCommandSetPinMode) IsFirmataCommandSetPinMode() {}

func (m *_FirmataCommandSetPinMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataCommandSetPinMode) deepCopy() *_FirmataCommandSetPinMode {
	if m == nil {
		return nil
	}
	_FirmataCommandSetPinModeCopy := &_FirmataCommandSetPinMode{
		m.FirmataCommandContract.(*_FirmataCommand).deepCopy(),
		m.Pin,
		m.Mode,
	}
	_FirmataCommandSetPinModeCopy.FirmataCommandContract.(*_FirmataCommand)._SubType = m
	return _FirmataCommandSetPinModeCopy
}

func (m *_FirmataCommandSetPinMode) String() string {
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
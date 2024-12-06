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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusADU is the corresponding interface of ModbusADU
type ModbusADU interface {
	ModbusADUContract
	ModbusADURequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsModbusADU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusADU()
	// CreateBuilder creates a ModbusADUBuilder
	CreateModbusADUBuilder() ModbusADUBuilder
}

// ModbusADUContract provides a set of functions which can be overwritten by a sub struct
type ModbusADUContract interface {
	// GetResponse() returns a parser argument
	GetResponse() bool
	// IsModbusADU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusADU()
	// CreateBuilder creates a ModbusADUBuilder
	CreateModbusADUBuilder() ModbusADUBuilder
}

// ModbusADURequirements provides a set of functions which need to be implemented by a sub struct
type ModbusADURequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetDriverType returns DriverType (discriminator field)
	GetDriverType() DriverType
}

// _ModbusADU is the data-structure of this message
type _ModbusADU struct {
	_SubType interface {
		ModbusADUContract
		ModbusADURequirements
	}

	// Arguments.
	Response bool
}

var _ ModbusADUContract = (*_ModbusADU)(nil)

// NewModbusADU factory function for _ModbusADU
func NewModbusADU(response bool) *_ModbusADU {
	return &_ModbusADU{Response: response}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusADUBuilder is a builder for ModbusADU
type ModbusADUBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ModbusADUBuilder
	// WithArgResponse sets a parser argument
	WithArgResponse(bool) ModbusADUBuilder
	// AsModbusTcpADU converts this build to a subType of ModbusADU. It is always possible to return to current builder using Done()
	AsModbusTcpADU() ModbusTcpADUBuilder
	// AsModbusRtuADU converts this build to a subType of ModbusADU. It is always possible to return to current builder using Done()
	AsModbusRtuADU() ModbusRtuADUBuilder
	// AsModbusAsciiADU converts this build to a subType of ModbusADU. It is always possible to return to current builder using Done()
	AsModbusAsciiADU() ModbusAsciiADUBuilder
	// Build builds the ModbusADU or returns an error if something is wrong
	PartialBuild() (ModbusADUContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ModbusADUContract
	// Build builds the ModbusADU or returns an error if something is wrong
	Build() (ModbusADU, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusADU
}

// NewModbusADUBuilder() creates a ModbusADUBuilder
func NewModbusADUBuilder() ModbusADUBuilder {
	return &_ModbusADUBuilder{_ModbusADU: new(_ModbusADU)}
}

type _ModbusADUChildBuilder interface {
	utils.Copyable
	setParent(ModbusADUContract)
	buildForModbusADU() (ModbusADU, error)
}

type _ModbusADUBuilder struct {
	*_ModbusADU

	childBuilder _ModbusADUChildBuilder

	err *utils.MultiError
}

var _ (ModbusADUBuilder) = (*_ModbusADUBuilder)(nil)

func (b *_ModbusADUBuilder) WithMandatoryFields() ModbusADUBuilder {
	return b
}

func (b *_ModbusADUBuilder) WithArgResponse(response bool) ModbusADUBuilder {
	b.Response = response
	return b
}

func (b *_ModbusADUBuilder) PartialBuild() (ModbusADUContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusADU.deepCopy(), nil
}

func (b *_ModbusADUBuilder) PartialMustBuild() ModbusADUContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusADUBuilder) AsModbusTcpADU() ModbusTcpADUBuilder {
	if cb, ok := b.childBuilder.(ModbusTcpADUBuilder); ok {
		return cb
	}
	cb := NewModbusTcpADUBuilder().(*_ModbusTcpADUBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ModbusADUBuilder) AsModbusRtuADU() ModbusRtuADUBuilder {
	if cb, ok := b.childBuilder.(ModbusRtuADUBuilder); ok {
		return cb
	}
	cb := NewModbusRtuADUBuilder().(*_ModbusRtuADUBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ModbusADUBuilder) AsModbusAsciiADU() ModbusAsciiADUBuilder {
	if cb, ok := b.childBuilder.(ModbusAsciiADUBuilder); ok {
		return cb
	}
	cb := NewModbusAsciiADUBuilder().(*_ModbusAsciiADUBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ModbusADUBuilder) Build() (ModbusADU, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForModbusADU()
}

func (b *_ModbusADUBuilder) MustBuild() ModbusADU {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusADUBuilder) DeepCopy() any {
	_copy := b.CreateModbusADUBuilder().(*_ModbusADUBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ModbusADUChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusADUBuilder creates a ModbusADUBuilder
func (b *_ModbusADU) CreateModbusADUBuilder() ModbusADUBuilder {
	if b == nil {
		return NewModbusADUBuilder()
	}
	return &_ModbusADUBuilder{_ModbusADU: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusADU(structType any) ModbusADU {
	if casted, ok := structType.(ModbusADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusADU) GetTypeName() string {
	return "ModbusADU"
}

func (m *_ModbusADU) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ModbusADU) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_ModbusADU) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ModbusADUParse[T ModbusADU](ctx context.Context, theBytes []byte, driverType DriverType, response bool) (T, error) {
	return ModbusADUParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), driverType, response)
}

func ModbusADUParseWithBufferProducer[T ModbusADU](driverType DriverType, response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ModbusADUParseWithBuffer[T](ctx, readBuffer, driverType, response)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ModbusADUParseWithBuffer[T ModbusADU](ctx context.Context, readBuffer utils.ReadBuffer, driverType DriverType, response bool) (T, error) {
	v, err := (&_ModbusADU{Response: response}).parse(ctx, readBuffer, driverType, response)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ModbusADU) parse(ctx context.Context, readBuffer utils.ReadBuffer, driverType DriverType, response bool) (__modbusADU ModbusADU, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ModbusADU
	switch {
	case driverType == DriverType_MODBUS_TCP: // ModbusTcpADU
		if _child, err = new(_ModbusTcpADU).parse(ctx, readBuffer, m, driverType, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusTcpADU for type-switch of ModbusADU")
		}
	case driverType == DriverType_MODBUS_RTU: // ModbusRtuADU
		if _child, err = new(_ModbusRtuADU).parse(ctx, readBuffer, m, driverType, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusRtuADU for type-switch of ModbusADU")
		}
	case driverType == DriverType_MODBUS_ASCII: // ModbusAsciiADU
		if _child, err = new(_ModbusAsciiADU).parse(ctx, readBuffer, m, driverType, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusAsciiADU for type-switch of ModbusADU")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [driverType=%v]", driverType)
	}

	if closeErr := readBuffer.CloseContext("ModbusADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusADU")
	}

	return _child, nil
}

func (pm *_ModbusADU) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ModbusADU, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusADU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusADU")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ModbusADU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusADU")
	}
	return nil
}

////
// Arguments Getter

func (m *_ModbusADU) GetResponse() bool {
	return m.Response
}

//
////

func (m *_ModbusADU) IsModbusADU() {}

func (m *_ModbusADU) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusADU) deepCopy() *_ModbusADU {
	if m == nil {
		return nil
	}
	_ModbusADUCopy := &_ModbusADU{
		nil, // will be set by child
		m.Response,
	}
	return _ModbusADUCopy
}
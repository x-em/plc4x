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

// ErrorReportingData is the corresponding interface of ErrorReportingData
type ErrorReportingData interface {
	ErrorReportingDataContract
	ErrorReportingDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsErrorReportingData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingData()
	// CreateBuilder creates a ErrorReportingDataBuilder
	CreateErrorReportingDataBuilder() ErrorReportingDataBuilder
}

// ErrorReportingDataContract provides a set of functions which can be overwritten by a sub struct
type ErrorReportingDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() ErrorReportingCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() ErrorReportingCommandType
	// IsErrorReportingData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingData()
	// CreateBuilder creates a ErrorReportingDataBuilder
	CreateErrorReportingDataBuilder() ErrorReportingDataBuilder
}

// ErrorReportingDataRequirements provides a set of functions which need to be implemented by a sub struct
type ErrorReportingDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() ErrorReportingCommandType
}

// _ErrorReportingData is the data-structure of this message
type _ErrorReportingData struct {
	_SubType interface {
		ErrorReportingDataContract
		ErrorReportingDataRequirements
	}
	CommandTypeContainer ErrorReportingCommandTypeContainer
}

var _ ErrorReportingDataContract = (*_ErrorReportingData)(nil)

// NewErrorReportingData factory function for _ErrorReportingData
func NewErrorReportingData(commandTypeContainer ErrorReportingCommandTypeContainer) *_ErrorReportingData {
	return &_ErrorReportingData{CommandTypeContainer: commandTypeContainer}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorReportingDataBuilder is a builder for ErrorReportingData
type ErrorReportingDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(commandTypeContainer ErrorReportingCommandTypeContainer) ErrorReportingDataBuilder
	// WithCommandTypeContainer adds CommandTypeContainer (property field)
	WithCommandTypeContainer(ErrorReportingCommandTypeContainer) ErrorReportingDataBuilder
	// AsErrorReportingDataGeneric converts this build to a subType of ErrorReportingData. It is always possible to return to current builder using Done()
	AsErrorReportingDataGeneric() ErrorReportingDataGenericBuilder
	// Build builds the ErrorReportingData or returns an error if something is wrong
	PartialBuild() (ErrorReportingDataContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ErrorReportingDataContract
	// Build builds the ErrorReportingData or returns an error if something is wrong
	Build() (ErrorReportingData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorReportingData
}

// NewErrorReportingDataBuilder() creates a ErrorReportingDataBuilder
func NewErrorReportingDataBuilder() ErrorReportingDataBuilder {
	return &_ErrorReportingDataBuilder{_ErrorReportingData: new(_ErrorReportingData)}
}

type _ErrorReportingDataChildBuilder interface {
	utils.Copyable
	setParent(ErrorReportingDataContract)
	buildForErrorReportingData() (ErrorReportingData, error)
}

type _ErrorReportingDataBuilder struct {
	*_ErrorReportingData

	childBuilder _ErrorReportingDataChildBuilder

	err *utils.MultiError
}

var _ (ErrorReportingDataBuilder) = (*_ErrorReportingDataBuilder)(nil)

func (b *_ErrorReportingDataBuilder) WithMandatoryFields(commandTypeContainer ErrorReportingCommandTypeContainer) ErrorReportingDataBuilder {
	return b.WithCommandTypeContainer(commandTypeContainer)
}

func (b *_ErrorReportingDataBuilder) WithCommandTypeContainer(commandTypeContainer ErrorReportingCommandTypeContainer) ErrorReportingDataBuilder {
	b.CommandTypeContainer = commandTypeContainer
	return b
}

func (b *_ErrorReportingDataBuilder) PartialBuild() (ErrorReportingDataContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ErrorReportingData.deepCopy(), nil
}

func (b *_ErrorReportingDataBuilder) PartialMustBuild() ErrorReportingDataContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ErrorReportingDataBuilder) AsErrorReportingDataGeneric() ErrorReportingDataGenericBuilder {
	if cb, ok := b.childBuilder.(ErrorReportingDataGenericBuilder); ok {
		return cb
	}
	cb := NewErrorReportingDataGenericBuilder().(*_ErrorReportingDataGenericBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ErrorReportingDataBuilder) Build() (ErrorReportingData, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForErrorReportingData()
}

func (b *_ErrorReportingDataBuilder) MustBuild() ErrorReportingData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ErrorReportingDataBuilder) DeepCopy() any {
	_copy := b.CreateErrorReportingDataBuilder().(*_ErrorReportingDataBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ErrorReportingDataChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateErrorReportingDataBuilder creates a ErrorReportingDataBuilder
func (b *_ErrorReportingData) CreateErrorReportingDataBuilder() ErrorReportingDataBuilder {
	if b == nil {
		return NewErrorReportingDataBuilder()
	}
	return &_ErrorReportingDataBuilder{_ErrorReportingData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingData) GetCommandTypeContainer() ErrorReportingCommandTypeContainer {
	return m.CommandTypeContainer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_ErrorReportingData) GetCommandType() ErrorReportingCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastErrorReportingCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorReportingData(structType any) ErrorReportingData {
	if casted, ok := structType.(ErrorReportingData); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingData); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingData) GetTypeName() string {
	return "ErrorReportingData"
}

func (m *_ErrorReportingData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ErrorReportingData) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_ErrorReportingData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ErrorReportingDataParse[T ErrorReportingData](ctx context.Context, theBytes []byte) (T, error) {
	return ErrorReportingDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingDataParseWithBufferProducer[T ErrorReportingData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ErrorReportingDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ErrorReportingDataParseWithBuffer[T ErrorReportingData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_ErrorReportingData{}).parse(ctx, readBuffer)
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

func (m *_ErrorReportingData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__errorReportingData ErrorReportingData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsErrorReportingCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[ErrorReportingCommandTypeContainer](ctx, "commandTypeContainer", "ErrorReportingCommandTypeContainer", ReadEnum(ErrorReportingCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[ErrorReportingCommandType](ctx, "commandType", (*ErrorReportingCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ErrorReportingData
	switch {
	case 0 == 0: // ErrorReportingDataGeneric
		if _child, err = new(_ErrorReportingDataGeneric).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ErrorReportingDataGeneric for type-switch of ErrorReportingData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}

	if closeErr := readBuffer.CloseContext("ErrorReportingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingData")
	}

	return _child, nil
}

func (pm *_ErrorReportingData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ErrorReportingData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorReportingData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorReportingData")
	}

	if err := WriteSimpleEnumField[ErrorReportingCommandTypeContainer](ctx, "commandTypeContainer", "ErrorReportingCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[ErrorReportingCommandTypeContainer, uint8](ErrorReportingCommandTypeContainer.GetValue, ErrorReportingCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ErrorReportingData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorReportingData")
	}
	return nil
}

func (m *_ErrorReportingData) IsErrorReportingData() {}

func (m *_ErrorReportingData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorReportingData) deepCopy() *_ErrorReportingData {
	if m == nil {
		return nil
	}
	_ErrorReportingDataCopy := &_ErrorReportingData{
		nil, // will be set by child
		m.CommandTypeContainer,
	}
	return _ErrorReportingDataCopy
}
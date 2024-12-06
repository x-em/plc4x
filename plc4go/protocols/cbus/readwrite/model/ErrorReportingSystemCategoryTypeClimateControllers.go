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

// ErrorReportingSystemCategoryTypeClimateControllers is the corresponding interface of ErrorReportingSystemCategoryTypeClimateControllers
type ErrorReportingSystemCategoryTypeClimateControllers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ErrorReportingSystemCategoryType
	// GetCategoryForType returns CategoryForType (property field)
	GetCategoryForType() ErrorReportingSystemCategoryTypeForClimateControllers
	// IsErrorReportingSystemCategoryTypeClimateControllers is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingSystemCategoryTypeClimateControllers()
	// CreateBuilder creates a ErrorReportingSystemCategoryTypeClimateControllersBuilder
	CreateErrorReportingSystemCategoryTypeClimateControllersBuilder() ErrorReportingSystemCategoryTypeClimateControllersBuilder
}

// _ErrorReportingSystemCategoryTypeClimateControllers is the data-structure of this message
type _ErrorReportingSystemCategoryTypeClimateControllers struct {
	ErrorReportingSystemCategoryTypeContract
	CategoryForType ErrorReportingSystemCategoryTypeForClimateControllers
}

var _ ErrorReportingSystemCategoryTypeClimateControllers = (*_ErrorReportingSystemCategoryTypeClimateControllers)(nil)
var _ ErrorReportingSystemCategoryTypeRequirements = (*_ErrorReportingSystemCategoryTypeClimateControllers)(nil)

// NewErrorReportingSystemCategoryTypeClimateControllers factory function for _ErrorReportingSystemCategoryTypeClimateControllers
func NewErrorReportingSystemCategoryTypeClimateControllers(categoryForType ErrorReportingSystemCategoryTypeForClimateControllers) *_ErrorReportingSystemCategoryTypeClimateControllers {
	_result := &_ErrorReportingSystemCategoryTypeClimateControllers{
		ErrorReportingSystemCategoryTypeContract: NewErrorReportingSystemCategoryType(),
		CategoryForType:                          categoryForType,
	}
	_result.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorReportingSystemCategoryTypeClimateControllersBuilder is a builder for ErrorReportingSystemCategoryTypeClimateControllers
type ErrorReportingSystemCategoryTypeClimateControllersBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(categoryForType ErrorReportingSystemCategoryTypeForClimateControllers) ErrorReportingSystemCategoryTypeClimateControllersBuilder
	// WithCategoryForType adds CategoryForType (property field)
	WithCategoryForType(ErrorReportingSystemCategoryTypeForClimateControllers) ErrorReportingSystemCategoryTypeClimateControllersBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ErrorReportingSystemCategoryTypeBuilder
	// Build builds the ErrorReportingSystemCategoryTypeClimateControllers or returns an error if something is wrong
	Build() (ErrorReportingSystemCategoryTypeClimateControllers, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorReportingSystemCategoryTypeClimateControllers
}

// NewErrorReportingSystemCategoryTypeClimateControllersBuilder() creates a ErrorReportingSystemCategoryTypeClimateControllersBuilder
func NewErrorReportingSystemCategoryTypeClimateControllersBuilder() ErrorReportingSystemCategoryTypeClimateControllersBuilder {
	return &_ErrorReportingSystemCategoryTypeClimateControllersBuilder{_ErrorReportingSystemCategoryTypeClimateControllers: new(_ErrorReportingSystemCategoryTypeClimateControllers)}
}

type _ErrorReportingSystemCategoryTypeClimateControllersBuilder struct {
	*_ErrorReportingSystemCategoryTypeClimateControllers

	parentBuilder *_ErrorReportingSystemCategoryTypeBuilder

	err *utils.MultiError
}

var _ (ErrorReportingSystemCategoryTypeClimateControllersBuilder) = (*_ErrorReportingSystemCategoryTypeClimateControllersBuilder)(nil)

func (b *_ErrorReportingSystemCategoryTypeClimateControllersBuilder) setParent(contract ErrorReportingSystemCategoryTypeContract) {
	b.ErrorReportingSystemCategoryTypeContract = contract
	contract.(*_ErrorReportingSystemCategoryType)._SubType = b._ErrorReportingSystemCategoryTypeClimateControllers
}

func (b *_ErrorReportingSystemCategoryTypeClimateControllersBuilder) WithMandatoryFields(categoryForType ErrorReportingSystemCategoryTypeForClimateControllers) ErrorReportingSystemCategoryTypeClimateControllersBuilder {
	return b.WithCategoryForType(categoryForType)
}

func (b *_ErrorReportingSystemCategoryTypeClimateControllersBuilder) WithCategoryForType(categoryForType ErrorReportingSystemCategoryTypeForClimateControllers) ErrorReportingSystemCategoryTypeClimateControllersBuilder {
	b.CategoryForType = categoryForType
	return b
}

func (b *_ErrorReportingSystemCategoryTypeClimateControllersBuilder) Build() (ErrorReportingSystemCategoryTypeClimateControllers, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ErrorReportingSystemCategoryTypeClimateControllers.deepCopy(), nil
}

func (b *_ErrorReportingSystemCategoryTypeClimateControllersBuilder) MustBuild() ErrorReportingSystemCategoryTypeClimateControllers {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ErrorReportingSystemCategoryTypeClimateControllersBuilder) Done() ErrorReportingSystemCategoryTypeBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewErrorReportingSystemCategoryTypeBuilder().(*_ErrorReportingSystemCategoryTypeBuilder)
	}
	return b.parentBuilder
}

func (b *_ErrorReportingSystemCategoryTypeClimateControllersBuilder) buildForErrorReportingSystemCategoryType() (ErrorReportingSystemCategoryType, error) {
	return b.Build()
}

func (b *_ErrorReportingSystemCategoryTypeClimateControllersBuilder) DeepCopy() any {
	_copy := b.CreateErrorReportingSystemCategoryTypeClimateControllersBuilder().(*_ErrorReportingSystemCategoryTypeClimateControllersBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateErrorReportingSystemCategoryTypeClimateControllersBuilder creates a ErrorReportingSystemCategoryTypeClimateControllersBuilder
func (b *_ErrorReportingSystemCategoryTypeClimateControllers) CreateErrorReportingSystemCategoryTypeClimateControllersBuilder() ErrorReportingSystemCategoryTypeClimateControllersBuilder {
	if b == nil {
		return NewErrorReportingSystemCategoryTypeClimateControllersBuilder()
	}
	return &_ErrorReportingSystemCategoryTypeClimateControllersBuilder{_ErrorReportingSystemCategoryTypeClimateControllers: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetParent() ErrorReportingSystemCategoryTypeContract {
	return m.ErrorReportingSystemCategoryTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetCategoryForType() ErrorReportingSystemCategoryTypeForClimateControllers {
	return m.CategoryForType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeClimateControllers(structType any) ErrorReportingSystemCategoryTypeClimateControllers {
	if casted, ok := structType.(ErrorReportingSystemCategoryTypeClimateControllers); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeClimateControllers); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeClimateControllers"
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).getLengthInBits(ctx))

	// Simple field (categoryForType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ErrorReportingSystemCategoryType, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (__errorReportingSystemCategoryTypeClimateControllers ErrorReportingSystemCategoryTypeClimateControllers, err error) {
	m.ErrorReportingSystemCategoryTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeClimateControllers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeClimateControllers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	categoryForType, err := ReadEnumField[ErrorReportingSystemCategoryTypeForClimateControllers](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForClimateControllers", ReadEnum(ErrorReportingSystemCategoryTypeForClimateControllersByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'categoryForType' field"))
	}
	m.CategoryForType = categoryForType

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeClimateControllers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeClimateControllers")
	}

	return m, nil
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeClimateControllers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeClimateControllers")
		}

		if err := WriteSimpleEnumField[ErrorReportingSystemCategoryTypeForClimateControllers](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForClimateControllers", m.GetCategoryForType(), WriteEnum[ErrorReportingSystemCategoryTypeForClimateControllers, uint8](ErrorReportingSystemCategoryTypeForClimateControllers.GetValue, ErrorReportingSystemCategoryTypeForClimateControllers.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
			return errors.Wrap(err, "Error serializing 'categoryForType' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeClimateControllers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeClimateControllers")
		}
		return nil
	}
	return m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) IsErrorReportingSystemCategoryTypeClimateControllers() {
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) deepCopy() *_ErrorReportingSystemCategoryTypeClimateControllers {
	if m == nil {
		return nil
	}
	_ErrorReportingSystemCategoryTypeClimateControllersCopy := &_ErrorReportingSystemCategoryTypeClimateControllers{
		m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).deepCopy(),
		m.CategoryForType,
	}
	_ErrorReportingSystemCategoryTypeClimateControllersCopy.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType)._SubType = m
	return _ErrorReportingSystemCategoryTypeClimateControllersCopy
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) String() string {
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
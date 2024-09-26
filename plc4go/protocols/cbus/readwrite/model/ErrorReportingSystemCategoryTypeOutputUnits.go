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

// ErrorReportingSystemCategoryTypeOutputUnits is the corresponding interface of ErrorReportingSystemCategoryTypeOutputUnits
type ErrorReportingSystemCategoryTypeOutputUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ErrorReportingSystemCategoryType
	// GetCategoryForType returns CategoryForType (property field)
	GetCategoryForType() ErrorReportingSystemCategoryTypeForOutputUnits
	// IsErrorReportingSystemCategoryTypeOutputUnits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingSystemCategoryTypeOutputUnits()
}

// _ErrorReportingSystemCategoryTypeOutputUnits is the data-structure of this message
type _ErrorReportingSystemCategoryTypeOutputUnits struct {
	ErrorReportingSystemCategoryTypeContract
	CategoryForType ErrorReportingSystemCategoryTypeForOutputUnits
}

var _ ErrorReportingSystemCategoryTypeOutputUnits = (*_ErrorReportingSystemCategoryTypeOutputUnits)(nil)
var _ ErrorReportingSystemCategoryTypeRequirements = (*_ErrorReportingSystemCategoryTypeOutputUnits)(nil)

// NewErrorReportingSystemCategoryTypeOutputUnits factory function for _ErrorReportingSystemCategoryTypeOutputUnits
func NewErrorReportingSystemCategoryTypeOutputUnits(categoryForType ErrorReportingSystemCategoryTypeForOutputUnits) *_ErrorReportingSystemCategoryTypeOutputUnits {
	_result := &_ErrorReportingSystemCategoryTypeOutputUnits{
		ErrorReportingSystemCategoryTypeContract: NewErrorReportingSystemCategoryType(),
		CategoryForType:                          categoryForType,
	}
	_result.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return ErrorReportingSystemCategoryClass_OUTPUT_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) GetParent() ErrorReportingSystemCategoryTypeContract {
	return m.ErrorReportingSystemCategoryTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) GetCategoryForType() ErrorReportingSystemCategoryTypeForOutputUnits {
	return m.CategoryForType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeOutputUnits(structType any) ErrorReportingSystemCategoryTypeOutputUnits {
	if casted, ok := structType.(ErrorReportingSystemCategoryTypeOutputUnits); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeOutputUnits); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeOutputUnits"
}

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).getLengthInBits(ctx))

	// Simple field (categoryForType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ErrorReportingSystemCategoryType, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (__errorReportingSystemCategoryTypeOutputUnits ErrorReportingSystemCategoryTypeOutputUnits, err error) {
	m.ErrorReportingSystemCategoryTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeOutputUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeOutputUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	categoryForType, err := ReadEnumField[ErrorReportingSystemCategoryTypeForOutputUnits](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForOutputUnits", ReadEnum(ErrorReportingSystemCategoryTypeForOutputUnitsByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'categoryForType' field"))
	}
	m.CategoryForType = categoryForType

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeOutputUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeOutputUnits")
	}

	return m, nil
}

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeOutputUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeOutputUnits")
		}

		if err := WriteSimpleEnumField[ErrorReportingSystemCategoryTypeForOutputUnits](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForOutputUnits", m.GetCategoryForType(), WriteEnum[ErrorReportingSystemCategoryTypeForOutputUnits, uint8](ErrorReportingSystemCategoryTypeForOutputUnits.GetValue, ErrorReportingSystemCategoryTypeForOutputUnits.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
			return errors.Wrap(err, "Error serializing 'categoryForType' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeOutputUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeOutputUnits")
		}
		return nil
	}
	return m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) IsErrorReportingSystemCategoryTypeOutputUnits() {
}

func (m *_ErrorReportingSystemCategoryTypeOutputUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

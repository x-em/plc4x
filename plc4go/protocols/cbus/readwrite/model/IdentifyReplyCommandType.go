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

// IdentifyReplyCommandType is the corresponding interface of IdentifyReplyCommandType
type IdentifyReplyCommandType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetUnitType returns UnitType (property field)
	GetUnitType() string
	// IsIdentifyReplyCommandType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandType()
}

// _IdentifyReplyCommandType is the data-structure of this message
type _IdentifyReplyCommandType struct {
	IdentifyReplyCommandContract
	UnitType string
}

var _ IdentifyReplyCommandType = (*_IdentifyReplyCommandType)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandType)(nil)

// NewIdentifyReplyCommandType factory function for _IdentifyReplyCommandType
func NewIdentifyReplyCommandType(unitType string, numBytes uint8) *_IdentifyReplyCommandType {
	_result := &_IdentifyReplyCommandType{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		UnitType:                     unitType,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandType) GetAttribute() Attribute {
	return Attribute_Type
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandType) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandType) GetUnitType() string {
	return m.UnitType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandType(structType any) IdentifyReplyCommandType {
	if casted, ok := structType.(IdentifyReplyCommandType); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandType); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandType) GetTypeName() string {
	return "IdentifyReplyCommandType"
}

func (m *_IdentifyReplyCommandType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Simple field (unitType)
	lengthInBits += 64

	return lengthInBits
}

func (m *_IdentifyReplyCommandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandType IdentifyReplyCommandType, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unitType, err := ReadSimpleField(ctx, "unitType", ReadString(readBuffer, uint32(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitType' field"))
	}
	m.UnitType = unitType

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandType")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandType")
		}

		if err := WriteSimpleField[string](ctx, "unitType", m.GetUnitType(), WriteString(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitType' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandType")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandType) IsIdentifyReplyCommandType() {}

func (m *_IdentifyReplyCommandType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

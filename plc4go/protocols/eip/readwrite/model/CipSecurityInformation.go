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

// CipSecurityInformation is the corresponding interface of CipSecurityInformation
type CipSecurityInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CommandSpecificDataItem
	// GetTodoImplement returns TodoImplement (property field)
	GetTodoImplement() []uint8
	// IsCipSecurityInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipSecurityInformation()
}

// _CipSecurityInformation is the data-structure of this message
type _CipSecurityInformation struct {
	CommandSpecificDataItemContract
	TodoImplement []uint8
}

var _ CipSecurityInformation = (*_CipSecurityInformation)(nil)
var _ CommandSpecificDataItemRequirements = (*_CipSecurityInformation)(nil)

// NewCipSecurityInformation factory function for _CipSecurityInformation
func NewCipSecurityInformation(todoImplement []uint8) *_CipSecurityInformation {
	_result := &_CipSecurityInformation{
		CommandSpecificDataItemContract: NewCommandSpecificDataItem(),
		TodoImplement:                   todoImplement,
	}
	_result.CommandSpecificDataItemContract.(*_CommandSpecificDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipSecurityInformation) GetItemType() uint16 {
	return 0x0086
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipSecurityInformation) GetParent() CommandSpecificDataItemContract {
	return m.CommandSpecificDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipSecurityInformation) GetTodoImplement() []uint8 {
	return m.TodoImplement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipSecurityInformation(structType any) CipSecurityInformation {
	if casted, ok := structType.(CipSecurityInformation); ok {
		return casted
	}
	if casted, ok := structType.(*CipSecurityInformation); ok {
		return *casted
	}
	return nil
}

func (m *_CipSecurityInformation) GetTypeName() string {
	return "CipSecurityInformation"
}

func (m *_CipSecurityInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CommandSpecificDataItemContract.(*_CommandSpecificDataItem).getLengthInBits(ctx))

	// Implicit Field (itemLength)
	lengthInBits += 16

	// Array field
	if len(m.TodoImplement) > 0 {
		lengthInBits += 8 * uint16(len(m.TodoImplement))
	}

	return lengthInBits
}

func (m *_CipSecurityInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipSecurityInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CommandSpecificDataItem) (__cipSecurityInformation CipSecurityInformation, err error) {
	m.CommandSpecificDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipSecurityInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipSecurityInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemLength, err := ReadImplicitField[uint16](ctx, "itemLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemLength' field"))
	}
	_ = itemLength

	todoImplement, err := ReadCountArrayField[uint8](ctx, "todoImplement", ReadUnsignedByte(readBuffer, uint8(8)), uint64(itemLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'todoImplement' field"))
	}
	m.TodoImplement = todoImplement

	if closeErr := readBuffer.CloseContext("CipSecurityInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipSecurityInformation")
	}

	return m, nil
}

func (m *_CipSecurityInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipSecurityInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipSecurityInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipSecurityInformation")
		}
		itemLength := uint16(uint16(len(m.GetTodoImplement())))
		if err := WriteImplicitField(ctx, "itemLength", itemLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemLength' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "todoImplement", m.GetTodoImplement(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'todoImplement' field")
		}

		if popErr := writeBuffer.PopContext("CipSecurityInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipSecurityInformation")
		}
		return nil
	}
	return m.CommandSpecificDataItemContract.(*_CommandSpecificDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipSecurityInformation) IsCipSecurityInformation() {}

func (m *_CipSecurityInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

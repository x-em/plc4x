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

// ModbusPDUWriteMultipleHoldingRegistersResponse is the corresponding interface of ModbusPDUWriteMultipleHoldingRegistersResponse
type ModbusPDUWriteMultipleHoldingRegistersResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// IsModbusPDUWriteMultipleHoldingRegistersResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUWriteMultipleHoldingRegistersResponse()
}

// _ModbusPDUWriteMultipleHoldingRegistersResponse is the data-structure of this message
type _ModbusPDUWriteMultipleHoldingRegistersResponse struct {
	ModbusPDUContract
	StartingAddress uint16
	Quantity        uint16
}

var _ ModbusPDUWriteMultipleHoldingRegistersResponse = (*_ModbusPDUWriteMultipleHoldingRegistersResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUWriteMultipleHoldingRegistersResponse)(nil)

// NewModbusPDUWriteMultipleHoldingRegistersResponse factory function for _ModbusPDUWriteMultipleHoldingRegistersResponse
func NewModbusPDUWriteMultipleHoldingRegistersResponse(startingAddress uint16, quantity uint16) *_ModbusPDUWriteMultipleHoldingRegistersResponse {
	_result := &_ModbusPDUWriteMultipleHoldingRegistersResponse{
		ModbusPDUContract: NewModbusPDU(),
		StartingAddress:   startingAddress,
		Quantity:          quantity,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetFunctionFlag() uint8 {
	return 0x10
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteMultipleHoldingRegistersResponse(structType any) ModbusPDUWriteMultipleHoldingRegistersResponse {
	if casted, ok := structType.(ModbusPDUWriteMultipleHoldingRegistersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteMultipleHoldingRegistersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetTypeName() string {
	return "ModbusPDUWriteMultipleHoldingRegistersResponse"
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUWriteMultipleHoldingRegistersResponse ModbusPDUWriteMultipleHoldingRegistersResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteMultipleHoldingRegistersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteMultipleHoldingRegistersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startingAddress, err := ReadSimpleField(ctx, "startingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingAddress' field"))
	}
	m.StartingAddress = startingAddress

	quantity, err := ReadSimpleField(ctx, "quantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'quantity' field"))
	}
	m.Quantity = quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteMultipleHoldingRegistersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteMultipleHoldingRegistersResponse")
	}

	return m, nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteMultipleHoldingRegistersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteMultipleHoldingRegistersResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "startingAddress", m.GetStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "quantity", m.GetQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteMultipleHoldingRegistersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteMultipleHoldingRegistersResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) IsModbusPDUWriteMultipleHoldingRegistersResponse() {
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

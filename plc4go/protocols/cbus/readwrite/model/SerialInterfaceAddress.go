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

// SerialInterfaceAddress is the corresponding interface of SerialInterfaceAddress
type SerialInterfaceAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAddress returns Address (property field)
	GetAddress() byte
	// IsSerialInterfaceAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSerialInterfaceAddress()
}

// _SerialInterfaceAddress is the data-structure of this message
type _SerialInterfaceAddress struct {
	Address byte
}

var _ SerialInterfaceAddress = (*_SerialInterfaceAddress)(nil)

// NewSerialInterfaceAddress factory function for _SerialInterfaceAddress
func NewSerialInterfaceAddress(address byte) *_SerialInterfaceAddress {
	return &_SerialInterfaceAddress{Address: address}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SerialInterfaceAddress) GetAddress() byte {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSerialInterfaceAddress(structType any) SerialInterfaceAddress {
	if casted, ok := structType.(SerialInterfaceAddress); ok {
		return casted
	}
	if casted, ok := structType.(*SerialInterfaceAddress); ok {
		return *casted
	}
	return nil
}

func (m *_SerialInterfaceAddress) GetTypeName() string {
	return "SerialInterfaceAddress"
}

func (m *_SerialInterfaceAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (address)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SerialInterfaceAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SerialInterfaceAddressParse(ctx context.Context, theBytes []byte) (SerialInterfaceAddress, error) {
	return SerialInterfaceAddressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SerialInterfaceAddressParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SerialInterfaceAddress, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SerialInterfaceAddress, error) {
		return SerialInterfaceAddressParseWithBuffer(ctx, readBuffer)
	}
}

func SerialInterfaceAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SerialInterfaceAddress, error) {
	v, err := (&_SerialInterfaceAddress{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_SerialInterfaceAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__serialInterfaceAddress SerialInterfaceAddress, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SerialInterfaceAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SerialInterfaceAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("SerialInterfaceAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SerialInterfaceAddress")
	}

	return m, nil
}

func (m *_SerialInterfaceAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SerialInterfaceAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SerialInterfaceAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SerialInterfaceAddress")
	}

	if err := WriteSimpleField[byte](ctx, "address", m.GetAddress(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'address' field")
	}

	if popErr := writeBuffer.PopContext("SerialInterfaceAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SerialInterfaceAddress")
	}
	return nil
}

func (m *_SerialInterfaceAddress) IsSerialInterfaceAddress() {}

func (m *_SerialInterfaceAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

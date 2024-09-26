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

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const TPKTPacket_PROTOCOLID uint8 = 0x03

// TPKTPacket is the corresponding interface of TPKTPacket
type TPKTPacket interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPayload returns Payload (property field)
	GetPayload() COTPPacket
	// IsTPKTPacket is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTPKTPacket()
}

// _TPKTPacket is the data-structure of this message
type _TPKTPacket struct {
	Payload COTPPacket
	// Reserved Fields
	reservedField0 *uint8
}

var _ TPKTPacket = (*_TPKTPacket)(nil)

// NewTPKTPacket factory function for _TPKTPacket
func NewTPKTPacket(payload COTPPacket) *_TPKTPacket {
	if payload == nil {
		panic("payload of type COTPPacket for TPKTPacket must not be nil")
	}
	return &_TPKTPacket{Payload: payload}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TPKTPacket) GetPayload() COTPPacket {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_TPKTPacket) GetProtocolId() uint8 {
	return TPKTPacket_PROTOCOLID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTPKTPacket(structType any) TPKTPacket {
	if casted, ok := structType.(TPKTPacket); ok {
		return casted
	}
	if casted, ok := structType.(*TPKTPacket); ok {
		return *casted
	}
	return nil
}

func (m *_TPKTPacket) GetTypeName() string {
	return "TPKTPacket"
}

func (m *_TPKTPacket) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Implicit Field (len)
	lengthInBits += 16

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_TPKTPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TPKTPacketParse(ctx context.Context, theBytes []byte) (TPKTPacket, error) {
	return TPKTPacketParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func TPKTPacketParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TPKTPacket, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TPKTPacket, error) {
		return TPKTPacketParseWithBuffer(ctx, readBuffer)
	}
}

func TPKTPacketParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TPKTPacket, error) {
	v, err := (&_TPKTPacket{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_TPKTPacket) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__tPKTPacket TPKTPacket, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TPKTPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TPKTPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolId, err := ReadConstField[uint8](ctx, "protocolId", ReadUnsignedByte(readBuffer, uint8(8)), TPKTPacket_PROTOCOLID, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolId' field"))
	}
	_ = protocolId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	len, err := ReadImplicitField[uint16](ctx, "len", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'len' field"))
	}
	_ = len

	payload, err := ReadSimpleField[COTPPacket](ctx, "payload", ReadComplex[COTPPacket](COTPPacketParseWithBufferProducer[COTPPacket]((uint16)(uint16(len)-uint16(uint16(4)))), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("TPKTPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TPKTPacket")
	}

	return m, nil
}

func (m *_TPKTPacket) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TPKTPacket) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TPKTPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TPKTPacket")
	}

	if err := WriteConstField(ctx, "protocolId", TPKTPacket_PROTOCOLID, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'protocolId' field")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}
	len := uint16(uint16(m.GetPayload().GetLengthInBytes(ctx)) + uint16(uint16(4)))
	if err := WriteImplicitField(ctx, "len", len, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'len' field")
	}

	if err := WriteSimpleField[COTPPacket](ctx, "payload", m.GetPayload(), WriteComplex[COTPPacket](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}

	if popErr := writeBuffer.PopContext("TPKTPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TPKTPacket")
	}
	return nil
}

func (m *_TPKTPacket) IsTPKTPacket() {}

func (m *_TPKTPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// TunnelingResponseDataBlock is the corresponding interface of TunnelingResponseDataBlock
type TunnelingResponseDataBlock interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetSequenceCounter returns SequenceCounter (property field)
	GetSequenceCounter() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
	// IsTunnelingResponseDataBlock is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTunnelingResponseDataBlock()
}

// _TunnelingResponseDataBlock is the data-structure of this message
type _TunnelingResponseDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
	Status                 Status
}

var _ TunnelingResponseDataBlock = (*_TunnelingResponseDataBlock)(nil)

// NewTunnelingResponseDataBlock factory function for _TunnelingResponseDataBlock
func NewTunnelingResponseDataBlock(communicationChannelId uint8, sequenceCounter uint8, status Status) *_TunnelingResponseDataBlock {
	return &_TunnelingResponseDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter, Status: status}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TunnelingResponseDataBlock) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_TunnelingResponseDataBlock) GetSequenceCounter() uint8 {
	return m.SequenceCounter
}

func (m *_TunnelingResponseDataBlock) GetStatus() Status {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTunnelingResponseDataBlock(structType any) TunnelingResponseDataBlock {
	if casted, ok := structType.(TunnelingResponseDataBlock); ok {
		return casted
	}
	if casted, ok := structType.(*TunnelingResponseDataBlock); ok {
		return *casted
	}
	return nil
}

func (m *_TunnelingResponseDataBlock) GetTypeName() string {
	return "TunnelingResponseDataBlock"
}

func (m *_TunnelingResponseDataBlock) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TunnelingResponseDataBlock) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TunnelingResponseDataBlockParse(ctx context.Context, theBytes []byte) (TunnelingResponseDataBlock, error) {
	return TunnelingResponseDataBlockParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TunnelingResponseDataBlockParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TunnelingResponseDataBlock, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TunnelingResponseDataBlock, error) {
		return TunnelingResponseDataBlockParseWithBuffer(ctx, readBuffer)
	}
}

func TunnelingResponseDataBlockParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TunnelingResponseDataBlock, error) {
	v, err := (&_TunnelingResponseDataBlock{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_TunnelingResponseDataBlock) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__tunnelingResponseDataBlock TunnelingResponseDataBlock, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TunnelingResponseDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TunnelingResponseDataBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	sequenceCounter, err := ReadSimpleField(ctx, "sequenceCounter", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceCounter' field"))
	}
	m.SequenceCounter = sequenceCounter

	status, err := ReadEnumField[Status](ctx, "status", "Status", ReadEnum(StatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	if closeErr := readBuffer.CloseContext("TunnelingResponseDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TunnelingResponseDataBlock")
	}

	return m, nil
}

func (m *_TunnelingResponseDataBlock) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TunnelingResponseDataBlock) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TunnelingResponseDataBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TunnelingResponseDataBlock")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "sequenceCounter", m.GetSequenceCounter(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'sequenceCounter' field")
	}

	if err := WriteSimpleEnumField[Status](ctx, "status", "Status", m.GetStatus(), WriteEnum[Status, uint8](Status.GetValue, Status.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}

	if popErr := writeBuffer.PopContext("TunnelingResponseDataBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TunnelingResponseDataBlock")
	}
	return nil
}

func (m *_TunnelingResponseDataBlock) IsTunnelingResponseDataBlock() {}

func (m *_TunnelingResponseDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

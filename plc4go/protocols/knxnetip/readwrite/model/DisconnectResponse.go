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

// DisconnectResponse is the corresponding interface of DisconnectResponse
type DisconnectResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
	// IsDisconnectResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDisconnectResponse()
}

// _DisconnectResponse is the data-structure of this message
type _DisconnectResponse struct {
	KnxNetIpMessageContract
	CommunicationChannelId uint8
	Status                 Status
}

var _ DisconnectResponse = (*_DisconnectResponse)(nil)
var _ KnxNetIpMessageRequirements = (*_DisconnectResponse)(nil)

// NewDisconnectResponse factory function for _DisconnectResponse
func NewDisconnectResponse(communicationChannelId uint8, status Status) *_DisconnectResponse {
	_result := &_DisconnectResponse{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		CommunicationChannelId:  communicationChannelId,
		Status:                  status,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DisconnectResponse) GetMsgType() uint16 {
	return 0x020A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DisconnectResponse) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DisconnectResponse) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_DisconnectResponse) GetStatus() Status {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDisconnectResponse(structType any) DisconnectResponse {
	if casted, ok := structType.(DisconnectResponse); ok {
		return casted
	}
	if casted, ok := structType.(*DisconnectResponse); ok {
		return *casted
	}
	return nil
}

func (m *_DisconnectResponse) GetTypeName() string {
	return "DisconnectResponse"
}

func (m *_DisconnectResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DisconnectResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DisconnectResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__disconnectResponse DisconnectResponse, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DisconnectResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DisconnectResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	status, err := ReadEnumField[Status](ctx, "status", "Status", ReadEnum(StatusByValue, ReadUnsignedByte(readBuffer, uint8(8))), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	if closeErr := readBuffer.CloseContext("DisconnectResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DisconnectResponse")
	}

	return m, nil
}

func (m *_DisconnectResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DisconnectResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DisconnectResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DisconnectResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
		}

		if err := WriteSimpleEnumField[Status](ctx, "status", "Status", m.GetStatus(), WriteEnum[Status, uint8](Status.GetValue, Status.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if popErr := writeBuffer.PopContext("DisconnectResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DisconnectResponse")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DisconnectResponse) IsDisconnectResponse() {}

func (m *_DisconnectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

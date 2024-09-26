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

// EipPacket is the corresponding interface of EipPacket
type EipPacket interface {
	EipPacketContract
	EipPacketRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsEipPacket is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEipPacket()
}

// EipPacketContract provides a set of functions which can be overwritten by a sub struct
type EipPacketContract interface {
	// GetSessionHandle returns SessionHandle (property field)
	GetSessionHandle() uint32
	// GetStatus returns Status (property field)
	GetStatus() uint32
	// GetSenderContext returns SenderContext (property field)
	GetSenderContext() []byte
	// GetOptions returns Options (property field)
	GetOptions() uint32
	// IsEipPacket is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEipPacket()
}

// EipPacketRequirements provides a set of functions which need to be implemented by a sub struct
type EipPacketRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommand returns Command (discriminator field)
	GetCommand() uint16
	// GetPacketLength returns PacketLength (discriminator field)
	GetPacketLength() uint16
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// _EipPacket is the data-structure of this message
type _EipPacket struct {
	_SubType      EipPacket
	SessionHandle uint32
	Status        uint32
	SenderContext []byte
	Options       uint32
}

var _ EipPacketContract = (*_EipPacket)(nil)

// NewEipPacket factory function for _EipPacket
func NewEipPacket(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipPacket {
	return &_EipPacket{SessionHandle: sessionHandle, Status: status, SenderContext: senderContext, Options: options}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EipPacket) GetSessionHandle() uint32 {
	return m.SessionHandle
}

func (m *_EipPacket) GetStatus() uint32 {
	return m.Status
}

func (m *_EipPacket) GetSenderContext() []byte {
	return m.SenderContext
}

func (m *_EipPacket) GetOptions() uint32 {
	return m.Options
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEipPacket(structType any) EipPacket {
	if casted, ok := structType.(EipPacket); ok {
		return casted
	}
	if casted, ok := structType.(*EipPacket); ok {
		return *casted
	}
	return nil
}

func (m *_EipPacket) GetTypeName() string {
	return "EipPacket"
}

func (m *_EipPacket) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (command)
	lengthInBits += 16

	// Implicit Field (packetLength)
	lengthInBits += 16

	// Simple field (sessionHandle)
	lengthInBits += 32

	// Simple field (status)
	lengthInBits += 32

	// Array field
	if len(m.SenderContext) > 0 {
		lengthInBits += 8 * uint16(len(m.SenderContext))
	}

	// Simple field (options)
	lengthInBits += 32

	return lengthInBits
}

func (m *_EipPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func EipPacketParse[T EipPacket](ctx context.Context, theBytes []byte, response bool) (T, error) {
	return EipPacketParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func EipPacketParseWithBufferProducer[T EipPacket](response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := EipPacketParseWithBuffer[T](ctx, readBuffer, response)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func EipPacketParseWithBuffer[T EipPacket](ctx context.Context, readBuffer utils.ReadBuffer, response bool) (T, error) {
	v, err := (&_EipPacket{}).parse(ctx, readBuffer, response)
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

func (m *_EipPacket) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (__eipPacket EipPacket, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadDiscriminatorField[uint16](ctx, "command", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}

	packetLength, err := ReadImplicitField[uint16](ctx, "packetLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'packetLength' field"))
	}
	_ = packetLength

	sessionHandle, err := ReadSimpleField(ctx, "sessionHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionHandle' field"))
	}
	m.SessionHandle = sessionHandle

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	senderContext, err := readBuffer.ReadByteArray("senderContext", int(int32(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'senderContext' field"))
	}
	m.SenderContext = senderContext

	options, err := ReadSimpleField(ctx, "options", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'options' field"))
	}
	m.Options = options

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child EipPacket
	switch {
	case command == 0x0001 && response == bool(false): // NullCommandRequest
		if _child, err = new(_NullCommandRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NullCommandRequest for type-switch of EipPacket")
		}
	case command == 0x0001 && response == bool(true): // NullCommandResponse
		if _child, err = new(_NullCommandResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NullCommandResponse for type-switch of EipPacket")
		}
	case command == 0x0004 && response == bool(false): // ListServicesRequest
		if _child, err = new(_ListServicesRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ListServicesRequest for type-switch of EipPacket")
		}
	case command == 0x0004 && response == bool(true) && packetLength == uint16(0): // NullListServicesResponse
		if _child, err = new(_NullListServicesResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NullListServicesResponse for type-switch of EipPacket")
		}
	case command == 0x0004 && response == bool(true): // ListServicesResponse
		if _child, err = new(_ListServicesResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ListServicesResponse for type-switch of EipPacket")
		}
	case command == 0x0063 && response == bool(false): // EipListIdentityRequest
		if _child, err = new(_EipListIdentityRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type EipListIdentityRequest for type-switch of EipPacket")
		}
	case command == 0x0063 && response == bool(true): // EipListIdentityResponse
		if _child, err = new(_EipListIdentityResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type EipListIdentityResponse for type-switch of EipPacket")
		}
	case command == 0x0065 && response == bool(false): // EipConnectionRequest
		if _child, err = new(_EipConnectionRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type EipConnectionRequest for type-switch of EipPacket")
		}
	case command == 0x0065 && response == bool(true) && packetLength == uint16(0): // NullEipConnectionResponse
		if _child, err = new(_NullEipConnectionResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NullEipConnectionResponse for type-switch of EipPacket")
		}
	case command == 0x0065 && response == bool(true): // EipConnectionResponse
		if _child, err = new(_EipConnectionResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type EipConnectionResponse for type-switch of EipPacket")
		}
	case command == 0x0066: // EipDisconnectRequest
		if _child, err = new(_EipDisconnectRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type EipDisconnectRequest for type-switch of EipPacket")
		}
	case command == 0x006F: // CipRRData
		if _child, err = new(_CipRRData).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipRRData for type-switch of EipPacket")
		}
	case command == 0x0070: // SendUnitData
		if _child, err = new(_SendUnitData).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SendUnitData for type-switch of EipPacket")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [command=%v, response=%v, packetLength=%v]", command, response, packetLength)
	}

	if closeErr := readBuffer.CloseContext("EipPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipPacket")
	}

	return _child, nil
}

func (pm *_EipPacket) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child EipPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("EipPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EipPacket")
	}

	if err := WriteDiscriminatorField(ctx, "command", m.GetCommand(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'command' field")
	}
	packetLength := uint16(uint16(uint16(m.GetLengthInBytes(ctx))) - uint16(uint16(24)))
	if err := WriteImplicitField(ctx, "packetLength", packetLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'packetLength' field")
	}

	if err := WriteSimpleField[uint32](ctx, "sessionHandle", m.GetSessionHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'sessionHandle' field")
	}

	if err := WriteSimpleField[uint32](ctx, "status", m.GetStatus(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}

	if err := WriteByteArrayField(ctx, "senderContext", m.GetSenderContext(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'senderContext' field")
	}

	if err := WriteSimpleField[uint32](ctx, "options", m.GetOptions(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'options' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("EipPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EipPacket")
	}
	return nil
}

func (m *_EipPacket) IsEipPacket() {}

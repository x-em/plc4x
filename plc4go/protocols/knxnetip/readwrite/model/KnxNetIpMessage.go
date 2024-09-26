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
const KnxNetIpMessage_PROTOCOLVERSION uint8 = 0x10

// KnxNetIpMessage is the corresponding interface of KnxNetIpMessage
type KnxNetIpMessage interface {
	KnxNetIpMessageContract
	KnxNetIpMessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsKnxNetIpMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKnxNetIpMessage()
}

// KnxNetIpMessageContract provides a set of functions which can be overwritten by a sub struct
type KnxNetIpMessageContract interface {
	// IsKnxNetIpMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKnxNetIpMessage()
}

// KnxNetIpMessageRequirements provides a set of functions which need to be implemented by a sub struct
type KnxNetIpMessageRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetMsgType returns MsgType (discriminator field)
	GetMsgType() uint16
}

// _KnxNetIpMessage is the data-structure of this message
type _KnxNetIpMessage struct {
	_SubType KnxNetIpMessage
}

var _ KnxNetIpMessageContract = (*_KnxNetIpMessage)(nil)

// NewKnxNetIpMessage factory function for _KnxNetIpMessage
func NewKnxNetIpMessage() *_KnxNetIpMessage {
	return &_KnxNetIpMessage{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_KnxNetIpMessage) GetProtocolVersion() uint8 {
	return KnxNetIpMessage_PROTOCOLVERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastKnxNetIpMessage(structType any) KnxNetIpMessage {
	if casted, ok := structType.(KnxNetIpMessage); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpMessage); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpMessage) GetTypeName() string {
	return "KnxNetIpMessage"
}

func (m *_KnxNetIpMessage) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8

	// Const Field (protocolVersion)
	lengthInBits += 8
	// Discriminator Field (msgType)
	lengthInBits += 16

	// Implicit Field (totalLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *_KnxNetIpMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func KnxNetIpMessageParse[T KnxNetIpMessage](ctx context.Context, theBytes []byte) (T, error) {
	return KnxNetIpMessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func KnxNetIpMessageParseWithBufferProducer[T KnxNetIpMessage]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := KnxNetIpMessageParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func KnxNetIpMessageParseWithBuffer[T KnxNetIpMessage](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_KnxNetIpMessage{}).parse(ctx, readBuffer)
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

func (m *_KnxNetIpMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__knxNetIpMessage KnxNetIpMessage, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	headerLength, err := ReadImplicitField[uint8](ctx, "headerLength", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'headerLength' field"))
	}
	_ = headerLength

	protocolVersion, err := ReadConstField[uint8](ctx, "protocolVersion", ReadUnsignedByte(readBuffer, uint8(8)), KnxNetIpMessage_PROTOCOLVERSION, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolVersion' field"))
	}
	_ = protocolVersion

	msgType, err := ReadDiscriminatorField[uint16](ctx, "msgType", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'msgType' field"))
	}

	totalLength, err := ReadImplicitField[uint16](ctx, "totalLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'totalLength' field"))
	}
	_ = totalLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child KnxNetIpMessage
	switch {
	case msgType == 0x0201: // SearchRequest
		if _child, err = new(_SearchRequest).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SearchRequest for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0202: // SearchResponse
		if _child, err = new(_SearchResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SearchResponse for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0203: // DescriptionRequest
		if _child, err = new(_DescriptionRequest).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DescriptionRequest for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0204: // DescriptionResponse
		if _child, err = new(_DescriptionResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DescriptionResponse for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0205: // ConnectionRequest
		if _child, err = new(_ConnectionRequest).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionRequest for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0206: // ConnectionResponse
		if _child, err = new(_ConnectionResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionResponse for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0207: // ConnectionStateRequest
		if _child, err = new(_ConnectionStateRequest).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionStateRequest for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0208: // ConnectionStateResponse
		if _child, err = new(_ConnectionStateResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionStateResponse for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0209: // DisconnectRequest
		if _child, err = new(_DisconnectRequest).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DisconnectRequest for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x020A: // DisconnectResponse
		if _child, err = new(_DisconnectResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DisconnectResponse for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x020B: // UnknownMessage
		if _child, err = new(_UnknownMessage).parse(ctx, readBuffer, m, totalLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type UnknownMessage for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0310: // DeviceConfigurationRequest
		if _child, err = new(_DeviceConfigurationRequest).parse(ctx, readBuffer, m, totalLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DeviceConfigurationRequest for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0311: // DeviceConfigurationAck
		if _child, err = new(_DeviceConfigurationAck).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DeviceConfigurationAck for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0420: // TunnelingRequest
		if _child, err = new(_TunnelingRequest).parse(ctx, readBuffer, m, totalLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TunnelingRequest for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0421: // TunnelingResponse
		if _child, err = new(_TunnelingResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TunnelingResponse for type-switch of KnxNetIpMessage")
		}
	case msgType == 0x0530: // RoutingIndication
		if _child, err = new(_RoutingIndication).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RoutingIndication for type-switch of KnxNetIpMessage")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [msgType=%v]", msgType)
	}

	if closeErr := readBuffer.CloseContext("KnxNetIpMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpMessage")
	}

	return _child, nil
}

func (pm *_KnxNetIpMessage) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child KnxNetIpMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("KnxNetIpMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for KnxNetIpMessage")
	}
	headerLength := uint8(uint8(6))
	if err := WriteImplicitField(ctx, "headerLength", headerLength, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'headerLength' field")
	}

	if err := WriteConstField(ctx, "protocolVersion", KnxNetIpMessage_PROTOCOLVERSION, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'protocolVersion' field")
	}

	if err := WriteDiscriminatorField(ctx, "msgType", m.GetMsgType(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'msgType' field")
	}
	totalLength := uint16(uint16(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "totalLength", totalLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'totalLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("KnxNetIpMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for KnxNetIpMessage")
	}
	return nil
}

func (m *_KnxNetIpMessage) IsKnxNetIpMessage() {}

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

// DeviceConfigurationAck is the corresponding interface of DeviceConfigurationAck
type DeviceConfigurationAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetDeviceConfigurationAckDataBlock returns DeviceConfigurationAckDataBlock (property field)
	GetDeviceConfigurationAckDataBlock() DeviceConfigurationAckDataBlock
	// IsDeviceConfigurationAck is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeviceConfigurationAck()
}

// _DeviceConfigurationAck is the data-structure of this message
type _DeviceConfigurationAck struct {
	KnxNetIpMessageContract
	DeviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock
}

var _ DeviceConfigurationAck = (*_DeviceConfigurationAck)(nil)
var _ KnxNetIpMessageRequirements = (*_DeviceConfigurationAck)(nil)

// NewDeviceConfigurationAck factory function for _DeviceConfigurationAck
func NewDeviceConfigurationAck(deviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock) *_DeviceConfigurationAck {
	if deviceConfigurationAckDataBlock == nil {
		panic("deviceConfigurationAckDataBlock of type DeviceConfigurationAckDataBlock for DeviceConfigurationAck must not be nil")
	}
	_result := &_DeviceConfigurationAck{
		KnxNetIpMessageContract:         NewKnxNetIpMessage(),
		DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeviceConfigurationAck) GetMsgType() uint16 {
	return 0x0311
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeviceConfigurationAck) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationAck) GetDeviceConfigurationAckDataBlock() DeviceConfigurationAckDataBlock {
	return m.DeviceConfigurationAckDataBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationAck(structType any) DeviceConfigurationAck {
	if casted, ok := structType.(DeviceConfigurationAck); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationAck); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationAck) GetTypeName() string {
	return "DeviceConfigurationAck"
}

func (m *_DeviceConfigurationAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (deviceConfigurationAckDataBlock)
	lengthInBits += m.DeviceConfigurationAckDataBlock.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DeviceConfigurationAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeviceConfigurationAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__deviceConfigurationAck DeviceConfigurationAck, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceConfigurationAckDataBlock, err := ReadSimpleField[DeviceConfigurationAckDataBlock](ctx, "deviceConfigurationAckDataBlock", ReadComplex[DeviceConfigurationAckDataBlock](DeviceConfigurationAckDataBlockParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceConfigurationAckDataBlock' field"))
	}
	m.DeviceConfigurationAckDataBlock = deviceConfigurationAckDataBlock

	if closeErr := readBuffer.CloseContext("DeviceConfigurationAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationAck")
	}

	return m, nil
}

func (m *_DeviceConfigurationAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceConfigurationAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationAck")
		}

		if err := WriteSimpleField[DeviceConfigurationAckDataBlock](ctx, "deviceConfigurationAckDataBlock", m.GetDeviceConfigurationAckDataBlock(), WriteComplex[DeviceConfigurationAckDataBlock](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceConfigurationAckDataBlock' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeviceConfigurationAck")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeviceConfigurationAck) IsDeviceConfigurationAck() {}

func (m *_DeviceConfigurationAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

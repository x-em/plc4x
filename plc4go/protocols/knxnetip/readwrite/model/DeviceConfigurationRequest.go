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

// DeviceConfigurationRequest is the corresponding interface of DeviceConfigurationRequest
type DeviceConfigurationRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetDeviceConfigurationRequestDataBlock returns DeviceConfigurationRequestDataBlock (property field)
	GetDeviceConfigurationRequestDataBlock() DeviceConfigurationRequestDataBlock
	// GetCemi returns Cemi (property field)
	GetCemi() CEMI
	// IsDeviceConfigurationRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeviceConfigurationRequest()
}

// _DeviceConfigurationRequest is the data-structure of this message
type _DeviceConfigurationRequest struct {
	KnxNetIpMessageContract
	DeviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock
	Cemi                                CEMI

	// Arguments.
	TotalLength uint16
}

var _ DeviceConfigurationRequest = (*_DeviceConfigurationRequest)(nil)
var _ KnxNetIpMessageRequirements = (*_DeviceConfigurationRequest)(nil)

// NewDeviceConfigurationRequest factory function for _DeviceConfigurationRequest
func NewDeviceConfigurationRequest(deviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock, cemi CEMI, totalLength uint16) *_DeviceConfigurationRequest {
	if deviceConfigurationRequestDataBlock == nil {
		panic("deviceConfigurationRequestDataBlock of type DeviceConfigurationRequestDataBlock for DeviceConfigurationRequest must not be nil")
	}
	if cemi == nil {
		panic("cemi of type CEMI for DeviceConfigurationRequest must not be nil")
	}
	_result := &_DeviceConfigurationRequest{
		KnxNetIpMessageContract:             NewKnxNetIpMessage(),
		DeviceConfigurationRequestDataBlock: deviceConfigurationRequestDataBlock,
		Cemi:                                cemi,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeviceConfigurationRequest) GetMsgType() uint16 {
	return 0x0310
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeviceConfigurationRequest) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationRequest) GetDeviceConfigurationRequestDataBlock() DeviceConfigurationRequestDataBlock {
	return m.DeviceConfigurationRequestDataBlock
}

func (m *_DeviceConfigurationRequest) GetCemi() CEMI {
	return m.Cemi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationRequest(structType any) DeviceConfigurationRequest {
	if casted, ok := structType.(DeviceConfigurationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationRequest) GetTypeName() string {
	return "DeviceConfigurationRequest"
}

func (m *_DeviceConfigurationRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (deviceConfigurationRequestDataBlock)
	lengthInBits += m.DeviceConfigurationRequestDataBlock.GetLengthInBits(ctx)

	// Simple field (cemi)
	lengthInBits += m.Cemi.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DeviceConfigurationRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeviceConfigurationRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage, totalLength uint16) (__deviceConfigurationRequest DeviceConfigurationRequest, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceConfigurationRequestDataBlock, err := ReadSimpleField[DeviceConfigurationRequestDataBlock](ctx, "deviceConfigurationRequestDataBlock", ReadComplex[DeviceConfigurationRequestDataBlock](DeviceConfigurationRequestDataBlockParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceConfigurationRequestDataBlock' field"))
	}
	m.DeviceConfigurationRequestDataBlock = deviceConfigurationRequestDataBlock

	cemi, err := ReadSimpleField[CEMI](ctx, "cemi", ReadComplex[CEMI](CEMIParseWithBufferProducer[CEMI]((uint16)(uint16(totalLength)-uint16((uint16(uint16(6))+uint16(deviceConfigurationRequestDataBlock.GetLengthInBytes(ctx)))))), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cemi' field"))
	}
	m.Cemi = cemi

	if closeErr := readBuffer.CloseContext("DeviceConfigurationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationRequest")
	}

	return m, nil
}

func (m *_DeviceConfigurationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceConfigurationRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationRequest")
		}

		if err := WriteSimpleField[DeviceConfigurationRequestDataBlock](ctx, "deviceConfigurationRequestDataBlock", m.GetDeviceConfigurationRequestDataBlock(), WriteComplex[DeviceConfigurationRequestDataBlock](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceConfigurationRequestDataBlock' field")
		}

		if err := WriteSimpleField[CEMI](ctx, "cemi", m.GetCemi(), WriteComplex[CEMI](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'cemi' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeviceConfigurationRequest")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_DeviceConfigurationRequest) GetTotalLength() uint16 {
	return m.TotalLength
}

//
////

func (m *_DeviceConfigurationRequest) IsDeviceConfigurationRequest() {}

func (m *_DeviceConfigurationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

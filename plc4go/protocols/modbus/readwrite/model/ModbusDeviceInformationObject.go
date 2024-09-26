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

// ModbusDeviceInformationObject is the corresponding interface of ModbusDeviceInformationObject
type ModbusDeviceInformationObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetObjectId returns ObjectId (property field)
	GetObjectId() uint8
	// GetData returns Data (property field)
	GetData() []byte
	// IsModbusDeviceInformationObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusDeviceInformationObject()
}

// _ModbusDeviceInformationObject is the data-structure of this message
type _ModbusDeviceInformationObject struct {
	ObjectId uint8
	Data     []byte
}

var _ ModbusDeviceInformationObject = (*_ModbusDeviceInformationObject)(nil)

// NewModbusDeviceInformationObject factory function for _ModbusDeviceInformationObject
func NewModbusDeviceInformationObject(objectId uint8, data []byte) *_ModbusDeviceInformationObject {
	return &_ModbusDeviceInformationObject{ObjectId: objectId, Data: data}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusDeviceInformationObject) GetObjectId() uint8 {
	return m.ObjectId
}

func (m *_ModbusDeviceInformationObject) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusDeviceInformationObject(structType any) ModbusDeviceInformationObject {
	if casted, ok := structType.(ModbusDeviceInformationObject); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusDeviceInformationObject); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusDeviceInformationObject) GetTypeName() string {
	return "ModbusDeviceInformationObject"
}

func (m *_ModbusDeviceInformationObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectId)
	lengthInBits += 8

	// Implicit Field (objectLength)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ModbusDeviceInformationObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusDeviceInformationObjectParse(ctx context.Context, theBytes []byte) (ModbusDeviceInformationObject, error) {
	return ModbusDeviceInformationObjectParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusDeviceInformationObjectParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusDeviceInformationObject, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusDeviceInformationObject, error) {
		return ModbusDeviceInformationObjectParseWithBuffer(ctx, readBuffer)
	}
}

func ModbusDeviceInformationObjectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusDeviceInformationObject, error) {
	v, err := (&_ModbusDeviceInformationObject{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ModbusDeviceInformationObject) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__modbusDeviceInformationObject ModbusDeviceInformationObject, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusDeviceInformationObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusDeviceInformationObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectId, err := ReadSimpleField(ctx, "objectId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectId' field"))
	}
	m.ObjectId = objectId

	objectLength, err := ReadImplicitField[uint8](ctx, "objectLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectLength' field"))
	}
	_ = objectLength

	data, err := readBuffer.ReadByteArray("data", int(objectLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ModbusDeviceInformationObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusDeviceInformationObject")
	}

	return m, nil
}

func (m *_ModbusDeviceInformationObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusDeviceInformationObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusDeviceInformationObject"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusDeviceInformationObject")
	}

	if err := WriteSimpleField[uint8](ctx, "objectId", m.GetObjectId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectId' field")
	}
	objectLength := uint8(uint8(len(m.GetData())))
	if err := WriteImplicitField(ctx, "objectLength", objectLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectLength' field")
	}

	if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("ModbusDeviceInformationObject"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusDeviceInformationObject")
	}
	return nil
}

func (m *_ModbusDeviceInformationObject) IsModbusDeviceInformationObject() {}

func (m *_ModbusDeviceInformationObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
